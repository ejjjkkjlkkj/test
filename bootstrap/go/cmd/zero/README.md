# ZERO bootstrap CLI

The CLI exposes the deterministic semantic reference pipeline without adding
network, database, OS integration, or hardware dependencies.

## Build

```
go build ./cmd/zero
```

## Input

One canonical ZERO record per stdin line.

## Example

```
echo 'RECORD|1|OBSERVE|id|7|2026-10-07T12:00:00Z|source|target|payload|proof' | ./zero -actor=actor -operation=OBSERVE
```

Output is one JSON receipt per input record.

The CLI is a reference/bootstrap executable. It is not a native CPU runtime
and does not claim hardware execution.
