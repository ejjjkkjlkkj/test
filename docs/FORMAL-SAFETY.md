# ZERO — Formal Safety Contract

Status: REQUIRED GATE

## 1. What can be proved

ZERO SHALL NOT claim "infallible" from testing alone.

The target is a machine-checkable chain:

SPECIFICATION -> INVARIANTS -> EXECUTION MODEL -> IMPLEMENTATION -> TEST CAMPAIGN -> INDEPENDENT REPRODUCTION

A test demonstrates an observed case. A proof establishes a property for every state covered by its formal assumptions.

## 2. Safety invariants

I0 — No implicit authority
A trace cannot acquire a capability, device, memory region, network peer, or external effect merely by existing.

I1 — Explicit effects
Every externally observable effect belongs to an explicit effect boundary.

I2 — No hidden mutation
A transition cannot mutate state outside its declared write-set.

I3 — Unknown preservation
UNKNOWN cannot silently become a concrete value.

I4 — History integrity
A transition cannot rewrite immutable historical evidence.

I5 — Isolation
A trace cannot read or write another protection domain without an explicit relation granting that access.

I6 — Deterministic replay
Given identical initial state, input trace, authority set, and execution policy, the verified core has one permitted transition relation.

I7 — Failure containment
An invalid transition is rejected without committing its partial effects.

I8 — Promotion safety
A discovered primitive is never promoted to CORE solely because it passed a benchmark.

I9 — Resource bounds
Every executable transition has an explicit resource accounting model. Exhaustion is a defined result, not undefined behavior.

I10 — Kernel boundary
Unverified native code is outside the trusted proof boundary and cannot be counted as proof of the ZERO core.

## 3. Security properties

For every two domains A and B:

- confidentiality: B cannot observe A unless an explicit information-flow edge exists;
- integrity: B cannot modify A unless an explicit authority edge exists;
- availability: B cannot consume unbounded shared resources without an enforced budget.

These properties are stated as non-interference / capability constraints, not as test expectations.

## 4. Trust boundary

The minimum trusted computing base SHALL be:

1. ZERO formal semantics;
2. proof checker;
3. verified kernel implementation;
4. verified compiler/translator, when compilation is used;
5. verified hardware/ISA assumptions.

Everything else is evidence or an untrusted input.

## 5. Adversarial requirement

For every claimed invariant, the campaign must include:

- direct positive cases;
- minimal counterexamples;
- mutation of one condition at a time;
- deletion of history;
- reordered inputs;
- duplicated inputs;
- contradictory inputs;
- resource exhaustion;
- hostile external effects;
- cross-domain access attempts;
- replay divergence attempts.

## 6. Verdict vocabulary

PROVEN = the formal proof checker accepted the theorem under explicit assumptions.

EXHAUSTIVELY-TESTED = every case in a finite declared state space was executed.

EMPIRICALLY-STRESS-TESTED = a declared number of generated cases executed.

REPRODUCED = an independent implementation reproduced the result.

NOT-PROVEN = any required proof obligation is missing.

ZERO MUST NEVER turn STRESS-TESTED into PROVEN.

## 7. Ultimate claim

The strongest defensible claim is:

"ZERO satisfies the stated security/safety specification for the formally modeled core under the stated machine and environment assumptions."

Absolute real-world infallibility is not a meaningful test result because the physical machine, compiler, firmware, hardware, specification, and environment introduce assumptions outside the language semantics.
