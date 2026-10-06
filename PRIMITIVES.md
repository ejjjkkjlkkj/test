# Machine Primitives

The primitive vocabulary is intentionally smaller than a conventional operating-system API.

## State

A machine is a changing state.

STATE contains observable machine facts without prescribing a particular CPU vendor.

## Transition

A transition changes state.

TRANSITION has:

- preconditions
- operation
- resulting state
- observable effects
- safety boundary

## Event

An event is an observable state transition or external occurrence.

Events preserve:

- source
- time/order
- meaning
- affected objects
- resulting state

## Storage

Storage represents persistent or transient machine state.

It is not inherently a file system.

## Communication

Communication represents information crossing a machine boundary.

It is not inherently a network socket, device API, or process API.

## Capability

A capability states what an entity can safely observe or perform.

Capabilities are explicit and cannot be inferred solely from a presentation modality.

## Identity

Every machine entity has a stable semantic identity independent of its physical address when possible.

## Relation

Relations connect entities without forcing them into a particular data structure.

These primitives are sufficient to construct higher-level machine semantics.
