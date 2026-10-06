# ZERO — Memory Model

Memory is semantic storage before it is treated as an address space.

A memory region has identity, extent, access properties, observation, relations, ownership/security boundary, persistence characteristic, and uncertainty when applicable.

`ADDRESS != IDENTITY`

A physical or virtual address is a location observation, not the identity of stored meaning. An entity may move while preserving identity.

A read produces an observation. A write is a transition requiring explicit capability. Invalid access is a semantic failure event.

Memory visible to the semantic machine remains inspectable according to capability and security constraints, even when its raw physical representation has no direct meaning.
