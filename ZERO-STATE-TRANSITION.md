# ZERO — transitions d'état

Le noyau manipule des transitions explicites.

TRANSITION {
  object
  before
  operation
  preconditions
  after
  events
  observations
  result
  proof
}

## Invariant

Une transition ne peut pas déclarer son nouvel état sans produire une relation explicite avec l'état précédent.

## Échec

En cas d'échec :

before reste observable
after décrit l'état réellement obtenu
result = FAILED
proof conserve l'environnement et la cause

Un échec ne doit jamais être converti en succès silencieux.

## Simulation

SIMULATE produit une transition marquée SIMULATED.

EXECUTE produit une transition correspondant à une action effectivement demandée au dispositif ou au moteur.

Les deux sont incompatibles dans la même preuve.
