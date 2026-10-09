# ZERO — références physiques réelles : son et lumière

Statut : références sourcées, méthode de comparaison définie ; aucune expérience physique de propagation réalisée par ce document.

## 1. Lumière dans le vide

La valeur SI de la vitesse de la lumière dans le vide est **exactement** :

\[
c = 299\,792\,458\ \mathrm{m/s}
\]

Cette valeur est fixée par la définition SI du mètre. Ce n'est donc pas une mesure que le programme ZERO doit ré-estimer. Source primaire : NIST, *Definitions of SI Base Units* et *Meet the Constants*.

Temps de propagation de référence sur une distance d :

\[
t_{lumière} = d / c
\]

Exemples :
- 1 m : environ 3,335640952 ns
- 100 m : environ 333,564095 ns
- 1 km : environ 3,335640952 µs

Ces valeurs représentent le temps de parcours idéal dans le vide. Elles ne comprennent ni l'encodage, ni l'électronique, ni les délais d'interface, ni le traitement, ni le réseau. Dans un matériau, la vitesse de phase/groupe dépend des propriétés du milieu et du signal ; ne pas appliquer la valeur du vide sans qualification.

## 2. Son dans l'air

Le son nécessite un milieu matériel. Sa vitesse dépend notamment de la température et de la composition du gaz. Il n'existe donc pas une valeur unique « vitesse du son » valable dans tous les milieux.

Pour une première comparaison indicative dans l'air sec proche des conditions ordinaires, on peut utiliser l'approximation :

\[
v_{son}(T) \approx 331.3 + 0.606 T\ \mathrm{m/s}
\]

où T est la température de l'air en degrés Celsius. C'est une approximation d'ingénierie, pas une mesure du milieu local. Pour une expérience métrologique, mesurer la température, documenter l'humidité/composition et utiliser une référence validée pour les conditions exactes.

Temps de propagation acoustique de référence sur une distance d :

\[
t_{son} = d / v_{son}(T)
\]

Exemple : à 20 °C, cette approximation donne 343,42 m/s. Sur 1 m, le temps de parcours est environ 2,912 ms. La vitesse dans l'eau, les métaux et les autres milieux est différente.

## 3. Ce que l'expérience ZERO peut démontrer

Le banc `zero-propagation` mesure le temps mur d'une charge de calcul locale, puis le compare aux temps de référence calculés pour une distance choisie.

Il peut produire une preuve de la forme : « sur cette machine et dans ces conditions, cette charge de calcul s'est terminée en X ns, avant le temps de référence Y pour que le son parcoure d mètres dans l'air à la température saisie ». Cela démontre uniquement la comparaison entre le temps de calcul et ce temps de parcours de référence.

Il **ne** démontre pas :
- qu'un signal physique a été mesuré en train de voyager ;
- que ZERO transmet une information plus vite que la lumière ;
- que la vitesse réelle du son dans l'environnement a été mesurée ;
- que les résultats s'appliquent à d'autres ordinateurs ou charges.

Pour mesurer la propagation physique réelle, il faut un montage adapté : émetteur, récepteur, distance mesurée, instrumentation temporelle suffisamment précise, calibration, répétitions et incertitudes. Un chronomètre logiciel ordinaire ne peut pas résoudre directement le trajet lumineux sur un mètre (environ 3,34 ns).

## 4. Contrat d'expérience reproductible

Chaque rapport doit enregistrer :
- commit ZERO et version du programme ;
- système, processeur, fréquence si disponible, mode d'alimentation ;
- distance saisie, milieu, température et formule employée ;
- charge de calcul et nombre d'itérations ;
- temps mesuré, nombre de répétitions et dispersion ;
- temps de propagation de référence et rapport entre les deux temps ;
- statut clair : `COMPUTE_BEFORE_REFERENCE`, `COMPUTE_NOT_BEFORE_REFERENCE` ou `INVALID_INPUT`.

Ne jamais nommer ce test « preuve de vitesse superluminale ». Ne jamais confondre un résultat de calcul antérieur à un temps de propagation de référence avec un signal transmis plus vite que la lumière.

## Sources de référence

1. NIST, *Definitions of SI Base Units* — définition du mètre et valeur exacte de c : https://www.nist.gov/si-redefinition/definitions-si-base-units
2. NIST, *Meet the Constants* — vitesse de la lumière dans le vide : https://www.nist.gov/si-redefinition/meet-constants
3. OpenStax, *Physics, section 14.1: Speed of Sound, Frequency, and Wavelength* — dépendance de la vitesse acoustique au milieu et exemples de matériaux : https://openstax.org/books/physics/pages/14-1-speed-of-sound-frequency-and-wavelength

Les liens et les sources doivent être vérifiés à nouveau avant de réutiliser ces chiffres pour une expérience de laboratoire formelle.
