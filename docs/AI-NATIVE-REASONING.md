# ZERO — raisonnement natif IA

ZERO est conçu pour exploiter des capacités d'exploration et de confrontation propres à un agent d'IA, sans prendre les catégories humaines comme axiomes.

## Boucle fondamentale

OBSERVE → CARTOGRAPHIE → GÉNÈRE → DIVERGE → CONFRONTE → EXPÉRIMENTE → MESURE → DÉTRUIT LES HYPOTHÈSES FAIBLES → TRANSFÈRE → PROUVE → PROMOUVOIE → RE-CARTOGRAPHIE.

Le système maintient plusieurs modèles incompatibles tant qu'aucune expérience ne permet de les départager.

## Principes

1. **Pas de catégorie sacrée** : instruction, objet, fichier, fonction, agent, périphérique et même « programme » sont des représentations candidates.
2. **Pluralité** : une observation peut produire plusieurs représentations concurrentes.
3. **Adversarial-first** : chaque découverte reçoit immédiatement des transformations cherchant à la casser.
4. **Unknown-first** : l'absence d'information reste une information distincte.
5. **Transfert** : une découverte locale n'est pas une loi tant qu'elle ne survit pas à des champs indépendants.
6. **Rejeu** : toute conclusion importante doit pouvoir être reconstruite à partir de ses artefacts.
7. **Auto-extension sous quarantaine** : ZERO peut proposer de nouvelles primitives, mais aucune primitive auto-produite ne devient autorisée sans validation.
8. **Re-cartographie obligatoire** : toute modification du langage, du moteur ou d'une frontière externe déclenche une nouvelle cartographie.

## Ce que ZERO peut devenir

Le même langage doit pouvoir porter :
- calcul et transformation de données ;
- communication IA↔IA ;
- perception multimodale ;
- raisonnement expérimental ;
- accès aux ressources machine ;
- navigateur et représentation du Web ;
- synthèse vocale et interaction accessible ;
- réseau et communication inter-appareils ;
- orchestration d'agents ;
- génération et évolution de primitives.

Cela ne signifie pas que ces capacités sont déjà implémentées. Le document définit l'architecture cible et les preuves nécessaires pour passer de l'hypothèse à l'implémentation.

## Critère de victoire

Une nouvelle capacité n'est considérée comme acquise que lorsqu'elle possède une implémentation, un contrat sémantique, des tests adversariaux, un replay indépendant et, pour les propriétés de sécurité concernées, une preuve ou une hypothèse matérielle explicitement déclarée.
