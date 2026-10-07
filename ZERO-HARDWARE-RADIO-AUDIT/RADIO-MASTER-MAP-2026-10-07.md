# ZERO HARDWARE RADIO — CARTOGRAPHIE ET PROTOCOLE COMPLET
Date de travail : 2026-10-07

## 1. But
Transformer l'audit Windows en preuve radio structurée, sans confondre matériel réel, pilote, interface, virtuel, fantôme, service ou composant non-radio.

## 2. Matériel radio identifié

### Wi-Fi
- Fabricant : MediaTek
- Modèle : MT7921 Wi-Fi 6
- Bus : PCI
- PNP : PCI\\VEN_14C3&DEV_7961&SUBSYS_46801A3B&REV_00\\4&303B37F5&0&0012
- Service : mtkwlex
- État PnP : OK
- Interface : Wi-Fi
- État interface : connecté
- Radio : 802.11ax
- Bande observée : 5 GHz
- Conclusion : HARDWARE_RADIO = PASS

### Bluetooth
- Fabricant : MediaTek
- Modèle : MediaTek Bluetooth Adapter
- Bus : USB
- PNP : USB\\VID_13D3&PID_3563&MI_00\\7&1D754FA2&0&0000
- Service : BTHUSB
- État PnP : OK
- Conclusion : HARDWARE_RADIO = PASS

### WWAN / cellulaire
- Aucun périphérique modem WWAN matériel identifié dans ce fichier.
- La section WWAN contient principalement des composants Realtek Audio.
- Conclusion : WWAN_HARDWARE = NOT_PROVEN.

## 3. Couches à ne jamais confondre

HARDWARE_RADIO
→ périphérique physique identifiable sur PCI/USB/SDIO/M.2/ACPI selon le matériel.

RADIO_DRIVER
→ pilote attaché au matériel.

RADIO_SERVICE
→ service Windows associé.

RADIO_INTERFACE
→ interface réseau visible par l'OS.

VIRTUAL_RADIO
→ Wi-Fi Direct, SWD\\RADIO, adaptateur virtuel, etc.

PHANTOM_RADIO
→ périphérique résiduel avec CM_PROB_PHANTOM.

NON_RADIO
→ audio, composants Realtek audio, endpoints, APO, ASIO, etc.

## 4. Faux positifs détectés

La section RADIO mélange :
- MediaTek MT7921 ;
- MediaTek Bluetooth ;
- SWD\\RADIO ;
- Wi-Fi Direct ;
- Bluetooth PAN ;
- composants Realtek Audio.

Les composants suivants ne doivent pas être comptés comme radios matérielles :
- Realtek Hardware Support Application
- Realtek(R) Audio
- Realtek Audio Universal Service
- Realtek Asio Component
- Realtek Audio Effects Component
- Realtek OVWrap2 Component
- endpoints microphone/haut-parleurs.

## 5. Périphériques fantômes

Le rapport identifie :
- Realtek 8812BU Wireless LAN 802.11ac USB NIC — CM_PROB_PHANTOM
- Wi-Fi 2 — CM_PROB_PHANTOM
- Microsoft Wi-Fi Direct Virtual Adapter #3 — CM_PROB_PHANTOM
- Microsoft Wi-Fi Direct Virtual Adapter #4 — CM_PROB_PHANTOM

Ils sont à classer PHANTOM_RADIO et ne constituent pas une preuve de matériel actuellement présent.

## 6. Satellite / radio externe

La catégorie SATELLITE doit être traitée séparément.

Un ordinateur ne doit pas être déclaré « satellite capable » simplement parce qu'il possède Wi-Fi/Bluetooth/GNSS ou Internet.

Preuve SATELLITE_HARDWARE attendue :
- périphérique GNSS/GPS identifiable ;
- modem satellite identifiable ;
- périphérique USB/PCIe série ou réseau associé ;
- identifiant PNP matériel ;
- pilote associé ;
- interface ou port de communication ;
- éventuellement terminal satellite externe.

Preuve SATELLITE_LINK attendue :
- liaison satellite réellement établie ;
- métriques ou état de session ;
- équipement/terminal identifié.

Dans le rapport actuel :
- SATELLITE_HARDWARE = NOT_PROVEN
- SATELLITE_LINK = NOT_PROVEN

Aucun satellite ne doit être inventé à partir de l'adresse IP, du Wi-Fi ou des routes VMware.

## 7. Cellulaire / WWAN

Même méthode que satellite :
- rechercher les périphériques modem ;
- identifier PCI/USB/ACPI ;
- relever VID/PID ou VEN/DEV ;
- relever le service/pilote ;
- relever les interfaces WWAN ;
- distinguer modem réel et adaptateur virtuel.

État actuel :
WWAN_HARDWARE = NOT_PROVEN.

## 8. GNSS / GPS

Catégorie indépendante :
- GNSS USB ;
- GNSS PCIe ;
- GNSS série ;
- capteur GNSS intégré ;
- récepteur externe.

Le fichier actuel ne fournit pas de preuve matérielle GNSS.

GNSS_HARDWARE = NOT_PROVEN.

## 9. NFC / RFID

Catégorie indépendante :
- USB ;
- PCIe ;
- I2C ;
- SPI ;
- ACPI ;
- lecteur externe.

Aucune preuve NFC/RFID matérielle dans ce rapport.

NFC_RFID_HARDWARE = NOT_PROVEN.

## 10. Infrared / IR

Catégorie indépendante :
- IR camera ;
- Consumer IR ;
- récepteur IR ;
- périphérique USB/ACPI associé.

Aucune preuve matérielle IR dans le rapport.

IR_HARDWARE = NOT_PROVEN.

## 11. Radio courte portée

Bluetooth est confirmé.

