# EXPERIMENT — non-human machine intelligence

This directory is an isolated experiment. It does not modify the existing Rust project.

Constraint:
- Python standard library only.
- No pretrained model.
- No external AI/API.
- No human cognitive architecture.
- No fixed task-specific solver.

Idea:
The machine does not begin with concepts, goals, language, or a predefined reasoning algorithm.

It receives streams of symbols and continuously creates, combines, mutates, tests, rejects, and preserves transformations according to one primitive rule:

**a transformation survives when it makes future observations more predictable without being explicitly told what prediction means.**

The experiment is deliberately strange. The useful architecture is expected to emerge from repeated selection rather than from a human-designed chain of reasoning.

Run:

    python experiment.py

The program prints the surviving internal structures and their behavior.
