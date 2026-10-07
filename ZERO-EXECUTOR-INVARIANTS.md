# ZERO — invariants of the minimal executor

## Validation

The executor has explicit tests for zero-size memory rejection, bounded READ/WRITE, READ returning a copy, deterministic PC advancement, rejected instructions leaving PC unchanged, and failed memory operations leaving PC unchanged.

## Boundary

This is still a bootstrap executor. It proves controlled deterministic execution in the Go reference implementation. It does not prove native CPU execution, hardware execution, network transport, or physical-device support.

Next architectural step: model authorization and event emission explicitly while preserving the invariant that failed or rejected operations do not mutate machine state.
