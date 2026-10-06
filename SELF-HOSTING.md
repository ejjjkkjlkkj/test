# ZERO — Auto-hébergement

L'objectif final est que ZERO puisse construire son propre environnement.

## Progression

ZERO spécification
→ représentation minimale
→ exécuteur minimal
→ pont machine
→ noyau ZERO
→ outils ZERO
→ compilateur / transformateur ZERO
→ environnement ZERO
→ applications ZERO

Un outil temporaire externe peut servir à fabriquer une première réalisation.

Il ne devient pas une dépendance conceptuelle du langage.

## Critère

ZERO est réellement auto-hébergé lorsque son environnement essentiel peut être défini, inspecté et construit depuis ZERO lui-même.

## Accessibilité

Les outils de construction doivent eux-mêmes respecter les invariants.

Le compilateur, le débogueur, le chargeur, le navigateur et les outils système ne deviennent jamais des exceptions d'accessibilité.
