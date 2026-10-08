# CHANGELOG

Toutes les modifications importantes de ZERO sont documentées ici.

## 0.1.0-rc.1

### Validation

- Bootstrap Go validé par CI.
- Cœur Go validé par CI.
- Validation native ZERO validée par CI.
- Validations de format, vérité, exécution, machine, moteur de référence, support physique, portabilité de vérité et accessibilité validées par CI.
- Contrôle de propreté des composants locaux intégré à la CI.
- Workflow de release avec contrôle exact de `VERSION` et du tag.

### Architecture

- Représentation canonique RECORD déterministe.
- Encodage et décodage binaire ZERO.
- ISA minimale et mémoire bornée.
- Séparation entre autorisation ISA et autorisation sémantique.
- Séparation entre sémantique et transport.
- Projections d'accessibilité conservant la même sémantique.
- Échelle de preuve explicite, sans promotion implicite vers le matériel ou les liaisons physiques.

### Limites connues

Cette version candidate ne constitue pas une preuve de fonctionnement matériel, RF, radio, satellite ou de liaison physique réelle.
