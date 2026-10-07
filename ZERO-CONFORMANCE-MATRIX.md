# ZERO — conformance matrix

This file is the single status matrix for the current bootstrap milestone.

| Requirement | Definition | Implementation | Automated test | Observed proof |
|---|---|---|---|---|
| Canonical RECORD | YES | YES | YES | NOT_OBSERVED |
| UTF-8 escaping | YES | YES | YES | NOT_OBSERVED |
| Decode/encode round-trip | YES | YES | YES | NOT_OBSERVED |
| Missing required field rejection | YES | YES | YES | NOT_OBSERVED |
| UNKNOWN preservation | YES | YES | YES | NOT_OBSERVED |
| Deterministic transition | YES | YES | YES | NOT_OBSERVED |
| Contradiction rejection | YES | YES | YES | NOT_OBSERVED |
| Deferred != delivered | YES | YES | YES | NOT_OBSERVED |
| Transport semantic invariance | YES | CONTRACT/VECTORS | PENDING | NOT_OBSERVED |
| Accessibility projection invariance | YES | CONTRACT | PENDING | NOT_OBSERVED |
| Sequence validation over a stream | YES | NOT_IMPLEMENTED | PENDING | NOT_OBSERVED |
| Unknown-field preservation | YES | NOT_IMPLEMENTED | PENDING | NOT_OBSERVED |
| Serializable event | YES | NOT_IMPLEMENTED | PENDING | NOT_OBSERVED |
| Serializable result | YES | NOT_IMPLEMENTED | PENDING | NOT_OBSERVED |
| Physical execution | YES | NOT_IMPLEMENTED | PENDING | NOT_PROVEN |
| Real radio | YES | NOT_IMPLEMENTED | PENDING | NOT_PROVEN |
| Satellite link | YES | NOT_IMPLEMENTED | PENDING | NOT_PROVEN |

## Rule

No row may be promoted to OBSERVED or PROVEN solely because its specification exists.

A green software test proves only the behavior exercised by that test.

## Bootstrap status

The Go implementation under bootstrap/go is a replaceable reference implementation. It is not the normative ZERO language.
