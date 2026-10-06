# ZERO — Bootstrap Contract

The bootstrap is the smallest boundary between physical execution and the ZERO semantic machine.

## It must establish

1. machine identity
2. initial machine state
3. storage regions
4. execution position
5. capability boundary
6. event channel
7. observation channel
8. transition mechanism

## It must not assume

- an operating system
- a process
- a filesystem
- a window
- a DOM
- an accessibility API
- a screen reader
- a conventional high-level runtime

## Output

The bootstrap publishes a semantic machine snapshot. The first language layer can immediately inspect, navigate, observe, request a transition, and receive its result.

A host may emulate this contract for development, but it is not the target architecture.

## First proof

`machine -> bootstrap -> semantic state -> observation -> transition -> event -> new semantic state`

This is the first executable semantic cycle.
