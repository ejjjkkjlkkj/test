# ZERO — Machine distribuée

ZERO ne s'arrête pas à une machine.

Un ensemble de machines constitue un espace sémantique distribué.

PC
<-> iPhone
<-> tablette
<-> autre PC
<-> serveur
<-> périphérique
<-> future machine ZERO

Chaque machine conserve son identité et ses autorisations.

## Identité

MACHINE_ID != NETWORK_ADDRESS

Une connexion réseau ne fusionne jamais deux machines.

La relation distribuée possède :

SOURCE
TARGET
RELATION
CAPABILITIES
AUTHORIZATION
STATE
EVENTS
UNCERTAINTY

## Échange

Une machine ZERO peut publier :

- présence ;
- capacités ;
- objets ;
- événements ;
- demandes ;
- résultats ;
- changements d'état.

La machine distante choisit ce qu'elle expose.

## Continuité

Une communication interrompue devient un état sémantique :

CONNECTED
DISCONNECTED
DEGRADED
RECONNECTING
UNKNOWN

La rupture ne doit pas produire un objet incohérent ou silencieusement perdu.

## Accessibilité

Les objets distants restent soumis aux mêmes invariants que les objets locaux.

La distribution ne crée pas une seconde accessibilité.
