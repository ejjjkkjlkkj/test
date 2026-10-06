# ZERO — Semantic I/O

Les entrées-sorties physiques ne sont pas directement l'interface du langage.

## Entrée

Le chemin normal est :

`PHYSICAL -> OBSERVATION -> EVENT -> INTENT -> TRANSITION`

Une touche n'est donc pas directement une commande.

Un microphone n'est pas directement une application.

Un capteur n'est pas directement une valeur brute.

Le noyau conserve d'abord ce qui est observé et son contexte.

## Sortie

Le chemin normal est :

`STATE -> SEMANTIC RESULT -> PROJECTION -> PHYSICAL OUTPUT`

Une sortie audio n'est donc pas la signification.

Une image n'est pas la signification.

Le braille n'est pas la signification.

Ce sont des projections.

## Équivalence

Plusieurs modalités peuvent converger vers la même intention.

Plusieurs modalités peuvent également projeter le même résultat.

Cela permet une véritable équivalence d'interaction sans imposer une interface principale.

## Périphériques inconnus

Un périphérique dont le protocole n'est pas encore compris produit quand même une observation native :

`EXISTENCE + IDENTITY? + BOUNDARY? + RAW OBSERVATION + UNCERTAINTY`

Le système ne transforme pas automatiquement l'inconnu en « inaccessible ».
