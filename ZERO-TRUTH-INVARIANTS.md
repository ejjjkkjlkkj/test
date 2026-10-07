# ZERO — invariants de vérité formelle

## 1. Règle absolue

ZERO sépare les vérités formelles des observations, hypothèses et prédictions.

Une opération déterministe définie par un système formel ne peut pas produire simultanément deux résultats différents dans le même contexte formel.

Exemple arithmétique :

`1 + 1 = 2`

Dans l'arithmétique définie par ZERO, le résultat `3` est INVALID.

Une IA, une observation bruitée ou une probabilité ne peut pas remplacer cette règle.

## 2. Hiérarchie

Les catégories suivantes ne sont jamais interchangeables :

`DEFINITION != OBSERVATION != FACT != INFERENCE != PREDICTION != DECISION != ACTION != RESULT != PROOF`

Une hypothèse ne devient pas un fait parce qu'elle possède une confiance élevée.

Une prédiction ne devient pas une preuve parce qu'elle semble probable.

## 3. Déterminisme

Pour un système formel F, un état initial S et une entrée X :

`EXECUTE(F,S,X)`

doit produire un résultat déterminé par les règles de F.

Si deux résultats incompatibles sont produits avec exactement F, S et X identiques, le système doit signaler une contradiction ou une erreur d'implémentation.

Il est interdit de choisir silencieusement un résultat pour résoudre la contradiction.

## 4. Vérification

Toute opération formelle critique doit pouvoir être soumise à une vérification indépendante de son générateur.

Le vérificateur ne doit pas reprendre la conclusion du générateur comme preuve.

Structure minimale :

`INPUT → RULES → RESULT → INDEPENDENT CHECK → VALID / INVALID`

## 5. Contre-exemple

Pour une affirmation A, ZERO peut rechercher un contre-exemple.

Si un contre-exemple satisfait les règles applicables et contredit A :

`A → REFUTED`

L'IA ne peut pas modifier A pour éviter le contre-exemple sans produire une nouvelle définition explicitement versionnée.

## 6. Inconnu

Lorsqu'aucune règle ou observation suffisante ne permet de conclure :

`UNKNOWN`

ou :

`NOT_PROVEN`

est conservé.

ZERO interdit le remplissage silencieux des données manquantes.

## 7. Preuve

Les niveaux de preuve restent séparés :

`DEFINED → IMPLEMENTED → TESTED → SIMULATED → QEMU → HARDWARE → RF_PROVEN → SATELLITE_LINK_PROVEN`

Un niveau supérieur ne peut être attribué par simple déclaration.

## 8. IA

L'IA peut :

- générer une hypothèse ;
- générer un test ;
- détecter une contradiction ;
- proposer une décision ;
- exécuter une opération autorisée ;
- observer le résultat.

L'IA ne peut pas :

- modifier une vérité formelle pour satisfaire son raisonnement ;
- transformer une hypothèse en fait ;
- transformer une simulation en matériel ;
- transformer une absence de preuve en preuve ;
- supprimer une contradiction sans trace.

## 9. Accessibilité

VOICE, BRAILLE, KEYBOARD, DISPLAY, TOUCH, POINTER, NETWORK et AUTOMATION projettent le même résultat sémantique.

Une projection accessible ne peut pas modifier la vérité de l'état source.

## 10. Test minimal obligatoire

Le noyau de validation doit au minimum démontrer :

- une règle valide est acceptée ;
- une règle contradictoire est rejetée ;
- une hypothèse reste une hypothèse ;
- une preuve ne monte pas sans évidence ;
- un résultat inconnu reste inconnu ;
- une contradiction est conservée et signalée.

## 11. Statut

INVARIANTS : DEFINED
EXECUTION : NOT_EXECUTED
HARDWARE : NOT_PROVEN
