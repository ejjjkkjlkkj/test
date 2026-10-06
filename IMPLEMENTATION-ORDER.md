# ZERO — Ordre d'implémentation

La spécification doit maintenant descendre progressivement vers une machine exécutable.

## Étape 1 — Représentation minimale

Définir l'encodage physique minimal de :

- identité ;
- sens ;
- état ;
- relation ;
- capacité ;
- observation ;
- événement ;
- incertitude.

## Étape 2 — Transition

Définir comment le CPU effectue :

OBSERVE
READ
WRITE
STEP
REQUEST
NEXT_EVENT

sans dépendre d'un runtime général.

## Étape 3 — I/O machine

Brancher le modèle sémantique sur :

- CPU ;
- mémoire ;
- GPU ;
- moniteur ;
- audio ;
- clavier ;
- stockage ;
- réseau ;
- firmware.

## Étape 4 — Interaction

Construire un noyau commun pour clavier, parole, braille, affichage et autres modalités.

## Étape 5 — Réseau ZERO

Implémenter la découverte, l'identité, l'authentification, les capacités, les événements et les transitions distantes.

## Étape 6 — Contenu

Construire les objets texte, image, photo, graphique, tableau, vidéo, carte, formulaire, jeu, CAPTCHA et contenu inconnu.

## Étape 7 — Navigateur

Le navigateur ZERO doit être une application native du langage, pas un simple habillage d'un navigateur existant.

## Étape 8 — Auto-hébergement

Le langage doit progressivement pouvoir construire ses propres outils et son propre environnement.

## Règle

Aucune étape ne peut introduire une couche d'accessibilité séparée qui récupère après coup une information détruite.
