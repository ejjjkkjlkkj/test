# ZERO — exécution universelle

## 1. Objectif

Un programme ZERO ne doit pas être conçu pour un ordinateur particulier.

Le même programme sémantique doit pouvoir être exécuté sur toute réalisation qui fournit les capacités nécessaires :

PC
IPHONE
IPAD
ANDROID
TABLETTE
WATCH
TV
AUTOMOBILE
SERVEUR
EMBARQUÉ
AUTRE MACHINE
NŒUD RÉSEAU
et, lorsque la réalisation le permet, une machine très contrainte ou un environnement distribué.

Le programme ne change pas de sens lorsqu'il change de machine.

## 2. Règle fondamentale

La portabilité ZERO est :

PROGRAMME SÉMANTIQUE
→ FORMAT CANONIQUE
→ CAPACITÉS REQUISES
→ RÉALISATION LOCALE
→ EXÉCUTION
→ OBSERVATION
→ RÉSULTAT

Le programme ne dépend pas directement d'une API propre à une plateforme pour définir son sens.

Une plateforme fournit une réalisation des capacités.

## 3. Identité du programme

L'identité d'un programme est indépendante de :

- CPU ;
- architecture processeur ;
- système d'exploitation ;
- écran ;
- résolution ;
- clavier ;
- souris ;
- écran tactile ;
- haut-parleur ;
- moteur vocal ;
- braille ;
- stockage ;
- protocole de transport ;
- adresse réseau ;
- appareil physique.

Un même PROGRAM_ID peut donc être exécuté sur plusieurs machines.

## 4. Contrat de réalisation

Chaque réalisation déclare uniquement des capacités observées :

DEVICE
IDENTITY
ARCHITECTURE
EXECUTION_CAPABILITIES
INPUT_CAPABILITIES
OUTPUT_CAPABILITIES
COMMUNICATION_CAPABILITIES
STORAGE_CAPABILITIES
ACCESSIBILITY_CAPABILITIES
SECURITY_CAPABILITIES
PROOF

Une capacité absente n'est jamais inventée.

Une capacité inconnue reste UNKNOWN.

## 5. Adaptation sans changement de sens

Exemple :

PROGRAM
    OBJECT = NOTE
    OPERATION = CREATE
    INPUT = "Bonjour"

La réalisation PC peut utiliser ses ressources locales.

La réalisation iPhone peut utiliser ses ressources locales.

La réalisation Android peut utiliser ses ressources locales.

La réalisation embarquée peut utiliser ses propres ressources.

Le résultat sémantique doit rester équivalent.

L'adaptation porte sur la réalisation, jamais sur la signification du programme.

## 6. Applications

Une application ZERO est donc un programme sémantique, pas une application liée à une seule interface graphique.

Une application peut posséder plusieurs projections :

VOICE
BRAILLE
KEYBOARD
DISPLAY
TOUCH
POINTER
AUDIO
HAPTIC
NETWORK
AUTOMATION

Aucune projection ne devient la source de vérité.

## 7. Réseau et exécution distante

Une application peut être :

LOCAL
REMOTE
DISTRIBUTED

Un objet peut être observé sur une autre machine avec le même modèle :

IDENTITY
MEANING
STATE
CAPABILITIES
EVENTS
OBSERVATIONS
RESULT
PROOF

Le transport réseau est une réalisation de COMMUNICATE.

La réussite de la connexion ne prouve pas que l'opération distante a réussi.

L'opération distante doit produire une observation et un résultat.

## 8. Migration

Un programme peut passer d'une réalisation à une autre :

DISCOVER
→ IDENTIFY
→ CHECK_CAPABILITIES
→ LOAD
→ EXECUTE
→ OBSERVE
→ VERIFY

La migration ne modifie pas silencieusement l'état sémantique.

Si une capacité manque :

MISSING_CAPABILITY

Si une capacité est inconnue :

UNKNOWN_CAPABILITY

Si l'état ne peut pas être restauré :

MIGRATION_REJECTED

## 9. Compatibilité

La compatibilité n'est pas :

"l'application semble fonctionner".

La compatibilité est :

SAME_PROGRAM_ID
+ SAME_SEMANTIC_INPUT
+ REQUIRED_CAPABILITIES_AVAILABLE
+ EXECUTION
+ OBSERVATION
+ EQUIVALENT_SEMANTIC_RESULT
+ ACCESSIBILITY_TESTED
+ PROOF

## 10. Ordinateur très contraint

Une réalisation très limitée peut exécuter une partie du programme si les capacités requises existent.

Elle ne doit jamais prétendre exécuter ce qu'elle ne peut pas exécuter.

Un programme complet peut être décomposé en opérations sémantiques distribuées, mais chaque opération conserve :

IDENTITY
SEQUENCE
STATE
EVENT
RESULT
PROOF

## 11. Application sans écran

L'absence d'écran ne rend pas le programme inaccessible.

La même application peut utiliser :

VOICE
BRAILLE
AUDIO
KEYBOARD
NETWORK
AUTOMATION

selon les capacités réellement disponibles.

## 12. Application sans haut-parleur

VOICE n'est pas supposée disponible.

Le moteur recherche les projections disponibles et requises.

Si une modalité obligatoire est absente, l'instance n'est pas ACCEPTED comme accessible.

Elle reste NOT_PROVEN ou REJECTED selon le contrat d'acceptation.

## 13. Compatibilité future

Une nouvelle machine ne nécessite pas une nouvelle définition du langage.

Elle doit fournir une réalisation du contrat :

SEMANTIC PROGRAM
→ CAPABILITIES
→ EXECUTION
→ OBSERVATION
→ ACCESSIBILITY
→ PROOF

Ainsi une nouvelle architecture CPU, un nouveau système ou un nouvel appareil peuvent rejoindre ZERO sans devenir une nouvelle source de vérité.

## 14. Règle absolue

ZERO ne promet pas :

"fonctionne partout".

ZERO exige une preuve par réalisation.

La propriété recherchée est :

SAME SEMANTICS EVERYWHERE
+
LOCAL REALIZATION
+
OBSERVED RESULT
+
ACCESSIBLE RESULT
+
PROOF

Un appareil non testé reste NON_PROUVÉ.

## 15. Statut

UNIVERSAL MODEL = DEFINED
CROSS-DEVICE EXECUTION = NOT_IMPLEMENTED
PC = NOT_PROVEN
IPHONE = NOT_PROVEN
ANDROID = NOT_PROVEN
EMBEDDED = NOT_PROVEN
DISTRIBUTED EXECUTION = NOT_PROVEN
