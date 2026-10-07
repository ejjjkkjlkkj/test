# ZERO — CARTOGRAPHIE VIVANTE DU PROJET

Date: 2026-10-07
Branche: experiment/non-human-zero-v2

## 0. Principe

Cette cartographie décrit le système réel du dépôt, ses frontières, ses dépendances,
ses propriétés, ses preuves et ses zones encore non prouvées. Elle n'est pas une
simple liste de fichiers : elle décrit les flux entre représentation, exécution,
découverte, navigateur, accessibilité, sécurité et machine.

## 1. Couches

### L0 — Langage / représentation
- ZERO.LANGUAGE — primitives, traces, UNKNOWN, auto-extension.
- ZERO.FIELD — substrat expérimental.
- ZERO.EXECUTION — protocole d'exécution.
- ZERO.FIRST-PROGRAM — premier programme expérimental.
- ZERO.LANGUAGE-TEST — tests du langage.

### L1 — Exploration / découverte
- src/generator.rs — génération.
- src/operator.rs — opérateurs expérimentaux.
- src/operator_lang.rs — programmes composés.
- src/discovery.rs — découverte de relations/opérateurs.
- src/search.rs — recherche.
- src/experiment.rs — orchestration expérimentale.
- src/challenge.rs — confrontation.
- src/evaluator.rs — évaluation.
- src/world.rs — observation.
- src/core.rs — boucle de résolution.
- src/memory.rs — mémoire d'expériences.
- src/archive.rs — archivage.

### L2 — Browser / accessibilité
- ZERO.BROWSER — objectif navigateur.
- ZERO.HTML — représentation HTML.
- ZERO.MULTIMODAL — texte, image, graphique et autres modalités.
- ZERO.CAPTCHA — représentation des challenges visuels/complexes.
- browser-core/zero_dom.* — arbre sémantique.
- browser-core/zero_core.* — navigation/focus.
- browser-core/zero_reader.* — sortie lecteur.
- browser-core/zero_media.* — média.
- browser-core/zero_media_test.c — tests média.

### L3 — Exécution native
- browser-core/main.c
- browser-core/Makefile
- browser-core/zero_campaign.c

Cette couche est actuellement un prototype C et ne constitue pas encore une
implémentation formellement raffinée du modèle ZERO.

### L4 — Sécurité / adversaire
- ZERO.ADversarial
- ZERO.FAILURE
- ZERO.MUTATION
- ZERO.EXTREMES
- ZERO.COVERAGE
- ZERO.CROSSINGS
- ZERO.MULTIVERSE
- ZERO.BEYOND-HUMAN
- ZERO.NOVELTY
- ZERO.BENCHMARK

### L5 — Preuve
- proof/ZeroSafety.v
- docs/FORMAL-SAFETY.md
- docs/PROOF-OBLIGATIONS.md
- ZERO.PROOF
- .github/workflows/zero-formal-proof.yml
- .github/workflows/zero-billion-campaign.yml

## 2. Flux de données

```
FIELD / INPUT
     ↓
TRACE / UNKNOWN
     ↓
OBSERVE → TRANSFORM → COLLIDE / SPLIT / MERGE / REINJECT
     ↓
CANDIDATE
     ↓
ADVERSARY + TRANSFER + REPLAY
     ↓
DISCOVERY
     ↓
PROMOTION GATE
     ↓
ZERO CORE
     ↓
EFFECT BOUNDARY
     ├── DOM / navigation
     ├── reader / accessibility
     ├── media
     ├── memory
     ├── network
     ├── GPU / display
     └── external devices
```

## 3. Trust boundary

Trusted only after proof:
- formal semantics;
- proof checker;
- verified core;
- verified translation/compiler;
- explicit hardware/ISA assumptions.

Untrusted by default:
- discovered rules;
- generated programs;
- web content;
- HTML;
- images;
- CAPTCHA;
- network;
- external devices;
- native adapters;
- experimental operators.

## 4. Property map

Current baseline:
I0–I10 in docs/FORMAL-SAFETY.md.

Next stronger layer:
P01 authority monotonicity
P02 complete mediation
P03 capability non-forgeability
P04 information-flow non-interference
P05 state/history immutability
P06 deterministic replay
P07 crash atomicity
P08 resource non-amplification
P09 temporal isolation
P10 provenance preservation
P11 semantic refinement
P12 compiler preservation
P13 device-effect containment
P14 network-peer containment
P15 multimodal accessibility preservation
P16 unknown non-collapse
P17 self-extension quarantine
P18 promotion monotonicity
P19 cross-instance isolation
P20 proof-artifact/source binding
P21 fault containment
P22 recovery convergence
P23 side-channel policy
P24 composition closure

## 5. Current blockers

B1 implementation ↔ formal model refinement missing.
B2 fixed in this revision: event state must be owned by ZeroCore, never global.
B3 executable transition semantics must exactly match ZERO.EXECUTION.
B4 formal coverage of memory, devices, network, GPU/display, parser and hostile effects missing.
B5 compiler/translator correctness missing.
B6 hardware/ISA assumptions not yet connected to the proof.
B7 browser semantics are not yet the same thing as the experimental discovery engine.
B8 accessibility properties are specified conceptually but not yet universally proved.

## 6. Promotion rule

A property cannot become CORE because of benchmarks, fuzzing or a large test count.
Promotion requires:
SPECIFIED → FORMALLY PROVED → REFINED TO IMPLEMENTATION → ADVERSARIAL → REPRODUCED.

## 7. Cartography status

The project is broad but currently heterogeneous: experimental Rust discovery,
C browser prototype, and formal Rocq model are separate evidence domains.
The next architectural goal is to make one semantic contract span all three.
