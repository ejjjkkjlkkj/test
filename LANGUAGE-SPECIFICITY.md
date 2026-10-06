# ZERO — Spécificité du langage machine-accessibilité

ZERO n'est pas un langage généraliste auquel on ajoute l'accessibilité.

Sa différence fondamentale est que **machine, langage et accessibilité sont une seule architecture**.

## 1. Principe fondateur

Dans un langage classique :

`machine -> runtime -> programme -> interface -> accessibilité`

Dans ZERO :

`machine -> sens -> langage -> objet -> interaction`

L'accessibilité n'arrive jamais après la création de l'objet.

Elle est une propriété de l'objet dès sa naissance.

## 2. Le sens précède la représentation

ZERO ne définit pas d'abord :

- pixels;
- texte affiché;
- son;
- coordonnées;
- boutons graphiques;
- touches;
- fichiers;
- widgets.

Il définit d'abord :

- identité;
- existence;
- structure;
- état;
- relations;
- signification;
- capacités;
- événements;
- observations;
- actions sûres.

Les représentations physiques sont des projections de ce modèle.

## 3. Un objet ZERO est nativement accessible

Tout objet possède au minimum :

`IDENTITY + MEANING + STRUCTURE + STATE + RELATIONS + CAPABILITIES + EVENTS + OBSERVATIONS + PROJECTIONS`

Il n'existe donc pas d'objet « normal » auquel il faudrait ensuite ajouter une version accessible.

## 4. L'interaction est indépendante de la modalité

Une action possède une identité sémantique.

Exemple :

`ACTIVATE(object)`

peut être produite par :

- clavier;
- parole;
- braille;
- pointeur;
- écran tactile;
- automatisation;
- une future modalité inconnue.

Ces entrées ne créent pas des actions différentes.

Elles expriment la même intention.

## 5. La machine elle-même est accessible

ZERO ne commence pas à la fenêtre ou à l'application.

Le modèle inclut :

- processeur;
- mémoire;
- périphériques;
- bus;
- stockage;
- affichage;
- audio;
- entrées;
- firmware;
- démarrage;
- énergie;
- sécurité;
- erreurs;
- périphériques inconnus.

Un périphérique inconnu reste un objet observable avec une identité, un état, des relations et une incertitude explicite.

## 6. Accessibilité sans interprétation tardive

Une couche inférieure ne peut pas produire volontairement une représentation opaque si elle possède déjà l'information sémantique correspondante.

Règle :

`information physique -> information sémantique -> projection`

et jamais :

`information physique -> représentation opaque -> tentative de récupération du sens`

Cette règle est centrale pour les images, graphiques, interfaces dynamiques, documents, jeux, CAPTCHA et objets futurs.

## 7. L'inconnu est un état normal

ZERO ne définit pas « connu » et « inaccessible » comme synonymes.

Un objet peut être :

`KNOWN`

ou

`PARTIALLY-KNOWN`

ou

`UNKNOWN`

mais dans les trois cas il reste représenté.

L'inconnu expose au minimum :

- existence;
- frontière;
- identité;
- observations disponibles;
- relations observables;
- actions sûres;
- incertitude.

Le système peut ensuite apprendre ou découvrir davantage sans changer l'identité de l'objet.

## 8. Accessibilité et sécurité sont compatibles

ZERO ne donne pas automatiquement tous les droits à l'utilisateur.

Il distingue :

`OBSERVE`

de

`ACT`

et :

`ACT`

de

`AUTHORIZE`.

Une information refusée pour des raisons de sécurité reste représentée comme refusée, avec sa raison sémantique lorsque celle-ci peut être exposée.

## 9. Interruption native

Toute interaction peut être interrompue sans perdre son identité sémantique.

L'interruption est un événement :

`RUNNING -> INTERRUPTED`

et non une disparition silencieuse.

Cela rend native l'interruption de la parole, de la navigation, d'une opération longue, d'un média ou d'une action machine.

## 10. Accessibilité des objets futurs

Le langage ne doit pas avoir besoin d'une nouvelle API pour chaque nouveau type d'objet.

Un nouvel objet doit pouvoir être construit à partir du même noyau :

`identity + meaning + structure + state + relation + capability + event + observation`

Sa projection peut être nouvelle sans modifier le principe du langage.

## 11. Propriété la plus importante

La propriété recherchée n'est pas :

**« tout est décrit pour l'utilisateur »**.

Elle est :

**« tout ce que le système représente possède déjà une voie sémantique d'observation et d'action équivalente »**.

C'est cette propriété qui fait de ZERO un langage machine-accessibilité, et non un langage avec un système d'accessibilité.

## 12. Test de spécificité

Une conception n'appartient pas au cœur ZERO si elle exige :

- un mode accessibilité séparé;
- un écran lecteur obligatoire;
- une API d'accessibilité externe pour récupérer le sens;
- une hiérarchie visuelle comme source de vérité;
- une modalité privilégiée;
- une conversion tardive d'une représentation opaque;
- une connaissance préalable de tous les types d'objets.

Si elle échoue à l'un de ces tests, elle doit être déplacée hors du noyau ou redessinée.

## 13. Formule

La spécificité de ZERO peut être résumée ainsi :

`MACHINE = SEMANTICS = LANGUAGE = ACCESSIBLE INTERACTION`

Ce n'est pas une égalité physique.

C'est une règle d'architecture : aucune de ces dimensions ne peut être conçue indépendamment des autres.
