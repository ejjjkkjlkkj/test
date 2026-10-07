# ZERO — validateur canonique

## 1. Rôle

Le validateur est une fonction sémantique du noyau ZERO.

Il ne dépend d'aucun langage hôte.

Entrée :

RECORD brut

Sortie :

VALID ou INVALID

avec :

- code ;
- position ;
- cause ;
- observation ;
- preuve.

Le validateur ne modifie jamais l'entrée.

## 2. Pipeline

VALIDATE(RECORD)

1. vérifier le préfixe RECORD ;
2. décoder les neuf champs ;
3. vérifier la version ;
4. vérifier le TYPE ;
5. vérifier IDENTITY selon le TYPE ;
6. vérifier SEQUENCE par rapport au flux ;
7. vérifier TIME ;
8. vérifier SOURCE et TARGET ;
9. décoder PAYLOAD ;
10. vérifier les contraintes du TYPE ;
11. vérifier PROOF ;
12. produire le résultat.

Aucune étape ne doit transformer UNKNOWN en valeur connue.

## 3. Résultat

Résultat valide :

RESULT=VALID
STATE=ACCEPTED

Résultat invalide :

RESULT=INVALID
STATE=REJECTED

Un rejet conserve :

- l'entrée originale ;
- l'étape de rejet ;
- la position ;
- la cause ;
- le type d'erreur ;
- la preuve de validation.

## 4. Codes

Les codes initiaux sont :

FORMAT.EMPTY
FORMAT.PREFIX
FORMAT.FIELD_COUNT
FORMAT.ESCAPE
FORMAT.VERSION
FORMAT.TYPE
FORMAT.IDENTITY
FORMAT.SEQUENCE
FORMAT.TIME
FORMAT.SOURCE
FORMAT.TARGET
FORMAT.PAYLOAD
FORMAT.PROOF
FORMAT.CANONICAL

Les codes sont des identités sémantiques, pas de simples nombres opaques.

## 5. Séquence

Le validateur conserve la dernière séquence observée par flux.

Conditions :

- première séquence >= 0 ;
- toute séquence suivante > précédente ;
- aucune diminution ;
- aucun doublon.

Un flux interrompu ne permet pas d'inventer les séquences manquantes.

## 6. Types inconnus

Un TYPE inconnu n'est pas automatiquement invalide.

Si son enveloppe est correcte, il est conservé comme :

TYPE=UNKNOWN_TYPE
RAW_TYPE=<valeur originale>

Le payload est conservé sans interprétation forcée.

## 7. Champs inconnus

Un champ supplémentaire à l'intérieur d'un payload connu est conservé lorsqu'il respecte la syntaxe.

Un champ obligatoire absent est une erreur.

Un champ inconnu ne doit jamais remplacer un champ obligatoire.

## 8. Canonicalisation

Après validation :

canonical_record = canonicalize(record)

La canonicalisation :

- décode les échappements ;
- normalise les booléens ;
- normalise les entiers ;
- normalise TIME ;
- ordonne les clés lorsque le type le requiert ;
- réencode selon les règles du format.

## 9. Invariants

I1 : VALID implique neuf champs exactement.

I2 : INVALID n'est jamais converti silencieusement en VALID.

I3 : UNKNOWN reste UNKNOWN.

I4 : la canonicalisation ne change pas le sens.

I5 : la projection d'accessibilité ne change pas le RECORD.

I6 : PROOF ne peut pas être augmenté par le validateur.

I7 : le validateur n'exécute aucune opération décrite dans PAYLOAD.

I8 : VALIDATION et EXECUTION sont deux étapes différentes.

## 10. Preuve du validateur

Une validation produit elle-même une observation :

OBSERVATION
identity=VALIDATION:<record-identity>
state=VALID ou INVALID
cause=<code si INVALID>
source=ZERO.VALIDATOR
proof=TESTED lorsque le validateur a effectivement été exécuté.

La preuve TESTED ne doit être attribuée que par une exécution réelle.

## 11. Statut

SPECIFICATION : DEFINED

ALGORITHM : DEFINED

REFERENCE IMPLEMENTATION : NOT_IMPLEMENTED

EXECUTION : NOT_EXECUTED

HARDWARE : NOT_PROVEN
