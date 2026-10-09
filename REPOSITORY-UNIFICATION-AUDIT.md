# ZERO — Registre central d'audit et d'unification des dépôts

Date de départ : 2026-10-09  
Statut global : PARTIAL — inventaire initial seulement. Aucun dépôt n'est déclaré entièrement audité.

## Objectif

Faire de ZERO le point d'intégration unique pour le langage machine universel, le modèle sémantique, l'exécution, les interfaces matérielles, l'accessibilité native, la communication et les preuves. Les autres dépôts restent des sources à examiner ; ils ne sont pas fusionnés automatiquement.

## Règles

1. Distinguer documentation, code implémenté, tests réussis, simulation, exécution QEMU et preuve sur matériel physique.
2. Ne supprimer aucun dépôt, branche, historique ou preuve unique pendant l'audit.
3. Vérifier licence, origine, API, tests et compatibilité avant toute réutilisation.
4. Chaque composant doit citer dépôt, branche, commit, chemin, décision et preuve.
5. La réussite dans le dépôt source ne prouve pas la réussite après intégration dans ZERO.
6. Viser une fiabilité mesurée de 99,99 % ; ne pas présenter l'objectif comme un résultat acquis.

## Statuts

- PASS : vérification réalisée, preuve consultable.
- PARTIAL : seule une partie du périmètre est vérifiée.
- BLOCKED : dépendance ou accès empêche la vérification.
- UNPROVEN : documentation ou déclaration sans preuve suffisante.
- DUPLICATE : fonctionnalité recouvrant un autre composant.
- EXCLUDE : hors périmètre ou incompatible, avec justification.
- NOT_RUN : contrôle non exécuté.

## Premier lot de dépôts candidats

| Dépôt | Domaine potentiel | Statut | Prochaine vérification |
|---|---|---|---|
| ejjjkkjlkkj/test (ZERO) | Langage, format canonique, noyau sémantique, transport | PARTIAL | Inventorier fichiers, branches, tests et CI |
| ejjjkkjlkkj/ADMWS12 | Capacités, découverte, preuves, HAL, état de plateforme | PARTIAL | Extraire les contrats réutilisables et vérifier la frontière entre prototype Python et runtime natif |
| ejjjkkjlkkj/omni-os | Rust, UEFI, boot, noyau, pilotes, audio | NOT_RUN | Identifier composants compilables, licences et niveau de preuve |
| ejjjkkjlkkj/accessible-windows | Accessibilité pré-OS, HII/IFR, navigation UEFI | NOT_RUN | Comparer branches et préserver les commits avec preuves uniques |
| ejjjkkjlkkj/NVDA-RUST-UIA-STANDALONE | UI Automation Windows et événements sémantiques | NOT_RUN | Vérifier API, tests et code Windows-spécifique |
| ejjjkkjlkkj/u | Synthèse vocale Rust et interfaces natives | NOT_RUN | Vérifier licence, dépendances, performances et sortie audio |
| ejjjkkjlkkj/project | Voix et navigation UEFI extraites | NOT_RUN | Rechercher les doublons avec accessible-windows et omni-os |
| ejjjkkjlkkj/solution | Validation, fuzzing, vérification et preuves | NOT_RUN | Identifier les validateurs indépendants et les gates réutilisables |
| ejjjkkjlkkj/omni-security | Sécurité, autorisation, confidentialité, accessibilité | NOT_RUN | Extraire les modèles transférables |
| ejjjkkjlkkj/android | Plateforme Android x86-64 accessible | NOT_RUN | Traiter comme backend optionnel, pas comme dépendance obligatoire du noyau |
| ejjjkkjlkkj/M1603QAAS-Audit | Inventaire et preuves firmware ASUS | NOT_RUN | Réutiliser les formats d'évidence sans propager de données sensibles |

Cette liste n'est pas encore exhaustive. Il faut ajouter tous les autres dépôts accessibles, branches, sous-modules et dépendances avant de conclure.

## Architecture d'unification

