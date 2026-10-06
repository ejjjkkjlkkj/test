# ZERO — Transition Contract

A transition is the primitive that changes semantic machine state.

## Contract

`T(S, I, C) -> (S', E, R)`

Where S is current state, I is intent, C is available capability, S' is resulting state, E is produced events, and R is the result: success, failure, refusal, or interruption.

## Rules

1. No hidden state mutation.
2. No implicit capability escalation.
3. Every externally observable change produces an event.
4. Failure is a semantic result, never an opaque crash.
5. Interrupted transitions remain observable as interrupted.
6. Uncertainty is data.
7. Equivalent interactions target the same semantic transition regardless of modality.

Keyboard, speech, braille, visual control, automation, and future modalities may all express the same intent such as `ACTIVATE(object-id)`.

Accessibility never weakens security.

The bootstrap must execute and record one transition without requiring an operating-system process model, GUI toolkit, accessibility API, or high-level language runtime.