Bluetooth LE est visible via plusieurs BTHLEDEVICE, mais ces entrées sont des services/profils et ne doivent pas être comptées comme plusieurs cartes Bluetooth.

Résultat :
- Bluetooth hardware : PASS
- Bluetooth stack : PASS
- BLE services : PRESENT
- nombre de radios physiques Bluetooth : 1 matériel confirmé par l'identifiant USB.

## 12. Wi-Fi Direct

Les Microsoft Wi-Fi Direct Virtual Adapter sont des interfaces virtuelles.

Ils démontrent une capacité logicielle au-dessus du matériel Wi-Fi mais ne constituent pas de nouvelles cartes radio.

## 13. Réseau virtuel VMware

Les nombreux VMnet :
- sont des adaptateurs virtuels ;
- ne sont pas du matériel radio ;
- ne doivent jamais entrer dans le compteur radio matériel.

## 14. Adresse MAC

L'interface Wi-Fi expose :
14:13:33:e1:0a:47

Cette donnée caractérise l'interface réseau mais ne constitue pas, seule, une preuve indépendante du composant physique.

La preuve forte vient de la chaîne PCI/PnP du MT7921.

## 15. Niveau de preuve

### Niveau P0
Nom générique d'interface seulement.
→ faible.

### Niveau P1
Interface + état réseau.
→ preuve de fonctionnement logiciel/réseau.

### Niveau P2
PNPDeviceID avec bus PCI/USB.
→ preuve matérielle forte au niveau OS.

### Niveau P3
PNP + fabricant/modèle + pilote/service cohérent.
→ preuve matérielle renforcée.

### Niveau P4
PNP + bus + pilote + interface + fonctionnement + corrélation indépendante.
→ preuve très forte.

### Niveau P5
Preuve physique indépendante de l'OS.
→ nécessaire pour une démonstration physique complète.

État actuel :
- Wi-Fi : P4 environ
- Bluetooth : P3/P4
- WWAN : P0
- GNSS : P0
- Satellite : P0
- NFC/RFID : P0
- IR : P0

## 16. Cartographie finale actuelle

| Domaine | Matériel | Pilote/stack | Interface | Verdict |
|---|---|---|---|---|
| Wi-Fi | MT7921 PCI | mtkwlex | Wi-Fi | PASS |
| Bluetooth | MediaTek USB | BTHUSB | Bluetooth | PASS |
| Wi-Fi Direct | non séparé | vwifimp | virtuel | VIRTUAL |
| Bluetooth PAN | non séparé | BTH stack | PAN | VIRTUAL/LOGICAL |
| WWAN | non identifié | — | — | NOT PROVEN |
| GNSS/GPS | non identifié | — | — | NOT PROVEN |
| Satellite | non identifié | — | — | NOT PROVEN |
| NFC/RFID | non identifié | — | — | NOT PROVEN |
| IR | non identifié | — | — | NOT PROVEN |

## 17. Règle de décision

Ne jamais conclure :
« radio matérielle présente »
sur la seule base de :
- nom d'interface ;
- adresse IP ;
- route ;
- service Windows ;
- SWD\\RADIO ;
- adaptateur virtuel ;
- profil Bluetooth ;
- MAC seule.

Conclusion HARDWARE_RADIO = PASS seulement si un identifiant matériel et un bus physique cohérent sont présents.

## 18. Audit complémentaire à exécuter

La prochaine collecte doit ajouter, toujours en lecture seule :
1. PNPDeviceID complet de tous les périphériques radio.
2. Bus type et emplacement.
3. Hardware IDs.
4. Compatible IDs.
5. Driver provider/version/date.
6. Service du pilote.
7. État PnP et problème ConfigManager.
8. ACPI/PCI/USB parentage.
9. interfaces réseau associées.
10. Wi-Fi capabilities.
11. Bluetooth radio/device information.
12. WWAN devices.
13. GNSS/GPS devices.
14. NFC/RFID devices.
15. IR devices.
16. USB tree pertinent.
17. PCI tree pertinent.
18. périphériques fantômes.
19. adaptateurs virtuels.
20. résumé machine-readable PASS/FAIL/NOT_PROVEN.

## 19. Satellite : extension de protocole

Pour une future preuve satellite, le protocole doit rechercher trois niveaux distincts :

SATELLITE_HARDWARE
→ terminal/récepteur/modem matériel.

SATELLITE_DRIVER
→ pilote attaché.

SATELLITE_LINK
→ liaison effectivement active.

Internet via Wi-Fi, Ethernet, 4G/5G ou VPN n'est jamais considéré comme preuve satellite.

## 20. Sécurité

Travail conçu en lecture seule :
- aucune suppression ;
- aucune désactivation ;
- aucune installation ;
- aucun flash ;
- aucune écriture firmware ;
- aucune modification NVRAM ;
- aucune modification réseau ;
- aucun redémarrage automatique.

## 21. Verdict global

HARDWARE RADIO:
- Wi-Fi = CONFIRMED
- Bluetooth = CONFIRMED

OTHER RADIO:
- WWAN = NOT PROVEN
- GNSS = NOT PROVEN
- SATELLITE = NOT PROVEN
- NFC/RFID = NOT PROVEN
- IR = NOT PROVEN

VIRTUAL:
- Wi-Fi Direct = PRESENT
- VMware VMnet = PRESENT

PHANTOM:
- Realtek 8812BU = PRESENT AS PHANTOM
- Wi-Fi 2 = PRESENT AS PHANTOM
- Wi-Fi Direct #3/#4 = PRESENT AS PHANTOM

PHYSICAL INDEPENDENT PROOF:
- NOT PRESENT IN THIS WINDOWS AUDIT

Cette cartographie remplace le comptage brut des « radio candidates » comme base de décision.