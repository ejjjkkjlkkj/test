# ZERO — ÉTAT DE REPRISE

Date de mise à jour: 2026-10-08
Dépôt: ejjjkkjlkkj/test
Branche de référence: main
HEAD de référence: c329857826e426b390712df7db6cd4c6ac7f1d5c

## But
Construire ZERO depuis zéro: machine, langage, communication, monde et accessibilité native sont une seule architecture sémantique.

## Contraintes
- Les langages classiques servent uniquement au bootstrap/validation; ils ne sont pas la fondation conceptuelle de ZERO.
- Pas de DOM, API d'accessibilité ou représentation graphique comme source de vérité.
- Aucune modalité privilégiée.
- Accessibilité et sécurité restent distinctes.
- Les objets inconnus restent représentés et inspectables.
- Erreurs, refus, interruptions et incertitudes sont sémantiques.
- Aucune preuve matérielle n'est déduite des tests bootstrap.
- Aucun composant IA local n'appartient au main release-bound.

## État validé
- représentation canonique RECORD déterministe;
- représentation binaire native ZERO avec encodeur/décodeur Go;
- ISA minimale et mémoire bornée;
- autorisation ISA et autorisation sémantique distinctes;
- événements et transitions déterministes;
- exécution canonique et native via le même pipeline;
- CLI bootstrap stdin/stdout;
- tests unitaires, déterministes et fuzz de non-panique;
- CI reproductible sur push, pull request vers main et déclenchement manuel;
- contrôle de propreté des composants locaux intégré à la validation.

## État de validation au 2026-10-08
- Bootstrap CI: PASS.
- zero-truth-portability.ps1: PASS.
- zero-machine-format.ps1: corrigé et exécuté dans la CI.
- Validation complète: actuellement bloquée uniquement par une référence textuelle interdite dans le rappel d'état.
- Release: NON VALIDÉE.

## Frontière de preuve
Le PASS concerne le bootstrap logiciel reproductible.
NON PROUVÉ: CPU natif, noyau/hyperviseur, isolation matérielle, sécurité physique, radio et fonctionnement matériel sans bootstrap.

## Règle de reprise
Toute nouvelle fonctionnalité doit avoir une implémentation identifiable, un test exécutable, une preuve CI correspondante lorsque applicable et une frontière de preuve explicite.
