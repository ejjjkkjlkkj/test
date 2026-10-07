# ZERO — représentation binaire native minimale

## Statut

FORMAT = DEFINED
BOOTSTRAP ENCODER = IMPLEMENTED (Go)
BOOTSTRAP DECODER = IMPLEMENTED (Go)
ROUND_TRIP = TESTED (Go)
NATIVE-TO-SEMANTIC EXECUTION = IMPLEMENTED (Go)
CPU EXECUTION = NOT_IMPLEMENTED
HARDWARE = NOT_PROVEN

## But

Matérialiser la représentation binaire du RECORD ZERO sans faire de Go le langage fondamental de ZERO.

Le format binaire transporte le même RECORD que `ZERO-MACHINE-FORMAT.md`. Il ne crée pas une seconde sémantique.

## Encodage canonique

En-tête :

- 4 octets ASCII : `ZER0`
- 1 octet : version du format, actuellement `1`
- après l'en-tête, les champs sont encodés dans l'ordre normatif ci-dessous.

Chaque champ variable est encodé comme :

`uint32_be longueur + octets UTF-8`

Ordre des champs :

1. VERSION du RECORD
2. TYPE
3. IDENTITY
4. SEQUENCE : uint64 big-endian
5. TIME
6. SOURCE
7. TARGET
8. PAYLOAD
9. PROOF

## Invariants

- aucune donnée n'est déduite d'une projection ;
- aucun champ n'est supprimé parce qu'il est inconnu ;
- l'identité sémantique n'est pas une adresse physique ;
- `SIMULATED` ne devient jamais `HARDWARE` ;
- le décodage refuse les longueurs impossibles, les données tronquées et les octets finaux ;
- `decode(encode(record)) = record` ;
- un RECORD natif décodé conserve la même représentation canonique avant exécution sémantique.

## Validation actuelle

Le bootstrap Go vérifie le round-trip, le déterminisme, les troncatures, l'absence de panique du décodeur et l'équivalence entre représentation native et représentation canonique.

`ExecuteNativeRecord` réutilise ensuite exactement le même pipeline d'autorisation, d'exécution et de transition que `ExecuteRecord`.

## Limite

Cette représentation est une matérialisation de données et un point d'entrée du bootstrap. Elle ne constitue pas encore une ISA CPU native ni une preuve matérielle.