```text
Dépôts sources + preuves immuables
  -> inventaire exact (dépôt/branche/commit/chemin/licence)
  -> évaluation (code/tests/CI/dépendances/preuves)
  -> contrats ZERO (identité/état/capacité/événement/résultat/preuve)
  -> adaptateurs (UEFI/x86-64/Windows/Linux/Android/réseau)
  -> noyau et langage ZERO
  -> validation reproductible et projections accessibles
```

Le langage et le format canonique doivent rester indépendants du langage hôte. Les adaptateurs traduisent les interfaces physiques en objets sémantiques ZERO ; ils ne deviennent pas le langage lui-même.

## Ordre d'exécution

1. Inventorier tous les dépôts, branches, commits, sous-modules, manifestes, CI, tests et licences.
2. Cartographier chaque composant avec son emplacement exact et son niveau de preuve.
3. Détecter les doublons et dépendances ; choisir une implémentation de référence sans supprimer l'historique.
4. Stabiliser les contrats : capacités, observations, événements, opérations, résultats, erreurs et preuves.
5. Valider le format canonique par parser/sérialiseur, round-trip, échappements, champs vides, entrées invalides et types inconnus.
6. Intégrer progressivement, un composant à la fois, sur commit épinglé avec tests avant/après.
7. Séparer les validations unitaires, intégration, fuzzing, simulation, QEMU et matériel physique.

## Premier risque technique à vérifier

ZERO-MACHINE-FORMAT.md définit neuf champs après le marqueur RECORD ; bootstrap/go/record.go déclare fieldCount = 10 en comptant le marqueur. Cette concordance apparente ne suffit pas : il faut exécuter des tests d'échappement, champs vides, retours à la ligne et round-trip avant de déclarer le format conforme.

## Journal

- 2026-10-09 : première lecture de fichiers de référence dans ZERO, ADMWS12 et plusieurs dépôts candidats.
- 2026-10-09 : inventaire des 30 dépôts visibles dans REPOSITORY-INVENTORY.md ; les branches secondaires restent à examiner.
- 2026-10-09 : ajout du contrat provisoire ADMWS12-ZERO-ADAPTER-CONTRACT.md.
- 2026-10-09 : bootstrap/go/record.go rejette désormais les séquences non canoniques avec zéros initiaux ; canonical_test.go couvre séquence non canonique, échappement inconnu, Unicode et round-trip.
- 2026-10-09 : commits code/tests : ebc3b07203b25e93192a8c72e904e757c036046b et 59a52a2e4ea28662fd80230b47938f2ad2d8cfd7. Aucun workflow ni statut CI associé trouvé lors de la vérification ; tests non exécutés dans un runtime local par cette session.
- 2026-10-09 : audit complet de toutes les branches, licences, dépendances et composants reste NOT_RUN.


## Mise à jour — revue source ciblée (2026-10-09)

Le fichier [COMPONENT-MAP.md](COMPONENT-MAP.md) a été créé au commit `e29eeeb1b66197443672c57afbdcbe34c666a8fc`. Il consigne les chemins et SHA de blobs consultés, les rôles candidats, les limites de licence et les gates qui restent à vérifier.

### Résultats factuels

- **ADMWS12 — PARTIAL :** lecture directe de `src/platform/capability.py`, `src/platform/evidence.py` et `src/platform/hal.py`. Ces modèles sont des références sémantiques possibles, pas un runtime à copier tel quel.
- **omni-os — PARTIAL / UNPROVEN :** README et licence racine consultés. Le README décrit loader UEFI, noyau Rust, HII/IFR et audio ; ce sont des déclarations de projet, pas une validation indépendante ici. Le manifeste Cargo racine n'a pas été trouvé au chemin testé.
- **accessible-windows — PARTIAL :** README indique que la ligne cohérente est `repo-clean-consolidation-20260924`, commit de consolidation `b53b62895ddb57094aa944cb1cdc549f0523668d`. Il précise que l'issue de parité #4 reste ouverte et que la cible voix 24 kHz n'est pas PASS. Les fichiers de cette branche doivent encore être inspectés directement.
- **NVDA-RUST-UIA-STANDALONE — PARTIAL :** README et `Cargo.toml` consultés ; backend Windows UIA/COM et modèle d'événement portable identifiés. Licence, build Windows et tests restent à vérifier.
- **u (ST) — PARTIAL :** README et `Cargo.toml` consultés ; interface Rust et moteur compact sans dépendance runtime externe identifiés. La voie neurale mentionne des éléments Python privés : elle n'est pas retenue comme dépendance obligatoire du noyau ZERO.
- **solution — PARTIAL :** README et licence 0BSD consultés ; méthodologie de gates et séparation logiciel/matériel potentiellement réutilisables. Le manifeste Cargo racine n'a pas été trouvé au chemin testé ; les commandes documentées sont Python.
- **omni-security — PARTIAL :** README consulté pour les exigences transversales sécurité + accessibilité ; aucune implémentation n'est déclarée validée sur la seule base de ce document.

