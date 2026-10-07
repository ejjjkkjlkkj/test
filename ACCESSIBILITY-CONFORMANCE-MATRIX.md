# ZERO — matrice de conformité accessibilité

Cette matrice vérifie que l'accessibilité est native au même modèle que le reste du système.

| Domaine | Exigence | Critère PASS |
|---|---|---|
| Existence | découverte | objet observable |
| Identité | distinction | identité stable |
| Sens | signification | sens indépendant de l'affichage |
| Structure | organisation | relations observables |
| État | état courant | état lisible |
| Navigation | accès | aucune modalité imposée |
| Action | opérations | capacités et autorisations explicites |
| Résultat | retour | succès/refus/échec/interruption observables |
| Événement | changement | événement conservé |
| Inconnu | information partielle | UNKNOWN conservé |
| Sécurité | droits | accessibilité sans privilège supplémentaire |
| Voix | projection | même source sémantique |
| Braille | projection | même source sémantique |
| Clavier | interaction | même intention |
| Affichage | projection | même source sémantique |
| Tactile | interaction | même intention |
| Réseau | projection | même objet sémantique |
| Automatisation | interaction | même opération |
| Multimodalité | équivalence | aucune modalité privilégiée |

## Règle PASS

Un composant n'est PASS que si le comportement a été effectivement exécuté et observé.

La présence d'une spécification ne suffit pas.

## États

DEFINED
IMPLEMENTED
TESTED
PASS
FAIL
BLOCKED
NOT_PROVEN

## Interdictions

ACCESSIBLE ≠ uniquement vocal

ACCESSIBLE ≠ uniquement graphique

ACCESSIBLE ≠ API externe

ACCESSIBLE ≠ OCR ajouté après coup

ACCESSIBLE ≠ droit supplémentaire

ACCESSIBLE ≠ preuve matérielle

## Statut

MATRICE : DEFINED
EXECUTION : NOT_EXECUTED
