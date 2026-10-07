# ZERO — IA locale utilisable

Le dépôt fournit maintenant un client Ollama local sans dépendance externe.

## Pré-requis Windows

- Go 1.22+
- Ollama installé et lancé
- un modèle Ollama disponible localement

Le client utilise par défaut :

- URL : http://127.0.0.1:11434
- modèle : qwen3.8-coder:latest

Ces valeurs sont surchargeables avec `ZERO_AI_HOST` et `ZERO_AI_MODEL`.

## Construire

Depuis `bootstrap/go` :

```powershell
go test ./...
go build -o zero-ai.exe ./cmd/zero-ai
```

## Utiliser

Prompt unique :

```powershell
.\zero-ai.exe --prompt "Analyse le dépôt et liste uniquement les éléments non prouvés."
```

Mode interactif :

```powershell
.\zero-ai.exe
```

Autre modèle :

```powershell
.\zero-ai.exe --model "qwen2.5-coder:14b"
```

Autre serveur compatible API Ollama :

```powershell
$env:ZERO_AI_HOST="http://127.0.0.1:11434"
$env:ZERO_AI_MODEL="qwen2.5-coder:14b"
.\zero-ai.exe
```

## Limites

Cette intégration donne une IA locale opérationnelle pour le dialogue et l'analyse. Elle n'autorise pas automatiquement l'IA à exécuter des opérations ZERO : l'autorisation et l'exécution restent dans le noyau.

Le statut de preuve de l'IA est donc :

- client local : IMPLEMENTED
- transport Ollama : IMPLEMENTED
- tests du client : TESTED
- exécution locale sur la machine utilisateur : à valider sur la machine avec Ollama et le modèle choisi
