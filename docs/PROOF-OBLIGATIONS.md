# ZERO Proof Obligations

Each obligation must eventually have a machine-checkable witness.

| ID | Property | Required witness |
|---|---|---|
| P01 | transition preserves state validity | formal invariant proof |
| P02 | invalid transition cannot commit | proof + negative tests |
| P03 | unknown never becomes guessed | proof + metamorphic tests |
| P04 | history is append-only | proof + mutation tests |
| P05 | capability isolation | non-interference proof |
| P06 | information-flow isolation | non-interference proof |
| P07 | replay relation is stable | semantic equivalence proof |
| P08 | resource budget is enforced | bound proof |
| P09 | compiler preserves semantics | compiler correctness proof |
| P10 | native boundary cannot bypass policy | boundary proof |
| P11 | promoted primitives satisfy the contract | promotion proof |
| P12 | proof artifacts correspond to exact source | cryptographic digest binding |

A proof is invalid for ZERO if it depends on a source revision different from
the revision being claimed.

## Required evidence chain

SOURCE_DIGEST
  -> SPEC_DIGEST
  -> PROOF_ARTIFACT
  -> EXECUTION_ARTIFACT
  -> TEST_ARTIFACT
  -> REPRODUCTION_ARTIFACT

Any broken link changes the verdict to NOT-PROVEN.
