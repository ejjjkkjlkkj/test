# ZERO — expérience d'intelligence sans modèle

Cette expérience repart de zéro et ne dépend d'aucun modèle, aucune API d'IA et aucun framework d'IA.

Langage : Python standard uniquement.

## Idée

Au lieu de demander « quelle réponse ressemble à une bonne réponse ? », le système cherche des **actions qui modifient son propre espace d'hypothèses**, observe ce que ces actions produisent, puis conserve uniquement les transformations qui ouvrent des possibilités nouvelles ou résolvent une contrainte.

Le moteur ne contient :
- ni réseau neuronal ;
- ni embeddings ;
- ni tokenizer LLM ;
- ni dataset d'entraînement ;
- ni appel à un modèle externe ;
- ni algorithme d'IA importé comme cœur.

Le cœur expérimental est volontairement petit :
1. perception ;
2. fragmentation ;
3. recombinaison ;
4. perturbation ;
5. confrontation ;
6. conséquence ;
7. mémoire des transformations utiles.

L'objectif n'est pas de prétendre être supérieur à l'humain ou à Claude. L'objectif est de produire une mécanique suffisamment différente pour pouvoir être testée.

## Lancer

```text
python zero.py
python zero.py "un problème à explorer"
```

Chaque expérience produit une trace lisible dans `runs/`.

## Règle

Ne pas optimiser trop tôt. Une idée qui échoue est une donnée expérimentale.
