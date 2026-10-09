# ZERO: architecture for sub-millisecond decisions and complete repository unification

Status: DESIGN_REQUIRED / NOT_IMPLEMENTED
Scope: all repositories and branches visible to the authenticated account
Safety: read-only discovery; no source deletion, overwrite, firmware write, or automatic merge

## Requirement clarification

A complete scan, semantic audit, test run, and merge of every file cannot finish in 0.1 ms (100 microseconds). That interval is shorter than normal network round trips and cannot guarantee even one remote GitHub API request. GitHub also applies API rate limits and can throttle excessive concurrency. A claim that the entire operation completes in 0.1 ms would be false.

ZERO can target <= 0.1 ms for a *bounded local lookup* against an already-built in-memory index, subject to measurement on the target machine. Full inventory, analysis, builds, and tests must run asynchronously and incrementally. The system must report both latency and completeness; it must never equate a fast cache hit with a completed audit.

## Architecture

1. **Cold-start discovery (not latency bounded):**
   - Enumerate all accessible repositories, branches, tags, submodules, default branches, and commit SHAs.
   - Record API pagination, permissions, rate-limit state, inaccessible repositories, and incomplete/truncated trees.
   - Fetch Git objects using Git protocol/API without modifying remote refs.
   - Store immutable raw manifests and source snapshots keyed by repository, ref, commit SHA, path, and blob SHA.
   - Mark missing or inaccessible data as BLOCKED, never silently omit it.

2. **Content-addressed deduplication:**
   - Key identical Git blobs by object SHA; compute a separate cryptographic digest for exported source bytes where appropriate.
   - Preserve every origin mapping (repository, branch, commit, path, license, provenance).
   - Never discard duplicates merely because their content is identical; retain all provenance and licensing evidence.

3. **Incremental updates:**
   - Use webhooks where available and scheduled reconciliation as a safety net.
   - On a push, compare old and new commit/tree IDs; process changed paths and new blobs only.
   - Invalidate only affected dependency/build/test results.
   - Use bounded worker pools and exponential backoff for API throttling. Do not fan out unlimited concurrent requests.
   - Persist a checkpoint after every completed repository/ref so interrupted runs resume safely.

4. **Semantic audit and component graph:**
   - Detect languages, build systems, tests, licenses, generated files, vendored dependencies, security-sensitive code, hardware/platform abstractions, accessibility APIs, and duplicate implementations.
   - Build a dependency graph and component catalog with evidence pointers to exact commit/path/line ranges.
   - Label findings PASS / FAIL / PARTIAL / BLOCKED / NOT_RUN. Inventory alone is not a semantic audit.

5. **Unification pipeline:**
   - Map each candidate component to a canonical ZERO interface.
   - Classify: REUSE_AS_IS, ADAPTER_REQUIRED, MERGE_CANDIDATE, CONFLICT, LICENSE_REVIEW, or REJECT_WITH_EVIDENCE.
   - Stage integrations in isolated branches/worktrees; preserve upstream sources and provenance.
   - Run format, compile, unit, integration, accessibility, security, and regression tests before any integration is accepted.
   - Never automatically overwrite conflicting implementations. Require explicit policy for licensing conflicts and architecture decisions.

6. **0.1 ms hot path:**
   - Keep a compact index resident in memory (for example, Rust service + immutable snapshot / hash maps).
   - A hot-path request may only perform a bounded lookup: key -> cached status, component ID, provenance pointer, and freshness marker.
   - No network calls, disk scans, parsing, builds, tests, or LLM inference are permitted in the 100-microsecond path.
   - Measure p50/p95/p99 and worst observed latency with a monotonic clock on the actual target machine. Define whether 100 microseconds includes IPC and serialization.
   - Return STALE / UNKNOWN if background inventory is incomplete or the index is older than the relevant repository head.

## Completeness contract

A global status may be COMPLETE only when:
- every accessible repository and requested ref is enumerated;
- every tree is non-truncated or supplemented by a complete local Git walk;
- inaccessible resources and API failures are resolved or explicitly excluded by user policy;
- every tracked file has an inventory record and provenance;
- all files have passed the configured semantic-analysis stages, or are individually marked with a reason;
- integration candidates have test evidence and conflict/license decisions.

The <= 0.1 ms target applies only to hot-path cached lookups, not to global completion. No 99.99% reliability or latency guarantee may be claimed until measured over a defined workload and test duration.

## Required measurements

Record: repository/ref counts; tracked-file and unique-blob counts; bytes indexed; inaccessible refs; truncated trees; retries; API rate-limit remaining/reset; semantic files processed/total; tests passed/failed/not run; cache age; cold-start duration; incremental-update duration; hot-path p50/p95/p99/max; and exact machine/runtime versions.

## Implementation order

1. Correctness-first inventory and checkpoint/resume.
2. Local complete Git trees for any API-truncated refs.
3. Content-addressed source index with provenance.
4. Incremental diff processor and dependency graph.
5. Rust in-memory lookup service and benchmark harness.
6. Per-component adapters and staged integration.
7. CI matrix for independent repository/ref analysis, with bounded concurrency.
8. End-to-end completeness gate; publish artifacts and evidence.

## Current evidence

This document is a design contract only. The existing PowerShell inventory script and GitHub Actions inventory workflow have not been proven by a successful run in this repository. No global semantic audit, complete unification, or 100-microsecond benchmark has been performed.