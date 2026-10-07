# ZERO — contrat central : vérité, accessibilité, indépendance matérielle

Ce contrat est normatif pour le cœur.

## A. Vérité

La définition formelle est indépendante de toute réalisation physique.

Pour toute réalisation conforme au même système formel :

`1 + 1 = 2`

La réalisation ne peut pas remplacer ce résultat par `3`.

Un résultat différent est une erreur ou une contradiction à rejeter, jamais une nouvelle vérité.

## B. Indépendance matérielle

Le cœur ne reçoit aucune identité matérielle comme paramètre de sa sémantique.

CPU, GPU, architecture, constructeur, OS, écran, clavier, réseau, stockage, alimentation et périphériques sont des capacités de réalisation.

Ils peuvent modifier les ressources disponibles et les performances. Ils ne modifient pas le sens du programme.

Une capacité absente produit un état explicite `MISSING_CAPABILITY` ou `REJECTED`. Elle n'est jamais simulée silencieusement.

## C. Accessibilité native

VOICE, BRAILLE, KEYBOARD, DISPLAY, TOUCH, POINTER, NETWORK et AUTOMATION sont des projections du même résultat sémantique.

Aucune projection n'est source de vérité.

L'accessibilité n'est pas un plugin du cœur et ne doit pas dépendre d'une interface graphique particulière.

## D. Preuve

`DEFINED != IMPLEMENTED != TESTED != SIMULATED != QEMU != HARDWARE != RF_PROVEN != SATELLITE_LINK_PROVEN`

Une preuve ne monte qu'à partir d'une évidence correspondant réellement au niveau demandé.

## E. Déterminisme

À système formel, état initial et entrée identiques, le résultat sémantique doit être identique.

Les informations physiques variables, comme l'heure de mesure ou la latence, ne doivent pas être utilisées pour modifier la vérité formelle.

## F. Compatibilité

Une réalisation est conforme si elle conserve :

`PROGRAM_ID + SEMANTICS + STATE + EVENTS + RESULTS + PROOF`

L'adaptateur matériel peut changer l'exécution physique sans changer ces éléments sémantiques.

## G. Inconnu

UNKNOWN et NOT_PROVEN sont des résultats valides.

Le système ne complète jamais silencieusement une donnée absente.

## H. Règle de sécurité

Le cœur ne doit jamais exécuter une action simplement parce qu'une IA la demande. Capacité et autorisation restent obligatoires.

## I. Objectif de fiabilité

La fiabilité d'ingénierie visée est **99,99 %**. Elle concerne l'exécution et la robustesse du système dans les conditions couvertes par les preuves.

Elle ne constitue pas une probabilité appliquée à la vérité mathématique. `1 + 1 = 2` reste un invariant formel indépendant du niveau de fiabilité physique.

## J. Critère de non-dépendance

Le cœur est considéré matériellement indépendant si ses tests de vérité, sérialisation, projection d'accessibilité et logique sémantique passent sans lire une identité matérielle particulière.

Ce contrat est permanent pour la version normative correspondante.\n\nUne panne physique peut empêcher une exécution, mais elle ne peut ni définir ni modifier la sémantique.
