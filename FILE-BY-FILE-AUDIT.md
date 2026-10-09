# ZERO — Audit fichier par fichier et registre de conservation

Date du registre : 2026-10-09  
Dépôt cible : `ejjjkkjlkkj/test` (ZERO), branche `main`  
Statut : **PARTIAL — audit en cours, aucun dépôt déclaré entièrement audité.**

## Règles non négociables

1. Ne supprimer ni dépôt, ni branche, ni commit, ni fichier source, ni preuve historique pendant l'audit.
2. Aucun code n'est considéré intégré parce qu'un README le décrit. Chaque élément doit pointer vers dépôt, branche, commit, chemin et blob SHA.
3. Ne pas écraser les variantes : les comparer et choisir une référence seulement après examen de leurs différences.
4. Séparer explicitement : documentation, code, tests définis, tests exécutés, CI, simulation, QEMU/VM, matériel physique.
5. Pour chaque fichier : classer `KEEP_SOURCE`, `CANDIDATE`, `DUPLICATE_CANDIDATE`, `EXCLUDE_WITH_REASON`, ou `NOT_REVIEWED`. Aucune suppression automatique.
6. Les licences doivent être vérifiées au niveau du fichier et du dépôt avant copie ou dérivation.
7. Objectif qualité : 99,99 % mesuré ; ce n'est pas un résultat actuel.
8. Le connecteur disponible ici permet de lire des chemins connus et de rechercher du contenu, mais ne fournit pas d'énumération récursive fiable de tous les arbres Git. Les entrées ci-dessous sont donc une piste de travail explicite, pas une fausse déclaration d'exhaustivité. Pour achever le vrai inventaire fichier par fichier, il faut obtenir les arbres complets de chaque branche, sous-modules et fichiers suivis.

## Registre des dépôts accessibles

La liste ci-dessous provient de l'inventaire GitHub accessible le 2026-10-09. La branche est la branche par défaut indiquée par GitHub, pas nécessairement la branche canonique de développement.

