# ZERO — Forme initiale du langage

La syntaxe n'est pas encore figée.

Le noyau doit d'abord être défini par les relations sémantiques qu'une représentation peut exprimer.

## Forme conceptuelle

OBSERVE target
REQUEST intent ON target
RELATE source WITH relation target
READ target
WRITE target VALUE value
STEP
NEXT EVENT
PROJECT object TO modality

Cette forme est une notation de conception, pas une syntaxe finale.

## Propriété

Une syntaxe future doit pouvoir exprimer directement :

- machine ;
- objet ;
- relation ;
- événement ;
- capacité ;
- transition ;
- observation ;
- communication distante ;
- projection.

Elle ne doit pas imposer un modèle de classe, de fenêtre, de DOM ou d'application.

## Extensibilité

Un programme doit pouvoir rencontrer un type inconnu et continuer à l'observer sans devoir être recompilé pour chaque nouveau type d'objet.

UNKNOWN est donc une valeur sémantique valide.
