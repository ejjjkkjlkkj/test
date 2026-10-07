# ZERO — politique absolue du code

## Règle universelle

Tout code du dépôt, sans exception, doit être :

FUNCTIONAL + ACCESSIBLE + PROVEN

Cette règle inclut le noyau, langage, ISA, runtime, objets, événements, capacités, I/O, stockage, réseau, navigateur, radio, satellite, boot, projections, tests, automatisation et IA.

Pour l'IA, elle inclut modèle, inférence, mémoire, planification, agent, outils et boucle autonome.

## Refus obligatoire

FUNCTIONAL + NOT_ACCESSIBLE = REJECTED
ACCESSIBLE + NOT_FUNCTIONAL = REJECTED
FUNCTIONAL + ACCESSIBLE + NOT_PROVEN = REJECTED
HYPOTHESIS = NEVER_PROOF

Compilation, documentation, modèle, simulation, QEMU, test non exécuté, sortie attendue ou affirmation d'une IA ne constituent pas à eux seuls une preuve de fonctionnement.

## Bout en bout

La preuve suit le chemin réel applicable :

INPUT → PROCESSING → STATE → OUTPUT → OBSERVATION → ACCESSIBILITY → RESULT → PROOF

Pour l'IA :

INPUT → REPRESENTATION → EXECUTION → ACTION → OBSERVATION → VERIFICATION → RESULT → ACCESSIBLE_PROJECTION → PROOF

Une couche interne PASS ne suffit jamais pour déclarer toute la chaîne PASS.

## IA

L'IA est soumise exactement au même contrat que tout autre code.

Elle ne peut pas transformer :
- hypothèse en fait ;
- inférence en fait ;
- prédiction en résultat observé ;
- confiance en preuve ;
- simulation en matériel ;
- QEMU en matériel réel ;
- texte généré en exécution ;
- sa propre affirmation en preuve.

Les états restent séparés :

OBSERVATION / FACT / HYPOTHESIS / DECISION / ACTION / RESULT / PROOF

Et :

GENERATED ≠ EXECUTED ≠ OBSERVED ≠ VERIFIED ≠ PROVEN

## Accessibilité

L'accessibilité est une propriété exécutable.

Chaque fonctionnalité acceptée doit exposer son sens par les modalités requises et conserver exactement identité, opération, données, état, résultat, erreurs et preuve.

Une projection qui échoue fait échouer l'unité.

Une accessibilité uniquement déclarée ou uniquement théorique ne permet pas un PASS physique.

## Vérité déterministe

L'IA ne peut pas modifier une vérité déterministe.

1 + 1 = 2.

1 + 1 = 3 est une erreur/contradiction, jamais une autre vérité.

## Gate

Une unité n'est ACCEPTÉE que si toutes ces conditions sont TRUE :

EXECUTED
OBSERVED
FUNCTIONAL
ACCESSIBILITY_TESTED
ACCESSIBLE
PROVEN

FALSE, UNKNOWN, NOT_PROVEN ou BLOCKED = REFUS.

## Preuve

Les niveaux restent séparés :

DEFINED → IMPLEMENTED → TESTED → SIMULATED → QEMU → HARDWARE → RF_PROVEN → SATELLITE_LINK_PROVEN

Un niveau supérieur exige l'observation correspondant à ce niveau.

## Progression

En cas d'échec : conserver l'échec, conserver la preuve disponible, trouver la première rupture observable, corriger, réexécuter, observer, tester l'accessibilité, puis seulement élever le statut.

Aucune hypothèse ne franchit le gate.
