# ZERO — Communication native avec la machine

ZERO ne considère pas le CPU, le moniteur, le GPU, l'audio, le clavier ou les autres périphériques comme des ressources externes à contrôler au moyen d'une API.

Ils font partie du monde sémantique de la machine.

## 1. Principe

Le langage doit pouvoir communiquer avec :

- CPU et unités d'exécution ;
- mémoire ;
- GPU et unités graphiques ;
- moniteur et sorties d'affichage ;
- audio et sorties/entrées sonores ;
- clavier, pointeur, tactile et autres entrées ;
- stockage ;
- réseau ;
- caméra et capteurs ;
- firmware et état de démarrage ;
- énergie et thermique ;
- bus et périphériques ;
- périphériques inconnus.

Cette communication ne signifie pas que le langage obtient automatiquement tous les droits.

Elle signifie que la machine expose d'abord son état, ses capacités et ses événements sous une forme sémantique.

## 2. Direction de communication

Le chemin minimal est :

PHYSIQUE
-> OBSERVATION
-> ÉVÉNEMENT
-> INTENTION
-> TRANSITION
-> NOUVEL ÉTAT
-> PROJECTION
-> PHYSIQUE

Exemple écran :

MONITEUR
-> état du périphérique
-> capacités d'affichage
-> surface sémantique
-> demande d'affichage
-> transition
-> nouvel état
-> projection visuelle

Exemple clavier :

CLAVIER
-> signal physique
-> observation
-> événement de touche
-> intention
-> transition
-> résultat

Exemple GPU :

GPU
-> identité + état + capacités
-> événements et observations
-> ressources graphiques sémantiques
-> intention de calcul ou de projection
-> transition
-> résultat observable

## 3. CPU

Le CPU n'est pas seulement une suite d'instructions opaque.

ZERO doit pouvoir représenter au niveau approprié :

- existence de l'unité d'exécution ;
- état d'exécution ;
- contexte courant ;
- événements ;
- capacités ;
- interruptions ;
- relations avec mémoire et périphériques ;
- résultats ;
- erreurs et arrêts ;
- incertitude lorsque certaines informations ne sont pas disponibles.

L'observation d'un CPU n'implique pas le droit de modifier son état.

## 4. GPU

Le GPU est un objet machine.

ZERO peut distinguer :

- identité ;
- présence ;
- capacités ;
- mémoire associée ;
- unités disponibles ;
- files ou travaux ;
- surfaces ;
- événements ;
- état ;
- sécurité ;
- résultats.

Une image affichée n'est donc pas nécessairement réduite à des pixels pour le langage.

Le langage peut conserver le sens de la scène et produire ensuite une projection visuelle.

## 5. Moniteur

Un moniteur est également un objet sémantique.

ZERO peut représenter :

- présence ;
- identité connue ou inconnue ;
- dimensions observées ;
- capacités ;
- état ;
- connexion ;
- surface de sortie ;
- événements ;
- relations avec GPU et autres sorties ;
- incertitude.

La sortie visuelle est une projection, pas la source unique du sens.

## 6. Accessibilité native

La même transition sémantique doit pouvoir être projetée vers :

- affichage ;
- parole ;
- braille ;
- clavier ;
- pointeur ;
- tactile ;
- automatisation ;
- toute future modalité.

Aucune de ces modalités n'est propriétaire du sens.

Un utilisateur qui n'utilise pas le moniteur visuellement doit pouvoir atteindre les mêmes objets, états, actions, résultats et refus que celui qui l'utilise visuellement, sous réserve des capacités et de la sécurité.

## 7. Sécurité

Communication native != accès illimité.

Chaque opération vérifie sa capacité.

Ainsi :

OBSERVE(CPU) != MODIFY(CPU)
OBSERVE(GPU) != CONTROL(GPU)
OBSERVE(MONITOR) != RECONFIGURE(MONITOR)

Un refus est lui-même un résultat sémantique et accessible.

## 8. Règle fondamentale

Aucun périphérique connu ne doit devenir opaque uniquement parce qu'une couche intermédiaire ne possède pas encore d'interface d'accessibilité.

Un périphérique inconnu doit au minimum exposer :

EXISTENCE
+ IDENTITÉ SI CONNUE
+ FRONTIÈRE OBSERVÉE
+ CAPACITÉS OBSERVÉES
+ ÉVÉNEMENTS OBSERVÉS
+ INCERTITUDE

Ainsi, l'accessibilité commence avant l'interface utilisateur.
