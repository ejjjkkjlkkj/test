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

## Exécution

```text
cargo run -- solve "..."
cargo run -- experiment
cargo test
```
