# ZERO — contrat de fiabilité 99,99 %

## 1. Séparation obligatoire

Deux notions ne doivent jamais être mélangées :

- **Vérité formelle** : par exemple `1 + 1 = 2`.
- **Fiabilité physique/logicielle** : objectif d'ingénierie fixé à **99,99 %**.

Le 99,99 % ne signifie pas que la vérité mathématique vaut 99,99 %.

La vérité formelle reste la référence du système.

## 2. Objectif

Le système vise une fiabilité globale de 99,99 % dans les conditions et hypothèses explicitement couvertes par ses preuves, tests et mécanismes de tolérance aux fautes.

Les conditions non couvertes ne sont jamais transformées en preuve.

## 3. Échec contrôlé

Une panne, corruption, capacité absente ou contradiction doit produire un état explicite :

- `FAILURE`
- `UNKNOWN`
- `REJECTED`
- `MISSING_CAPABILITY`

Elle ne doit jamais modifier une vérité formelle.

## 4. Indépendance matérielle

CPU, GPU, architecture, OS, mémoire, écran, clavier, réseau et autres périphériques sont des moyens d'exécution.

Ils peuvent réduire la disponibilité ou la performance.

Ils ne peuvent pas redéfinir la sémantique.

## 5. Accessibilité

VOICE, BRAILLE, KEYBOARD, DISPLAY, TOUCH, POINTER, NETWORK et AUTOMATION doivent représenter le même résultat sémantique.

Une panne d'une modalité ne doit pas modifier le résultat sémantique.

## 6. Mesure

Le projet doit publier séparément :

1. couverture des tests ;
2. propriétés formellement vérifiées ;
3. fautes injectées et fautes tolérées ;
4. conditions physiques couvertes ;
5. conditions non couvertes ;
6. preuves réelles obtenues.

Aucune valeur de fiabilité ne doit être inventée à partir d'un simple nombre de tests.

## 7. Règle fondamentale

`99,99 % de fiabilité` est notre **objectif d'ingénierie**.

`1 + 1 = 2` est notre **invariant mathématique**.

Ces deux affirmations restent strictement séparées.
