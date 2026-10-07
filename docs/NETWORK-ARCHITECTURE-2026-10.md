# Network architecture — hardware-independent semantic layer

Date: 2026-10-07

## Objective

Build one logical communication fabric whose meaning does not depend on the physical carrier.

The semantic layer must remain valid when a carrier is absent, delayed, partitioned, replaced, or physically unavailable.

A physical link is therefore a capability, not an authority.

## Transport families investigated

- I2P: decentralized encrypted overlay network.
- DTN / Bundle Protocol: store-and-forward communication for intermittent connectivity.
- NASA ION-DTN: open-source DTN implementation.
- MeshDTN: transport-independent bundles designed for LoRa, BLE and LAN.
- DARKSOL: research on running cloud-native workloads over disrupted space links using DTN.
- D2D satellite: direct satellite-to-device connectivity.
- SatNOGS / TinyGS: community satellite infrastructure.

These systems solve different parts of the problem. None alone provides the complete target architecture.

## Core rule

The following are semantic:
- identity
- operation
- payload
- proof level
- sequence
- accessibility projection

The following are transport metadata:
- carrier
- route
- link availability
- delivery state
- latency
- retransmission
- physical device

A transport failure MUST NOT rewrite a semantic result.

Example:

    1 + 1
      -> 2
      -> semantic result = VALID

If no route exists:

    delivery = FAILED

The result remains:

    1 + 1 = 2

## Dark / disconnected operation

The architecture explicitly supports operation when the normal Internet is unavailable.

A node may:
1. create a semantic envelope;
2. store it locally;
3. wait for a contact opportunity;
4. forward it through any compatible carrier;
5. receive it later through another carrier.

This is compatible with DTN's store-and-forward model.

The system must not equate NO_LINK with INVALID_MESSAGE.

## Carrier independence

The same envelope can be carried through SOFTWARE, MESH, DTN, I2P, RADIO, SATELLITE or D2D without changing its semantic identity or payload.

The implementation in core/transport.go intentionally contains no network, radio, OS, device, GPU, satellite or runtime dependency.

## Routing direction

    semantic core
          |
    immutable envelope
          |
    capability selection
          |
    mesh / DTN / I2P / satellite
          |
      destination

Routing chooses a carrier. It never chooses the truth.

## Reliability

The engineering target remains 99.99% execution reliability.

This is not a claim that mathematical truth has a 99.99% probability of being true. Mathematical invariants are normative; 99.99% applies to engineered execution and delivery behavior.

## Evidence boundary

Current repository work proves only the semantic/architecture properties implemented and tested in core.

It does NOT yet prove global physical coverage, free global Internet access, satellite D2D interoperability, RF transmission, I2P interoperability, DTN interoperability, or real-world delivery.

Those require separate evidence and must never be silently promoted to HARDWARE, RF_PROVEN, or SATELLITE_LINK_PROVEN.

## External research references

IETF DTN: https://www.rfc-editor.org/rfc/rfc4838.html
NASA ION-DTN: https://github.com/nasa-jpl/ION-DTN
NASA DTN: https://www.nasa.gov/communicating-with-missions/delay-disruption-tolerant-networking/
I2P: https://i2p.net/
MeshDTN: https://gitlab.com/meshdtn/meshdtn
DARKSOL / D3TN: https://d3tn.com/
