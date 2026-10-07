# Modèle de preuve satellite

## États

DEFINED = spécification uniquement.
IMPLEMENTED = code présent.
TESTED = test automatisé exécuté avec résultat attendu.
SIMULATED = scénario simulé réussi.
QEMU = exécution démontrée sous QEMU.
HARDWARE = matériel réellement identifié.
RF_PROVEN = chaîne RF réellement mesurée.
SATELLITE_LINK_PROVEN = échange réel avec un satellite démontré.

## Interdictions

SIMULATED ≠ HARDWARE
QEMU ≠ HARDWARE
HARDWARE ≠ RF_PROVEN
RF_PROVEN ≠ SATELLITE_LINK_PROVEN

Chaque preuve doit conserver : identifiant, source, date, environnement, méthode, mesure, résultat et hash lorsque pertinent.