| # | Dépôt | Branche par défaut | Rôle candidat | Statut audit | Conservation / prochaine action |
|---:|---|---|---|---|---|
| 1 | `NVDA-UPSTREAM-COMPLETE` | `master` | Source NVDA volumineuse | NOT_REVIEWED | Conserver intact ; cartographier origine upstream, sous-modules, licence, patchs et branches avant toute intégration. |
| 2 | `NVDA-SUBMODULE-.vscode` | `nvda-pinned` | Configuration d'éditeur | NOT_REVIEWED | Conserver comme dépendance potentielle ; vérifier si configuration ou code requis. |
| 3 | `NVDA-SUBMODULE-include-cppjieba` | `master` | Segmentation texte | NOT_REVIEWED | Vérifier provenance, licence, versions et dépendances amont. |
| 4 | `NVDA-SUBMODULE-include-cppjieba-deps-limonp` | `ci-windows-2022` | Dépendance de cppjieba | NOT_REVIEWED | Préserver la référence épinglée ; vérifier compatibilité/licence. |
| 5 | `NVDA-SUBMODULE-include-detours` | `4.0.1` | Instrumentation Windows | NOT_REVIEWED | Vérifier usage réel et licence ; ne pas importer dans le noyau portable. |
| 6 | `NVDA-SUBMODULE-include-espeak` | `android` | Synthèse vocale | NOT_REVIEWED | Comparer fork et upstream, voix, licences et build. |
| 7 | `NVDA-SUBMODULE-include-ia2` | `mark` | API Accessibilité Windows | NOT_REVIEWED | Vérifier IDL/headers et provenance. |
| 8 | `NVDA-SUBMODULE-include-javaAccessBridge32` | `jab64` | Access Bridge Java | NOT_REVIEWED | Vérifier architecture, artefacts et provenance. |
| 9 | `NVDA-SUBMODULE-include-liblouis` | `archive/zstanecic_tables` | Braille / tables | NOT_REVIEWED | Préserver tables et historique ; comparer à upstream avant sélection. |
| 10 | `NVDA-SUBMODULE-include-nsis` | `add-Process-docs` | Installation Windows | NOT_REVIEWED | Séparer code d'installateur et documentation ; vérifier patchs. |
| 11 | `NVDA-SUBMODULE-include-nvda-cldr` | `main` | Données locales / CLDR | NOT_REVIEWED | Vérifier données, licence, génération et version. |
| 12 | `NVDA-SUBMODULE-include-nvda-mathcat` | `main` | Lecture mathématique | NOT_REVIEWED | Vérifier origine, API et jeux de tests. |
| 13 | `NVDA-SUBMODULE-include-sonic` | `master` | Traitement de parole | NOT_REVIEWED | Comparer au fork speedy et mesurer les effets audio. |
| 14 | `NVDA-SUBMODULE-include-sonic-speedy` | `main` | Variante traitement de parole | NOT_REVIEWED | Ne pas fusionner avec sonic avant diff et tests comparatifs. |
| 15 | `NVDA-SUBMODULE-include-w3c-aria-practices` | `2021-11_Note` | Référence sémantique ARIA | NOT_REVIEWED | Traiter comme référence documentaire ; contrôler version et licence. |
| 16 | `NVDA-SUBMODULE-include-wil` | `copilot/port-crash-resistant-event-handler` | Utilitaires Windows | NOT_REVIEWED | Vérifier la branche non standard et la compatibilité. |
| 17 | `NVDA-SUBMODULE-miscDeps` | `brlapi37` | Dépendances mixtes NVDA | NOT_REVIEWED | Décomposer par sous-composant, licence et architecture. |
| 18 | `NVDA-RUST-UIA-STANDALONE` | `main` | Lecteur d'écran Rust / UIA | PARTIAL | README et Cargo.toml lus précédemment ; auditer chaque module, tests Windows, licence et API. |
| 19 | `android` | `main` | OS Android x86-64 accessible | PARTIAL | README lu ; considérer comme backend/plateforme distincte, pas dépendance obligatoire du noyau ZERO. |
| 20 | `accessible-windows` | `main` | Orchestration matérielle historique | PARTIAL | Lire aussi `repo-clean-consolidation-20260924`; préserver les deux histoires et vérifier les fichiers de code de chaque branche. |
| 21 | `serveur` | `main` | Rôle à déterminer | NOT_REVIEWED | Inspecter chaque fichier et historique avant attribution de rôle. |
| 22 | `project` | `main` | Voix + navigation UEFI extraites | PARTIAL | README lu ; source déclarée `uefi-realtime-screenreader-20260919` / commit `4ae932d8a8dff533cfeb92156d49634fd060f84b`. Racine Cargo non trouvée aux chemins testés. |
| 23 | `solution` | `main` | Vérification, fuzzing, preuves | PARTIAL | README et licence 0BSD rapportés dans l'audit précédent ; commandes décrites en Python. |
| 24 | `omni-security` | `main` | Sécurité + accessibilité | PARTIAL | README lu ; fichier LICENSE racine non trouvé au chemin testé. Ne pas supposer une licence. |
| 25 | `omni-os` | `main` | Plateforme OS, UEFI, noyau, voix | PARTIAL | README, LICENSE, `os/Cargo.toml`, `os/Cargo.lock`, `os/rust-toolchain.toml`, CI et fichiers de voix/lecteur d'écran inspectés partiellement. Voir constats ci-dessous. |
| 26 | `UTM-Python316` | `work-20260929-093312` | Rôle à déterminer | NOT_REVIEWED | Inspecter branche et fichiers ; évaluer Python comme outil hors runtime ZERO si pertinent. |
| 27 | `ADMWS12` | `main` | Modèles de capacités, preuves, HAL et recherche IA | PARTIAL | README et `src/platform/capability.py`, `evidence.py`, `hal.py` consultés précédemment ; extraire contrats, ne pas importer le runtime Python tel quel. |
| 28 | `u` | `main` | Synthèse vocale Rust ST | PARTIAL | README et Cargo.toml lus ; séparer backend compact Rust de backend neural et ses dépendances privées. |
| 29 | `test` (ZERO) | `main` | Langage canonique, modèle sémantique, bootstrap | PARTIAL | Auditer chaque fichier, tests, format, sécurité et CI avant de déclarer le noyau conforme. |
| 30 | `M1603QAAS-Audit` | `main` | Audit firmware ASUS / preuves | NOT_REVIEWED | Privé ; conserver les preuves, éviter de copier des données sensibles dans les dépôts publics. |

## Fichiers effectivement consultés dans cette passe

### ZERO — `ejjjkkjlkkj/test`, branche `main`

