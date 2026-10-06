# ZERO — Représentation machine native

## Statut
Première étape de matérialisation. Ce document complète les invariants existants.

## 1. Principe
ZERO commence par une représentation de l'état de la machine, pas par une syntaxe de programme.
La représentation doit exprimer une entité, son identité, son état, ses observations, ses changements possibles, ses capacités, ses relations, ses événements, son incertitude et le résultat d'une action.
Une adresse ou position physique n'est jamais l'identité sémantique.

## 2. Noyau minimal
Un état ZERO minimal contient: ENTITY, IDENTITY, STATE, RELATIONS, CAPABILITIES, OBSERVATIONS, EVENTS, RESULT, UNCERTAINTY.
Les projections visuelles, vocales, braille ou autres dérivent de cet état et ne sont jamais la source de vérité.

## 3. Premier état
Le bootstrap représente au minimum la machine, l'exécution actuelle, une région de stockage, le canal d'observation, le canal d'événement et les capacités autorisées.

## 4. Première transition
Première expérience: OBSERVE(target).
Résultat attendu: RESULT + OBSERVATION + EVENT.
Deuxième expérience: REQUEST(intent), avec vérification des capacités avant toute modification.

## 5. Critère de matérialisation
Un état doit pouvoir être chargé, observé, soumis à une intention, vérifié, transformé, publié comme événement, observé à nouveau et interrompu sans perdre son sens.

## 6. Interdictions
La première réalisation ne doit pas dépendre d'un système de fenêtres, DOM, API d'accessibilité, lecteur d'écran externe, modèle de processus OS ou langage existant comme définition de ZERO.

## 7. Suite
ZERO STATE -> ZERO EXECUTOR -> MACHINE BRIDGE -> REAL MACHINE -> I/O -> NETWORK -> WORLD
