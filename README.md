# ZERO — Universal Accessible Machine Language

This project starts from the machine itself.

The objective is not to make an inaccessible machine accessible later through an application, screen reader, browser, or accessibility API.

The objective is to design a language and machine model in which accessibility exists from the first layer.

## Core thesis

> Accessibility begins at the machine.

The stack is:

MACHINE
-> MACHINE SEMANTICS
-> LANGUAGE
-> RUNTIME
-> APPLICATION
-> CONTENT
-> USER

Every layer must preserve semantic accessibility.

## Universal requirement

Anything the system can represent, execute, observe, display, hear, receive, transmit, store, or control must have an intrinsic semantic representation and an equivalent interaction path.

This includes hardware, firmware, operating environments, applications, websites, images, photographs, graphics, video, audio, games, CAPTCHA, unknown objects, and future technologies.

## No afterthought

Accessibility is not a plugin.
It is not a screen-reader API.
It is not a browser feature.
It is not a UI layer.

It is a property of the machine-language model itself.

## Current objective

Design the machine semantics and language core first, then implement the smallest native execution model capable of proving the invariants.