| Chemin | Blob SHA | Résultat de la lecture | Décision provisoire |
|---|---|---|---|
| `go.mod` | `b071d4b865b5ca89f52c6a88fac1495f3d570b11` | Module `zero`, Go 1.23 | CANDIDATE — bootstrap Go ; compiler/tester réellement avant PASS. |
| `ZERO-MACHINE-FORMAT.md` | `af65d80f9785ef3fd212de0acff949b1d17bef64` | Format documenté comme représentation canonique ; document précise qu'il n'est pas encore un protocole réseau ni un exécutable | CANDIDATE — contrat à comparer au parser et aux implémentations indépendantes. |
| `bootstrap/go/record.go` | `bb41fdded18a68194f73e6a9e2c7b6420fdc0ca9` | Encode/Decode, échappements `\\`, `\|`, `\n`, `\r`, `\t`, rejet des séquences non canoniques | CANDIDATE — tests réels requis ; champs vides et limites doivent être vérifiés. |
| `bootstrap/go/engine.go` | `6666f92e2a2371304cfd5064cc5dea22f8b2268a` | Transition déterministe, refuse UNKNOWN/CONTRADICTION | CANDIDATE — vérifier toutes les branches de résultat et l'autorisation. |
| `bootstrap/go/canonical_test.go` | `5d6bb0e6062686e53a75fb169e7651cf95e56dd7` | Tests de canonicalité, échappement, Unicode et round-trip | TESTS_DEFINED ; NOT_RUN dans cette session. |
| `.github/workflows/go.yml` | `2f0ee4d4a72387f6386ae266293b1705d3d9886c` | Workflow push/PR/manual prévu pour Go test et vet | CI_CONFIGURED ; aucun run associé n'a encore été vérifié ici. |

### omni-os — `ejjjkkjlkkj/omni-os`, branche `main`

| Chemin | Blob SHA | Résultat de la lecture | Décision provisoire |
|---|---|---|---|
| `README.md` | consulté ; SHA à enregistrer lors de la prochaine lecture | Décrit OS x86-64 accessible, UEFI, noyau Rust, voix et navigation | SOURCE-CLAIM ; ce n'est pas une preuve de build/boot. |
| `LICENSE` | `01506c469a30ac12c42a12e7fae25c5e33a3289f` | Licence permissive 0BSD-like | LICENCE-OBSERVED ; vérifier les licences de chaque source incorporée. |
| `os/Cargo.toml` | `0f9e42aed8dc6ac1279f1b0105e29f4653ca1a4c` | Workspace Rust avec de nombreux crates `aw-*`, exclusion explicite de `boot/uefi`, `kernel/x86_64`, `fuzz` | CANDIDATE ; l'exclusion signifie que les builds workspace ne prouvent pas ces composants. |
| `os/Cargo.lock` | `5e20ce68a7e258e4f2518b6b6ae96aa0dc3b89d4` | Lockfile de workspace ; packages locaux et dépendances | CANDIDATE ; vérifier cohérence avec manifestes et CI. |
| `os/rust-toolchain.toml` | `222d8cb2b39f9e82ad282798bfbf7311adc0ca26` | Rust 1.98.1, cible x86_64-unknown-none | TOOLCHAIN_DECLARED ; non validé localement. |
| `os/crates/aw-voice/Cargo.toml` | `48511d693fc15414d441fae99651b28ddce06e54` | Crate `aw-voice`, edition 2024, dépendance libm | CANDIDATE ; vérifier licence héritée, tests et provenance réelle de la voix. |
| `os/crates/aw-voice/src/lib.rs` | `025b11dbf0de8a107d995b1391505bd2e4deb6bf` | Déclare réutiliser la source `voice-st` via inclusion, plutôt que copier la source | DUPLICATE_CANDIDATE ; vérifier l'inclusion exacte, le chemin, l'historique et l'identité des octets. |
| `os/crates/aw-screen-reader/Cargo.toml` | `01049275d586dc616888f32207cff459e1cac89a` | Crate screen reader dépendant de `aw-accessibility` | CANDIDATE ; lire source/tests et distinguer sémantique de sortie vocale. |
| `tools/integrity/verify.py` | `e3708fd707180d3162d357b7e7e6e0f77a87d1c3` | Gate de provenance, branches d'archive, secrets et tailles | CANDIDATE comme contrôle de conservation ; c'est un outil Python, pas runtime ZERO. |
| `.github/workflows/ci.yml` | `92abbfeb07cb30c19ca96730595c3b2fbdc6c5e8` | CI intégrité + tests workspace/UEFI/kernel annoncés | CI_CONFIGURED ; consulter les runs récents et les étapes exactes avant PASS. |
| `.github/workflows/integrity-daily.yml` | `04fa7ee0db71eb3fb845b334ff136d8b5fa3045c` | Audit quotidien des branches archive | CANDIDATE comme mécanisme de conservation ; vérifier les résultats d'exécution. |

