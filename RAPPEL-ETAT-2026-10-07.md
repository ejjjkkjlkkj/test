# RAPPEL-ETAT-2026-10-07

## Dépôt
- Repository: ejjjkkjlkkj/test
- Branche: main
- Objectif: dépôt complet propre, cohérent, fonctionnel et préparé pour release.
- Ollama / IA locale: retiré de main. Ne pas réintroduire.
- Ne pas supprimer ni réécrire l'historique sans demande explicite.

## Derniers travaux confirmés
- 73dd16d — renforcement de l'équivalence sémantique universelle.
- 5a9f1f5 — identité des projections d'accessibilité.
- ac9c423 — profils de portabilité de la vérité.
- 8cce8e7 — validation du format machine.

## Blocage actuel
Deux fichiers doivent être réparés avant toute déclaration PASS:
1. tests/zero-machine-format.ps1
   - la section finale autour de INVALID_PROOF_ACCEPTED est actuellement corrompue.
2. tests/zero-truth-portability.ps1
   - $modalities est déclaré après Project-Truth alors que Project-Truth l'utilise.
   - avec StrictMode, l'ordre doit être corrigé.

## Validation
- Les modifications récentes ont été écrites sur main.
- Aucun statut CI exploitable n'est actuellement retourné pour le dernier commit.
- Ne pas déclarer RELEASE/PASS tant que les deux scripts ne sont pas corrigés puis exécutés.
- Après correction: exécuter toute la chaîne de tests définie par .github/workflows/zero-validation.yml et vérifier les résultats réels.
- Vérifier également le release-gate avant toute release.

## Prochaine séquence
1. Corriger zero-machine-format.ps1.
2. Corriger l'ordre de déclaration de zero-truth-portability.ps1.
3. Relire les deux fichiers après écriture.
4. Rechercher TODO/FIXME/panic(/placeholder et références Ollama/LOCAL-AI.
5. Vérifier workflows et version.
6. Exécuter la validation complète.
7. Seulement après PASS réel: préparer la release.

## Règle
Ce fichier est un rappel d'état, pas une preuve de validation. Toute validation doit être fondée sur une exécution réelle.
