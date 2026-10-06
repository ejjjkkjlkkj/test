# TEST — Experimental Intelligence Core

Projet expérimental construit from scratch.

## Intention

Explorer une architecture d'intelligence qui ne cherche pas à reproduire directement les méthodes humaines ou les architectures LLM classiques.

Le moteur travaille par :
- expériences concurrentes ;
- mémoire d'épisodes ;
- hypothèses révisables ;
- contradictions explicites ;
- recherche de transformations ;
- sélection par conséquences observées ;
- auto-révision du comportement.

Il ne prétend pas être supérieur à Claude ou à l'humain par déclaration : la supériorité doit être démontrée par des tests reproductibles.

## Contraintes

- zéro modèle pré-entraîné ;
- zéro API d'IA externe ;
- zéro copie d'algorithme existant comme cœur du système ;
- pas de dépendance à un théorème mathématique ;
- expérimentation mesurable et réversible.

## Première cible

Construire un agent capable de trouver une réponse utile lorsqu'aucune procédure explicite ne lui est fournie, puis de conserver la trace de ce qui a fonctionné ou échoué.

## Langage machine

La branche `experiment/non-human-zero-v2` développe un langage machine natif décrit par `ZERO.LANGUAGE`.

Le langage privilégie les transformations, les traces, les collisions, les contradictions, la réinjection et le transfert plutôt que des instructions humaines prédéfinies.

Fichiers centraux :
- `ZERO.FIELD` — substrat initial ;
- `ZERO.LANGUAGE` — grammaire et comportement du langage ;
- `ZERO.PROGRAMS` — programmes-seeds ;
- `ZERO.LANGUAGE-TEST` — falsification du langage ;
- `ZERO.PROOF` — protocole de preuve ;
- `ZERO.REALITY-PLAN` — validation externe.

Aucune supériorité n'est considérée comme acquise. Le seul verdict acceptable est expérimental et reproductible.
