# ZERO — ISA sémantique minimale

ISA = operations semantiques observables.

Instructions minimales:

DECODE
VALIDATE
CREATE
SET
RELATE
CAPABILITY
EMIT
OBSERVE
TRANSITION
PROJECT
COMMUNICATE
STORE
LOAD
HALT

Cycle:
FETCH -> DECODE -> VALIDATE -> CAPABILITY -> AUTHORIZATION -> TRANSITION -> EVENT -> OBSERVATION -> RESULT

OBSERVE ne modifie jamais l'etat.

Une action exige une cible, les capacites requises et une autorisation.

Une instruction inconnue n'est jamais executee par defaut. Elle peut seulement etre observee, stockee ou transmise si les capacites correspondantes existent.

Le mapping vers le materiel reste externe:
ZERO ISA -> semantique machine -> implementation -> ressources physiques

Statut:
ISA=DEFINED
ENCODING=NOT_IMPLEMENTED
EXECUTOR=NOT_IMPLEMENTED
HARDWARE=NOT_PROVEN
