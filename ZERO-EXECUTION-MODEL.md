# ZERO — modèle d'exécution

ZERO ne doit pas commencer par un compilateur classique.

Le premier noyau exécutable est un moteur capable de prendre une représentation ZERO, de créer des objets sémantiques, de faire circuler des événements et de produire des observations déterministes.

## Cycle

DECODE → VALIDATE → RESOLVE → EXECUTE → EMIT → OBSERVE → PROVE

## Objet d'exécution

Chaque opération possède :

- operation_id
- actor
- target
- intent
- inputs
- authorization
- state
- observations
- events
- result
- proof

## Règle

Une opération sans autorisation nécessaire ne peut pas passer de OBSERVE à ACT.

Une opération inconnue ne doit jamais être interprétée silencieusement comme une opération sûre.

## Déterminisme

À entrée identique, le noyau doit produire une structure d'état équivalente, indépendamment de la modalité utilisée pour demander l'action.
