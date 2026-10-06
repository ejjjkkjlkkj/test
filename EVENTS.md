# ZERO — Event Model

Events explain machine change.

## Event shape

Each event contains source identity, sequence/order, time information, meaning, affected identities, resulting observations, and interruption status.

The model distinguishes physical occurrence, observation, and semantic publication. This prevents false precision about timing.

## No silent changes

If semantic state changes and the machine can observe that change, the semantic layer can represent it as an event.

## Input

`physical signal -> observation -> semantic event -> intent -> transition`

## Output

`semantic state -> projection -> physical output`

Speech, braille, keyboard, visual, audio, haptic, and future modalities consume the same semantic event model. No modality receives privileged meaning.
