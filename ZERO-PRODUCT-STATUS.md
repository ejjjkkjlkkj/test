# ZERO — état du produit bootstrap

## État

Le dépôt contient un produit de référence exécutable en Go couvrant :

- représentation canonique déterministe ;
- représentation binaire native de référence ;
- ISA minimale ;
- mémoire bornée ;
- compteur de programme déterministe ;
- autorisation ISA ;
- autorisation sémantique explicite ;
- événements d'exécution ;
- état sémantique et transitions ;
- gestion explicite des contradictions ;
- pipeline complet DECODE → VALIDATE → CANONICALIZE → AUTHORIZE → EXECUTE → TRANSITION ;
- CLI stdin/stdout produisant des reçus JSON ;
- tests unitaires, déterministes et fuzz de non-panique pour les décodeurs.

## Frontière de preuve

PASS signifie ici : comportement reproductible du bootstrap logiciel testé par CI.

NON PROUVÉ :

- exécution native CPU ;
- privilèges noyau/hyperviseur ;
- isolation matérielle ;
- sécurité physique ;
- communication radio ;
- fonctionnement matériel sans logiciel de bootstrap.

Aucune de ces propriétés n'est déduite des tests Go.

## Règle de reproductibilité

Une même entrée canonique, une même opération autorisée et un même état initial
doivent produire le même reçu et le même état suivant.

Les contradictions sont des résultats déterministes non-transitionnels.
Les opérations inconnues sont refusées avant exécution sémantique.
