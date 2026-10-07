# ZERO — moteur de référence abstrait

## 1. Objet

Le moteur de référence définit le comportement observable minimal du premier noyau ZERO.

Il ne choisit ni CPU, ni langage hôte, ni système d'exploitation.

Il reçoit un flux de RECORD, le valide, conserve l'état sémantique et produit des RESULT, EVENT et OBSERVATION.

## 2. État du moteur

ENGINE {
  version
  sequence
  objects
  capabilities
  events
  observations
  results
  proof
}

Le moteur conserve l'état précédent. Une observation ne remplace pas l'objet.

## 3. Cycle

INPUT
→ DECODE
→ VALIDATE
→ RESOLVE
→ TRANSITION
→ EMIT
→ OBSERVE
→ RESULT

Une entrée invalide s'arrête à VALIDATE.

Une entrée valide mais inconnue est conservée sans interprétation forcée.

## 4. CREATE

CREATE d'un OBJECT valide :

1. vérifier l'identité ;
2. vérifier qu'elle n'est pas déjà utilisée dans le même espace ;
3. créer l'objet ;
4. enregistrer son état initial ;
5. enregistrer ses capacités déclarées ;
6. produire l'événement correspondant ;
7. produire l'observation ;
8. produire RESULT=COMPLETED.

## 5. OBSERVE

OBSERVE ne modifie pas l'objet.

Il retourne au minimum :

IDENTITY
MEANING
STATE
CAPABILITIES
EVENTS_RELEVANT
PROOF

## 6. ACTION

Une opération d'action doit posséder :

INTENT
TARGET
INPUTS
REQUIRED_CAPABILITIES
AUTHORIZATION

Absence de capacité ou d'autorisation :

RESULT=REJECTED

Le moteur ne tente pas une action de remplacement.

## 7. UNKNOWN

Un objet ou type inconnu reste présent avec :

IDENTITY
RAW_TYPE ou RAW_DATA
STATE=UNKNOWN
OBSERVATIONS
UNCERTAINTY
PROOF

UNKNOWN n'est jamais converti automatiquement en ERROR.

## 8. Interruption

Une opération interrompue produit :

RESULT=INTERRUPTED

avec :

PARTIAL_STATE
CAUSE
OBSERVATIONS
NEXT_VALID_STATE

Aucune partie non observée ne doit être inventée.

## 9. Déterminisme

À entrée identique et état initial identique :

même état sémantique
mêmes identités
même ordre logique des événements
même résultat

Le temps physique peut différer lorsqu'il est mesuré par l'environnement ; il ne doit pas être utilisé pour masquer une différence sémantique.

## 10. Séparation des preuves

Le moteur ne transforme jamais :

SIMULATED → HARDWARE
QEMU → HARDWARE
HARDWARE → RF_PROVEN
RF_PROVEN → SATELLITE_LINK_PROVEN

Une preuve supérieure doit provenir d'une observation ou mesure correspondant réellement à ce niveau.

## 11. Accessibilité

Le moteur expose la même structure à toutes les projections :

VOICE
BRAILLE
KEYBOARD
DISPLAY
NETWORK
AUTOMATION

La projection ne possède aucune copie concurrente de l'état.

## 12. Statut

MODEL : DEFINED

REFERENCE IMPLEMENTATION : NOT_IMPLEMENTED

EXECUTION : NOT_EXECUTED

HARDWARE : NOT_PROVEN
