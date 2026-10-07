# ZERO — architecture de transport indépendante

## Statut

DEFINED. Cette spécification sépare le sens d'un message de son moyen de transport et de sa livraison.

## 1. Chaîne normative

TRUTH → MESSAGE → TRANSPORT → DELIVERY → OBSERVATION → PROOF

Ces niveaux ne sont pas interchangeables.

- TRUTH : règle ou état sémantique défini.
- MESSAGE : représentation structurée du sens à transmettre.
- TRANSPORT : mécanisme utilisé pour déplacer le message.
- DELIVERY : résultat de la tentative de remise.
- OBSERVATION : fait effectivement constaté.
- PROOF : évidence permettant d'attribuer un niveau de preuve.

## 2. Transports représentables

Le modèle peut représenter, lorsque les capacités existent :
SOFTWARE
MESH
DTN
I2P
RADIO
SATELLITE
D2D

Cette liste décrit des classes de transport, pas des connexions actuellement disponibles.

## 3. Invariance sémantique

Pour un message M et deux transports T1 et T2, l'encodage doit préserver le même sens sémantique de M.

Un changement de transport ne peut pas modifier silencieusement :
- l'identité de la source ;
- l'identité de la destination ;
- le type du message ;
- le contenu sémantique ;
- la séquence ;
- l'intégrité déclarée ;
- l'état de preuve.

## 4. Échec de livraison

États minimaux :
CREATED → QUEUED → SENT → DELIVERED
ou
CREATED → QUEUED → FAILED

Une absence de connectivité produit un état observable tel que UNAVAILABLE, DEFERRED ou FAILED. Elle ne produit jamais DELIVERED par défaut.

## 5. DTN / stockage-retransmission

Un transport de type DTN peut conserver un message et tenter sa retransmission ultérieurement. Cela modifie l'état de livraison, pas le sens du message.

MESSAGE ≠ DELIVERY

## 6. Dark / disconnected

Un environnement intermittent ou déconnecté est un état du transport. Le runtime doit pouvoir conserver message, identité, séquence, timestamp, état de livraison et preuve disponible.

La reconnexion permet une nouvelle tentative sans réécrire rétroactivement l'historique.

## 7. Satellite / radio

RADIO et SATELLITE sont des catégories de réalisation physique ou de liaison. Leur simple représentation dans ZERO ne constitue pas une preuve RF ou satellite.

Niveaux de preuve :
DEFINED → IMPLEMENTED → TESTED → SIMULATED → HARDWARE → RF_PROVEN → SATELLITE_LINK_PROVEN

## 8. D2D et mesh

Une topologie multi-sauts peut transmettre un même message via plusieurs nœuds. Les nœuds intermédiaires appartiennent au chemin de transport ; ils ne deviennent pas automatiquement l'auteur du message.

MESSAGE.SOURCE != NEXT_HOP

## 9. I2P

I2P peut être représenté comme une réalisation de transport lorsque ses capacités sont effectivement disponibles. La présence d'une abstraction I2P ne prouve ni l'accès au réseau I2P ni la disponibilité d'un routeur.

## 10. Règle d'accessibilité

Le résultat de transport reste un objet sémantique accessible par VOICE, BRAILLE, KEYBOARD, DISPLAY, TOUCH, NETWORK ou AUTOMATION. La projection peut varier ; le sens ne varie pas.

## 11. Test minimal

Pour un même message M :
1. construire M ;
2. sélectionner plusieurs transports ;
3. encoder M ;
4. tester succès, panne et différé ;
5. vérifier que le message reste sémantiquement identique ;
6. vérifier que DELIVERED n'est obtenu que par une observation de livraison correspondante.

## 12. Limite de preuve

Cette spécification établit une architecture et des invariants testables. Elle ne prétend pas établir une communication radio, D2D, I2P ou satellite réelle sans observation physique correspondante.
