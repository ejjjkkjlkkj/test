# ZERO — Lot d'audit automatisé 01 : arbres récursifs et adaptateur ADMWS12

Date : 2026-10-09
Cible : `ejjjkkjlkkj/test` (ZERO), branche `main`
Statut global : **PARTIAL**. Inventaire exact de la branche par défaut pour les dix dépôts ci-dessous ; revue sémantique fichier par fichier et audit de toutes les branches encore incomplets.

## 1. Inventaire récursif vérifié

Source : endpoint Git Trees GitHub avec `recursive=1`. Tous les arbres ci-dessous ont retourné `truncated=false`. Les SHA sont ceux des arbres observés pendant ce lot.

| Dépôt | Branche examinée | SHA arbre | Entrées | Fichiers | Dossiers | Statut |
|---|---|---|---:|---:|---:|---|
| `ejjjkkjlkkj/test` (ZERO) | main | `214a0a8e94948ab6b7fbeba850456db67e9f8f2f` | 189 | 175 | 14 | INVENTORIED |
| `ejjjkkjlkkj/ADMWS12` | main | `a7f38d456d05b35d962f217264aabb594ed04e3c` | 84 | 61 | 23 | INVENTORIED |
| `ejjjkkjlkkj/omni-os` | main | `b8db99ec0e90c41f7bb757eac67c3a516ec41bc6` | 1560 | 1241 | 319 | INVENTORIED |
| `ejjjkkjlkkj/accessible-windows` | main | `d809e75d5aeb1260c143bf013539bdf9a12f38d6` | 7 | 4 | 3 | INVENTORIED — arbre principal minimal, autres branches à auditer |
| `ejjjkkjlkkj/NVDA-RUST-UIA-STANDALONE` | main | `c732a3d88f65cbed34b21e2af6063222ec218061` | 85 | 52 | 33 | INVENTORIED |
| `ejjjkkjlkkj/u` | main | `1b0fbdb676273602ca38583dc4b1cf8da80865b9` | 118 | 89 | 29 | INVENTORIED |
| `ejjjkkjlkkj/project` | main | `1c0852cd6c3135ca7f40bf25785a980c55606d20` | 45 | 35 | 10 | INVENTORIED |
| `ejjjkkjlkkj/solution` | main | `5e8947eba7a6734ee1cbe806f7abdcef782eb985` | 318 | 284 | 34 | INVENTORIED |
| `ejjjkkjlkkj/omni-security` | main | `db62100fafd50f3e3f74895afcc1dfd0bc070b1c` | 71 | 58 | 13 | INVENTORIED |
| `ejjjkkjlkkj/android` | main | `3d708dbc9271e1ccdc493b93a72f177dd83ae5f9` | 111 | 87 | 24 | INVENTORIED |
| **Total** | | | **2588** | **2086** | **502** | **10 arbres complets, branches secondaires non couvertes** |

## 2. Revue concrète ADMWS12 → ZERO

Fichiers sources relus directement dans les dépôts :
- `ADMWS12/src/platform/capability.py` — blob `852c6e852918f71c804ed5a9f766900bcc79bb2f`.
- `ADMWS12/src/platform/evidence.py` — blob `45f70957a19995685f85d9a55ee64487fc3cf0da`.
- `ADMWS12/src/platform/hal.py` — blob `dc4f853248baee4368c65e5861e1eedad2707f39`.
- `ZERO/bootstrap/go/admws12_mapping.go` — blob `760ca2514ff6421bf1e53cc200d17b66d0cbd474`.
- `ZERO/bootstrap/go/canonical_test.go` — blob `a4e751998059a139fb99c1805f08eb5e3d55eda8`.
- `ZERO/.github/workflows/go.yml` — blob `6e74bbc951594dfde78a38b30f0ebbbee84ab489`.

### Constat corrigé

L'adaptateur Go **existe déjà** sur `main` sous la fonction `MapADMWS12Record`. Le précédent statut `NOT_IMPLEMENTED` est obsolète. Il accepte les cinq états de capacité et les cinq états de preuve connus, refuse le type et les états inconnus, refuse une identité vide, et attribue systématiquement `Proof: UNPROVEN`. Les tests correspondants sont définis dans `canonical_test.go`.

Ce contrôle est une inspection du code, pas une exécution. Le workflow `ZERO Go validation` définit `go test ./...` et `go vet ./...` dans `bootstrap/go`, mais ce lot n'a pas exécuté ces commandes et n'a pas obtenu de résultat CI : **TESTS NOT_RUN / UNPROVEN**.

### Point de vigilance technique

Le payload est assemblé en paires séparées par `|`, alors que le format canonique documente aussi ce caractère comme séparateur de RECORD et l'échappement doit être interprété au niveau de l'enveloppe. Les tests couvrent des valeurs contenant `|` et une conversion aller-retour, mais une revue complémentaire doit vérifier si le payload structuré a une grammaire non ambiguë indépendante du découpage du RECORD. Ne pas déclarer le contrat d'intégration complet avant ce contrôle.

## 3. Décisions de conservation

- Aucun dépôt, branche, commit ou fichier source supprimé ou déplacé.
- Aucun code importé automatiquement.
- Les arbres sont des instantanés de branches par défaut ; les branches historiques et de travail doivent être inventoriées séparément.
- Les fichiers binaires volumineux, sous-modules, licences et dépendances transitives nécessitent un contrôle dédié.
- Les résultats source/CI historiques ne valent pas preuve physique. QEMU/VM ne vaut pas matériel réel.

## 4. Prochaine séquence automatisable

1. Générer et versionner le manifeste complet `path + blob SHA + taille + type` pour chaque arbre, sans remplacer les registres existants.
2. Paginer l'inventaire des dépôts et branches accessibles, puis collecter les arbres récursifs de chaque branche.
3. Classer automatiquement les fichiers par manifeste, langage, licence, tests, workflows, sous-modules et taille ; laisser `NOT_REVIEWED` tant qu'un fichier n'a pas été lu ou contrôlé par une règle explicite.
4. Comparer les fichiers candidats par blob SHA, puis diff sémantique si les blobs diffèrent.
5. Exécuter la validation Go par CI et consigner le run exact, commit, jobs et résultats avant de changer les statuts.
6. Continuer l'audit fichier par fichier par priorité d'intégration, sans supprimer les variantes ni déclarer de validation physique non réalisée.

## Limites

Ce lot n'est pas l'audit exhaustif de tous les dépôts : il couvre dix arbres de branche par défaut. Les autres dépôts, branches secondaires, tags, sous-modules, licences fichier par fichier, dépendances et exécutions CI restent à traiter.
