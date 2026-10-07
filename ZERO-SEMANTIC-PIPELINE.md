# ZERO — pipeline sémantique de référence

Le bootstrap Go fournit maintenant un pipeline déterministe complet au niveau
sémantique :

`DECODE → VALIDATE → CANONICALIZE → AUTHORIZE → EXECUTE → TRANSITION`

## Sortie

`EngineReceipt` conserve :

- représentation canonique ;
- décision d'autorisation ;
- résultat ;
- état précédent ;
- état suivant.

## Garanties

- une entrée invalide est rejetée avant exécution ;
- une demande non autorisée est rejetée avant exécution sémantique ;
- une identité ne peut pas faire évoluer l'état d'une autre identité ;
- les résultats UNKNOWN et CONTRADICTION ne créent pas de transition implicite ;
- une même entrée et un même état produisent le même état suivant ;
- la représentation canonique est conservée dans la sortie.

## Limite de preuve

Ceci constitue un produit de référence exécutable au niveau bootstrap Go.
Il ne constitue pas une implémentation CPU native, un hyperviseur, une
garantie de sécurité système ou une preuve matérielle.
