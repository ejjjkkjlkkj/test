# ZERO — architecture canonique

## 1. Source unique du sens

Le noyau ZERO suit une seule chaîne sémantique :

MACHINE
→ STATE
→ OBJECT
→ RELATION
→ CAPABILITY
→ OPERATION
→ TRANSITION
→ EVENT
→ OBSERVATION
→ RESULT
→ PROOF
→ PROJECTION

Cette chaîne est normative.

Aucune projection, interface, modalité ou liaison réseau ne devient une seconde source de vérité.

## 2. Trois niveaux

### Niveau A — Sémantique

OBJECT, STATE, RELATION, CAPABILITY, OPERATION, TRANSITION, EVENT, OBSERVATION, RESULT, PROOF.

### Niveau B — Représentation

RECORD et format machine canonique.

### Niveau C — Réalisation

CPU, mémoire, bus, pilotes, périphériques, stockage, réseau, radio, affichage, audio, autres machines.

Le niveau C réalise le niveau A. Il ne redéfinit pas son sens.

## 3. Accessibilité

L'accessibilité est une projection du même état sémantique :

VOICE
BRAILLE
KEYBOARD
DISPLAY
TOUCH
POINTER
NETWORK
AUTOMATION
FUTURE_MODALITY

Aucune de ces projections n'est privilégiée.

## 4. Communication

Une machine locale et une machine distante utilisent le même modèle :

OBJECT + EVENT + OBSERVATION + OPERATION + RESULT + PROOF

Le transport peut changer.

Le sens ne change pas.

## 5. Inconnu

UNKNOWN est un état sémantique valide.

Un objet inconnu conserve :

IDENTITY
RAW_INFORMATION
OBSERVATIONS
RELATIONS
UNCERTAINTY
SAFE_CAPABILITIES
PROOF

Il ne doit jamais être supprimé simplement parce que son type n'est pas connu.

## 6. Sécurité

Le modèle distingue :

OBSERVE
ACT
AUTHORIZE

Une projection accessible ne reçoit aucun droit supplémentaire.

## 7. Preuve

Les niveaux sont strictement séparés :

DEFINED
IMPLEMENTED
TESTED
SIMULATED
QEMU
HARDWARE
RF_PROVEN
SATELLITE_LINK_PROVEN

Une couche ne peut pas augmenter elle-même son niveau de preuve.

## 8. Règle de cohérence

Tout nouveau composant ZERO doit répondre à quatre questions :

1. Quel objet ou état sémantique représente-t-il ?
2. Quelles observations produit-il ?
3. Quelles capacités et opérations expose-t-il ?
4. Quelle preuve démontre son fonctionnement ?

S'il introduit une seconde représentation concurrente du sens, il est hors architecture.

## 9. Statut

ARCHITECTURE : DEFINED
IMPLEMENTATION COMPLETE : NO
FULL MACHINE EXECUTION : NOT_EXECUTED
HARDWARE : NOT_PROVEN
