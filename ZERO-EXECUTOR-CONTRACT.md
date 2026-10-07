# ZERO — exécuteur minimal

## Statut

CONTRACT = DEFINED
BOOTSTRAP IMPLEMENTATION = IMPLEMENTED
BOOTSTRAP TESTS = DEFINED (execution status supplied by CI)
HARDWARE = NOT_PROVEN

## Référence normative

L'unité d'entrée est le RECORD défini par `ZERO-MACHINE-FORMAT.md`.

Il contient exactement, dans cet ordre :

1. VERSION
2. TYPE
3. IDENTITY
4. SEQUENCE
5. TIME
6. SOURCE
7. TARGET
8. PAYLOAD
9. PROOF

L'implémentation bootstrap Go sous `bootstrap/go` sert uniquement à vérifier cette représentation. Go n'est pas le langage fondamental de ZERO.

## Pipeline

INPUT
→ PARSE
→ VALIDATE
→ TRANSITION
→ EVENT
→ RESULT

Une erreur d'une étape ne peut pas être transformée silencieusement en succès.

## Validation

Le validateur rejette au minimum :

- un champ obligatoire absent ;
- une séquence non numérique ;
- un enregistrement mal formé ;
- un échappement invalide ;
- une contradiction interne.

Une donnée absente reste absente. Une donnée inconnue reste UNKNOWN.

## Transition

`STATE + INPUT + RULE_VERSION → NEW_STATE`

Pour une même entrée, un même état et une même version de règle, le résultat doit être identique.

Une contradiction produit :

`CONTRADICTION`

et ne peut jamais produire arbitrairement `DELIVERED`, `SUCCESS` ou un autre état favorable.

## Événement et résultat

Une transition acceptée doit conserver :

- l'identité de la source ;
- la séquence ;
- l'état précédent ;
- l'état suivant ;
- le résultat ;
- le niveau de preuve.

## Transport

Le transport est extérieur au sens du RECORD.

`SEMANTICS(message)` ne change pas quand le transport change.

Les états de livraison restent distincts :

`DELIVERED`
`DEFERRED`
`FAILED`
`UNKNOWN`

Ni DEFERRED ni FAILED ne signifient DELIVERED.

## Accessibilité

Une projection peut cibler :

VOICE
BRAILLE
KEYBOARD
DISPLAY
TOUCH
POINTER
NETWORK
AUTOMATION

Toutes les projections doivent recevoir la même structure sémantique. Aucune projection ne peut augmenter la preuve ou modifier le résultat.

## Reproductibilité

La validation bootstrap doit être reproductible à partir :

- du dépôt ;
- de la version de Go déclarée ;
- des fichiers du bootstrap ;
- des vecteurs versionnés ;
- des tests versionnés.

Le workflow GitHub Actions correspondant ne constitue pas une preuve matérielle.

## Preuve

`DEFINED != IMPLEMENTED != TESTED != EXECUTED != OBSERVED != PROVEN`

Un test logiciel réussi prouve uniquement le comportement testé de cette implémentation.

Il ne prouve pas automatiquement :

- CPU universel ;
- autre architecture ;
- téléphone ;
- radio ;
- réseau réel ;
- satellite ;
- fonctionnement hors ligne physique ;
- fonctionnement aérien ;
- accessibilité de tout matériel.

## Objectif suivant

Faire correspondre les vecteurs de transport et de round-trip au RECORD canonique, puis ajouter :

1. vérification de séquence sur flux ;
2. conservation explicite des champs inconnus ;
3. génération d'événements ;
4. résultat sérialisable ;
5. tests croisés encodeur/décodeur ;
6. tests d'accessibilité sans changement sémantique.
