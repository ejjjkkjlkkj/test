# ZERO bootstrap validator

This directory is a bootstrap/reference implementation of the language-neutral ZERO machine format.

It is NOT the ZERO foundation and does not redefine ZERO semantics.

Purpose:

- deterministic parsing;
- canonical encoding;
- validation;
- round-trip verification;
- minimal deterministic transition tests.

The implementation must remain replaceable by a native ZERO implementation.

Validation status must distinguish:

DEFINED
IMPLEMENTED
TESTED
EXECUTED
OBSERVED
PROVEN

A passing bootstrap test proves only the tested software implementation. It does not prove hardware, radio, satellite, or universal physical execution.
