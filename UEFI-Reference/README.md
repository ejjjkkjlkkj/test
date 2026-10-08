# UEFI Reference Collection

Cette branche regroupe les références utiles au développement UEFI/PI et à l'étude de l'accessibilité firmware.

## Déjà présent

- `UEFI-Vocabulaire/UEFI_Spec_Final_2.11.pdf` — UEFI Specification 2.11.

## Références officielles à récupérer localement

### Spécifications UEFI Forum
- UEFI Specification 2.11
- UEFI Shell Specification 2.2
- UEFI Platform Initialization Specification 1.10
- ACPI 6.6 lorsque nécessaire
- UEFI Self-Certification Tests (SCT) lorsque nécessaire

Les spécifications UEFI Forum restent des références officielles. Ne pas les recopier massivement dans ce dépôt sans vérifier leurs conditions de redistribution.

### Implémentation open source TianoCore
- https://github.com/tianocore/edk2
- https://github.com/tianocore-docs/Docs
- https://github.com/tianocore-docs/edk2-UefiDriverWritersGuide
- https://github.com/tianocore-docs/edk2-VfrSpecification
- https://github.com/tianocore-docs/edk2-InfSpecification
- https://github.com/tianocore-docs/edk2-DscSpecification
- https://github.com/tianocore-docs/edk2-DecSpecification
- https://github.com/tianocore-docs/edk2-UniSpecification
- https://github.com/tianocore-docs/EDK_II_Secure_Coding_Guide
- https://github.com/tianocore-docs/Understanding_UEFI_Secure_Boot_Chain

## Zones EDK II prioritaires pour ADMWS12

- `MdePkg/Include/Uefi/`
- `MdePkg/Include/Protocol/`
- HII / IFR / VFR
- `MdeModulePkg/`
- `ShellPkg/`
- `SecurityPkg/`
- `OvmfPkg/`
- `EmulatorPkg/`

EDK II sert ici de référence d'implémentation et de validation conceptuelle, pas de base à copier pour le design propriétaire ADMWS12.

## Règle de validation

Aucune modification BIOS/NVRAM/firmware physique n'est requise par cette collection. Les tests doivent rester séparés entre IMAGE, QEMU, USB et PHYSIQUE.
