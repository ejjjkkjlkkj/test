# Execution Model

The language starts below the operating-system abstraction.

## Layers

0. physical machine
1. processor-defined execution state
2. memory and machine state
3. minimal language machine
4. semantic objects
5. runtime services
6. applications

The project must explicitly distinguish these layers.

## No hidden host assumptions

The language specification must not require:

- a POSIX process
- a Windows process
- a JVM
- a JavaScript engine
- a Python interpreter
- a Rust runtime
- a C runtime
- an existing GUI toolkit

These may be used as temporary bootstrap tools, but they are not semantic dependencies.

## Native execution

The eventual runtime must be capable of starting from machine initialization and constructing the semantic environment itself.

## Determinism of meaning

The same machine state and semantic transition must preserve the same meaning regardless of whether the underlying implementation is physical or virtual.
