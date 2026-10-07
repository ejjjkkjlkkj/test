# DRIVER CONTRACT v0.1

## Purpose

Define our own driver abstraction. A driver is an adapter between a physical/OS device and the normalized capability protocol.

## Required interface

### Identify
Returns immutable identity and topology evidence.

### Probe
Performs read-only capability discovery.

### Start
Starts a controlled session only when explicitly requested.

### Stop
Stops the session without destroying device state.

### Read
Reads device state/data.

### Write
Writes only when the capability explicitly permits it and the caller has authorization.

### Subscribe
Subscribes to normalized device events.

### GetStatus
Returns health, state, latency and error information.

### GetCapabilities
Returns normalized capabilities and evidence.

### Diagnostics
Returns structured diagnostics without modifying the device.

## Driver rules

1. No vendor-specific semantics above the driver boundary.
2. No OCR as a required primitive for accessibility.
3. No GUI automation as the hardware abstraction.
4. No silent fallback from hardware control to unrelated software simulation.
5. Physical, virtual and phantom devices are distinct.
6. Driver claims must reference evidence.
7. A driver must expose deterministic error states.
8. All asynchronous operations must support cancellation.
9. Timing must be measurable.
10. Driver lifecycle must be reversible.

## Async contract

Every asynchronous operation exposes:

- operation_id
- start_time
- completion_time
- cancellation_state
- first_event_time
- final_state
- error

The same model will later be used for audio announcement latency and radio events.

## Initial driver families

- pci
- usb
- acpi
- audio
- input
- display
- network
- wifi
- bluetooth
- wwan
- gnss
- nfc_rfid
- infrared
- satellite

A family can report NOT_PROVEN without implementing unsafe probing.
