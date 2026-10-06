# ZERO — Accessibilité réseau native

L'accessibilité doit fonctionner lorsqu'un objet vient d'une autre machine.

## Principe

LOCAL OBJECT
et
REMOTE OBJECT

utilisent le même noyau sémantique.

La distance réseau ne doit pas créer une nouvelle catégorie d'objet inaccessible.

## Projection

REMOTE OBJECT
-> SEMANTIC OBJECT
-> UNIVERSAL NAVIGATION
-> VOICE / BRAILLE / KEYBOARD / VISUAL / TOUCH / AUTOMATION

L'utilisateur ne devrait pas avoir à connaître le protocole réseau pour comprendre l'objet.

## Exemple

Un téléphone expose un événement autorisé.

Le chemin est :

iPhone
-> réseau
-> événement distant
-> observation ZERO
-> objet sémantique
-> navigation
-> parole ou braille ou autre modalité

La couche réseau transporte l'information.

Elle ne définit pas son accessibilité.

## Limite importante

ZERO ne promet pas l'accès à des données ou fonctions que l'appareil distant ne rend pas disponibles.

L'accessibilité rend disponible ce qui est sémantiquement observable et autorisé.

Elle ne contourne ni authentification, ni permission, ni sécurité.
