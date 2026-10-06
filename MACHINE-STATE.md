# ZERO — Machine State

The machine is represented as semantic state, not as opaque registers or devices.

## State

A state contains at minimum:

- existence and identity
- storage
- execution position
- capabilities
- relations
- pending events
- observations
- time/order
- security state

No field is disposable merely because a host implementation does not use it.

## Transition

`STATE + INTENT -> STATE'`

A valid transition declares its source state, operation, required capabilities, resulting state, observable events, and safety result.

The semantic result exists even when the physical operation fails.

## Observation before operation

Observation is distinct from operation. Observation may reveal existence, identity, state, relations, capabilities, uncertainty, and recent events. Observation never implies permission to modify.

## Unknown entities

`UNKNOWN = identity + observation + boundary + uncertainty`

Discovery may extend this representation without replacing it.

## Nondeterminism

Physical nondeterminism enters explicitly as an event or observation. The semantic machine never silently invents a result.

## Accessibility invariant

Every state exposed to the language remains semantically inspectable and navigable. A physical representation may be discarded only after its semantic information has been preserved.

## Bootstrap target

The first executable layer establishes state, transition, storage, event, capability, identity, relation, and observation.
