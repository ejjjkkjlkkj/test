use std::env;
use std::path::PathBuf;
use std::time::Instant;
use zero_index::{build_index, write_snapshot, MemoryIndex};

fn main() {
    let mut args = env::args().skip(1);
    let mut output: Option<PathBuf> = None;
    let mut source_root: Option<PathBuf> = None;
    let mut manifests = Vec::new();
    let mut benchmark = false;
    let mut benchmark_report: Option<PathBuf> = None;

    while let Some(arg) = args.next() {
        match arg.as_str() {
            "--out" => output = args.next().map(PathBuf::from),
            "--source-root" => source_root = args.next().map(PathBuf::from),
            "--benchmark" => benchmark = true,
            "--benchmark-report" => benchmark_report = args.next().map(PathBuf::from),
            "--help" | "-h" => {
                println!("zero-index --out <snapshot.json> [--source-root <dir>] [--benchmark] [--benchmark-report <report.json>] <manifest.jsonl>...");
                println!("Read-only. Builds a local index from JSONL Git-tree manifests; does not access the network.");
                return;
            }
            _ if arg.starts_with('-') => {
                eprintln!("[FAIL] Unknown argument: {arg}");
                eprintln!("Use --help for usage.");
                return;
            }
            _ => manifests.push(PathBuf::from(arg)),
        }
    }

    if manifests.is_empty() {
        eprintln!("[FAIL] No JSONL manifests supplied.");
        eprintln!("Usage: zero-index --out snapshot.json <manifest.jsonl>...");
        return;
    }
    let output = output.unwrap_or_else(|| PathBuf::from("zero-index-snapshot.json"));
    let snapshot = match build_index(&manifests, source_root.as_deref()) {
        Ok(s) => s,
        Err(e) => {
            eprintln!("[FAIL] Index build failed: {e}");
            return;
        }
    };
    if let Err(e) = write_snapshot(&snapshot, &output) {
        eprintln!("[FAIL] Could not write snapshot: {e}");
        return;
    }

    println!("[INFO] Snapshot: {}", output.display());
    println!("[INFO] Records: {}", snapshot.records.len());
    println!("[INFO] Unique blobs: {}", snapshot.unique_blob_count);
    println!("[INFO] Duplicate blob occurrences: {}", snapshot.duplicate_blob_count);
    println!("[INFO] Incomplete/unverified: {}", snapshot.incomplete_count);
    println!("[STATUS] {}", snapshot.status);

    if benchmark {
        let index = MemoryIndex::from_snapshot(&snapshot);
        let keys: Vec<String> = snapshot.records.iter()
            .map(|r| format!("{}\\0{}\\0{}", r.origin.repository, r.origin.branch, r.origin.path))
            .collect();
        if keys.is_empty() {
            println!("[BENCH] NOT_RUN: no file records to query.");
            return;
        }
        let rounds = 100_000usize;
        let mut durations = Vec::with_capacity(rounds);
        let total_start = Instant::now();
        for i in 0..rounds {
            let key = &keys[i % keys.len()];
            let start = Instant::now();
            let result = index.lookup_path_key_count(key);
            std::hint::black_box(result);
            durations.push(start.elapsed().as_nanos() as u64);
        }
        let total = total_start.elapsed();
        durations.sort_unstable();
        let percentile = |p: usize| -> u64 {
            let idx = ((durations.len() - 1) * p) / 100;
            durations[idx]
        };
        println!("[BENCH] queries={rounds}");
        println!("[BENCH] p50_ns={}", percentile(50));
        println!("[BENCH] p95_ns={}", percentile(95));
        println!("[BENCH] p99_ns={}", percentile(99));
        println!("[BENCH] max_ns={}", durations[durations.len() - 1]);
        println!("[BENCH] total_ms={}", total.as_millis());
        println!("[BENCH] threshold_100us={}", if percentile(99) <= 100_000 { "PASS_ON_THIS_RUN" } else { "FAIL_ON_THIS_RUN" });
        println!("[BENCH] NOTE=process-local microbenchmark only; excludes freshness reconciliation, IPC, disk, network, and concurrent load.");
        if let Some(report_path) = benchmark_report {
            let report = serde_json::json!({
                "schema_version": 1,
                "status": if percentile(99) <= 100_000 { "PASS_ON_THIS_RUN" } else { "FAIL_ON_THIS_RUN" },
                "scope": "process-local in-memory lookup only",
                "queries": rounds,
                "dataset_records": snapshot.records.len(),
                "p50_ns": percentile(50),
                "p95_ns": percentile(95),
                "p99_ns": percentile(99),
                "max_ns": durations[durations.len() - 1],
                "total_elapsed_ns": total.as_nanos(),
                "threshold_ns": 100_000,
                "exclusions": ["freshness reconciliation", "IPC", "disk", "network", "concurrent load"],
                "snapshot_status": snapshot.status,
                "snapshot_incomplete_count": snapshot.incomplete_count
            });
            if let Some(parent) = report_path.parent() {
                if let Err(error) = std::fs::create_dir_all(parent) {
                    eprintln!("[FAIL] Cannot create benchmark report directory: {error}");
                    return;
                }
            }
            match std::fs::write(&report_path, serde_json::to_vec_pretty(&report).unwrap_or_default()) {
                Ok(()) => println!("[BENCH] report={}", report_path.display()),
                Err(error) => eprintln!("[FAIL] Cannot write benchmark report: {error}"),
            }
        }
    }
}