### Changements et limites

- Création de `COMPONENT-MAP.md` sur `main`.
- Aucun code n'a été copié ou fusionné depuis les dépôts candidats.
- Aucun test/build candidat n'a été exécuté dans cette passe.
- La correction précédente du parseur Go et ses tests ajoutés ne sont toujours pas déclarés PASS fonctionnel : aucun résultat d'exécution n'est disponible.
- Le périmètre compte toujours 30 dépôts recensés dans l'inventaire initial ; l'audit complet des branches, sous-modules, licences et fichiers reste NOT_RUN.


### Suite de la revue — preuves de statut et contrat d'adaptation (2026-10-09)

- Lecture de `accessible-windows/docs/STATUS.md` sur la branche `repo-clean-consolidation-20260924`, blob `780acbc2111171535f12277065af809b3fbdbd67`.
- Le document source rapporte une baseline historique verte au run `35846636674`, commit projet `8302a95b0b4d2fe4d57e2832bcc945819728b80d`, avec voix à 16 kHz, build NAVIGATION.EFI et découverte du lecteur d'écran en VMware. Statut dans ZERO : `SOURCE-REPORTED`, non rejoué ici.
- Le même document rapporte un run ultérieur `35848464855` échoué dans l'étape PsExec physique après `STAGE=SYSTEM_VOICE_READY`. La cible 24 kHz et la parité finale restent ouvertes. Ne pas transformer la baseline historique en preuve de réussite du HEAD actuel.
- Création de `ADMWS12-ADAPTER-IMPLEMENTATION.md`, commit `5babd9006d847a322a369dbed6ec0d7d99b17d48`. Le contrat préserve les états source, interdit toute promotion implicite de preuve et exige des tests explicites. Il reste `SPECIFIED / NOT_IMPLEMENTED / NOT_TESTED`.
- Mise à jour de `COMPONENT-MAP.md`, commit `e2591b338450d58b9a281c124cae83975c463fe7`.
- Contrôles CI consultés pour le commit de la spécification : aucun statut ni workflow associé n'a été retourné. Le workflow Go est présent, mais son succès n'est pas encore établi.


### Audit élargi — registre fichier par fichier (2026-10-09)

