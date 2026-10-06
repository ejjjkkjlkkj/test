# Machine Foundation

Accessibility begins before the application, before the browser, and before the language runtime.

The machine itself is part of the language model.

## Principle

A program must not have to discover accessibility after the machine has already exposed inaccessible hardware or firmware state.

The stack is:

MACHINE
-> MACHINE SEMANTICS
-> LANGUAGE
-> RUNTIME
-> APPLICATION
-> CONTENT
-> USER MODALITY

Accessibility must survive every boundary.

## Machine object

Every machine resource exposed to the language is represented semantically:

- processor
- memory
- storage
- display
- graphics
- audio input/output
- keyboard
- pointing devices
- touch
- network
- camera
- sensors
- firmware
- boot state
- power state
- security state
- buses
- peripherals
- virtual devices
- unknown devices

A resource is never exposed only as an address, register, interrupt, device identifier, pixel buffer, waveform, or opaque driver object.

It has:

IDENTITY
CAPABILITIES
STATE
RELATIONS
EVENTS
OPERATIONS
OBSERVATIONS
SECURITY
ACCESSIBLE REPRESENTATION

## Firmware and boot

The accessibility contract begins at the earliest executable machine layer.

Boot, firmware configuration, errors, diagnostics, device discovery, security prompts, and recovery operations must participate in the same semantic model.

The operating system is therefore not the point where accessibility starts. It is only one consumer of machine semantics.

## Hardware independence

The model must not depend on one vendor, CPU, firmware implementation, display technology, audio codec, or input device.

Hardware-specific mechanisms are translated into stable machine semantics.

## Unknown hardware

An unknown device must remain discoverable.

The machine model progressively exposes:

1. existence
2. identity
3. capabilities
4. state
5. relationships
6. observable behavior
7. safe operations
8. uncertainty

Unknown hardware cannot silently become inaccessible merely because no driver-specific accessibility integration exists.

## Safety

Semantic accessibility does not authorize unsafe hardware operations.

Read, observe, diagnose, simulate, and operate are distinct capabilities.
