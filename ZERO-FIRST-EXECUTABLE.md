# ZERO — premier objectif exécutable

Objectif minimal :

1. créer un objet
2. lui attribuer une identité
3. lui attribuer une signification
4. lui attribuer un état
5. enregistrer une capacité
6. produire un événement
7. observer l'objet
8. conserver la preuve de l'opération

## Premier scénario

CREATE SATELLITE:test

SET MEANING "machine spatiale de test"

SET STATE SIMULATED

ADD CAPABILITY SATELLITE.TRACKING

EMIT SATELLITE.DISCOVERED

OBSERVE SATELLITE:test

Le résultat attendu est une observation structurée contenant l'identité, le sens, l'état, la capacité et l'événement.

## Limite

Ce scénario ne constitue pas une preuve de satellite réel.

Il constitue uniquement le premier test du noyau sémantique.
