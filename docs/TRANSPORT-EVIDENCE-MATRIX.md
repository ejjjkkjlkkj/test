# Transport evidence matrix

Updated: 2026-10-07

| Transport | Exists as technology | Integrated in core | Real-link proof in this repository |
|---|---:|---:|---:|
| SOFTWARE | YES | YES | NO |
| MESH | YES | YES (abstract capability) | NO |
| DTN | YES | YES (abstract capability) | NO |
| I2P | YES | YES (abstract capability) | NO |
| RADIO | YES | YES (abstract capability) | NO |
| SATELLITE | YES | YES (abstract capability) | NO |
| D2D | YES | YES (abstract capability) | NO |

## Important boundary

core deliberately models transports as capabilities and metadata.

It does not claim that selecting SATELLITE means that a satellite link has been established.

It does not claim that selecting RADIO proves RF transmission.

It does not claim that selecting I2P proves an I2P session.

It does not claim that selecting DTN proves end-to-end bundle delivery.

The proof ladder remains authoritative:

    UNKNOWN
      -> DEFINED
      -> IMPLEMENTED
      -> TESTED
      -> SIMULATED
      -> QEMU
      -> HARDWARE
      -> RF_PROVEN
      -> SATELLITE_LINK_PROVEN

A transport capability is not evidence.

## Why DTN is a strong architectural fit

IETF RFC 9171 defines Bundle Protocol v7 as a store-carry-forward overlay that can operate with intermittent connectivity and underlying constituent networks through convergence-layer adapters. That separation closely matches the project's semantic/physical boundary.

Source: https://www.rfc-editor.org/rfc/rfc9171.html

## Current external evidence

MeshDTN documents transport-independent bundle serialization across LoRa, BLE and LAN.

Source: https://gitlab.com/meshdtn/meshdtn/-/blob/main/MeshDTN-Protocol.md

These sources are external evidence for architectural compatibility only. They are not proof that this repository currently communicates over those transports.
