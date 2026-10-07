# ZERO — autorisation et émission déterministes

Cette étape ajoute une frontière explicite entre une demande sémantique et
l'exécution du bootstrap.

## Modèle

1. `ExecutionRequest` identifie l'opération, l'acteur et la cible.
2. `Authorize` décide avant toute mutation.
3. Une demande refusée ne modifie ni la mémoire ni le PC.
4. `ExecuteAuthorized` appelle ensuite l'exécuteur ISA existant.
5. Un `Event` déterministe expose l'opération et la transition PC observée.

## Règles validées

- acteur obligatoire ;
- cible obligatoire ;
- instruction inconnue refusée ;
- autorisation identique pour une entrée identique ;
- aucun effet mémoire ni avance du PC après refus ;
- événement cohérent avec le résultat d'exécution.

Cette implémentation est un modèle de référence Go. Elle ne constitue pas une
barrière de sécurité système, un hyperviseur, un mécanisme CPU natif ou une
preuve matérielle.
