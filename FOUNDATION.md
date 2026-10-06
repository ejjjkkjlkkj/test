# Foundation — Lowest Possible Layer

The language begins at the lowest layer that can be made independent of an existing operating system, runtime, framework, or accessibility stack.

The target boundary is:

PHYSICAL MACHINE
-> CPU EXECUTION
-> MEMORY
-> MINIMAL MACHINE PRIMITIVES
-> LANGUAGE CORE

The project must not begin at:

- application APIs
- desktop APIs
- browser APIs
- accessibility APIs
- operating-system widgets
- DOM
- screen readers
- existing programming-language runtimes

## What "lowest possible" means

We do not assume that the machine is a collection of ready-made abstractions.

The foundation defines the smallest semantic vocabulary required to observe and safely control computation:

- state
- storage
- transition
- event
- capability
- relation
- identity
- time
- communication

Everything higher is constructed from these primitives.

## Accessibility at the foundation

The lowest layer must already preserve enough meaning for every higher layer to remain accessible.

Accessibility therefore cannot be reconstructed after information has been discarded.

## Bootstrapping

The first implementation may use an existing host only as a temporary laboratory.

The host is not part of the language definition.

The long-term target is a self-hosting machine environment in which the language can establish its own semantic machine boundary.
