# ZERO — Minimal Machine Operations

The initial operation vocabulary is deliberately tiny.

`OBSERVE(x)` — return the representable semantic observation of x.

`RELATE(x, relation)` — follow a declared relation without assuming a UI tree.

`REQUEST(intent)` — request a semantic transition.

`NEXT_EVENT()` — return the next published event according to machine ordering.

`READ(region, position, extent)` — observe storage.

`WRITE(region, position, value)` — request a storage transition subject to capability.

`STEP()` — advance the machine execution boundary by one semantic step.

These are semantic operations. A physical implementation may map them to processor instructions, firmware calls, device operations, or a host emulator, but that mapping must not alter their meaning.
