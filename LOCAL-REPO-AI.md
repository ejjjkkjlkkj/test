# Local repository AI

Le dépôt fournit maintenant une analyse locale en lecture seule avec Ollama.

## Construire

    powershell
    cd bootstrap/go
    go test ./...
    go build -o zero-ai-repo.exe ./cmd/zero-ai-repo

## Utiliser sous Windows

Depuis la racine du dépôt :

    powershell
    .\\bootstrap\\go\\zero-ai-repo.exe --repo . --prompt "Audit complet : faits, erreurs, éléments non prouvés, risques et prochaines validations."

Le contexte est limité à 256 KiB au total et 64 KiB par fichier. Les répertoires .git, node_modules, vendor et les fichiers binaires courants sont exclus.

Le programme ne modifie aucun fichier et ne donne aucun privilège d'exécution à l’IA. Toute future action doit passer par l’autorisation ZERO existante.
