# ZERO — langage machine universel (spécification fondatrice)

Statut : **SPECIFICATION EN CONSTRUCTION — pas encore une preuve d'implémentation universelle**  
Nom de la base : **0 (ZERO)**

## 1. Mission non négociable

Construire dans le projet **0** un langage et une base d'exécution universels qui puissent décrire, compiler, vérifier et exécuter des opérations sur des matériels et logiciels différents, du bare metal aux systèmes hôtes, sans supposer un seul processeur, système d'exploitation ou fournisseur.

« Universel » est un objectif d'interopérabilité démontré par des cibles et tests concrets, pas une propriété à déclarer sans preuve.

## 2. Sens précis de « plus rapide que le son ou la lumière »

Le projet doit viser les meilleures performances mesurables et publier les preuves. Il faut distinguer trois grandeurs :

1. **Temps de calcul** : cycles, instructions, latence et débit d'un programme.
2. **Temps de propagation** : durée nécessaire à un signal pour parcourir une distance donnée dans un milieu.
3. **Temps de bout en bout** : calcul + mémoire + E/S + réseau + propagation.

0 peut viser un temps de calcul plus court que celui d'une implémentation de référence, un débit supérieur et une latence inférieure sur un test défini. Il peut aussi exploiter le parallélisme, la localité, la compilation spécialisée et des algorithmes meilleurs.

Aucun rapport de benchmark ne doit prétendre qu'un logiciel transmet une information plus vite que la lumière dans le vide ou qu'un signal matériel dépasse localement la vitesse du son sans une expérience physique appropriée. La vitesse de la lumière dans le vide est exactement 299 792 458 m/s dans le SI. Une performance logicielle exceptionnelle ne constitue pas une violation de cette limite physique. Si « plus vite » désigne un résultat de calcul, une latence, un débit ou un temps de réaction, le rapport doit nommer précisément la métrique.

## 3. Architecture cible de 0

- **ZERO-IR** : représentation intermédiaire stable, typée, versionnée et indépendante de l'ISA.
- **ZERO-Core** : sémantique minimale, déterministe lorsque requis, avec erreurs et effets explicites.
- **ZERO-Backend** : générateurs séparés pour chaque ISA/ABI/plateforme prise en charge.
- **ZERO-BareMetal** : démarrage et accès matériels propres à chaque cible, sans supposer que toutes les cartes partagent les mêmes registres ou firmwares.
- **ZERO-Host** : intégration aux systèmes d'exploitation et aux logiciels existants par des interfaces explicites.
- **ZERO-Verify** : tests différentiels, tests de propriétés, fuzzing, vérification des invariants et comparaison de sorties entre backends.
- **ZERO-Perf** : bancs de mesure reproductibles, avec baselines, distributions de latence, débit, consommation lorsque mesurable et limites documentées.
- **ZERO-Access** : interfaces sémantiques et pilotables au clavier, conçues pour les technologies d'assistance.
- **ZERO-Bridge** : adaptateurs d'import/export pour réutiliser les composants existants sans effacer leur provenance ni leur licence.

## 4. Contrat de preuve

Chaque capacité reçoit un statut explicite :

- **PASS** : preuve reproductible jointe et critères satisfaits.
- **FAIL** : test exécuté et critère non satisfait.
- **PARTIAL** : une partie seulement est démontrée.
- **BLOCKED** : dépendance ou environnement empêche le test.
- **UNPROVEN** : affirmation non encore démontrée.
- **NOT_RUN** : test non exécuté.

Chaque rapport de performance doit enregistrer au minimum : commit exact, version des outils, matériel réel ou émulation/VM, système, options de compilation, corpus et son hash, commandes exactes, nombre de répétitions, échauffement, p50/p95/p99/max, débit, consommation si mesurée, résultats bruts et comparaison à une baseline équivalente. Les résultats QEMU/VM doivent être étiquetés comme tels et ne prouvent pas une performance bare metal.

## 5. Étapes de réalisation

1. Écrire la grammaire initiale et la sémantique de ZERO-IR.
2. Créer un parseur, un validateur et une suite de tests de conformité.
3. Compiler un petit ensemble d'opérations vers une première cible documentée.
4. Ajouter un interpréteur de référence afin de comparer les résultats avec le backend compilé.
5. Produire des tests croisés et conserver les résultats dans le dépôt.
6. Ajouter les backends et plateformes un par un, chacun avec une matrice de capacités et des preuves.
7. Auditer les dépôts sources, conserver les hashes, licences, chemins et origines; n'intégrer que par des adaptateurs traçables.
8. Mesurer les performances sur matériel identifié, puis optimiser le chemin démontré le plus coûteux.

## 6. Règles de non-régression et d'intégrité

- Ne jamais marquer une cible « prise en charge » avant un test de conformité reproductible.
- Ne jamais confondre inventaire, compilation, test réussi, mesure et preuve physique.
- Ne jamais supprimer ou écraser un composant source pendant l'unification sans décision explicite, historique et vérification des licences.
- Ne jamais intégrer de secrets, jetons ou clés dans les fichiers, journaux ou artefacts.
- Préserver la provenance de chaque fichier importé : dépôt, branche, commit, chemin, blob SHA-256 ou Git blob SHA, licence et statut de vérification.
- Séparer strictement les opérations en lecture seule des opérations modifiant le matériel ou le firmware.

## 7. Critère d'acceptation de la fondation

La fondation 0 n'est validée que lorsque la grammaire, la sémantique, le parseur/validateur, les tests de conformité et au moins une chaîne de compilation de bout en bout sont présents, exécutés et accompagnés de résultats reproductibles. L'universalité totale et les objectifs de performance restent **UNPROVEN** jusqu'à ce que les cibles correspondantes soient effectivement testées.
