# ZERO — indépendance du support physique

## Principe

Le support physique n'est pas une contrainte sémantique du langage ni du modèle IA.

ZERO doit pouvoir être réalisé sur toute machine physique capable de fournir les ressources et capacités minimales nécessaires, notamment :

PC
SMARTPHONE
TABLETTE
SERVEUR
EMBARQUÉ
OBJET MOBILE
ROBOT
DRONE
AÉRONEF
PLATEFORME AÉRIENNE
VÉHICULE
MACHINE SPATIALE
AUTRE SUPPORT PHYSIQUE

Le support peut être fixe, mobile, aérien, spatial ou distribué.

## IA

Le modèle IA ZERO utilise le même contrat sur chaque support :

INPUT
→ REPRESENTATION
→ EXECUTION
→ ACTION
→ OBSERVATION
→ VERIFICATION
→ RESULT
→ ACCESSIBLE_PROJECTION
→ PROOF

Le modèle ne doit pas supposer :

- écran ;
- clavier ;
- souris ;
- alimentation permanente ;
- réseau permanent ;
- GPS/GNSS ;
- stockage permanent ;
- connexion à un serveur ;
- puissance CPU particulière ;
- GPU ;
- architecture processeur particulière ;
- présence humaine.

Une capacité absente est ABSENT.
Une capacité inconnue est UNKNOWN.
Une capacité insuffisante provoque une exécution partielle ou un refus explicite.

## Fonctionnement aérien

Une réalisation aérienne peut fonctionner sans connexion permanente au sol.

Le runtime doit pouvoir :

- exécuter localement ;
- conserver son état ;
- fonctionner avec une connectivité intermittente ;
- enregistrer les observations ;
- continuer les opérations autorisées ;
- synchroniser ultérieurement ;
- préserver les séquences et identités ;
- ne jamais transformer une absence de communication en succès fictif.

La perte du réseau est un état observable, pas une destruction du sens.

## Support physique arbitraire

Le même programme sémantique peut être chargé sur plusieurs supports.

Le programme conserve :

PROGRAM_ID
SEMANTICS
IDENTITIES
STATE
EVENTS
RESULTS
PROOF

La réalisation fournit :

HARDWARE
MEMORY
COMPUTE
INPUT
OUTPUT
COMMUNICATION
STORAGE
ACCESSIBILITY
POWER
SECURITY

Le runtime adapte uniquement l'exécution physique.

## Limites réelles

L'indépendance du support ne signifie pas qu'une machine peut exécuter une opération impossible.

Exemple :

CAPABILITY_REQUIRED = GPU.COMPUTE
CAPABILITY_PRESENT = FALSE

Résultat :

REJECTED
MISSING_CAPABILITY

Il est interdit de simuler silencieusement cette capacité et de déclarer l'opération réussie.

## IA embarquée

Une IA embarquée peut fonctionner :

LOCAL_ONLY
OFFLINE
INTERMITTENT_NETWORK
REMOTE_ASSISTED
DISTRIBUTED

Le mode d'exécution est observable.

Une réponse produite localement et une réponse produite à distance doivent conserver leur provenance.

## Accessibilité

L'accessibilité reste obligatoire sur tous les supports.

Un support sans écran peut utiliser les capacités disponibles :

VOICE
BRAILLE
AUDIO
HAPTIC
NETWORK
AUTOMATION
autres projections réellement disponibles.

L'absence d'une modalité obligatoire interdit l'acceptation comme accessible pour cette configuration.

## Preuve

Les catégories suivantes restent séparées :

MODEL_DEFINED
RUNTIME_IMPLEMENTED
EXECUTED
OBSERVED
ACCESSIBILITY_TESTED
PROVEN

Un test sur PC ne constitue pas une preuve physique pour un aéronef.

Un test simulé d'aéronef ne constitue pas une preuve d'aéronef réel.

## Statut

PHYSICAL_INDEPENDENCE_MODEL = DEFINED
AIRBORNE_RUNTIME = NOT_IMPLEMENTED
MULTI_SUPPORT_RUNTIME = NOT_PROVEN
PHYSICAL_AI = NOT_PROVEN
