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

    while let Some(arg) = args.next() {
        match arg.as_str() {
            "--out" => output = args.next().map(PathBuf::from),
            "--source-root" => source_root = args.next().map(PathBuf::from),
            "--benchmark" => benchmark = true,
            "--help" | "-h" => {
                println!("zero-index --out <snapshot.json> [--source-root <dir>] [--benchmark] <manifest.jsonl>...");
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
        let keys: Vec<_> = snapshot.records.iter()
            .map(|r| (r.origin.repository.clone(), r.origin.branch.clone(), r.origin.path.clone()))
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
            let result = index.lookup_path(&key.0, &key.1, &key.2);
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
    }
}
