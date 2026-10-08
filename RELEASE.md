# ZERO — procédure de release

## Version courante

La version candidate déclarée par `VERSION` est `0.1.0-rc.1`.

## Préconditions

Une publication de release n'est autorisée que si :

1. `VERSION` contient une version non vide.
2. La CI de `main` passe complètement.
3. Les tests bootstrap et core passent.
4. La validation native ZERO passe.
5. Tous les validateurs PowerShell passent.
6. Le contrôle de propreté des composants locaux passe.
7. Le tag de release correspond exactement à `v<VERSION>`.
8. Les preuves restent limitées à leur niveau réellement démontré.

## Publication

Le workflow `.github/workflows/release-gate.yml` est déclenché par un tag `v*` ou manuellement.

Pour un tag, le workflow refuse toute divergence entre le tag et `VERSION`.

La création effective d'une release GitHub reste une opération distincte de la validation CI.

## Frontière de preuve

Un PASS logiciel ne constitue pas une preuve automatique de fonctionnement CPU natif, matériel, RF, radio, satellite ou liaison physique réelle.

Aucune release ne doit présenter une preuve logicielle comme une preuve matérielle.

## Règle de conservation

Les branches expérimentales et historiques ne sont pas fusionnées uniquement pour réduire le nombre de branches. Une intégration doit apporter une fonctionnalité vérifiable sans dégrader `main`.
