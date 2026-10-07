# ZERO — représentation binaire native minimale

## Statut

FORMAT = DEFINED  
BOOTSTRAP ENCODER = IMPLEMENTED (Go)  
BOOTSTRAP DECODER = IMPLEMENTED (Go)  
ROUND_TRIP = TESTED (Go)  
CPU EXECUTION = NOT_IMPLEMENTED  
HARDWARE = NOT_PROVEN

## But

Matérialiser l'étape 1 de `IMPLEMENTATION-ORDER.md` sans faire de Go le langage fondamental de ZERO.

Le format binaire transporte le même RECORD que `ZERO-MACHINE-FORMAT.md`. Il ne crée pas une seconde sémantique.

## Encodage canonique

En-tête :

- 4 octets ASCII : `ZER0`
- 1 octet : version du format, actuellement `1`
- après l'en-tête, les champs sont encodés dans l'ordre normatif ci-dessous ; le TYPE est une chaîne UTF-8 préfixée par sa longueur

Pour rester extensible, chaque champ variable est encodé comme :

`uint32_be longueur + octets UTF-8`

Ordre des champs :

1. TYPE
2. IDENTITY
3. SEQUENCE : uint64 big-endian
4. TIME
5. SOURCE
6. TARGET
7. PAYLOAD
8. PROOF

La version du RECORD est encodée comme premier champ variable après l'en-tête.

## Invariants

- aucune donnée n'est déduite d'une projection ;
- aucun champ n'est supprimé parce qu'il est inconnu ;
- l'identité sémantique n'est pas une adresse physique ;
- `SIMULATED` ne devient jamais `HARDWARE` ;
- le décodage doit refuser les longueurs impossibles et les données tronquées ;
- `decode(encode(record)) = record`.

## Position actuelle

Cette représentation est une matérialisation de données, pas encore une ISA CPU.

Étape suivante : définir le paquet d'instructions minimal et son état d'exécution, puis connecter ce modèle à une mémoire contrôlée dans un environnement de test.
