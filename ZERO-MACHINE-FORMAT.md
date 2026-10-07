# ZERO — format machine canonique

## 1. But

Ce document définit la première représentation sérialisable de ZERO.

Elle est indépendante du langage hôte, du système d'exploitation, du protocole réseau, de l'interface graphique et de la modalité d'accessibilité.

La représentation canonique sert à transporter exactement le même sens entre stockage, mémoire, réseau, moteur d'exécution et projections d'accessibilité.

Elle ne constitue pas encore un protocole réseau ni un exécutable.

## 2. Unité

L'unité fondamentale est un RECORD.

Un flux ZERO est une suite ordonnée de RECORD.

Chaque RECORD possède :

1. VERSION
2. TYPE
3. IDENTITY
4. SEQUENCE
5. TIME
6. SOURCE
7. TARGET
8. PAYLOAD
9. PROOF

L'ordre est obligatoire. Aucun champ implicite ne doit être ajouté par le lecteur.

## 3. Types

Les types initiaux sont :

- OBJECT
- EVENT
- OPERATION
- RESULT
- PROOF
- CAPABILITY
- OBSERVATION

Un type inconnu doit rester représentable comme UNKNOWN_TYPE avec ses données brutes conservées.

## 4. Format textuel canonique

La première forme normative est UTF-8, une ligne par RECORD.

Syntaxe :

`RECORD|version|type|identity|sequence|time|source|target|payload|proof`

Les neuf champs sont toujours présents.

Le séparateur `|` fait partie du format. Il est interdit dans les valeurs brutes ; une valeur qui doit contenir ce caractère doit utiliser l'échappement UTF-8 défini ci-dessous.

### Échappement

- `\\` devient `\\\\`
- `|` devient `\\|`
- retour à la ligne devient `\\n`
- retour chariot devient `\\r`
- tabulation devient `\\t`

Les champs vides sont valides lorsqu'ils sont autorisés par le type.

## 5. Valeurs

Les valeurs sont typées par leur contenu sémantique, mais leur enveloppe reste textuelle dans cette première représentation.

Types primitifs réservés :

- STRING
- INTEGER
- BOOLEAN
- BYTES
- TIME
- ID
- STATE
- UNKNOWN

Un champ structuré utilise une liste ordonnée de paires :

`key=value`

séparées par `;`.

L'ordre des paires est normatif lorsqu'elles appartiennent à une représentation canonique.

## 6. Identité

IDENTITY doit être stable pendant la durée de vie logique de l'objet.

Une adresse physique, une adresse mémoire, un chemin de fichier ou une adresse réseau ne doit pas être utilisée comme identité sémantique si elle peut changer.

## 7. Séquence

SEQUENCE est un entier non négatif.

Dans un flux donné :

`sequence(n+1) > sequence(n)`

Une réutilisation ou une diminution de séquence est une erreur de validation.

## 8. Temps

TIME utilise une valeur UTC normalisée au format :

`YYYY-MM-DDThh:mm:ss.sssZ`

Lorsque le temps exact n'est pas disponible, le champ peut contenir `UNKNOWN`. L'incertitude ne doit jamais être transformée en précision inventée.

## 9. Payload

PAYLOAD contient le sens spécifique au TYPE.

Exemple OBJECT :

`meaning=machine spatiale;state=SIMULATED;capabilities=SATELLITE.TRACKING`

Exemple EVENT :

`kind=SATELLITE.DISCOVERED;state_before=UNKNOWN;state_after=SIMULATED`

## 10. Preuve

PROOF doit distinguer au minimum :

- UNKNOWN
- DEFINED
- IMPLEMENTED
- TESTED
- SIMULATED
- QEMU
- HARDWARE
- RF_PROVEN
- SATELLITE_LINK_PROVEN

La présence d'un RECORD ne constitue pas une preuve de son contenu physique.

## 11. Canonicalisation

Deux représentations sont canoniquement identiques si :

1. elles ont exactement les mêmes neuf champs ;
2. les échappements sont normalisés ;
3. les clés structurées sont dans l'ordre canonique ;
4. les entiers n'ont pas de zéros non significatifs ;
5. BOOLEAN vaut exactement `TRUE` ou `FALSE` ;
6. TIME respecte le format UTC ;
7. le type et la version sont identiques.

La comparaison canonique doit être effectuée après décodage des échappements.

## 12. Validation

Un décodeur doit rejeter :

- moins de neuf champs ;
- version absente ;
- type absent ;
- identité absente lorsqu'elle est obligatoire ;
- séquence invalide ;
- échappement invalide ;
- temps malformé lorsqu'un temps est fourni ;
- type connu avec un payload incompatible.

Un champ inconnu d'un type connu doit être conservé si sa structure est valide.

## 13. Round-trip

Pour tout RECORD valide :

`decode(encode(record)) = record`

La canonicalisation doit également satisfaire :

`canonical(encode(decode(bytes))) = canonical(bytes)`

pour toute représentation canonique valide.

## 14. Séparation preuve / simulation

SIMULATED, QEMU, HARDWARE, RF_PROVEN et SATELLITE_LINK_PROVEN sont des états de preuve distincts.

Un RECORD SIMULATED ne peut pas être relu comme HARDWARE.

Une projection d'accessibilité ne peut pas augmenter le niveau de preuve.

## 15. Statut

FORMAT : DEFINED

SERIALIZER : NOT_IMPLEMENTED

DESERIALIZER : NOT_IMPLEMENTED

CANONICAL BYTE VALIDATOR : NOT_IMPLEMENTED

TEST EXECUTION : NOT_EXECUTED

HARDWARE : NOT_PROVEN
