# Emergent transformations — Gen 2

This experiment removes the fixed named strategy list from the discovery boundary.

The initial seed contains only mechanical primitives. Discovery composes those primitives into executable transformations. A composed transformation is new data: it is not a new Rust enum variant.

## Evidence

A candidate must:
- be absent from the initial seed;
- execute deterministically;
- change at least one observed state;
- survive an independent perturbation case.

Selection is behavioral. No output keyword is used.

## Important limitation

This is still a scaffold. The primitive operations are hand-written and the state model is symbolic. The next experiment must reduce those assumptions and make the discovered transformation space itself mutable.

Run:

cargo run -- emerge

The experiment must remain reproducible and reversible.
