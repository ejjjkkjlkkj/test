# ZERO — scénarios d'exécution normatifs

Ces scénarios décrivent les premières exécutions attendues du noyau.

Ils ne sont pas encore exécutés.

## X001 — création et observation

Entrées :

CREATE OBJECT `SATELLITE:test`
MEANING `machine spatiale de test`
STATE `SIMULATED`
CAPABILITY `SATELLITE.TRACKING`
EVENT `SATELLITE.DISCOVERED`

Sortie attendue :

RESULT=COMPLETED
OBSERVATION.IDENTITY=SATELLITE:test
OBSERVATION.STATE=SIMULATED
OBSERVATION.CAPABILITY=SATELLITE.TRACKING
OBSERVATION.EVENT=SATELLITE.DISCOVERED

## X002 — observation sans mutation

OBSERVE `SATELLITE:test`

Attendu :

RESULT=COMPLETED

Aucune transition d'état de l'objet.

## X003 — action sans capacité

REQUEST `TRACK` ON `SATELLITE:test`

sans `SATELLITE.TRACKING`.

Attendu :

RESULT=REJECTED
CAUSE=CAPABILITY.MISSING

Aucune action physique.

## X004 — action sans autorisation

CAPABILITY présente mais AUTHORIZATION absente.

Attendu :

RESULT=REJECTED
CAUSE=AUTHORIZATION.MISSING

## X005 — interruption

Une action passe de RUNNING à INTERRUPTED.

Attendu :

RESULT=INTERRUPTED
PARTIAL_STATE conservé
CAUSE conservée
NEXT_VALID_STATE conservé

## X006 — type futur

Un RECORD possède un type inconnu.

Attendu :

RECORD conservé
STATE=UNKNOWN
RAW_TYPE conservé
RESULT=ACCEPTED

## X007 — preuve

Une entrée SIMULATED ne peut produire qu'une preuve compatible avec son observation.

Attendu :

SIMULATED reste SIMULATED.

## Statut

DEFINED
EXECUTION=NOT_EXECUTED
HARDWARE=NOT_PROVEN
