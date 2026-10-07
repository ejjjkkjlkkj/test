# ZERO — noyau de vérité déterministe

## Statut

Le noyau de vérité formelle est indépendant du support physique.

Une vérité formelle ne dépend ni :
- du CPU ;
- du GPU ;
- de la mémoire disponible ;
- du système d'exploitation ;
- du téléphone ;
- du PC ;
- du réseau ;
- du stockage ;
- de l'écran ;
- de la voix ;
- d'un serveur ;
- d'une IA particulière.

## Invariant permanent

Pour la définition arithmétique normative ZERO :

`1 + 1 = 2`

et :

`1 + 1 = 3` → `INVALID`.

Cette règle n'est pas une prédiction, une observation, une préférence ou une sortie d'IA.

## Indépendance

Le même triplet :

`RULE=ARITHMETIC.ADD`
`INPUT=1,1`
`EXPECTED=2`

doit produire `2` sur toute réalisation qui exécute ce noyau formel.

Le support peut changer l'exécution physique, mais il ne peut pas changer la définition de la vérité.

## IA

L'IA peut observer, calculer, tester ou vérifier la règle.

Elle ne peut pas :
- modifier `1 + 1 = 2` ;
- choisir `3` pour satisfaire une hypothèse ;
- transformer une erreur en vérité ;
- modifier rétroactivement une preuve ;
- remplacer une règle déterministe par une probabilité.

Une contradiction déclenche `ERROR/INVALID`, jamais une nouvelle vérité.

## Versionnement

Une modification du système formel ne modifie pas silencieusement une version existante.

Une nouvelle règle doit :
1. recevoir une nouvelle identité/version ;
2. déclarer explicitement sa définition ;
3. être soumise aux tests ;
4. conserver la trace de la version précédente.

Ainsi, « pour toujours » signifie : une version normative déjà acceptée ne change pas silencieusement.

## Accessibilité

La représentation de la vérité doit rester identique pour les projections :

`VOICE`
`BRAILLE`
`KEYBOARD`
`DISPLAY`
`TOUCH`
`POINTER`
`NETWORK`
`AUTOMATION`

La projection ne peut jamais transformer `2` en `3`.

## Preuve

Ce document définit le contrat. La preuve d'exécution est fournie séparément par les tests exécutés.

`DEFINED != EXECUTED != OBSERVED != PROVEN`

## Règle de refus

Si le résultat observé d'une opération déterministe ne respecte pas sa définition :

`REJECTED`

Le système ne doit pas produire un PASS artificiel.
