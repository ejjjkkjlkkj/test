# ZERO — PROPRIÉTÉS SUPÉRIEURES

Objectif: dépasser le socle I0–I10 sans confondre expérimentation et preuve.

## A. Autorité

P01 — MONOTONIC AUTHORITY
Une transition ne peut qu'utiliser ou déléguer une autorité explicitement présente;
elle ne peut jamais créer silencieusement un privilège.

P02 — COMPLETE MEDIATION
Toute lecture, écriture, exécution, communication ou effet externe passe par une
frontière contrôlée. Aucun chemin de contournement implicite n'est valide.

P03 — NON-FORGEABILITY
Une représentation de capability invalide, dupliquée, altérée ou hors domaine est
rejetée sans effet partiel.

## B. État

P04 — HISTORY IMMUTABILITY
L'historique accepté ne peut être réécrit par une transition future.

P05 — ATOMICITY
Une transition échouée laisse l'état observable équivalent à l'état pré-transition.

P06 — CRASH CONSISTENCY
Après arrêt arbitraire, redémarrage et reprise, aucun état intermédiaire interdit ne
devient observable.

P07 — INSTANCE ISOLATION
Deux ZeroCore indépendants ne partagent aucun état observable ou mutable implicite.

## C. Information

P08 — NON-INTERFERENCE
Modifier une entrée secrète ne doit pas modifier les observations publiques autorisées.

P09 — UNKNOWN PRESERVATION
UNKNOWN reste distinct de toute valeur concrète jusqu'à preuve d'information.

P10 — PROVENANCE
Toute valeur promue conserve l'origine et les transformations nécessaires à son audit.

P11 — NO SILENT LOSS
Une transformation ne peut supprimer une information déclarée nécessaire sans
produire une preuve ou un verdict de perte explicite.

## D. Temps / concurrence

P12 — REPLAY DETERMINISM
Même état initial + mêmes entrées + mêmes autorités + même politique = même relation
de transitions autorisées.

P13 — TEMPORAL ISOLATION
Une action d'un domaine ne peut influencer un autre domaine avant l'instant où une
relation d'autorité le permet.

P14 — CONCURRENCY SAFETY
Toutes les interleavings autorisées par le modèle préservent les invariants.

## E. Ressources

P15 — RESOURCE NON-AMPLIFICATION
Une entrée de coût borné ne peut produire un coût non borné sans franchir une
frontière explicitement autorisée.

P16 — FAIR TERMINATION
L'épuisement de ressources produit un état défini et sûr; jamais un comportement
indéfini.

## F. Effets externes

P17 — EFFECT CONTAINMENT
Chaque effet CPU/mémoire/GPU/écran/audio/stockage/réseau/périphérique est attaché à
une capability et à un domaine.

P18 — DEVICE IDENTITY BINDING
Un périphérique externe ne peut être traité comme l'identité attendue sans preuve
d'attachement au canal autorisé.

P19 — NETWORK CONFINEMENT
Une communication ne peut atteindre qu'un pair, protocole et volume explicitement
autorisés.

## G. Compilation / machine

P20 — SEMANTIC REFINEMENT
L'implémentation ne peut produire un comportement interdit par la sémantique ZERO.

P21 — COMPILER PRESERVATION
Chaque transformation du compilateur conserve les propriétés pertinentes de la
sémantique source.

P22 — ISA BOUNDARY
Les hypothèses sur CPU, mémoire, DMA, interruptions et privilèges sont explicites;
aucune hypothèse implicite n'entre dans la preuve.

## H. Accessibilité

P23 — MODALITY CONSERVATION
Une information perceptible par une modalité source doit rester représentable dans
le modèle ZERO sans dépendre d'une modalité particulière.

P24 — SEMANTIC EQUIVALENCE
Texte, image, graphique, structure, événement et état doivent converger vers une
représentation sémantique commune lorsque leurs informations sont équivalentes.

P25 — ACCESSIBILITY NON-INTERFERENCE
L'activation du lecteur, de la parole, du braille ou du clavier ne doit pas détruire
les propriétés de sécurité du contenu sous-jacent.

P26 — UNKNOWN ACCESSIBILITY
Une information inconnue ou ambiguë reste annoncée comme inconnue/ambiguë; le moteur
ne fabrique pas une certitude pour rendre l'interface plus simple.

## I. Auto-extension

P27 — QUARANTINED SELF-EXTENSION
Une primitive découverte reste hors du noyau tant que sa spécification, sa preuve et
son test adversarial ne sont pas acceptés.

P28 — PROMOTION MONOTONICITY
Une propriété déjà prouvée ne peut pas devenir moins sûre par l'ajout d'une primitive.

P29 — SELF-MODIFICATION CONTAINMENT
Une expérience peut modifier son espace expérimental, jamais le noyau de confiance,
sans passer par le même gate de preuve.

## J. Résilience

P30 — FAULT CONTAINMENT
Une faute locale ne doit pas devenir une faute globale sans franchir une frontière
explicitement modélisée.

P31 — RECOVERY CONVERGENCE
Toute reprise autorisée converge vers un état sûr défini ou vers un arrêt sûr.

P32 — FAIL-CLOSED
En cas d'incertitude critique, de preuve manquante, de violation de budget ou de
désaccord de vérificateurs: aucune autorité nouvelle n'est accordée.

## K. Preuve

P33 — ARTIFACT BINDING
Un résultat de preuve est lié au digest exact de la spécification, du modèle et du
code qu'il certifie.

P34 — INDEPENDENT RECHECK
Les résultats critiques doivent pouvoir être vérifiés par un second chemin
indépendant.

P35 — COMPOSITION CLOSURE
Si deux composants sont prouvés compatibles selon un contrat de composition, leur
composition conserve les propriétés déclarées.

## Règle de statut

UNTESTED < OBSERVED < STRESS-TESTED < EXHAUSTIVE-FINITE < REPRODUCED < PROVEN.

Aucun statut inférieur ne peut être converti implicitement en PROVEN.

## Objectif final

Le projet ne cherche plus seulement une énorme quantité de tests. Il cherche à
transformer les propriétés en invariants universels, puis à relier ces invariants
du langage jusqu'aux effets machine.
