# ZERO — Inventaire des dépôts GitHub accessibles

Date de capture : 2026-10-09  
Compte examiné : `ejjjkkjlkkj`  
Résultat de l'énumération : 30 dépôts visibles dans la connexion GitHub au moment de la capture.  
Portée : noms, branches par défaut, taille déclarée et première classification. Ce fichier ne constitue pas encore l'audit du code, de toutes les branches ou des licences.

## Inventaire

| # | Dépôt | Branche par défaut | Taille GitHub (Ko) | Catégorie initiale | Priorité |
|---:|---|---|---:|---|---|
| 1 | `NVDA-UPSTREAM-COMPLETE` | `master` | 1,074,538 | Lecteur d'écran upstream / référence et dépendances | P2 |
| 2 | `NVDA-SUBMODULE-.vscode` | `nvda-pinned` | 25 | Configuration d'éditeur NVDA | P4 |
| 3 | `NVDA-SUBMODULE-include-cppjieba` | `master` | 9,742 | Dépendance linguistique NVDA | P3 |
| 4 | `NVDA-SUBMODULE-include-cppjieba-deps-limonp` | `ci-windows-2022` | 735 | Dépendance transitive NVDA | P3 |
| 5 | `NVDA-SUBMODULE-include-detours` | `4.0.1` | 841 | Dépendance système NVDA | P3 |
| 6 | `NVDA-SUBMODULE-include-espeak` | `android` | 78,724 | Synthèse vocale / dépendance NVDA | P2 |
| 7 | `NVDA-SUBMODULE-include-ia2` | `mark` | 203 | Accessibilité Windows / IAccessible2 | P2 |
| 8 | `NVDA-SUBMODULE-include-javaAccessBridge32` | `jab64` | 79 | Accessibilité Java | P3 |
| 9 | `NVDA-SUBMODULE-include-liblouis` | `archive/zstanecic_tables` | 99,933 | Braille et tables de traduction | P2 |
| 10 | `NVDA-SUBMODULE-include-nsis` | `add-Process-docs` | 3,620 | Installation / packaging | P4 |
| 11 | `NVDA-SUBMODULE-include-nvda-cldr` | `main` | 5,466 | Données de locale et internationalisation | P3 |
| 12 | `NVDA-SUBMODULE-include-nvda-mathcat` | `main` | 8,243 | Sémantique mathématique accessible | P2 |
| 13 | `NVDA-SUBMODULE-include-sonic` | `master` | 11,774 | Traitement vocal / dépendance NVDA | P3 |
| 14 | `NVDA-SUBMODULE-include-sonic-speedy` | `main` | 26,227 | Traitement vocal / dépendance NVDA | P3 |
| 15 | `NVDA-SUBMODULE-include-w3c-aria-practices` | `2021-11_Note` | 36,241 | Sémantique et pratiques ARIA | P2 |
| 16 | `NVDA-SUBMODULE-include-wil` | `copilot/port-crash-resistant-event-handler` | 2,913 | Bibliothèque Windows C++ | P4 |
| 17 | `NVDA-SUBMODULE-miscDeps` | `brlapi37` | 67,785 | Dépendances diverses NVDA | P3 |
| 18 | `NVDA-RUST-UIA-STANDALONE` | `main` | 348 | Noyau sémantique / UI Automation Rust | P1 |
| 19 | `android` | `main` | 49,574 | Plateforme Android x86-64 accessible | P3 |
| 20 | `accessible-windows` | `main` | 24,304 | Accessibilité UEFI et pré-OS | P1 |
| 21 | `serveur` | `main` | 12 | Serveur — rôle à déterminer | P4 |
| 22 | `project` | `main` | 550 | Voix et navigation UEFI | P1 |
| 23 | `solution` | `main` | 997 | Validation, tests, fuzzing et preuves | P1 |
| 24 | `omni-security` | `main` | 7,309 | Sécurité et accessibilité sécurisée | P1 |
| 25 | `omni-os` | `main` | 47,911 | OS, noyau Rust, UEFI, pilotes et audio | P1 |
| 26 | `UTM-Python316` | `work-20260929-093312` | 44 | Outils Python / UTM — rôle à confirmer | P4 |
| 27 | `ADMWS12` | `main` | 140 | Capacités, découverte, preuves et HAL | P1 |
| 28 | `u` | `main` | 17,327 | Synthèse vocale Rust | P1 |
| 29 | `test` (ZERO) | `main` | 16,348 | Langage machine, noyau sémantique et communication | P0 |
| 30 | `M1603QAAS-Audit` | `main` | 23,582 | Audit firmware privé et preuves ASUS | P2 |

## Décision d'architecture provisoire

- **ZERO (`test`) reste le dépôt d'intégration cible.**
- **ADMWS12** sert à analyser et préparer des contrats de capacités, découverte, état, preuves et adaptateurs. Il ne doit pas devenir un second noyau concurrent.
- **omni-os**, **accessible-windows**, **project**, **u** et **NVDA-RUST-UIA-STANDALONE** sont des sources candidates à analyser au niveau fichier/commit avant toute réutilisation.
- Les dépôts `NVDA-SUBMODULE-*` sont d'abord des dépendances upstream ou des références. Ne pas les copier intégralement dans ZERO.
- `solution` est candidat pour les validateurs et gates ; `omni-security` pour les contrats de sécurité et d'autorisation.
- `android` doit rester un backend/port optionnel, sans rendre Android obligatoire pour le langage ou le noyau.
- `M1603QAAS-Audit` est privé ; n'exporter aucune donnée sensible dans les dépôts publics.

## Limites de la capture

- La branche par défaut est recensée ; toutes les autres branches et leur contenu ne sont pas encore audités.
- La taille GitHub est une mesure indicative, pas une mesure de qualité ou de réutilisabilité.
- Les priorités sont des ordres de lecture, pas des validations.
- Aucun niveau PASS d'intégration n'est attribué par ce seul inventaire.
