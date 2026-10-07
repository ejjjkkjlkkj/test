# ZERO — invariants of the minimal executor

## Validation

The executor has explicit tests for:

- zero-size memory rejection;
- bounded READ/WRITE;
- READ returning a copy;
- deterministic PC advancement;
- rejected instructions leaving PC unchanged;
- failed memory operations leaving PC unchanged;
- deterministic uint64 PC wrap semantics.

The semantic layer additionally verifies authorization before semantic execution, identity-bound transitions, deterministic contradictions, and rejection of unknown semantic operations.

## Native representation boundary

A native `ZER0` record can now be decoded and passed into the same semantic execution pipeline as a canonical RECORD. This prevents the native representation from silently becoming a second execution model.

## Boundary

This remains a bootstrap executor. It proves controlled deterministic execution in the Go reference implementation. It does not prove native CPU execution, hardware execution, network transport, or physical-device support.

## Next architectural step

Extend conformance vectors and execution entry points while preserving one semantic pipeline and one explicit proof boundary.
