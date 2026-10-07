# ZERO — Machine-First Universal Accessible Language

ZERO est un langage-machine : il ne sépare pas artificiellement langage, machine, communication et accessibilité.

## Monde complet

CPU
↔ mémoire
↔ GPU
↔ moniteur
↔ audio
↔ entrées
↔ stockage
↔ firmware
↔ réseau
↔ autres machines
↔ applications
↔ contenu
↔ utilisateur

## Même modèle partout

Chaque entité peut être représentée par :

IDENTITY
MEANING
STATE
RELATIONS
CAPABILITIES
OBSERVATIONS
EVENTS
PROJECTIONS
SECURITY
UNCERTAINTY

L'accessibilité est une propriété intrinsèque de ce modèle.

## Communication

ZERO communique avec la machine locale et avec les machines distantes.

PC ↔ iPhone ↔ PC ↔ serveur ↔ périphérique ↔ machine ZERO.

La distance ne change pas le modèle.

## Interaction

PHYSIQUE
→ OBSERVATION
→ EVENT
→ INTENT
→ TRANSITION
→ RESULT
→ PROJECTION
→ PHYSIQUE

Voix, braille, clavier, affichage, tactile, réseau et futures modalités convergent vers le même noyau.

## Contenu

Le même système doit pouvoir représenter et rendre navigables :

texte, image, photo, graphique, tableau, vidéo, audio, carte, document, application, web, formulaire, jeu, CAPTCHA et objets inconnus.

## Objectif

Construire progressivement une machine et un langage qui ne perdent pas le sens entre le matériel et l'utilisateur.

Le projet ne consiste donc pas à ajouter une accessibilité à un langage.

Il consiste à rendre **le langage lui-même, la machine elle-même et leurs communications intrinsèquement accessibles**.

## Release candidate

Version courante : `0.1.0-rc.1`.

Le `main` release-bound est soumis aux tests bootstrap/core, à la validation native ZERO et aux validateurs de format, vérité, exécution universelle, machine, moteur de référence, support physique, portabilité de vérité et accessibilité. La validation effective dépend de l'exécution de la CI.

Le workflow `.github/workflows/release-gate.yml` rejoue ces contrôles lorsqu'un tag `v*` est publié. Le tag doit correspondre exactement au contenu de `VERSION`.

Les preuves logicielles ne constituent pas automatiquement des preuves matérielles, RF, radio, satellite ou de liaison physique réelle.
