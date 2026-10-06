# ZERO — Sémantique des périphériques

ZERO ne construit pas une accessibilité spéciale pour les périphériques.

Il rend les périphériques eux-mêmes observables, navigables et actionnables selon leurs capacités.

## Modèle

DEVICE
=
IDENTITY
+
STATE
+
CAPABILITIES
+
OBSERVATIONS
+
EVENTS
+
RELATIONS
+
SECURITY
+
PROJECTIONS
+
UNCERTAINTY

## Découverte

Lorsqu'un périphérique apparaît :

1. existence détectée ;
2. frontière déterminée si possible ;
3. identité récupérée si disponible ;
4. capacités observées ;
5. relations déterminées ;
6. événements publiés ;
7. état rendu navigable.

Aucune étape ne nécessite de connaître à l'avance tous les types de périphériques.

## CPU

Le CPU expose son état observable et ses événements selon les droits disponibles.

Le langage peut demander :

OBSERVE(CPU)
INSPECT(CPU)
EXECUTE(...)
INTERRUPT(...)
MODIFY(...)

Chaque opération est soumise à la capacité correspondante.

## GPU

Le GPU expose notamment :

- ressources ;
- capacités ;
- mémoire ;
- travaux ;
- surfaces ;
- événements ;
- erreurs ;
- résultats.

Une projection graphique peut être reliée à son objet sémantique d'origine.

## Moniteur

Le moniteur expose :

- présence ;
- identité ;
- connexion ;
- capacités ;
- état ;
- surface ;
- événements.

Le fait qu'une information soit affichée sur un moniteur ne rend pas cette information exclusivement visuelle.

## Audio

Un périphérique audio expose :

- entrée ou sortie ;
- capacités ;
- état ;
- flux ;
- événements ;
- latence observée ;
- erreurs ;
- relations avec les objets sonores.

La parole peut être une projection d'un résultat sémantique, et non un canal séparé du langage.

## Entrées

Clavier, pointeur, tactile, manette, capteur ou autre entrée suivent le même chemin :

SIGNAL PHYSIQUE
-> OBSERVATION
-> ÉVÉNEMENT
-> INTENTION

La modalité physique ne définit pas l'action.

## Périphérique inconnu

ZERO ne doit pas déclarer automatiquement :

UNKNOWN = INACCESSIBLE

Au contraire :

UNKNOWN
=
EXISTENCE
+
OBSERVATION
+
CAPABILITIES ?
+
RELATIONS ?
+
UNCERTAINTY

Le système peut ainsi continuer à communiquer avec un monde qu'il ne comprend pas encore complètement.

## Projection universelle

Un même état de périphérique peut être projeté vers plusieurs modalités.

Exemple :

ÉTAT GPU
-> visuel
-> vocal
-> braille
-> navigation clavier
-> automatisation

Les projections ne modifient pas le sens source.

## Principe final

Le langage ne « commande pas des périphériques » comme des boîtes noires.

Il communique avec des entités de la machine.

La distinction importante devient :

OBSERVABLE
vs
MODIFIABLE

et :

ACCESSIBLE
vs
AUTORISÉ

Cette séparation permet une accessibilité native sans supprimer la sécurité.
