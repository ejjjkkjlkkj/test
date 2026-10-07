# ZERO — modèle d'événements

L'événement est une primitive centrale du langage.

## Structure

EVENT {
  id
  time
  source
  target
  kind
  state_before
  state_after
  payload
  cause
  proof
}

## Propriétés

- immutable
- ordonné
- identifiable
- observable
- annulable uniquement par un nouvel événement compensatoire
- aucune suppression silencieuse

## Familles

DEVICE.*
RADIO.*
NETWORK.*
SATELLITE.*
ORBIT.*
LINK.*
FRAME.*
TELEMETRY.*
COMMAND.*
ACCESSIBILITY.*

## Accessibilité

Toute projection vocale, braille, clavier ou graphique consomme les mêmes événements sémantiques.
