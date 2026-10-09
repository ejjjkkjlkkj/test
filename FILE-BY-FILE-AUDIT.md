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
- Voice provenance comparison completed: `u/src/synth_inc.rs` blob `e446a31a488abf1e2af9dd4186c0d3e9abe55c36` (106,262 characters returned) differs from `omni-os/voice-st/src/synth_inc.rs` blob `2ec524e16552d5435c7d3b34da710f50b3ddf044` (113,267 characters returned); first content difference at byte offset 1706. Do not call them duplicates or merge them without a semantic/performance diff. `omni-os/os/crates/aw-voice/src/lib.rs` directly includes its own `voice-st` source.
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


## Fichiers supplémentaires consultés — lot 2 (2026-10-09)

### NVDA-RUST-UIA-STANDALONE — branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `Cargo.toml` | `cdc0d994e359eb09aa153b02ba103a383cded3f2` | Crate Rust 2024 ; dépendance Windows crate épinglée à une révision Git précise et features UI Automation | CANDIDATE ; vérifier les modules, build Windows et tests. |
| `src/lib.rs` | `852b5af11540af229863f8aacda0d68215db81c3` | Expose modules platform, presentation, semantic et réglage du moniteur | CANDIDATE ; comparer les contrats sémantiques avec ZERO sans importer le runtime Windows dans le noyau portable. |
| `LICENSE` | — | 404 au chemin racine testé | LICENSE_NOT_FOUND_AT_PATH ; ne pas supposer de licence avant examen des autres fichiers et métadonnées. |

### u — synthèse vocale ST, branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `Cargo.toml` | `e4624290a34c94373934ce3793c0def4664165c0` | Package `st` version `0.6.0-rc.3`, Rust 2021, crate `cdylib` + `rlib`, dépendances `libm` et `serde_json` | CANDIDATE ; vérifier fonctionnalités et licences de données/voix. |
| `src/lib.rs` | `50957bedec5f55495beff6563da81a5954371d02` | Modules audio, engine, frontend, FFI v1 ; synthèse liée à `synth_inc.rs` | CANDIDATE ; the file was compared directly to `omni-os/voice-st/src/synth_inc.rs` and is NOT byte-identical. |
| `tests/synth_tests.rs` | — | 404 au chemin testé | PATH_NOT_FOUND ; localiser le vrai dossier de tests via arborescence avant toute conclusion. |

### solution — vérification, branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `pyproject.toml` | `14ff7514f9de721da743fe9f698196963b3e80a0` | Package Python `omniexec-solution`, licence déclarée 0BSD, Python >=3.11, dépendances runtime déclarées vides | CANDIDATE pour méthodologie de vérification ; pas une dépendance obligatoire du runtime ZERO. |
| `src/omni/cli.py` | `f874d644a0e71de2c0ddf0a2f689372ebb0fdbb6` | CLI assemble modules ceiling, firmware, IFR, voice frontend/pipeline/quality | CANDIDATE ; inspecter chaque module et les tests avant de réutiliser un validateur. |
| `tests/test_integrity.py` | — | 404 au chemin testé | PATH_NOT_FOUND ; les tests ne sont pas déclarés absents, seulement non localisés. |

### ADMWS12 — modèles de plateforme, branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `src/platform/capability.py` | `852c6e852918f71c804ed5a9f766900bcc79bb2f` | CapabilityState distingue UNKNOWN, AVAILABLE, UNAVAILABLE, UNSUPPORTED, FAILED | CANDIDATE pour mapping sémantique ; conserver les états sans les promouvoir implicitement. |
| `src/platform/evidence.py` | `45f70957a19995685f85d9a55ee64487fc3cf0da` | EvidenceState inclut OBSERVED, ABSENT, UNSUPPORTED, FAILED, UNKNOWN | CANDIDATE ; préserver la distinction entre absence observée et inconnue. |
| `src/platform/hal.py` | `dc4f853248baee4368c65e5861e1eedad2707f39` | PlatformContext relie lifecycle, capabilities, discovery et état de plateforme | CANDIDATE pour spécification d'adaptateur ; aucune intégration Go n'est déclarée réalisée. |

