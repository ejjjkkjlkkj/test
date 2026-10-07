# ACCESSIBILITY STACK — Protocol v0.1

## 1. Envelope

Every message uses:

- protocol_version
- message_type
- timestamp
- source
- subject
- evidence[]
- payload
- safety_class

## 2. Device identity

A device record contains:

- stable_id
- pnp_id
- hardware_ids[]
- compatible_ids[]
- bus
- parent
- location
- manufacturer
- model
- driver
- service
- state
- problem_code
- evidence_level

## 3. Capability

A capability is:

CAPABILITY.<domain>.<name>

Examples:

- CAPABILITY.RADIO.WIFI
- CAPABILITY.RADIO.BLUETOOTH
- CAPABILITY.RADIO.WWAN
- CAPABILITY.RADIO.GNSS
- CAPABILITY.RADIO.SATELLITE
- CAPABILITY.RADIO.NFC
- CAPABILITY.RADIO.RFID
- CAPABILITY.RADIO.IR
- CAPABILITY.AUDIO.INPUT
- CAPABILITY.AUDIO.OUTPUT
- CAPABILITY.INPUT.KEYBOARD
- CAPABILITY.DISPLAY.OUTPUT

## 4. State

Allowed discovery states:

- PROVEN
- PRESENT
- AVAILABLE
- ACTIVE
- NOT_PROVEN
- ABSENT
- ERROR
- PHANTOM
- VIRTUAL
- UNKNOWN

## 5. Evidence

P0 = generic software/interface evidence
P1 = interface + operational state
P2 = PNP ID with PCI/USB/bus identity
P3 = PNP + manufacturer/model + driver/service
P4 = PNP + bus + driver + interface + operational correlation
P5 = independent physical evidence

## 6. Events

DEVICE_ADDED
DEVICE_REMOVED
DEVICE_STATE_CHANGED
CAPABILITY_CHANGED
INPUT_EVENT
AUDIO_EVENT
DISPLAY_EVENT
RADIO_STATE_CHANGED
NETWORK_STATE_CHANGED
FOCUS_CHANGED
SEMANTIC_CHANGED
ANNOUNCEMENT_REQUEST

## 7. Actions

Actions must declare:

- requested capability
- target stable_id
- parameters
- authorization
- safety_class
- expected effect

Discovery actions are read-only. Control actions are disabled by default in the audit stack.

## 8. Security

A driver must never silently escalate privileges, alter firmware, alter NVRAM, install itself, remove devices, reboot, or modify network configuration.

## 9. Accessibility invariant

One semantic event stream must feed all outputs. Speech, braille, keyboard and visual output must not maintain incompatible independent states.
