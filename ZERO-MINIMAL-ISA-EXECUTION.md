# ZERO — exécution machine minimale

## Statut

ISA MINIMALE = IMPLEMENTED (Go)
MEMOIRE CONTROLEE = IMPLEMENTED (Go)
EXECUTION DETERMINISTE = TESTED (Go)
CPU REEL = NOT_IMPLEMENTED
HARDWARE = NOT_PROVEN

## Instruction set initial

Le premier noyau exécutable expose exactement six opérations :

- OBSERVE
- READ
- WRITE
- STEP
- REQUEST
- NEXT_EVENT

Une instruction inconnue est rejetée. Le contenu de PAYLOAD n'est jamais exécuté comme code.

## Mémoire

La mémoire est un tableau borné contrôlé par le moteur.

READ et WRITE vérifient systématiquement les bornes.
Une lecture retourne une copie.
Une écriture ne modifie que la plage demandée.

## État

PC est un compteur monotone.
Chaque instruction acceptée avance PC exactement de 1.
Une instruction rejetée ne modifie pas PC.

## Preuve

Les tests vérifient :

1. lecture/écriture ;
2. rejet des dépassements ;
3. déterminisme du PC ;
4. rejet d'une instruction inconnue ;
5. absence d'exécution arbitraire du payload.

Cette étape constitue une exécution dans le bootstrap Go.
Elle ne constitue pas encore une exécution CPU native ni une preuve matérielle.
