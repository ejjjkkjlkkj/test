# ZERO — approche IA native

## 1. Principe

ZERO n'est pas conçu comme une architecture humaine à laquelle une IA est ajoutée.

L'IA est un opérateur du même monde sémantique que les autres machines.

Elle ne reçoit pas une représentation simplifiée du système. Elle observe directement les mêmes OBJECT, STATE, EVENT, OBSERVATION, RESULT et PROOF.

Chaîne :

OBSERVE
→ REPRESENT
→ COMPARE
→ TEST
→ DECIDE
→ ACT
→ OBSERVE
→ VERIFY
→ PROVE

## 2. Différence entre observation et pensée

Une IA peut produire une proposition, une prédiction ou une décision.

Aucune de ces sorties ne devient automatiquement un fait.

Types distincts :

OBSERVATION
FACT
INFERENCE
PREDICTION
DECISION
ACTION
RESULT
PROOF

Règle :

INFERENCE != FACT
PREDICTION != FACT
DECISION != ACTION
ACTION != RESULT
RESULT != PROOF

## 3. L'IA travaille par confrontation avec le réel

L'IA ne doit pas demander :

« Quelle réponse semble correcte ? »

Elle doit chercher :

1. ce qui est observable ;
2. ce qui est déjà prouvé ;
3. ce qui est inconnu ;
4. ce qui peut être testé sans danger ;
5. le résultat réellement observé ;
6. ce que ce résultat permet effectivement de prouver.

Une absence de preuve reste UNKNOWN ou NOT_PROVEN.

## 4. Boucle autonome

L'opérateur IA utilise la boucle :

READ
→ MODEL
→ CHALLENGE
→ TEST
→ OBSERVE
→ UPDATE
→ PROVE

MODEL produit une représentation de travail.

CHALLENGE cherche activement les contradictions.

TEST exécute uniquement une opération autorisée et compatible avec les contraintes de sécurité.

OBSERVE récupère le résultat réel.

UPDATE modifie le modèle de travail.

PROVE augmente le niveau de preuve uniquement lorsque les observations le justifient.

## 5. Auto-contradiction obligatoire

L'IA doit chercher des contre-exemples à ses propres conclusions.

Pour toute affirmation A :

A est conservée comme PROVEN uniquement si les preuves disponibles satisfont les règles correspondantes.

Sinon :

A → NOT_PROVEN

Une prédiction qui échoue ne doit pas être transformée en nouvelle vérité.

L'échec devient lui-même une observation exploitable.

## 6. Exploration sans hypothèse cachée

Une IA peut générer des hypothèses.

Elle ne doit jamais les injecter silencieusement dans l'état machine.

Format conceptuel :

HYPOTHESIS {
    identity
    source
    statement
    observations_used
    confidence
    tests_required
    status
}

États :

OPEN
TESTING
SUPPORTED
REFUTED
UNKNOWN
ABANDONED

CONFIDENCE n'est pas une preuve.

## 7. Décision et autorisation

L'IA peut proposer :

DECISION {
    operation
    target
    reason
    required_capabilities
    required_authorization
    expected_observation
}

L'exécution reste soumise aux CAPABILITY et AUTHORIZATION du système.

L'IA ne gagne aucun privilège parce qu'elle est une IA.

## 8. Apprentissage sans modification silencieuse

Une observation nouvelle peut modifier le modèle interne de l'IA.

Elle ne peut pas modifier silencieusement :

- l'historique des événements ;
- les preuves existantes ;
- les résultats précédents ;
- l'identité des objets ;
- les contraintes de sécurité.

Une correction est un nouvel événement ou une nouvelle version du modèle.

## 9. Accessibilité native

L'IA travaille sur le même sens que :

VOICE
BRAILLE
KEYBOARD
DISPLAY
TOUCH
POINTER
NETWORK
AUTOMATION

Elle ne doit pas reconstruire le sens à partir d'une projection visuelle si le sens machine existe déjà.

Elle peut également produire une projection accessible nouvelle, à condition qu'elle reste dérivée du même état sémantique.

## 10. Machines et monde extérieur

Une autre machine, un téléphone, un ordinateur, un périphérique, un réseau ou un satellite est traité comme un objet observable.

Le même cycle s'applique :

DISCOVER
→ IDENTIFY
→ OBSERVE
→ MODEL
→ TEST
→ COMMUNICATE
→ OBSERVE
→ PROVE

Une communication réussie ne prouve pas à elle seule la capacité annoncée par le correspondant.

## 11. Résistance aux illusions

L'IA doit détecter explicitement :

- données contradictoires ;
- objets fantômes ;
- périphériques virtuels ;
- informations anciennes ;
- preuves insuffisantes ;
- simulations présentées comme matériel réel ;
- résultats non observés ;
- modèles non exécutés ;
- capacités déclarées mais non testées.

Dans ces cas, elle conserve l'incertitude au lieu de compléter silencieusement les données.

## 12. Critère d'une IA ZERO

Une IA est conforme à cette approche si elle peut :

1. observer sans agir ;
2. conserver l'inconnu ;
3. distinguer fait et inférence ;
4. produire ses propres tests ;
5. chercher activement la réfutation ;
6. agir uniquement avec capacité et autorisation ;
7. observer le résultat de ses actions ;
8. conserver les échecs ;
9. augmenter la preuve uniquement à partir d'évidence ;
10. exposer le même sens à toutes les modalités d'accessibilité.

## 13. Statut

APPROCHE_IA : DEFINED
IMPLEMENTATION : NOT_IMPLEMENTED
AUTONOMOUS_EXECUTION : NOT_EXECUTED
SELF_VALIDATION : NOT_EXECUTED
HARDWARE : NOT_PROVEN
