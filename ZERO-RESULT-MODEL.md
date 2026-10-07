# ZERO — résultat canonique

Un résultat est une donnée sémantique produite par une validation, une observation ou une opération.

## Structure

RESULT {
  identity
  operation
  state
  observations
  events
  cause
  uncertainty
  proof
}

## États

- CREATED
- ACCEPTED
- REJECTED
- COMPLETED
- FAILED
- INTERRUPTED
- UNKNOWN

## Règles

Un résultat FAILED n'est pas équivalent à UNKNOWN.

Un résultat REJECTED signifie que la demande ou l'enregistrement a été explicitement refusé.

Un résultat UNKNOWN signifie que le noyau ne possède pas assez d'information.

Un résultat INTERRUPTED conserve l'état partiel réellement observé.

La preuve reste attachée au résultat.

## Accessibilité

Le même RESULT peut être projeté vers :

VOICE
BRAILLE
KEYBOARD
DISPLAY
NETWORK
AUTOMATION

Aucune projection ne modifie le RESULT.

## Statut

DEFINED
IMPLEMENTATION=NOT_IMPLEMENTED