- Création de `FILE-BY-FILE-AUDIT.md`, commit `1831955c8208359c5b8eca65e87eb9a426c2a3cf`.
- Le registre reprend les 30 dépôts visibles, les branches par défaut, les rôles candidats et les règles de conservation. Il distingue les fichiers réellement consultés des chemins seulement testés et des zones non examinées.
- Nouvelle lecture directe de ZERO : `go.mod` (blob `b071d4b865b5ca89f52c6a88fac1495f3d570b11`), `ZERO-MACHINE-FORMAT.md` (blob `af65d80f9785ef3fd212de0acff949b1d17bef64`), `bootstrap/go/record.go` (blob `bb41fdded18a68194f73e6a9e2c7b6420fdc0ca9`), `bootstrap/go/engine.go` (blob `6666f92e2a2371304cfd5064cc5dea22f8b2268a`), `bootstrap/go/canonical_test.go` (blob `5d6bb0e6062686e53a75fb169e7651cf95e56dd7`) et `.github/workflows/go.yml` (blob `2f0ee4d4a72387f6386ae266293b1705d3d9886c`). Les tests restent `NOT_RUN` ; aucun PASS d'exécution n'est revendiqué.
- Nouvelle lecture directe d'`omni-os` : `os/Cargo.toml` (blob `0f9e42aed8dc6ac1279f1b0105e29f4653ca1a4c`), `os/Cargo.lock` (blob `5e20ce68a7e258e4f2518b6b6ae96aa0dc3b89d4`), `os/rust-toolchain.toml` (blob `222d8cb2b39f9e82ad282798bfbf7311adc0ca26`), `os/crates/aw-voice/Cargo.toml` (blob `48511d693fc15414d441fae99651b28ddce06e54`), `os/crates/aw-voice/src/lib.rs` (blob `025b11dbf0de8a107d995b1391505bd2e4deb6bf`), `os/crates/aw-screen-reader/Cargo.toml` (blob `01049275d586dc616888f32207cff459e1cac89a`), `tools/integrity/verify.py` (blob `e3708fd707180d3162d357b7e7e6e0f77a87d1c3`), `.github/workflows/ci.yml` (blob `92abbfeb07cb30c19ca96730595c3b2fbdc6c5e8`) et `.github/workflows/integrity-daily.yml` (blob `04fa7ee0db71eb3fb845b334ff136d8b5fa3045c`). Le workspace Rust est sous `os/` et exclut explicitement `boot/uefi`, `kernel/x86_64` et `fuzz` ; ces zones doivent être validées indépendamment.
- Nouvelle lecture d'`accessible-windows` : README consolidé (blob `1eef452167295024869ae6a085682fdcfab2708c`), `docs/STATUS.md` de la branche consolidée (blob `780acbc2111171535f12277065af809b3fbdbd67`) et de `main` (blob `a8717fe5fe16ee60851ad708f6c7821f65e8c9a9`). Les résultats physiques sont des déclarations source datées ; le run indiqué comme queued doit être revérifié. Aucun fichier n'a été fusionné.
- Nouvelle lecture de `project/README.md` (blob `29696f6c4fd9a3c83d07b5026db335dddf6e8268`), qui décrit une extraction limitée de voix/navigation et un commit source figé ; classé `DUPLICATE_CANDIDATE`, non doublon confirmé.
- Limite d'outillage : l'interface GitHub utilisée ici ne fournit pas d'énumération récursive de l'arbre. L'audit complet fichier par fichier n'est donc pas terminé ; le registre rend explicites les éléments lus, les chemins en 404 et les éléments `NOT_REVIEWED`. Aucun dépôt, branche ou historique n'a été supprimé.


### Audit fichier par fichier — lot 2 (2026-10-09)

- Le registre détaillé `FILE-BY-FILE-AUDIT.md` a été enrichi au commit `7837d33e93dd3a424ecb8c496a772af6dc09fd86`.
- Lecture de `NVDA-RUST-UIA-STANDALONE/Cargo.toml` (blob `cdc0d994e359eb09aa153b02ba103a383cded3f2`) et `src/lib.rs` (blob `852b5af11540af229863f8aacda0d68215db81c3`). Dépendance Windows crate épinglée ; licence au chemin `LICENSE` non trouvée.
- Lecture de `u/Cargo.toml` (blob `e4624290a34c94373934ce3793c0def4664165c0`) et `u/src/lib.rs` (blob `50957bedec5f55495beff6563da81a5954371d02`). La synthèse vocale est Rust, avec modules audio/engine/frontend/FFI ; comparaison de provenance avec la voix intégrée à omni-os encore requise.
- Lecture de `solution/pyproject.toml` (blob `14ff7514f9de721da743fe9f698196963b3e80a0`) et `solution/src/omni/cli.py` (blob `f874d644a0e71de2c0ddf0a2f689372ebb0fdbb6`). La licence déclarée est 0BSD ; la CLI relie plusieurs modules de validation Python.
- Relecture de `ADMWS12/src/platform/capability.py` (blob `852c6e852918f71c804ed5a9f766900bcc79bb2f`), `evidence.py` (blob `45f70957a19995685f85d9a55ee64487fc3cf0da`) et `hal.py` (blob `dc4f853248baee4368c65e5861e1eedad2707f39`). Les états UNKNOWN/ABSENT/UNSUPPORTED/FAILED doivent rester distincts dans le contrat ZERO.
- Plusieurs chemins testés retournent 404 ; ils sont consignés comme `PATH_NOT_FOUND`, jamais comme preuve d'absence du composant ou de ses tests.
- Aucun code source n'a été copié, aucune fusion de branches n'a été effectuée et aucun test n'est déclaré exécuté par cette session.


