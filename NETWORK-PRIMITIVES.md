# ZERO — Primitives réseau minimales

Le réseau doit être représenté par des primitives sémantiques plutôt que par une dépendance à un protocole particulier.

## Noyau conceptuel

DISCOVER(scope)
OBSERVE(remote)
IDENTIFY(remote)
CONNECT(remote)
AUTHENTICATE(remote)
SEND(remote, intent)
RECEIVE()
SUBSCRIBE(remote, event)
DISCONNECT(remote)

Ces opérations ne constituent pas encore une syntaxe finale.

Elles définissent les capacités minimales que la machine-langage doit pouvoir exprimer.

## Règle

Une primitive réseau doit produire un résultat sémantique.

Elle ne doit pas simplement retourner :

SUCCESS / ERROR

mais pouvoir exposer :

RESULT
+ SOURCE
+ TARGET
+ STATE
+ EVENT
+ CAPABILITIES
+ UNCERTAINTY

## Transport

Le transport peut être remplacé ou étendu.

ZERO ne doit pas devenir dépendant d'un seul protocole pour définir son modèle mental.

La découverte, l'identité, l'autorisation, l'événement et la transition sont les concepts stables.

Le protocole est une réalisation.
