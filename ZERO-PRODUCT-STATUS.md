# ZERO — état du produit bootstrap

## État

Le dépôt contient un produit de référence exécutable en Go couvrant :

- représentation canonique déterministe ;
- représentation binaire native de référence ;
- décodage natif vers le même RECORD sémantique ;
- ISA minimale ;
- mémoire bornée ;
- compteur de programme déterministe ;
- autorisation ISA ;
- autorisation sémantique explicite ;
- événements d'exécution ;
- état sémantique et transitions ;
- gestion explicite des contradictions ;
- pipeline complet DECODE → VALIDATE → CANONICALIZE → AUTHORIZE → EXECUTE → TRANSITION ;
- exécution depuis RECORD canonique et depuis RECORD natif ;
- CLI stdin/stdout produisant des reçus JSON ;
- tests unitaires, déterministes et fuzz de non-panique pour les décodeurs.

## Validation CI

Le HEAD de la branche de travail est validé par :

- ZERO validation = PASS ;
- ZERO bootstrap validation = PASS.

Les échecs précédents du bootstrap provenaient de l'utilisation de l'autorisation ISA pour des opérations sémantiques. Le pipeline utilise maintenant explicitement `AuthorizeSemantic`.

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

Une même entrée canonique, une même opération autorisée et un même état initial doivent produire le même reçu et le même état suivant.

Les contradictions sont des résultats déterministes non-transitionnels.
Les opérations inconnues sont refusées avant exécution sémantique.

## Prochaine limite technique

La prochaine extension doit renforcer la conformance et les vecteurs d'exécution autour de cette représentation native, sans transformer le bootstrap Go en prétendue preuve CPU ou matérielle.