### android — branche `main`

| Chemin testé | Résultat | Statut |
|---|---|---|
| `README.md` | Lu au lot précédent ; décrit une distribution Android 17 x86-64 accessible en VM, ISO/installation et audio | SOURCE-CLAIM ; les artefacts et workflows doivent être vérifiés. |
| `Cargo.toml`, `LICENSE` | 404 aux chemins racine testés | PATH_NOT_FOUND ; ce dépôt n'est pas supposé être Rust et la licence reste à localiser. |

### Conclusions du lot 2

- Le modèle de capacité et d'évidence ADMWS12 peut informer les contrats ZERO, mais le code Python n'est pas intégré au runtime.
- La voix de `u` et celle annoncée par `omni-os` présentent un lien source explicite ; il faut prouver la provenance exacte et éviter de dupliquer ou de remplacer la version de référence sans tests comparatifs.
- Les crates UIA Windows sont spécifiques à leur plateforme ; les objets sémantiques doivent rester indépendants du backend.
- Les 404 indiquent uniquement que le chemin testé n'a pas été trouvé. Ils ne prouvent ni absence de tests ni absence de licence dans le dépôt.
- Aucun code source n'a été copié ; aucune branche, aucun fichier source et aucun historique n'ont été supprimés.


## Fichiers supplémentaires consultés — lot 3 (2026-10-09)

### NVDA-RUST-UIA-STANDALONE — branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `src/semantic.rs` | `a068aef2811b16eb1911c82dfebb932e8deeec92` | Définit un modèle de rôles sémantiques accessibles, incluant notamment Unknown, Application, Window, Dialog, Document, Heading, Paragraph et contrôles | CANDIDATE ; comparer le vocabulaire et les états au modèle canonique ZERO. |
| `src/presentation.rs` | `38212a887a2565c17258a1e1d97052d05d4414f9` | Définit une présentation vocale avec priorités Background, Normal, Focus et Urgent | CANDIDATE ; garder la politique de présentation séparée de la sémantique et de la machine. |
| `src/platform/mod.rs` | — | 404 au chemin testé | PATH_NOT_FOUND ; retrouver les modules exacts sans inférer leur absence. |
| `.github/workflows/ci.yml` | — | 404 au chemin testé | PATH_NOT_FOUND ; CI non vérifiée par ce chemin. |

### omni-security — branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `SECURITY.md` | `48534ea9c423ee388e2a44415a5ac205a8672a5e` | Politique de sécurité ; périmètre couvrant cryptographie, firmware, boot, réseau, confidentialité, supply chain, identité et chemin d'accessibilité de confiance | CANDIDATE comme exigences de sécurité ; ce n'est pas une preuve de conformité du code. |
| `LICENSE`, `LICENSE.md`, `docs/THREAT-MODEL.md` | — | 404 aux chemins testés | LICENSE/THREAT_MODEL_NOT_FOUND_AT_TESTED_PATHS ; licence et modèle de menace complets restent à localiser. |

### omni-os — manifestes de crates, branche `main`

| Chemin | Blob SHA | Observation | Statut / décision |
|---|---|---|---|
| `os/crates/aw-accessibility/Cargo.toml` | `0ce26e8ab778f8f5ca2a760b088460639dc1a3cf` | Crate accessible partagé dans le workspace, édition/licence/lints hérités | CANDIDATE ; inspecter les types et tests avant mapping vers ZERO. |
| `os/crates/aw-kernel-contract/Cargo.toml` | `a07c5548c38c22bd8fd02c7a108f1ad814d7edf2` | Crate de contrats kernel séparé | CANDIDATE ; comparer les contrats à ZERO sans importer des détails de plateforme. |
| `os/crates/aw-x86-platform/Cargo.toml` | `df15cf4e999d806cc38dd5cbadaf413f953ec34e` | Crate de plateforme x86, sans dépendances déclarées dans l'extrait consulté | CANDIDATE ; spécifique backend x86, pas langage universel. |
| `os/crates/aw-voice/src/tests.rs` | — | 404 au chemin testé | PATH_NOT_FOUND ; tests à localiser dans l'arborescence exacte. |

