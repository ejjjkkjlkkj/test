# ZERO — Capability Model

Capabilities define what an entity or execution context may safely observe or change.

## Dimensions

A capability may authorize observe, inspect, navigate, create, modify, execute, communicate, operate hardware, or administer security.

These are semantic rights, not UI permissions.

`CAN_OBSERVE != CAN_MODIFY`

`CAN_NAVIGATE != CAN_OPERATE`

`CAN_DESCRIBE != CAN_EXECUTE`

An accessible interface must never silently convert one into another.

A denied operation remains represented as intent, required capability, denial, and reason.

The first machine layer has an explicit capability boundary. No ambient authority is assumed.
