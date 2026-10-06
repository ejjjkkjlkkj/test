# ZERO — Conformité accessibilité native

Un composant ZERO est conforme seulement s'il respecte le noyau.

## Tests obligatoires

### Existence
L'objet peut-il être découvert ?

### Identité
Peut-il être distingué d'un autre objet ?

### Sens
Une information sémantique existe-t-elle indépendamment de son affichage ?

### Navigation
Peut-on atteindre l'objet sans imposer une modalité particulière ?

### Observation
Son état est-il observable ?

### Action
Les actions autorisées sont-elles exposées ?

### Résultat
Succès, refus, échec, interruption et état partiel sont-ils observables ?

### Multimodalité
La même intention peut-elle être exprimée par plusieurs modalités ?

### Inconnu
Un objet partiellement compris reste-t-il utilisable et inspectable ?

### Sécurité
L'accessibilité évite-t-elle d'accorder un droit supplémentaire ?

## Échec

Un composant qui dépend d'une couche externe pour rendre son sens accessible ne satisfait pas le noyau ZERO.
