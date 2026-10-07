# ZERO — contrat sémantique commun

## But

Ce document devient le point de raccord entre la définition du langage, le moteur expérimental, le noyau natif et les modèles formels.

## Unité

L'unité sémantique est une **trace** : une représentation observable accompagnée de relations, d'historique, de pression et d'âge lorsqu'ils sont nécessaires à l'effet observé.

## Transition

Une transition :

`(champ, trace, capacité, budget) -> (champ', historique', résultat)`

est valide seulement si :

- son autorité est explicitement disponible ;
- son budget est respecté ;
- les informations inconnues ne sont pas remplacées silencieusement ;
- l'historique reste cohérent ;
- les effets externes passent par leur frontière d'autorité ;
- le résultat est rejouable lorsque le contrat l'exige.

## Effets

Le calcul interne et l'effet externe sont séparés. Une demande d'effet ne vaut pas exécution. L'effet doit être autorisé, vérifié, puis engagé.

## Communication

La communication IA↔IA et machine↔machine est une transition entre deux champs. Un message peut transporter observation, hypothèse, inconnue, contradiction, proposition ou résultat. Il ne confère jamais une capacité par lui-même.

## Raffinement

Toute implémentation — Rust expérimentale, C natif, représentation binaire future ou adaptateur externe — doit fournir une relation de raffinement vers ce contrat. Tant que cette relation n'est pas démontrée, l'implémentation reste dans le statut expérimental.

## Propriétés centrales

P01 autorité monotone
P02 médiation complète
P03 capacités non forgeables
P04 historique immuable
P05 atomicité
P06 cohérence après crash
P07 isolation des instances
P08 non-interférence
P09 conservation de l'inconnu
P10 provenance
P12 rejeu déterministe
P14 sûreté de concurrence
P17 confinement des effets
P18 liaison identité/périphérique
P19 confinement réseau
P20 raffinement sémantique
P21 préservation par compilation
P22 frontière ISA
P23 conservation des modalités
P25 non-interférence d'accessibilité
P27 auto-extension en quarantaine
P28 promotion monotone
P29 auto-modification contenue
P32 fail-closed
P33 liaison des artefacts
P34 re-vérification indépendante
P35 clôture de composition

## Statut

Le contrat est une spécification. Il ne transforme pas automatiquement les implémentations existantes en systèmes prouvés. Chaque raccord doit produire des artefacts d'implémentation et de preuve.
