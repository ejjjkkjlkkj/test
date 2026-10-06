# ZERO — Machine-First Universal Accessible Language

The language starts as low as possible.

Not at the browser.
Not at the operating system.
Not at a GUI.
Not at a screen reader.
Not at an accessibility API.

The conceptual path is:

PHYSICAL MACHINE
-> CPU / MACHINE STATE
-> MINIMAL PRIMITIVES
-> LANGUAGE CORE
-> SEMANTIC MACHINE
-> RUNTIME
-> APPLICATION
-> CONTENT
-> USER

The language must preserve meaning from the bottom upward.

## Foundational rule

> If accessibility is added after a lower layer has discarded meaning, it is already too late.

Therefore accessibility is defined at the machine-language boundary itself.

## Goal

Build a language and machine environment in which anything representable by the system — including hardware, firmware, applications, content and unknown future objects — has an intrinsic semantic representation and equivalent interaction paths.

The current repository is a specification foundation. The next implementation target is the minimal machine primitive layer, not a conventional application runtime.
