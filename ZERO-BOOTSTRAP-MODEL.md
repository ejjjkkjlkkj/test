# ZERO — bootstrap minimal

## Sequence

RESET -> MACHINE.IDENTITY -> MACHINE.STATE -> MEMORY.OBSERVE -> BUS.OBSERVE -> DEVICE.DISCOVER -> CAPABILITY.DISCOVER -> EVENT.CHANNEL -> SEMANTIC.OBJECTS -> EXECUTION -> ACCESSIBILITY

Le bootstrap commence par OBSERVE.

Avant toute ecriture, la ressource doit etre identifiee, son etat observe, la capacite determinee et l'autorisation obtenue.

CPU, memoire, bus et peripheriques sont representes comme objets semantiques. La decouverte d'un peripherique ne prouve pas que toutes ses fonctions sont operationnelles.

Les projections d'accessibilite doivent pouvoir fonctionner sans interface graphique.

Securite: aucun effacement, formatage, modification firmware/NVRAM, partitionnement ou redemarrage automatique n'est implicite.

Chaque etape produit une observation, un resultat et une preuve. Une simulation reste SIMULATED et QEMU reste QEMU.

Statut:
BOOTSTRAP=DEFINED
IMPLEMENTATION=NOT_IMPLEMENTED
REAL_EXECUTION=NOT_EXECUTED
HARDWARE=NOT_PROVEN