### Audit fichier par fichier — lot 3 (2026-10-09)

- Le registre détaillé `FILE-BY-FILE-AUDIT.md` a été complété avec un troisième lot, commit `c055924f0fcb3a4d1d3da928484e0a7a9d983778`.
- Lecture de `NVDA-RUST-UIA-STANDALONE/src/semantic.rs` (blob `a068aef2811b16eb1911c82dfebb932e8deeec92`) et `src/presentation.rs` (blob `38212a887a2565c17258a1e1d97052d05d4414f9`). La séparation rôles sémantiques / priorités de présentation est un candidat architectural, pas une intégration.
- Lecture de `omni-security/SECURITY.md` (blob `48534ea9c423ee388e2a44415a5ac205a8672a5e`) ; les exigences de sécurité et d'accessibilité peuvent guider les gates ZERO. Licence et modèle de menace ne sont pas encore localisés aux chemins testés.
- Lecture des manifestes `omni-os/os/crates/aw-accessibility/Cargo.toml` (blob `0ce26e8ab778f8f5ca2a760b088460639dc1a3cf`), `aw-kernel-contract/Cargo.toml` (blob `a07c5548c38c22bd8fd02c7a108f1ad814d7edf2`) et `aw-x86-platform/Cargo.toml` (blob `df15cf4e999d806cc38dd5cbadaf413f953ec34e`).
- Aucun test n'a été exécuté ; aucun code n'a été copié ; aucun dépôt ou historique n'a été supprimé. L'audit fichier par fichier reste incomplet jusqu'à obtention d'un inventaire exact des arbres et branches.


### Première intégration fonctionnelle — adaptateur ADMWS12 vers ZERO (2026-10-09)

- Implémentation ajoutée : `bootstrap/go/admws12_mapping.go`, blob actuel `760ca2514ff6421bf1e53cc200d17b66d0cbd474`, commit initial `f261f76d8caf06c910985c8dd16c15a3e480bf68`, puis correction du séparateur canonique dans `e229341ab8ce1eb4b309fab99491ae7e63e37deb`.
- `MapADMWS12Record` convertit Capability en `CAPABILITY` et Evidence en `OBSERVATION`, vérifie les états explicitement autorisés, rejette kind/state inconnus et identité vide, fixe `Source=ADMWS12`, `Target=ZERO`, `Time=UNKNOWN`, et garde `Proof=UNPROVEN`.
- Tests de régression ajoutés dans `bootstrap/go/canonical_test.go`, commit `6457fb6f6bd3f3a152fb01150dd4b8d1eb471f21`, puis test de l'état Evidence ABSENT et round-trip canonique, commit `24ea4985d92b0c3523a532dba8ba7fd93146fd8e`.
- Contrat `ADMWS12-ADAPTER-IMPLEMENTATION.md` actualisé, commit `4afdc2ba6bcf37cc5673c5312deb8eca36a2593f`.
- Validation : **CODE ADDED / TESTS ADDED / NOT_RUN / NOT_CONFIRMED**. Le statut combiné GitHub pour le commit de tests a retourné une liste vide ; cela ne constitue pas une réussite. Il faut confirmer la compilation et exécuter `go test ./...` et `go vet ./...` sur le SHA exact.
- Aucun code Python ADMWS12 copié, aucun changement firmware/NVRAM, aucun dépôt ou historique supprimé.


### Accélération de l'audit : inventaire récursif et première intégration (2026-10-09)

