# ZERO — langage machine universel (spécification fondatrice)

Statut : **SPECIFICATION EN CONSTRUCTION — pas encore une preuve d'implémentation universelle**  
Nom de la base : **0 (ZERO)**

## 1. Mission non négociable

Construire dans le projet **0** un langage et une base d'exécution universels qui puissent décrire, compiler, vérifier et exécuter des opérations sur des matériels et logiciels différents, du bare metal aux systèmes hôtes, sans supposer un seul processeur, système d'exploitation ou fournisseur.

« Universel » est un objectif d'interopérabilité démontré par des cibles et tests concrets, pas une propriété à déclarer sans preuve.

## 2. Objectif de conception : dépasser les références sonores et lumineuses par le résultat mesuré

Le but ambitieux de 0 est de concevoir les calculs et les communications pour obtenir le résultat utile avant une référence physique définie, lorsque cela est possible. Pour chaque expérience, le projet doit choisir et publier une distance, un milieu, une tâche, un matériel et une méthode de mesure.

Le terme « dépasser la vitesse du son/de la lumière » doit être traduit en tests falsifiables distincts :

1. **Référence acoustique** : comparer le temps de calcul ou de réponse au temps qu'un son mettrait à parcourir une distance définie dans un milieu et à une température indiqués.
2. **Référence lumineuse** : comparer le temps de calcul ou de réponse au temps qu'un signal lumineux mettrait à parcourir une distance de référence définie. Préciser si la comparaison utilise le vide ou un milieu donné.
3. **Performance logicielle** : comparer la latence, le débit, le travail utile par joule et les temps de bout en bout à une baseline équivalente.
4. **Propagation réelle** : mesurer séparément la propagation du signal dans le support physique, sans la confondre avec le temps de calcul ou avec une comparaison de benchmark.

Cela permet de démontrer, par exemple, qu'un calcul local produit son résultat avant qu'un signal ait le temps de parcourir la distance de référence choisie. C'est un objectif expérimental réel, mais ce n'est pas la transmission d'information plus vite que la lumière dans le vide.

La vitesse de la lumière dans le vide est exactement 299 792 458 m/s dans le SI. 0 ne doit pas revendiquer une violation de cette limite physique. Il doit en revanche chercher, mesurer et publier toute amélioration légitime qui permet d'obtenir un résultat plus tôt que les systèmes comparés ou qu'une référence de propagation explicitement définie.

## 3. Principes pour gagner du temps par conception

- **Calcul près des données** : éviter de déplacer des données quand le calcul local est moins coûteux.
- **Moins de travail** : choisir de meilleurs algorithmes et éliminer les calculs redondants.
- **Parallélisme explicite** : décrire les dépendances afin que les opérations indépendantes puissent être exécutées en parallèle.
- **Latence bornée quand possible** : distinguer les opérations déterministes des opérations dont le temps dépend du système, du réseau ou du matériel.
- **Compilation par cible** : produire du code adapté à l'ISA, à l'ABI, à la mémoire et aux capacités réellement détectées.
- **Communication minimale** : réduire les allers-retours, copies, allocations, attentes et synchronisations.
- **Prédiction vérifiable** : la spéculation ou le pré-calcul ne compte comme réussite que si la réponse finale est correcte et le coût total est mesuré.
- **Pas de faux raccourcis** : les caches, résultats pré-calculés, jeux de données réduits et simulations doivent être déclarés dans les résultats.

## 4. Architecture cible de 0

