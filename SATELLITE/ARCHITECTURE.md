# Architecture satellite

## 1. Machine
Le satellite est modélisé comme un ensemble d'objets, capacités, événements, états et liens.

## 2. Couches

- DEVICE : matériel réellement observé
- RADIO : chaîne RF abstraite
- SATELLITE : identité, temps, position, orbite
- LINK : relation entre deux nœuds
- FRAME : unité de transport physique/logique
- PROTOCOL : règles d'échange
- TELEMETRY : état descendant
- COMMAND : contrôle montant
- NETWORK : acheminement multi-nœuds
- SEMANTICS : représentation machine lisible
- ACCESSIBILITY : exposition voix/braille/clavier/autres sorties

## 3. Types de nœuds

GROUND, SATELLITE, RELAY, SPACECRAFT, USER_DEVICE, NETWORK_NODE.

## 4. Règle

Aucun niveau supérieur ne peut déclarer une capacité inférieure comme prouvée sans preuve correspondante.
