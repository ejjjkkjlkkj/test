# ZERO — vecteurs de test

Ces fichiers sont des données de test normatives.

Ils ne constituent pas encore une preuve d'exécution.

## Objectifs

Les futurs encodeur et décodeur devront vérifier :

1. lecture d'un RECORD valide ;
2. conservation de l'identité ;
3. conservation de la séquence ;
4. conservation du payload ;
5. conservation de la preuve ;
6. échappement et déséchappement ;
7. rejet des RECORD incomplets ;
8. rejet des séquences invalides ;
9. conservation des champs inconnus ;
10. round-trip encode/decode.

## Statut

VECTEURS : DEFINED

EXECUTION : NOT_EXECUTED

SERIALIZER : NOT_IMPLEMENTED

DESERIALIZER : NOT_IMPLEMENTED


## Transport invariance

`transport-invariance.zero` vérifie que le transport ne change pas le sens d'un message, que `DEFERRED` et `FAILED` ne deviennent pas `DELIVERED`, et qu'une catégorie `SATELLITE` ne crée pas automatiquement une preuve satellite. Ces vecteurs restent des données de validation et ne constituent pas une preuve de connectivité réelle.