### accessible-windows — source consolidée

| Chemin | Branche | Blob SHA | Résultat de la lecture | Décision provisoire |
|---|---|---|---|---|
| `README.md` | `repo-clean-consolidation-20260924` | `1eef452167295024869ae6a085682fdcfab2708c` | Confirme le périmètre UEFI/HII/IFR/audio et que la parité reste ouverte | SOURCE-CLAIM ; pas une validation actuelle. |
| `docs/STATUS.md` | `repo-clean-consolidation-20260924` | `780acbc2111171535f12277065af809b3fbdbd67` | Rapporte un run physique vert historique et une sortie 16 kHz ; cible 24 kHz/parité ouverte | SOURCE-REPORTED ; ne pas présenter comme rejoué. |
| `docs/STATUS.md` | `main` | `a8717fe5fe16ee60851ad708f6c7821f65e8c9a9` | `main` est orchestration/compatibilité ; run 35990276163 indiqué queued au moment de mise à jour | SOURCE-REPORTED ; récupérer le statut réel du run avant toute conclusion. |
| `Cargo.toml` | `repo-clean-consolidation-20260924` | — | 404 au chemin racine testé | PATH_NOT_FOUND ; ne pas conclure qu'aucun manifeste n'existe dans les sous-répertoires. |
| `.github/workflows/ci.yml` | `repo-clean-consolidation-20260924` | — | 404 au chemin testé | PATH_NOT_FOUND ; rechercher les chemins workflow exacts. |

### project — extraction voix/navigation

| Chemin | Branche | Résultat de la lecture | Décision provisoire |
|---|---|---|---|
| `README.md` | `main` | Indique une extraction de `accessible-windows`, source `uefi-realtime-screenreader-20260919`, commit `4ae932d8a8dff533cfeb92156d49634fd060f84b`; exclut volontairement noyau, stockage, USB, etc. | DUPLICATE_CANDIDATE par rapport aux sources ; préserver comme snapshot tant que le diff n'est pas vérifié. |
| `Cargo.toml`, `voice/Cargo.toml`, `navigation/Cargo.toml` | `main` | 404 sur les chemins testés | PATH_NOT_FOUND ; arborescence non résolue, ne pas supposer le langage ou l'absence de code. |

## Constat technique de cette passe

- ZERO a des artefacts de format, parser Go, moteur de transition, tests et workflow. Le code et les tests sont lus ; **les tests n'ont pas été exécutés par cette session**.
- `omni-os` contient un workspace sous `os/`; la racine du dépôt n'a pas de `Cargo.toml` au chemin testé. Plusieurs composants boot/kernel sont explicitement exclus du workspace : leur validation doit être indépendante.
- `omni-os` et `u` semblent avoir une relation de partage de code de voix ; aucune fusion de copie ne doit être faite avant vérification octet par octet.
- `accessible-windows` a au moins deux lignes ayant des rôles différents : `main` pour l'orchestration et la branche consolidée pour la source UEFI. Ne pas fusionner les historiques.
- `project` est annoncé comme extraction partielle de `accessible-windows`, donc doublon candidat, pas doublon confirmé.
- Les sous-dépôts NVDA ne doivent pas être traités comme 16 projets indépendants à fusionner : les liens de provenance, versions, sous-modules et patchs doivent être établis avant tout choix.

## Prochaine séquence d'audit

1. Obtenir une liste récursive exacte des arbres suivis (toutes les branches et sous-modules) et consigner chaque chemin avec blob SHA, taille, type, licence et décision.
2. Examiner les manifestes, tests, workflows et licences de chacun des 30 dépôts, sans se limiter aux READMEs.
3. Pour chaque composant candidat, comparer le contenu exact entre les dépôts et leurs commits sources ; établir les doublons par diff, pas par ressemblance de nom.
4. Ajouter dans ZERO des contrats et adaptateurs seulement après que les tests de référence sont reproductibles.
5. Vérifier les statuts CI par run et SHA ; exécuter les tests localement quand un runtime est disponible.
6. Après chaque lot, mettre à jour ce registre et `REPOSITORY-UNIFICATION-AUDIT.md`. Aucun lot ne vaut validation matérielle sans preuve matérielle correspondante.
