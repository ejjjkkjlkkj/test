# Hypothèse de travail

Le projet n'essaie pas de rendre une machine « humaine ».

L'hypothèse est différente :

> Une intelligence plus forte peut émerger d'un système qui considère toute réponse comme une expérience provisoire, conserve les conséquences de ses expériences, provoque volontairement des contradictions et change de représentation quand une représentation cesse de produire des progrès.

## Ce qui est volontairement absent

- raisonnement fondé sur un théorème particulier ;
- réseau neuronal pré-entraîné ;
- prompt engineering comme mécanisme principal ;
- imitation d'un agent humain ;
- copie d'une architecture de modèle commercial.

## Critère de supériorité

« Mieux que Claude » ne sera jamais accepté comme une affirmation.

Le projet doit battre des références sur des suites de tâches inconnues selon :
1. taux de résolution ;
2. capacité à détecter ses propres erreurs ;
3. capacité à sortir d'une impasse ;
4. généralisation vers une tâche jamais vue ;
5. coût en temps et en ressources ;
6. stabilité après accumulation d'expérience.

Une victoire sur un seul test ne suffit pas.

## Direction expérimentale

Le noyau actuel utilise une boucle :

EXPÉRIENCE → TENSION → RÉVISION → CONFRONTATION → MÉMOIRE

La prochaine évolution doit chercher à remplacer progressivement les règles écrites à la main par des mécanismes capables de découvrir eux-mêmes de nouvelles opérations.
