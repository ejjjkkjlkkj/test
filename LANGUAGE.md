# Language Core

## Axiom 0

Every value has an intrinsic semantic representation.

## Axiom 1

Every semantic representation has an equivalent interaction path.

## Axiom 2

No modality is privileged.

## Axiom 3

Unknown objects must remain inspectable, navigable, and operable.

## Axiom 4

Accessibility cannot be disabled by application code.

## Universal object

An object is defined by:

IDENTITY
MEANING
STRUCTURE
STATE
RELATIONS
CAPABILITIES
EVENTS
OBSERVATIONS
PROJECTIONS

These are semantic properties, not accessibility metadata.

## Projection

A projection maps the same object into a modality:

speech(object)
braille(object)
keyboard(object)
visual(object)
audio(object)
haptic(object)
spatial(object)

A projection may differ in representation but must preserve the object's actionable meaning.

## Unknown-object rule

When an object has no known semantic decoder, the runtime must progressively expose:

1. existence
2. boundaries
3. structure
4. observable properties
5. relations
6. possible actions
7. confidence and uncertainty

The object must never collapse into an inaccessible opaque blob merely because its type is unknown.
