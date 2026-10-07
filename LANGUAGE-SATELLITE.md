# ZERO — langage machine satellite

Ce document définit l'extension satellite du langage machine-accessibilité ZERO.

Le satellite n'est pas traité comme une application réseau. Il est un ensemble de machines, d'états, de relations et de flux physiques représentés nativement par le langage.

## 1. Objet SATELLITE

Un objet satellite possède au minimum :

IDENTITY
MEANING
STATE
POSITION
TIME
ORBIT
CAPABILITIES
LINKS
EVENTS
OBSERVATIONS
AUTHORIZATION
PROOF

Exemple conceptuel :

SATELLITE {
  identity
  state
  position
  velocity
  orbit
  capabilities
  links
}

## 2. Objet LINK

LINK représente une relation réelle ou simulée entre deux nœuds.

LINK {
  source
  destination
  direction
  state
  start
  end
  latency
  rate
  loss
  observations
  proof
}

Un LINK n'est jamais considéré comme réel uniquement parce qu'il est déclaré dans un modèle.

## 3. Objet ORBIT

ORBIT représente une relation temporelle entre un objet spatial et un référentiel.

ORBIT {
  reference
  epoch
  position
  velocity
  uncertainty
  validity
  source
}

La position calculée et la position observée restent distinctes.

## 4. Objet FRAME

FRAME représente une unité structurée transportée par une liaison.

FRAME {
  source
  destination
  type
  sequence
  payload
  integrity
  timestamp
  observations
}

Le payload peut être inconnu sans rendre le FRAME inexistant.

## 5. TELEMETRY

TELEMETRY est une observation produite par une machine distante.

TELEMETRY {
  source
  time
  parameter
  value
  unit
  quality
  sequence
  proof
}

La qualité et l'origine font partie de la donnée.

## 6. COMMAND

COMMAND représente une intention d'action distante.

COMMAND {
  source
  destination
  operation
  parameters
  authorization
  state
  result
}

États :

CREATED
AUTHORIZED
QUEUED
SENT
RECEIVED
ACCEPTED
EXECUTING
COMPLETED
REJECTED
EXPIRED
CANCELLED
FAILED

Aucune commande distante n'est automatiquement autorisée.

## 7. Événements natifs

SATELLITE.DISCOVERED
SATELLITE.STATE_CHANGED
ORBIT.UPDATED
LINK.DETECTED
LINK.UP
LINK.DOWN
FRAME.RECEIVED
FRAME.INVALID
TELEMETRY.RECEIVED
COMMAND.SENT
COMMAND.ACCEPTED
COMMAND.REJECTED
COMMAND.COMPLETED

## 8. Accessibilité native

Les objets spatiaux produisent directement des événements sémantiques.

Le moteur peut donc demander :

OBSERVE(SATELLITE)
OBSERVE(ORBIT)
OBSERVE(LINK)
OBSERVE(TELEMETRY)
OBSERVE(COMMAND)

et projeter le résultat vers :

VOICE
BRAILLE
KEYBOARD
DISPLAY
NETWORK
AUTOMATION

Aucune interface graphique n'est nécessaire pour comprendre l'objet.

## 9. Preuve

Chaque capacité peut avoir un état indépendant :

UNKNOWN
DEFINED
IMPLEMENTED
SIMULATED
QEMU
HARDWARE
RF_PROVEN
SATELLITE_LINK_PROVEN

La preuve accompagne l'objet au lieu d'être ajoutée après coup.

## 10. Principe fondamental

ZERO ne confond jamais :

REPRESENTATION
OBSERVATION
SIMULATION
EXECUTION
PROOF

C'est la base du langage satellite.
