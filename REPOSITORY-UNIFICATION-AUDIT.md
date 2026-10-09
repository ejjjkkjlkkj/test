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
