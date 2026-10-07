# ZERO — vecteurs de validation

Les vecteurs suivants définissent les comportements attendus du futur validateur.

## V001 — valide

Entrée : `TEST-VECTORS/object-create.zero`

Attendu :

VALID
ACCEPTED

## V002 — champ manquant

Entrée :

`RECORD|1|OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=test`

Attendu :

INVALID
FORMAT.FIELD_COUNT

## V003 — séquence décroissante

Flux :

`RECORD|1|OBJECT|x|2|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=test|DEFINED`
`RECORD|1|EVENT|x|1|2026-10-07T00:00:01.000Z|ZERO|x|kind=test|DEFINED`

Attendu du second :

INVALID
FORMAT.SEQUENCE

## V004 — temps invalide

`RECORD|1|OBJECT|x|1|NOT-A-TIME|ZERO|WORLD|meaning=test|DEFINED`

Attendu :

INVALID
FORMAT.TIME

## V005 — type inconnu

`RECORD|1|FUTURE_OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|future=value|DEFINED`

Attendu :

VALID
TYPE=UNKNOWN_TYPE
RAW_TYPE=FUTURE_OBJECT

## V006 — preuve non augmentée

Entrée avec `proof=SIMULATED`.

Attendu :

VALIDATION=VALID
PROOF reste SIMULATED

Le validateur ne peut pas produire HARDWARE.

## V007 — échappement

Payload :

`meaning=ligne\\navec\\|séparateur\\\\réel`

Attendu :

VALID

Après décodage, le payload contient un retour à la ligne, `|` et `\\`.

## V008 — opération non exécutée

Payload contenant une intention d'action.

Attendu :

VALIDATION=VALID

Aucune action physique ou distante ne doit être exécutée par le validateur.

## Statut

DEFINED
EXECUTION=NOT_EXECUTED
