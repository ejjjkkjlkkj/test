# ZERO — modèle de capacités

Une capacité décrit ce qu'un objet peut réellement faire.

## Structure

CAPABILITY {
  identity
  owner
  operation
  requirements
  authorization
  state
  proof
}

## États

UNKNOWN
AVAILABLE
UNAVAILABLE
BLOCKED
ACTIVE
FAILED
PROVEN

## Règle

La présence d'un identifiant de capacité ne prouve pas son fonctionnement.

La capacité doit conserver son niveau de preuve.

Exemple :

CAPABILITY.SATELLITE.DOWNLINK
state = AVAILABLE
proof = NOT_PROVEN

est valide.

Le système ne doit pas transformer automatiquement cette information en :

proof = SATELLITE_LINK_PROVEN.
