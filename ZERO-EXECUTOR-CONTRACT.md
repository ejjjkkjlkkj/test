# ZERO — exécuteur minimal

## Statut

DEFINED / IMPLEMENTATION CONTRACT

## Pipeline

INPUT
→ PARSE
→ VALIDATE
→ TRANSITION
→ EVENT
→ RESULT

Chaque étape produit un état explicite. Une étape en erreur arrête la transition correspondante.

## Entrée

Un record ZERO doit contenir au minimum :

TYPE
ID
SOURCE
TARGET
SEQUENCE
STATE
PAYLOAD
PROOF

## Validation

Le validateur rejette :

- un champ obligatoire absent ;
- une séquence invalide ;
- un état inconnu ;
- une preuve déclarée sans définition correspondante ;
- une livraison déclarée DELIVERED sans observation de livraison ;
- une contradiction interne.

Le validateur ne complète pas silencieusement une donnée absente.

## Transition

Une transition est :

STATE + INPUT + RULE → NEW_STATE

La transition doit être déterministe pour un même état, une même entrée et une même version de règle.

Une contradiction produit :

ERROR
CONTRADICTION
INVALID

et ne choisit jamais arbitrairement un résultat.

## Événement

Chaque transition acceptée peut produire un événement structuré :

EVENT
ID
SOURCE
TARGET
SEQUENCE
PREVIOUS_STATE
NEW_STATE
RESULT
PROOF

L'événement conserve la provenance.

## Résultat

Un résultat possède au minimum :

STATUS
VALUE
SOURCE_EVENT
PROOF

Valeurs minimales :

SUCCESS
REJECTED
FAILED
DEFERRED
UNKNOWN
CONTRADICTION

## Transport

Le transport intervient après la construction du message sémantique.

EXECUTE(message) ne dépend pas du transport.

SEND(message, transport) produit un résultat de livraison sans modifier le message source.

## Invariant fondamental

Pour tout transport T :

SEMANTICS(SEND(M,T)) = SEMANTICS(M)

si SEND retourne un message ou une représentation de message.

En cas d'échec :

DELIVERY != DELIVERED

Le message original reste inchangé.

## Accessibilité

Les résultats de l'exécuteur sont projetables vers :

VOICE
BRAILLE
KEYBOARD
DISPLAY
TOUCH
NETWORK
AUTOMATION

La projection ne modifie jamais RESULT.

## Preuve

L'exécuteur ne peut annoncer qu'un niveau effectivement établi.

DEFINED n'est pas EXECUTED.
EXECUTED n'est pas OBSERVED.
OBSERVED n'est pas HARDWARE.
HARDWARE n'est pas RF_PROVEN.

## Premier objectif d'implémentation

Créer une implémentation minimale capable de :

1. lire un record ;
2. valider ses champs ;
3. effectuer une transition déterministe ;
4. produire un événement ;
5. produire un résultat ;
6. rejeter une contradiction ;
7. conserver UNKNOWN ;
8. ne pas confondre livraison et exécution.

Cette implémentation devra ensuite être testée par les vecteurs existants.
