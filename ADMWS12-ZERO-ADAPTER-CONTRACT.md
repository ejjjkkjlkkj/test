# Contrat d'adaptation ADMWS12 → ZERO

Statut : PROPOSITION DE CONTRAT — pas encore implémentée ni testée en intégration.  
Cible : intégrer les modèles logiciels d'ADMWS12 dans ZERO sans créer deux noyaux concurrents.

## 1. Responsabilités

- ZERO possède le format machine canonique, les identités sémantiques, le langage, les transitions, les résultats et le registre de preuves.
- ADMWS12 fournit des modèles réutilisables de capacités, découverte, évidence, état de plateforme et orchestration HAL.
- Un adaptateur traduit les types ADMWS12 en objets ZERO. Il ne transforme jamais une déclaration de disponibilité en preuve d'exécution physique.
- Aucun pilote matériel n'est considéré comme existant simplement parce qu'une interface HAL existe.

## 2. Correspondance des types

| ADMWS12 | ZERO | Règle |
|---|---|---|
| `Capability.name` | `identity` de la capacité + payload sémantique | Le nom seul n'est pas une identité globale ; construire une identité stable avec namespace et version |
| `Capability.state = UNKNOWN` | `UNKNOWN` | État inconnu conservé |
| `Capability.state = AVAILABLE` | `AVAILABLE` | Déclaration d'adaptateur uniquement ; ne vaut ni ACTIVE ni PROVEN |
| `Capability.state = UNAVAILABLE` | `UNAVAILABLE` | Capacité connue mais actuellement indisponible |
| `Capability.state = UNSUPPORTED` | `UNSUPPORTED` ou `UNAVAILABLE` avec raison typée | Garder la distinction explicite dans le payload si le noyau ZERO ne possède pas encore UNSUPPORTED |
| `Capability.state = FAILED` | `FAILED` | Associer l'erreur, la source et l'instant d'observation |
| `EvidenceState.OBSERVED` | `OBSERVATION` + niveau de preuve correspondant à la méthode réelle | OBSERVED ne signifie pas automatiquement HARDWARE |
| `EvidenceState.ABSENT` | observation d'absence | Ne pas conclure UNSUPPORTED sans test couvrant la capacité |
| `EvidenceState.UNSUPPORTED` | capacité non prise en charge avec preuve de la vérification | Conserver méthode et périmètre |
| `EvidenceState.FAILED` | résultat FAILED + observation d'erreur | Conserver cause et récupérabilité |
| `EvidenceState.UNKNOWN` | `UNKNOWN` | Ne pas inventer une valeur |
| `DiscoveryResult` | séquence d'observations/capacités + preuve de source | Rejeter doublons ou noms vides avant publication |
| `PlatformState.READY` | état de plateforme READY | Seulement si les préconditions obligatoires sont satisfaites |
| `PlatformState.DEGRADED` | état DEGRADED + capacités manquantes | La dégradation doit être observable et accessible |
| `PlatformState.FAILED` | état FAILED + cause | Ne pas convertir l'échec en absence |
| `PlatformContext.has(name)` | requête de capacité | True signifie uniquement que le contrat ADMWS12 la considère utilisable, pas que son effet physique a été prouvé |

## 3. Enveloppe de preuve minimale

Toute capacité importée doit porter au minimum :

- `schema_version`
- `identity`
- `capability_name`
- `state`
- `source_repository`
- `source_commit`
- `adapter_id`
- `observed_at` ou `UNKNOWN`
- `method`
- `scope`
- `proof_level`
- `error` (si applicable)

Le format de transport final doit être encodé dans le format canonique ZERO. Cette liste est le contenu sémantique requis, pas une nouvelle grammaire concurrente.

## 4. Séparation des niveaux de preuve

| Niveau | Ce qu'il permet d'affirmer |
|---|---|
| DEFINED | Le contrat est documenté |
| IMPLEMENTED | Le code existe |
| TESTED | Les tests définis ont réussi sur le code exact |
| EXECUTED | Le code a été exécuté dans l'environnement indiqué |
| SIMULATED | Le résultat vient d'une simulation |
| QEMU | Le résultat a été observé dans QEMU |
| HARDWARE | Le résultat a été observé sur matériel physique identifié |
| RF_PROVEN / SATELLITE_LINK_PROVEN | La preuve porte réellement sur le lien radio / satellite indiqué |

Une preuve QEMU ne devient jamais HARDWARE. Une capacité AVAILABLE ne devient jamais PROVEN par simple traduction.

## 5. Contrat d'erreur et d'accessibilité

Toute erreur exposée par l'adaptateur doit préserver :
- code stable ;
- cause ou `UNKNOWN` ;
- opération concernée ;
- état précédent et état observé ;
- capacités requises ;
- possibilité de récupération ;
- preuve/source de l'observation.

Les interfaces vocales, braille, clavier et visuelles sont des projections du même résultat sémantique. Elles ne doivent pas modifier la preuve ni masquer une transition échouée.

## 6. Conditions d'acceptation

L'adaptateur ne passe à PASS que lorsque :

1. les conversions d'états ont des tests exhaustifs ;
2. les états inconnus et non pris en charge sont préservés ;
3. les noms vides et doublons sont rejetés ;
4. le round-trip du format ZERO préserve les données sémantiques ;
5. les preuves de simulation/QEMU/matériel restent séparées ;
6. une erreur de découverte n'est pas transformée en absence ;
7. les tests tournent sur un commit identifié ;
8. le contrat est validé dans ZERO, pas seulement dans ADMWS12.

État actuel : DEFINED — intégration NOT_RUN — validation NOT_RUN.
