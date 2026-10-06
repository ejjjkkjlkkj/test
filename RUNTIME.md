# Native Runtime

The runtime sits on top of a semantic machine, not on top of an inaccessible hardware abstraction.

Pipeline:

PHYSICAL / VIRTUAL MACHINE
-> MACHINE SEMANTICS
-> LANGUAGE
-> SEMANTIC OBJECT GRAPH
-> EXECUTION
-> OBSERVATION
-> UNIVERSAL NAVIGATION
-> MODALITY PROJECTION

The runtime owns:

- machine semantic ingestion
- device discovery
- boot-state representation
- focus and navigation semantics
- object discovery
- state changes
- event ordering
- interruption
- equivalent actions
- semantic inspection
- uncertainty
- modality conversion

There is no separate accessibility subsystem.

## Machine events

Hardware events become semantic events before application logic consumes them.

Examples:

key press -> key event
audio arrival -> audio event
device insertion -> device event
display change -> visual-state event
firmware failure -> boot-failure event

The event retains meaning and provenance.

## Safety

Machine operations carry explicit capability and security semantics. Accessibility never means unrestricted hardware access.
