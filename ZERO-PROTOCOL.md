# ZERO — Protocole sémantique machine-à-machine

Le protocole ZERO transporte le sens nécessaire à une communication entre machines.

Il ne définit pas encore un format binaire final.

## Séquence minimale

HELLO
-> IDENTITY
-> CAPABILITIES
-> AUTHORIZATION
-> SESSION
-> EVENTS / REQUESTS / RESULTS
-> CLOSE

## Message sémantique

Chaque message doit pouvoir exprimer :

SOURCE
TARGET
IDENTITY
SEQUENCE
INTENT ou EVENT
PAYLOAD SEMANTIQUE
CAPABILITIES
STATE
RESULT
UNCERTAINTY

## Règles

1. Une machine ne doit pas prétendre posséder une capacité absente.
2. Un message doit rester interprétable même si une partie est inconnue.
3. L'ordre des événements doit être conservé lorsque le sens l'exige.
4. Une interruption doit être observable.
5. Une demande refusée produit un résultat sémantique.
6. L'authentification et l'autorisation sont distinctes.
7. Le transport réseau n'est pas le sens.

## Évolution

Une version future peut ajouter des champs sans rendre les anciens objets opaques.

UNKNOWN_FIELD
est une donnée, pas une erreur fatale.

## Sécurité

Le protocole n'autorise jamais une machine distante à dépasser les capacités explicitement accordées.
