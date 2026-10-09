# ZERO performance and physical-speed evidence contract

## Non-negotiable claim boundary

ZERO will not claim that software, data, or information travels faster than light in vacuum. The SI value is exactly 299,792,458 m/s; the meter is defined using that value. See NIST:
- https://www.nist.gov/si-redefinition/meet-constants
- https://www.nist.gov/si-redefinition/meter

A computer can reduce avoidable overhead, parallelize independent work, reuse prior results, and answer a cached local query very quickly. None of those changes establishes superluminal information transfer.

## Performance goals

1. Hot path: target <= 100 microseconds p99 for a preloaded, process-local lookup with precomposed key, measured in release mode.
2. Report p50, p95, p99, maximum, throughput, dataset size, CPU/runtime, and benchmark command.
3. Separately measure cold start, disk load, source hashing, repository discovery, network fetch, semantic analysis, builds, and tests.
4. Never blend the fast lookup result with end-to-end completion latency.
5. A benchmark PASS applies only to the measured run and workload. It is not a universal guarantee.
6. Never mark a repository/ref complete while branches, trees, blobs, or analysis steps are missing, inaccessible, truncated, stale, or untested.

## Proof artifacts required for a performance claim

- Source commit SHA and clean/dirty working-tree state.
- Compiler version and release profile.
- Machine model, CPU, RAM, OS, power mode, and process affinity where available.
- Exact command and input dataset SHA-256.
- Warm-up policy and at least 100,000 measured operations.
- p50/p95/p99/max latency, throughput, and repeated-run variance.
- Explicit inclusion/exclusion of timer overhead, allocations, IPC, disk, network, freshness checks, and contention.
- Raw machine-readable benchmark output preserved as an artifact.

## Whole-repository completeness contract

The global job must inventory all selected repositories, branches, tags and submodules; retain source provenance and licensing; reconcile each ref with its current head; process changes incrementally; and emit an explicit status for every file and analysis stage. A fast cache hit cannot stand in for a completed global audit.

## Current state

This contract is documentation, not experimental evidence. The Rust lookup benchmark workflow has been added, but no successful workflow run or local benchmark result is established by this file. The 100-microsecond target remains unverified until actual artifacts are collected and inspected.
