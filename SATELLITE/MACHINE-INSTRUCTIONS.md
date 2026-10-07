# ZERO — primitives machine satellite

Les opérations suivantes sont des primitives sémantiques, pas une syntaxe finale imposée.

DISCOVER node
OBSERVE node
READ node.property
WATCH node.event

LOCATE node
TRACK node
PREDICT node

OPEN link
CLOSE link
MEASURE link

RECEIVE frame
DECODE frame
VERIFY frame

READ telemetry
WATCH telemetry

CREATE command
AUTHORIZE command
SEND command
CANCEL command
OBSERVE command

ROUTE packet
STORE packet
FORWARD packet

Chaque opération produit un résultat structuré avec état, observations, erreur éventuelle et preuve.

## Interdiction

Une opération de simulation ne peut masquer une opération matérielle.

SIMULATE x

doit rester explicitement distincte de :

EXECUTE x
