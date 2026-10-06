# ZERO — Accessibility Kernel

Ce fichier définit le noyau minimal qui rend l'accessibilité native au langage.

## Le noyau

L'accessibilité native n'est pas un service parallèle.

Le noyau garantit six propriétés à chaque objet :

1. **Existence** — l'objet peut être distingué de son absence.
2. **Compréhension** — ce qui est connu de l'objet possède une représentation sémantique.
3. **Navigation** — l'objet peut être atteint par des relations.
4. **Observation** — son état peut être consulté selon les capacités autorisées.
5. **Action** — ses actions sûres sont représentées.
6. **Retour** — toute action produit un résultat ou un refus observable.

## Création d'objet

Créer un objet incomplet est autorisé.

Créer un objet sans identité ou sans possibilité d'observation n'est pas une forme normale du noyau.

## Projection

Une projection transforme le sens en une modalité.

`SEMANTIC OBJECT -> PROJECTION -> MODALITY`

La projection ne devient jamais la source de vérité.

## Exemple

Un graphique n'est pas :

`pixels`

Il est d'abord :

`graph + axes + series + values + relations + trends + state`

Les pixels, la parole, le braille ou une autre sortie sont des projections de cette même structure.

## Vérification

Pour tout objet `O`, le noyau doit pouvoir demander :

`EXISTS(O)`

`INSPECT(O)`

`NAVIGATE(O)`

`OBSERVE(O)`

`ACTIONS(O)`

`RESULT(O)`

sans dépendre d'une interface visuelle.

## Interdiction architecturale

Un composant externe ne peut pas être requis pour transformer après coup un objet opaque en objet accessible.

Cette transformation doit être impossible comme architecture normale du noyau.
