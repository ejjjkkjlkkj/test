# ZERO — ÉTAT DE REPRISE

Date de mise à jour: 2026-10-07
Dépôt: ejjjkkjlkkj/test
Branche de référence: main
HEAD de référence: 6a12c69a50dce25349c28f41d7b2eb70d172967b

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
- suppression du composant produit local Ollama/AI du main release-bound.

## Frontière de preuve
Le PASS actuel concerne le bootstrap logiciel reproductible.
NON PROUVÉ: CPU natif, noyau/hyperviseur, isolation matérielle, sécurité physique, radio et fonctionnement matériel sans bootstrap.

## Règle de reprise
Toute nouvelle fonctionnalité doit avoir une implémentation identifiable, un test exécutable, une preuve CI correspondante lorsque applicable et une frontière de preuve explicite.