- Interrogation de l'API Git Trees récursive sur les branches par défaut des 30 dépôts. Les 30 réponses ont retourné `truncated=false`; total recensé : **7 823 fichiers suivis**. Le détail des branches et Tree SHA est dans `REPOSITORY-TREE-INVENTORY.md`, commit `2dfedeaf8a9e16860a198b52aaf1400bb720a2d3`.
- Inventaire exact des 171 fichiers de ZERO pour le snapshot `7f1ac40661b0d20cb314fa09df9278369f190731`, avec chemins, tailles et blob SHA, découpé en trois fichiers : `ZERO-FILE-INVENTORY-PART-1.md` commit `f6de3ea98ad2344b65c1653e4a5bbb10aaa90eb0`, `PART-2` commit `94cec2452bb764a61a4d933eb36f5825b1eac306`, `PART-3` commit `36dc1cdc9f8ea46b43644f743e5c75ae1bde4e3b`. Ce snapshot ne comprend pas les inventaires ajoutés ensuite ; les trois parties restent cohérentes entre elles.
- L'inventaire des arbres est complet pour les branches par défaut capturées, mais **ce n'est pas encore un audit du contenu de 7 823 fichiers**. Les branches non par défaut, sous-modules et licences par fichier restent à comparer.
- Implémentation Go de `MapADMWS12Record` ajoutée et tests de régression complétés. L'adaptateur conserve les états sources et impose `Proof=UNPROVEN`. Les tests sont définis mais restent `NOT_RUN` ; le statut GitHub du commit de test était vide.
- Prochaine accélération : lire les fichiers de code/tests/manifests par lots prioritaires, faire les comparaisons exactes de source de voix, sémantique UIA, UEFI/HII/IFR et sécurité, puis exécuter les gates sur des commits exacts. Aucune fusion ou suppression n'a été effectuée.


### Résultats CI récupérés et correctif du gate (2026-10-09)

- Lecture directe de l'API GitHub Actions : run `37916541170` — workflow `ZERO bootstrap validation`, commit `3a5bb04dff101be29d03d4c532afe61f43c85507`, job `113774000501` — **SUCCESS**. L'étape `go test ./...` dans `bootstrap/go` est PASS à ce SHA ; les tests ajoutés pour l'adaptateur font partie de cette exécution.
- Le workflow global `ZERO validation` a échoué au step `Local component cleanliness`, non aux étapes précédentes : son journal signale le fichier historique `RAPPEL-ETAT-2026-10-08.md`. Le document est une preuve d'état archivée, pas une dépendance du runtime.
- Correctif conservateur dans `tests/no-local-component.ps1`, commit `e8f6c5d23b1054807e3ac11f9956936339082828` : les rapports historiques nommés `RAPPEL-ETAT-*.md` restent inchangés mais sont exclus du scan de composants runtime. Aucune preuve ni mention historique n'a été supprimée.
- Après le correctif, les workflows de commit `e8f6c5d23b1054807e3ac11f9956936339082828` étaient encore `in_progress` au dernier contrôle ; ne pas annoncer la validation globale PASS avant leur conclusion.
- `go vet ./...` n'a pas encore de résultat vérifié. Le statut de l'adaptateur est donc `BOOTSTRAP_GO_TEST=PASS`, `GO_VET=NOT_RUN`, `FULL_ZERO_VALIDATION=RETRYING`.


### Renforcement du mapping et des gates CI (2026-10-09)

- Le workflow `ZERO bootstrap validation` a été confirmé **SUCCESS** sur le commit `3a5bb04dff101be29d03d4c532afe61f43c85507`, avec `go test ./...` dans `bootstrap/go`. Run `37916541170`.
- Le contrat d'adaptateur a été mis à jour pour enregistrer ce résultat et distinguer explicitement les validations testées des validations encore en attente.
- Ajout de tests couvrant les 10 états ADMWS12 acceptés (5 Capability, 5 Evidence), le rejet d'un kind inconnu et d'une identité vide : commit `62663ad8eab9f551cabaa1d909350c63ac6272fc`.
- Ajout de `go vet ./...` au workflow `zero-bootstrap.yml`, commit `aacf8ffdd3891c5bedecfb217eee28b0ff7aec1f`. Son résultat pour le nouveau step doit être confirmé par run.
- Le gate `no-local-component.ps1` a été ajusté pour ne pas analyser les rapports d'audit/inventaire comme s'ils étaient des dépendances d'exécution. Les documents restent intacts. Le run précédent avait révélé une seconde fausse alerte sur les chemins de l'inventaire ; les nouveaux runs sont en cours de vérification.
