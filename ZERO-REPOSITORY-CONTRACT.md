# ZERO — contrat de conformité global du dépôt

## Règle

Le dépôt entier est soumis à une seule règle d'acceptation :

`FONCTIONNEL ET ACCESSIBLE`

Une fonctionnalité qui fonctionne mais n'est pas accessible est REFUSÉE.

Une fonctionnalité accessible mais non fonctionnelle est REFUSÉE.

Une spécification, un modèle, un vecteur ou une documentation ne constitue pas une implémentation fonctionnelle.

## États autorisés

Pour une unité exécutable :

- `FUNCTIONAL_ACCESSIBLE` : exécution réelle + résultat observé + accessibilité testée.
- `FUNCTIONAL_NOT_ACCESSIBLE` : REFUSÉ.
- `NOT_FUNCTIONAL` : REFUSÉ.
- `NOT_PROVEN` : REFUSÉ.
- `BLOCKED` : REFUSÉ pour l'acceptation.

## Preuve minimale

Une unité ne peut entrer dans `FUNCTIONAL_ACCESSIBLE` que si :

1. elle possède une implémentation exécutable ;
2. son test est exécuté ;
3. son résultat est observé ;
4. son comportement déterministe est vérifié lorsqu'il est applicable ;
5. son sens est exposé par les projections d'accessibilité requises ;
6. les projections conservent exactement le même sens ;
7. l'échec d'une projection fait échouer la validation globale.

## Portée

Cette règle s'applique à toute fonctionnalité du dépôt :

- noyau ZERO ;
- langage ;
- format machine ;
- exécution ;
- objets ;
- événements ;
- capacités ;
- I/O ;
- communication ;
- réseau ;
- navigateur ;
- accessibilité ;
- boot ;
- machines distribuées ;
- radio ;
- satellite ;
- IA ;
- stockage ;
- projections ;
- tests et outils qui prétendent valider une fonctionnalité.

## Interdiction

Il est interdit de déclarer une fonctionnalité fonctionnelle uniquement parce que :

- elle est documentée ;
- elle est modélisée ;
- elle est compilable ;
- un test existe mais n'a pas été exécuté ;
- une simulation réussit ;
- QEMU réussit ;
- une projection visuelle existe ;
- une IA affirme que cela fonctionne.

## Gate global

Le CI doit refuser le dépôt tant qu'une unité requise reste sans preuve fonctionnelle et accessible.

Le passage du gate ne doit jamais être obtenu en abaissant le niveau d'exigence.

## Statut actuel

Le dépôt n'est pas encore globalement fonctionnel et accessible.

Cette situation est volontairement REFUSÉE par le gate jusqu'à implémentation et exécution des unités restantes.
