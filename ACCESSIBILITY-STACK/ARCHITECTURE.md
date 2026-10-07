# ACCESSIBILITY STACK — Architecture v0.1

## Objective

Build an accessibility stack designed by us for our own hardware, with a hardware-first model and explicit evidence. The stack must not depend on a graphical toolkit or screen-reader API as its architectural foundation.

## Chain

HARDWARE → BUS → DEVICE → DRIVER → CAPABILITY → EVENT → SEMANTIC MODEL → ACCESSIBILITY ENGINE → OUTPUT

Each layer has a stable contract and can be replaced independently.

## Principles

- Read-only discovery before control.
- Never infer physical hardware from a software interface alone.
- Evidence is first-class data.
- NOT_PROVEN is different from ABSENT.
- Drivers expose capabilities, not vendor-specific UI.
- Accessibility semantics are independent of Windows UIA/MSAA/ARIA.
- Speech, braille, keyboard and future outputs consume the same semantic event stream.
- No firmware/NVRAM write, driver installation, device removal, reboot, or network reconfiguration during discovery.
- Physical proof and OS proof are distinct.

## Current machine evidence

- Wi-Fi: MediaTek MT7921, PCI PNP evidence, operational interface.
- Bluetooth: MediaTek Bluetooth Adapter, USB PNP evidence.
- Audio: Realtek audio stack present.
- WWAN/GNSS/NFC/RFID/IR/satellite: not proven by the current audit.

## Layer contracts

### Hardware
Identity, bus, topology, electrical/physical evidence.

### Driver
Binding, lifecycle, capabilities, events, diagnostics and safe control methods.

### Capability
A normalized statement such as RADIO.WIFI, RADIO.BLUETOOTH, AUDIO.OUTPUT, INPUT.KEYBOARD.

### Semantic model
Objects, relationships, states, actions, focus, notifications and changes.

### Accessibility engine
Converts semantic state/events into user-observable output.

### Output
Speech, braille, keyboard feedback, display, audio cues and machine-to-machine protocol.

## Non-goals

This document does not claim that an unobserved device exists. It defines the architecture required to discover and support it safely.
