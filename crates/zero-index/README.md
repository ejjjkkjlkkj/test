# ZERO index (Rust)

This crate builds an offline index from Git tree JSONL manifests and provides in-memory lookups while preserving repository/branch/commit/path/blob provenance.

## Build and test

```powershell
cargo test --manifest-path crates/zero-index/Cargo.toml
cargo build --release --manifest-path crates/zero-index/Cargo.toml
```

## Build a snapshot

Pass one or more `.jsonl` manifests produced by the ZERO inventory workflow or the multi-repository PowerShell inventory script:

```powershell
cargo run --release --manifest-path crates/zero-index/Cargo.toml -- --out audit-output\zero-index-snapshot.json audit-output\multi-repo-*\*\tree-*.jsonl
```

PowerShell does not expand recursive wildcard patterns uniformly for every external program. For many manifest files, collect paths first:

```powershell
$manifests = Get-ChildItem .\audit-output -Recurse -File -Filter *.jsonl | ForEach-Object FullName
cargo run --release --manifest-path crates/zero-index/Cargo.toml -- --out .\audit-output\zero-index-snapshot.json @manifests
```

## Benchmark the cached lookup path

```powershell
cargo run --release --manifest-path crates/zero-index/Cargo.toml -- --benchmark --out .\audit-output\zero-index-snapshot.json @manifests
```

The benchmark reports p50/p95/p99/max for 100,000 process-local lookups. The result only validates that particular run and workload; it does not measure cold-start, network, disk, inter-process communication, freshness checks, or concurrent load.

## Status semantics

- `INDEX_HIT`: an in-memory record was found.
- `NOT_FOUND`: no matching record exists in this snapshot.
- `UNVERIFIED_REQUIRES_HEAD_RECONCILIATION`: snapshot freshness is not established.
- `INVENTORIED_NOT_CONTENT_VERIFIED`: tree metadata is present, but local file bytes have not been hashed.
- `PARTIAL_INVENTORY_OR_CONTENT_UNVERIFIED`: one or more rows or content checks are incomplete.

An inventory is not a semantic code audit. The index does not merge, delete, modify, or push source files.
