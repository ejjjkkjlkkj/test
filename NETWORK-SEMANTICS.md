# ZERO — Communication native entre machines

Le réseau n'est pas une API ajoutée au langage.

Une autre machine est une entité du monde machine que ZERO peut découvrir, observer, identifier, communiquer avec et, si autorisé, faire agir.

## 1. Deux machines, un même espace sémantique

Exemple :

PC ZERO
<-> réseau local
<-> iPhone

ZERO représente l'iPhone comme un objet distant :

DEVICE
+ IDENTITY
+ NETWORK-IDENTITY
+ STATE
+ CAPABILITIES
+ OBSERVATIONS
+ EVENTS
+ RELATIONS
+ UNCERTAINTY

L'iPhone n'a pas besoin d'être construit avec ZERO pour être représentable.

## 2. Découverte

La découverte réseau produit progressivement :

EXISTENCE
-> PRÉSENCE RÉSEAU
-> IDENTITÉ SI DISPONIBLE
-> PROPRIÉTAIRE / AUTORITÉ SI DISPONIBLE
-> CAPACITÉS OBSERVÉES
-> SERVICES OBSERVÉS
-> RELATIONS
-> INCERTITUDE

La découverte ne doit pas supposer que tous les appareils sont connus à l'avance.

ZERO ne transforme pas une adresse réseau en identité permanente.

NETWORK ADDRESS != MACHINE IDENTITY

Une adresse peut changer alors que l'identité logique reste la même.

## 3. Communication

Le chemin sémantique devient :

LOCAL MACHINE
-> OBSERVE
-> DISCOVER REMOTE ENTITY
-> ESTABLISH COMMUNICATION
-> EXCHANGE SEMANTIC EVENTS
-> REQUEST
-> REMOTE TRANSITION
-> RESULT
-> OBSERVATION

Le transport physique ou protocolaire est une couche de réalisation.

Le langage manipule d'abord la relation et le sens.

## 4. Exemple PC -> iPhone

Si l'iPhone et le PC sont sur le même réseau :

1. ZERO observe les entités réseau visibles ;
2. un appareil correspondant à l'iPhone est découvert ;
3. ZERO représente l'appareil distant ;
4. les capacités réellement exposées sont découvertes ;
5. une relation sécurisée peut être établie ;
6. ZERO peut ensuite demander une opération autorisée ;
7. l'iPhone renvoie un résultat ou un refus ;
8. le résultat devient immédiatement un objet sémantique navigable.

Exemple conceptuel :

DISCOVER
-> DEVICE(iPhone)
-> CAPABILITIES
-> CONNECT
-> REQUEST
-> RESULT

Le protocole concret n'est pas le langage lui-même.

## 5. Accessibilité distante

L'accessibilité ne s'arrête pas au PC.

Un objet provenant de l'iPhone doit conserver les mêmes propriétés :

- identité ;
- sens ;
- structure ;
- état ;
- relations ;
- capacités ;
- événements ;
- observations ;
- incertitude ;
- projection.

Ainsi, ZERO peut projeter un objet distant en :

- parole ;
- braille ;
- clavier ;
- affichage ;
- automatisation ;
- autre modalité.

Il n'est pas nécessaire de fabriquer une deuxième couche d'accessibilité pour le réseau.

## 6. Sécurité

Être sur le même réseau ne signifie pas être autorisé.

Ces opérations restent distinctes :

OBSERVE
DISCOVER
CONNECT
AUTHENTICATE
REQUEST
MODIFY
ADMINISTER

Un appareil découvert peut donc être :

VISIBLE
mais
NON-AUTORISÉ

Un refus est lui-même un événement sémantique :

DENIED
+ SOURCE
+ TARGET
+ REQUIRED_CAPABILITY
+ REASON

ZERO ne doit jamais contourner l'autorisation simplement parce qu'un appareil est accessible physiquement sur le réseau.

## 7. Appareils ZERO et appareils externes

Deux catégories peuvent coexister :

### Machine ZERO

Elle comprend nativement le protocole sémantique ZERO.

### Machine externe

Elle ne connaît pas ZERO.

Dans ce cas, une frontière d'adaptation peut traduire les capacités réellement disponibles vers le modèle sémantique ZERO.

Mais l'adaptation ne doit jamais inventer une capacité absente.

UNKNOWN
reste
UNKNOWN.

## 8. Communication bidirectionnelle

La relation n'est pas uniquement :

PC -> iPhone

Elle peut être :

PC <-> iPhone
PC <-> autre PC
PC <-> tablette
PC <-> serveur
PC <-> périphérique réseau
PC <-> future machine ZERO

Les événements distants peuvent également revenir vers ZERO :

REMOTE EVENT
-> OBSERVATION
-> SEMANTIC EVENT
-> LOCAL STATE UPDATE

Le réseau devient donc une extension du monde observable de la machine, sans fusionner les identités des machines.

## 9. Règle fondamentale

> Une machine distante est un objet sémantique distant, pas une adresse IP.

ZERO doit pouvoir communiquer avec un monde hétérogène sans perdre :

- identité ;
- sens ;
- sécurité ;
- état ;
- événements ;
- incertitude ;
- accessibilité.
