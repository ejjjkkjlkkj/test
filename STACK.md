# Accessibility Stack

Accessibility is continuous from machine to user.

MACHINE
  ↓
MACHINE SEMANTICS
  ↓
LANGUAGE
  ↓
RUNTIME
  ↓
APPLICATION
  ↓
CONTENT
  ↓
UNIVERSAL INTERACTION
  ↓
MODALITY

There is no special accessibility layer inserted at the end.

## Boundary invariant

Every boundary must preserve:

- meaning
- state
- relations
- capabilities
- events
- safe operations
- accessible projections

If a boundary loses semantic information, the system is considered inaccessible by construction.

## Hardware to language

The language does not directly inherit arbitrary hardware APIs as its semantic foundation.

Instead, machine phenomena become first-class semantic objects.

This permits the same language model to operate across physical machines, virtual machines, embedded systems, and future architectures.
