# Native Runtime

The runtime is responsible for making intrinsic accessibility unavoidable.

Pipeline:

SOURCE
-> SEMANTIC OBJECT GRAPH
-> EXECUTION
-> OBSERVATION
-> UNIVERSAL NAVIGATION
-> MODALITY PROJECTION

The runtime owns:

- focus and navigation semantics
- object discovery
- state changes
- event ordering
- interruption
- equivalent actions
- semantic inspection
- uncertainty
- modality conversion

There is no separate accessibility subsystem.

## Interruption

Any active presentation can be interrupted without destroying semantic state.

## Equivalence

If an operation is possible visually, the semantic graph must expose an equivalent operation through the universal interaction model unless the operation itself is inherently unavailable.

## Security

Semantic exposure must respect the object's security boundary. Accessibility never means unauthorized disclosure.
