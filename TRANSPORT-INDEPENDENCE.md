# ZERO — indépendance du transport

## Statut

IMPLEMENTED (bootstrap Go) + TESTED (unit tests).

## Règle

Le transport est une propriété de livraison, pas une propriété de vérité sémantique.

Le même Envelope conserve son identité, sa séquence, sa source, sa cible, son contenu et son niveau de preuve lorsqu'il passe par un transport différent ou lorsque la livraison échoue.

## Routage

Le noyau reçoit uniquement des capacités déclaratives: transport, disponibilité, joignabilité, store-and-forward et classe de coût.

Il ne reçoit pas de handle radio, réseau, satellite ou système d'exploitation.

La sélection est déterministe et l'absence de route ne modifie pas la vérité ou le contenu sémantique.

## Frontière de preuve

Cette implémentation prouve uniquement l'invariance dans le bootstrap Go testé.

Elle ne prouve aucun transport réel, aucune radio, aucun lien satellite et aucune connectivité physique.
