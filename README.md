# ZERO — Universal Accessibility Language

This repository is restarted around one invariant:

> Every thing expressible, executed, observed, produced, or encountered by the language is intrinsically accessible.

Accessibility is not an API, library, annotation, plugin, screen-reader mode, or optional layer. It is part of the language model itself.

## Design constraints

- No JavaScript.
- No Python.
- No dependency on Rust, C, LLVM, DOM, UIA, MSAA, ARIA, Electron, Qt, GTK, or another existing accessibility foundation.
- The language is not modeled as a human-facing programming language first.
- Accessibility must remain valid for unknown and future object types.
- One semantic object may be projected to speech, braille, keyboard, visual, audio, haptic, or another interface without changing its meaning.

## First target

Build the language model first. Then build its native execution model. Then use it to construct an accessible browser and its universal perception/interaction engine.

Nothing is considered accessible merely because a screen reader can announce it.
