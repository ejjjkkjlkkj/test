# ZERO — Matrice de validation

Chaque version doit être testée sur plusieurs frontières.

| Domaine | Validation |
|---|---|
| CPU | observation, transition, interruption |
| mémoire | lecture, écriture autorisée, refus |
| GPU | découverte, état, capacité |
| moniteur | état, projection, changement |
| audio | entrée, sortie, interruption |
| clavier | événement, intention |
| stockage | identité, état, accès |
| réseau | découverte, identité, session |
| machine distante | capacités, événements, refus |
| texte | structure, navigation |
| image | scène, relations, incertitude |
| graphique | axes, séries, valeurs |
| tableau | lignes, colonnes, relations |
| vidéo | temps, événements |
| web | navigation, interaction |
| CAPTCHA | tâche, règles, résultat |
| inconnu | existence, observation, incertitude |
| accessibilité | multimodalité sans perte de sens |
| sécurité | aucun contournement par projection |

## Condition globale

Un seul échec d'un invariant fondamental bloque la validation de la version.
