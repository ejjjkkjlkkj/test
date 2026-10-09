# ADMWS12 Adapter — implementation contract

Status: IMPLEMENTED PARTIALLY, TESTS ADDED, NOT EXECUTED
Target: ZERO Go bootstrap, package `zero`
Date: 2026-10-09

## Purpose

Translate ADMWS12 platform observations into ZERO canonical records without importing the Python runtime or allowing source states to silently gain stronger semantics.

Source files reviewed on `ejjjkkjlkkj/ADMWS12@main`:
- `src/platform/capability.py`, blob `852c6e852918f71c804ed5a9f766900bcc79bb2f`
- `src/platform/evidence.py`, blob `45f70957a19995685f85d9a55ee64487fc3cf0da`
- `src/platform/hal.py`, blob `dc4f853248baee4368c65e5861e1eedad2707f39`

ZERO source contract:
- `bootstrap/go/record.go`, blob `bb41fdded18a68194f73e6a9e2c7b6420fdc0ca9`
- `bootstrap/go/engine.go`, blob `6666f92e2a2371304cfd5064cc5dea22f8b2268a`

## Mapping rules

| ADMWS12 source kind | Accepted source state | ZERO record type | Required normalized payload state |
|---|---|---|---|
| Capability | UNKNOWN | CAPABILITY | UNKNOWN |
| Capability | AVAILABLE | CAPABILITY | AVAILABLE |
| Capability | UNAVAILABLE | CAPABILITY | UNAVAILABLE |
| Capability | UNSUPPORTED | CAPABILITY | UNSUPPORTED |
| Capability | FAILED | CAPABILITY | FAILED |
| Evidence | OBSERVED | OBSERVATION | OBSERVED |
| Evidence | ABSENT | OBSERVATION | ABSENT |
| Evidence | UNSUPPORTED | OBSERVATION | UNSUPPORTED |
| Evidence | FAILED | OBSERVATION | FAILED |
| Evidence | UNKNOWN | OBSERVATION | UNKNOWN |

Reject empty identities, unknown source kinds, and unknown state values. Do not infer a state from a missing field. Preserve the original source state and source project identity in the payload.

## Proof and safety constraints

1. Set record `Source=ADMWS12` and `Target=ZERO`.
2. Set time to `UNKNOWN` when no trusted source timestamp exists.
3. Set record `Proof=UNPROVEN` for a translation-only adapter. A source claim of AVAILABLE or OBSERVED is not an independent ZERO proof.
4. Do not turn ABSENT into UNAVAILABLE without an explicit capability/evidence policy; they are distinct source semantics.
5. Never convert UNKNOWN, FAILED, or UNSUPPORTED into AVAILABLE.
6. Serialize the mapping payload deterministically. Prefer a fixed-field struct and standard JSON encoder, then test byte-stable output.
7. Keep this adapter in the ZERO Go package. Do not add a second platform state machine or copy ADMWS12's Python HAL wholesale.

## Required tests before implementation is considered complete

- Each accepted capability state and evidence state maps exactly as specified.
- Empty identity, unknown kind, and invalid state are rejected.
- Values containing Unicode, quotes, backslashes, and newlines survive canonical record Encode/Decode.
- Record proof remains UNPROVEN even for AVAILABLE/OBSERVED source states.
- Sequence values remain canonical and round-trip byte-stably.
- No test claims physical hardware verification.
- Run `go test ./...` and `go vet ./...` on the exact commit; attach successful CI evidence before changing status to PASS.

## Gate

Current status: `SPECIFIED / NOT_IMPLEMENTED / NOT_TESTED`.

This contract intentionally does not claim a working adapter. Implementation must follow only after the current ZERO Go test gate produces a recorded result.


## Implementation update (2026-10-09)

- Added `bootstrap/go/admws12_mapping.go` with `MapADMWS12Record`.
- The adapter accepts the explicitly enumerated Capability and Evidence states, rejects unknown kind/state and empty identity, fixes Source/Target/Time, and always sets Proof to `UNPROVEN`.
- Payload uses a fixed field order and ZERO's escaping function. Values are not promoted to proof.
- Added two regression checks to `bootstrap/go/canonical_test.go` for the basic capability mapping and rejection of an invalid state.
- Commits: adapter `f261f76d8caf06c910985c8dd16c15a3e480bf68`; canonical escaping adjustment `e229341ab8ce1eb4b309fab99491ae7e63e37deb`; regression checks `6457fb6f6bd3f3a152fb01150dd4b8d1eb471f21`.
- Validation: **NOT_RUN**. No local Go execution or successful GitHub Actions result has been observed for these commits. Do not mark the adapter PASS until `go test ./...` and `go vet ./...` succeed on the exact commit.
