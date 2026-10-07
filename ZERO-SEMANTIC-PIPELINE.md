# ZERO — pipeline sémantique de référence

Le bootstrap Go fournit un pipeline déterministe complet au niveau sémantique :

`DECODE → VALIDATE → CANONICALIZE → AUTHORIZE → EXECUTE → TRANSITION`

## Entrées

Deux représentations convergent vers le même pipeline :

- `ExecuteRecord` : RECORD canonique ;
- `ExecuteNativeRecord` : RECORD binaire `ZER0`.

La représentation native est donc une matérialisation de données et non une seconde sémantique.

## Sortie

`EngineReceipt` conserve :

- représentation canonique ;
- décision d'autorisation sémantique explicite ;
- résultat ;
- état précédent ;
- état suivant ;
- indicateur explicite de transition.

## Garanties

- une entrée invalide est rejetée avant exécution ;
- une demande non autorisée est rejetée avant exécution sémantique ;
- une identité ne peut pas faire évoluer l'état d'une autre identité ;
- les résultats UNKNOWN et CONTRADICTION ne créent pas de transition implicite ;
- une contradiction déterministe est rapportée sans modifier l'état ;
- les opérations ISA et les opérations sémantiques utilisent des politiques d'autorisation distinctes ;
- une même entrée et un même état produisent le même état suivant ;
- la représentation canonique est conservée dans la sortie ;
- l'entrée native et l'entrée canonique convergent vers la même logique d'exécution.

## Limite de preuve

Ceci constitue un produit de référence exécutable au niveau bootstrap Go.

Il ne constitue pas une implémentation CPU native, un hyperviseur, une garantie de sécurité système ou une preuve matérielle.
