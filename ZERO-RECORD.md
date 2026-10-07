# ZERO — enregistrement canonique

ZERO transporte des unités sémantiques indépendantes de toute API graphique, système d'exploitation ou langage hôte.

## Types

OBJECT
EVENT
OPERATION
RESULT
PROOF
CAPABILITY
OBSERVATION

## Enveloppe

RECORD {
  version
  type
  identity
  sequence
  time
  source
  target
  payload
  proof
}

## Règles

- version obligatoire
- type obligatoire
- identity obligatoire sauf pour un résultat global
- sequence monotone dans un flux
- payload interprété selon type
- preuve conservée avec l'enregistrement
- aucun champ visuel obligatoire

L'enregistrement doit pouvoir être stocké, transmis, journalisé ou projeté vers une modalité d'accessibilité sans modifier son sens.
