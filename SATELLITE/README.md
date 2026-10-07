# SATELLITE — architecture originale

Cette couche définit notre représentation machine des systèmes spatiaux.

Principe : les projets externes servent uniquement de références de recherche. Aucun projet externe ne constitue notre architecture.

Chaîne :

HARDWARE → BUS/RF → DRIVER → RADIO → SATELLITE → LINK → FRAME → PROTOCOL → TELEMETRY/COMMAND → NETWORK → SEMANTICS → ACCESSIBILITY

Les états de preuve sont indépendants :
DEFINED, IMPLEMENTED, TESTED, SIMULATED, QEMU, HARDWARE, RF_PROVEN, SATELLITE_LINK_PROVEN.

Une simulation ne constitue jamais une preuve de liaison satellite réelle.
