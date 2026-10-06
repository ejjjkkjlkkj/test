# ZERO — Exécution du langage

L'exécution ZERO ne suppose pas un runtime d'un autre langage.

## Cycle

FETCH
→ INTERPRET
→ RESOLVE
→ CHECK
→ TRANSITION
→ EVENT
→ OBSERVE
→ CONTINUE

L'exécution elle-même est un état sémantique.

## Instruction

Une instruction n'est pas seulement un opcode.

Elle possède :

INTENT
TARGETS
INPUTS
REQUIRED_CAPABILITIES
TRANSITION
RESULT

Le CPU réalise finalement cette intention par des opérations physiques.

## Erreur

Une erreur n'est jamais seulement un code opaque.

Elle est un objet :

CAUSE
LOCATION
STATE
EXPECTED
OBSERVED
CAPABILITY
RECOVERY
UNCERTAINTY

Elle peut donc être inspectée, annoncée, braillée, enregistrée ou transmise à une autre machine.

## Interruption

INTERRUPT est une transition normale du modèle.

Une action interrompue produit :

INTERRUPTED
+ PARTIAL_STATE
+ CAUSE
+ NEXT_VALID_STATE
