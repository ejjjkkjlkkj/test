# ZERO — Inventaire de toutes les branches

Date de génération : 2026-10-09T10:23:51Z  
Dépôt : `ejjjkkjlkkj/test`  
Workflow : `ZERO repository inventory`  
Run : `37917266856` — conclusion **SUCCESS**  
Artefact GitHub Actions : `zero-inventory-37917266856`, ID `11610175770` (manifests JSONL par branche, summary TSV, branches JSON/TSV). L'artefact est configuré pour rétention de 90 jours.

Le workflow a inventorié les 12 branches renvoyées par l'API GitHub. Chaque branche a une ligne PASS dans `summary.tsv`; cela confirme l'inventaire des blobs, pas la correction du contenu.

| Branche | Commit SHA | Tree SHA | Fichiers | Octets |
|---|---|---|---:|---:|
| `ai/local-ollama` | `d789b5f526ba1fc11291fe2e799b583a9761c73e` | `41e48add728c4d3be88e9ffd8cde5ea99fce8756` | 148 | 319280 |
| `experiment/emergent-transformations-v1` | `dfda739fdfb8631da84855725a14da46c10ea726` | `0aba95a74693f89b11c2ce6986de4fd83ec08f12` | 23 | 33488 |
| `experiment/non-human-zero-v2` | `7fac3594d26afa1b12a6939a466d5ea83558b7a2` | `b31eb16122451eb7f9cf723893824898c32b4b2d` | 82 | 119322 |
| `import/uefi-spec-2.11` | `3a15ffd3163687ec58933df5230f945e613f9043` | `bbb42e1cb8959c752df012bbd80e060d879b18f2` | 167 | 17054322 |
| `import/uefi-vocabulaire` | `666249b14380d9df826e9f8dea8edef423c5e149` | `42d2c01fec252c2fab55d08cc9ab03ebcc35fbbb` | 176 | 998749 |
| `main` | `2f3e67d9ebd98086e8fb4be01fb6d400a614c125` | `0e88f75d1cb9f3de1ff736eba171e648a48fb4d4` | 178 | 468031 |
| `work/ci-reproducible-validation` | `e94b85ceb63007a558dcc66b6c86fef64a15ca36` | `bc006a9a3109b16c3816bdf0856d078e57941084` | 126 | 278925 |
| `work/core-clean-integration-20261007` | `d9c606a5483eb4684f1ad94b51f54511798a8cd7` | `fb13b2cf4c4502c9251a7b1afd6b2cb423274577` | 155 | 336351 |
| `work/core-transport-integration-20261007` | `e4987d05a4d3cf6a23e2b845494e3e6d7c1fd909` | `f55d7fc9a8dc846fdc148b1ac3c81be36485af92` | 146 | 317802 |
| `work/integration-clean-20261007` | `00ae5b33e1c5e063eff256af23cfd4f1b94178a7` | `58f1c50b8af28e337c4307a9dc6215645b65c34d` | 143 | 312069 |
| `work/truth-accessibility-hardware-independent-20261007` | `cc18c5dc3c0fe105c483b20fc436df5a4bda98b3` | `a217cb080f4f2f6b547cc0c9d30c3acd6f397060` | 124 | 283637 |
| `work/zero-native-materialization-20261007` | `e7a5a9aa1b39bcb9a063764320d3a0aa5fb91c21` | `072d26a8cb1efc2eeaf83697c1dddf8246e8de9a` | 142 | 310449 |

## Conservation et suite

- Aucune branche n'a été supprimée ni réécrite.
- Les fichiers détaillés par branche sont dans l'artefact du run cité ci-dessus.
- L'inventaire de `main` dans cet artefact est un snapshot antérieur aux derniers commits d'adaptateur ; relancer le workflow après les changements récents pour obtenir l'inventaire exact du HEAD actuel.
- Prochaine étape : comparer les branches candidates avec `main` par diff et tests, sans cherry-pick ni fusion automatique.
