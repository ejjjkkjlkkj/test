# Inventaire récursif des arbres Git — 30 dépôts

- Date de capture : 2026-10-09.
- Méthode : GitHub Git Trees API avec `recursive=1`, branches par défaut visibles au moment de la capture.
- Statut des arbres listés : `truncated=false` pour chaque résultat récupéré.
- Cet inventaire confirme les chemins et blobs dans chaque snapshot ; il ne signifie pas que le contenu de chaque fichier a été lu ou audité.
- Les sous-modules Git doivent encore être vérifiés par rapport aux pointeurs et commits épinglés ; les arbres de leurs dépôts séparés sont inventoriés indépendamment.

Fichiers suivis inventoriés : **7823** sur les 30 branches capturées.

| Dépôt | Branche capturée | Fichiers | Tree SHA | Arbre tronqué |
|---|---|---:|---|---|
| `NVDA-UPSTREAM-COMPLETE` | `master` | 1472 | `dd2bccb30e4faad4bbdbfcb811a6c3a4acb2bc83` | NON |
| `NVDA-SUBMODULE-.vscode` | `nvda-pinned` | 6 | `22faeb775256e4e8fa9ba2d650da1666489ceaaf` | NON |
| `NVDA-SUBMODULE-include-cppjieba` | `master` | 75 | `8f171de5018e8478ff22ca58caacf579cba809c8` | NON |
| `NVDA-SUBMODULE-include-cppjieba-deps-limonp` | `ci-windows-2022` | 56 | `f8b5b96ef0e5ac416facef2a1891fbd0f4324cfe` | NON |
| `NVDA-SUBMODULE-include-detours` | `4.0.1` | 158 | `e5400b4ec59478cb0f435cf3b1338226bcbe28f6` | NON |
| `NVDA-SUBMODULE-include-espeak` | `android` | 2006 | `edf2aad65c9c7ec0f832964df990db2b50f4bb9e` | NON |
| `NVDA-SUBMODULE-include-ia2` | `mark` | 44 | `88d08dee9bc5f1fb7e7928c53d1c59c8cc27b580` | NON |
| `NVDA-SUBMODULE-include-javaAccessBridge32` | `jab64` | 4 | `541e1a81e49b1318f511d66520c5f881285f97e1` | NON |
| `NVDA-SUBMODULE-include-liblouis` | `archive/zstanecic_tables` | 489 | `e2d9f96f5c57fdf9dc6a72fcb7d2f1b410fd7f81` | NON |
| `NVDA-SUBMODULE-include-nsis` | `add-Process-docs` | 289 | `3c0f9cb59e5e6640935f1cde8c1fee0ded5b5b1d` | NON |
| `NVDA-SUBMODULE-include-nvda-cldr` | `main` | 6 | `c888fe0f855899ba9147bf0fca8dc22ab7334472` | NON |
| `NVDA-SUBMODULE-include-nvda-mathcat` | `main` | 154 | `6448f67b00d992c69a8eb27eb9c627493e7dca57` | NON |
| `NVDA-SUBMODULE-include-sonic` | `master` | 60 | `b93885dcb70aae50c6f76b0fe4e0868f029a077e` | NON |
| `NVDA-SUBMODULE-include-sonic-speedy` | `main` | 29 | `e05c9b6fa473881e9c2b5ddaa323435b5513012c` | NON |
| `NVDA-SUBMODULE-include-w3c-aria-practices` | `2021-11_Note` | 418 | `c83f7620b97478109438be4d81765f9ab0ed9957` | NON |
| `NVDA-SUBMODULE-include-wil` | `copilot/port-crash-resistant-event-handler` | 125 | `eb35e667d2a1fbb26d1efb25815b3397991abe4d` | NON |
| `NVDA-SUBMODULE-miscDeps` | `brlapi37` | 315 | `2d217eff63997389fe2271810c6d13173185121c` | NON |
| `NVDA-RUST-UIA-STANDALONE` | `main` | 52 | `c732a3d88f65cbed34b21e2af6063222ec218061` | NON |
| `android` | `main` | 87 | `3d708dbc9271e1ccdc493b93a72f177dd83ae5f9` | NON |
| `accessible-windows` | `main` | 4 | `d809e75d5aeb1260c143bf013539bdf9a12f38d6` | NON |
| `serveur` | `main` | 4 | `6a55bc7cf8b20b11e7ecf90c999902ed527454e2` | NON |
| `project` | `main` | 35 | `1c0852cd6c3135ca7f40bf25785a980c55606d20` | NON |
| `solution` | `main` | 284 | `5e8947eba7a6734ee1cbe806f7abdcef782eb985` | NON |
| `omni-security` | `main` | 58 | `db62100fafd50f3e3f74895afcc1dfd0bc070b1c` | NON |
| `omni-os` | `main` | 1241 | `b8db99ec0e90c41f7bb757eac67c3a516ec41bc6` | NON |
| `UTM-Python316` | `work-20260929-093312` | 22 | `00ee273a5e87ba160f3de6077e817fc5cef099d4` | NON |
| `ADMWS12` | `main` | 61 | `a7f38d456d05b35d962f217264aabb594ed04e3c` | NON |
| `u` | `main` | 89 | `1b0fbdb676273602ca38583dc4b1cf8da80865b9` | NON |
| `test` | `main` | 171 | `7f1ac40661b0d20cb314fa09df9278369f190731` | NON |
| `M1603QAAS-Audit` | `main` | 9 | `f353f4fbb6580ca41579e2db63dd0b2528551460` | NON |

## État du contrôle

- `TREE_INVENTORY`: PASS pour la complétude de la réponse API de ces snapshots (aucun arbre signalé tronqué).
- `CONTENT_AUDIT`: PARTIAL — il faut examiner les fichiers et les branches non par défaut, comparer les blobs et les licences.
- `BUILD_TEST`: NOT_RUN dans le cadre de cette capture.
- `MERGE`: aucune fusion ni suppression exécutée.

## Inventaire ZERO détaillé

Les chemins, tailles et blob SHA des 171 fichiers de ZERO sur le snapshot `7f1ac40661b0d20cb314fa09df9278369f190731` sont consignés dans `ZERO-FILE-INVENTORY-PART-1.md`, `ZERO-FILE-INVENTORY-PART-2.md` et `ZERO-FILE-INVENTORY-PART-3.md`. Ces trois parties partagent le même snapshot de référence.

## Priorité d'unification

1. ZERO bootstrap/format/machine records.
2. ADMWS12 capability/evidence semantics.
3. omni-os / u voice source provenance and hardware boundary.
4. NVDA-RUST-UIA-STANDALONE semantic model and Windows adapter.
5. accessible-windows UEFI/HII/IFR and voice path.
6. omni-security and solution verification gates.
7. Remaining NVDA upstream/submodule inventory and branch/patch comparison.