### Notes du lot 3

- Le prototype UIA distingue déjà le modèle sémantique de la présentation vocale. C'est une séparation utile à préserver dans ZERO, mais le code ne doit pas être copié avant comparaison des rôles, états, événements, tests et licence.
- `omni-security/SECURITY.md` peut alimenter les exigences transversales. L'absence de LICENSE au chemin racine testé empêche encore toute conclusion de réutilisation juridique.
- Aucun test n'a été exécuté, aucun statut CI n'a été validé, aucun fichier source n'a été copié et aucune branche n'a été fusionnée.


## Integration progress — ZERO-native adapters (2026-10-09)

### ADMWS12 → ZERO

- Implementation: `bootstrap/go/admws12_mapping.go`.
- Mapping: Capability states UNKNOWN/AVAILABLE/UNAVAILABLE/UNSUPPORTED/FAILED and Evidence states OBSERVED/ABSENT/UNSUPPORTED/FAILED/UNKNOWN are accepted exactly; unknown kind/state and blank identity are rejected.
- Payload: fixed-order escaped fields; decoder validates and reverses the payload.
- Safety: `Source=ADMWS12`, `Target=ZERO`, `Time=UNKNOWN`, `Proof=UNPROVEN`; no source state is upgraded into independent proof.
- Tests: all 10 source states, invalid kind/state/identity, escaped Unicode/delimiters and round-trip.

### NVDA-RUST-UIA-STANDALONE → ZERO

- Source files: `src/semantic.rs` blob `a068aef2811b16eb1911c82dfebb932e8deeec92` and `src/presentation.rs` blob `38212a887a2565c17258a1e1d97052d05d4414f9`.
- Implemented ZERO adapter: `bootstrap/go/uia_mapping.go`; maps identity, role, native role, accessible name and state names to `SEMANTIC_NODE` records.
- Forward compatibility: unknown role/state names are preserved rather than dropped.
- Privacy boundary: the adapter accepts no raw control value; only accessible name and semantic states are serialized.
- Proof: remains `UNPROVEN`; this does not prove the native UIA runtime is running or that a physical Windows desktop was observed.
- Tests: Unicode/delimiter round-trip, unknown state preservation, invalid identity and invalid state delimiters.

### CI evidence for the adapters

- `ZERO bootstrap validation` run `37917266841`, commit `72a78afb9e065be69a61295e82e512617e3dd6a0`: **SUCCESS**.
- `ZERO validation` run `37917266796`, same commit: **SUCCESS**.
- `ZERO Go validation` run `37917266847`, same commit: **SUCCESS**.
- The later documentation-only commit `2f3e67d9ebd98086e8fb4be01fb6d400a614c125` has its own global workflow in progress at the time of this update.


## Voice provenance comparison (2026-10-09)

- Exact content comparison completed between `u/src/synth_inc.rs` and `omni-os/voice-st/src/synth_inc.rs`: not identical; distinct blob SHAs and sizes, first difference at offset 1706. This resolves the earlier tentative relationship note: they are related candidate implementations, not a byte-identical shared source.
- `omni-os/os/crates/aw-voice/src/lib.rs` includes `../../../../voice-st/src/synth_inc.rs` directly. That establishes intra-repository source reuse for omni-os, not equivalence with the separate `u` repository.
- Next gate: compare APIs, golden corpora, license notices, output sample rate, mastering/resampling, latency and intelligibility before selecting a single voice backend. No voice source was copied or merged.
