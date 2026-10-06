# Accessibility Foundation

Accessibility begins at the lowest semantic boundary.

## Information preservation

A lower layer must not destroy information required to represent the object meaningfully at a higher layer.

## Universal representation

Every primitive has:

- semantic identity
- observable state
- safe capabilities
- events
- relations
- projections

## No privileged modality

The machine does not assume that vision is primary.

Pixels, sound, tactile output, keyboard input, speech, and other modalities are physical channels. Meaning exists above those channels.

## Failure accessibility

Machine initialization errors, hardware failures, security states, resource exhaustion, and recovery states are semantic machine states.

They must be observable and navigable through the same universal model.

## Unknown hardware

Unknown hardware is represented as an entity with measured observations and uncertainty, rather than as an inaccessible failure.

## Core rule

No lower layer may emit an opaque representation when a semantic representation can be preserved.