- **ZERO-IR** : représentation intermédiaire stable, typée, versionnée et indépendante de l'ISA.
- **ZERO-Core** : sémantique minimale, déterministe lorsque requis, avec erreurs et effets explicites.
- **ZERO-Backend** : générateurs séparés pour chaque ISA/ABI/plateforme prise en charge.
- **ZERO-BareMetal** : démarrage et accès matériels propres à chaque cible, sans supposer que toutes les cartes partagent les mêmes registres ou firmwares.
- **ZERO-Host** : intégration aux systèmes d'exploitation et aux logiciels existants par des interfaces explicites.
- **ZERO-Verify** : tests différentiels, tests de propriétés, fuzzing, vérification des invariants et comparaison de sorties entre backends.
- **ZERO-Perf** : bancs de mesure reproductibles, avec baselines, distributions de latence, débit, consommation lorsque mesurable et limites documentées.
- **ZERO-Access** : interfaces sémantiques et pilotables au clavier, conçues pour les technologies d'assistance.
- **ZERO-Bridge** : adaptateurs d'import/export pour réutiliser les composants existants sans effacer leur provenance ni leur licence.

## 5. Contrat de preuve

Chaque capacité reçoit un statut explicite :

- **PASS** : preuve reproductible jointe et critères satisfaits.
- **FAIL** : test exécuté et critère non satisfait.
- **PARTIAL** : une partie seulement est démontrée.
- **BLOCKED** : dépendance ou environnement empêche le test.
- **UNPROVEN** : affirmation non encore démontrée.
- **NOT_RUN** : test non exécuté.

Chaque rapport de performance doit enregistrer au minimum : commit exact, version des outils, matériel réel ou émulation/VM, système, options de compilation, corpus et son hash, commandes exactes, nombre de répétitions, échauffement, p50/p95/p99/max, débit, consommation si mesurée, résultats bruts et comparaison à une baseline équivalente.

Pour les références de propagation, enregistrer aussi la distance, le milieu, la température lorsque pertinente, la vitesse de référence retenue, le temps théorique calculé, le temps mesuré, l'incertitude instrumentale et les sources d'erreur. Les résultats QEMU/VM doivent être étiquetés comme tels et ne prouvent pas une performance bare metal. Un benchmark logiciel ne prouve pas à lui seul une vitesse de propagation physique.

## 6. Étapes de réalisation

1. Écrire la grammaire initiale et la sémantique de ZERO-IR.
2. Créer un parseur, un validateur et une suite de tests de conformité.
3. Compiler un petit ensemble d'opérations vers une première cible documentée.
4. Ajouter un interpréteur de référence afin de comparer les résultats avec le backend compilé.
5. Créer un premier banc reproductible comparant temps de calcul et temps de propagation théorique sur des distances définies.
6. Produire des tests croisés et conserver les résultats bruts dans le dépôt.
7. Ajouter les backends et plateformes un par un, chacun avec une matrice de capacités et des preuves.
8. Auditer les dépôts sources, conserver les hashes, licences, chemins et origines; n'intégrer que par des adaptateurs traçables.
9. Mesurer les performances sur matériel identifié, puis optimiser le chemin démontré le plus coûteux.

## 7. Règles de non-régression et d'intégrité

- Ne jamais marquer une cible « prise en charge » avant un test de conformité reproductible.
- Ne jamais confondre inventaire, compilation, test réussi, mesure et preuve physique.
- Ne jamais supprimer ou écraser un composant source pendant l'unification sans décision explicite, historique et vérification des licences.
- Ne jamais intégrer de secrets, jetons ou clés dans les fichiers, journaux ou artefacts.
- Préserver la provenance de chaque fichier importé : dépôt, branche, commit, chemin, blob SHA-256 ou Git blob SHA, licence et statut de vérification.
- Séparer strictement les opérations en lecture seule des opérations modifiant le matériel ou le firmware.

## 8. Critère d'acceptation de la fondation

La fondation 0 n'est validée que lorsque la grammaire, la sémantique, le parseur/validateur, les tests de conformité et au moins une chaîne de compilation de bout en bout sont présents, exécutés et accompagnés de résultats reproductibles. Chaque revendication de performance doit être attachée à une expérience reproductible. L'universalité totale et les objectifs de performance restent **UNPROVEN** jusqu'à ce que les cibles correspondantes soient effectivement testées.
