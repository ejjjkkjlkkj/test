# PROTOCOL MATRIX v0.1

| Domain | Discovery source | Driver family | Capability | Current status |
|---|---|---|---|---|
| Wi-Fi | PCI/PNP + network | wifi | RADIO.WIFI | PROVEN |
| Bluetooth | USB/PNP + BTH stack | bluetooth | RADIO.BLUETOOTH | PROVEN |
| Audio | PNP + audio endpoints | audio | AUDIO.INPUT/OUTPUT | PROVEN |
| WWAN | PCI/USB/ACPI | wwan | RADIO.WWAN | NOT_PROVEN |
| GNSS | USB/PCI/ACPI/serial | gnss | RADIO.GNSS | NOT_PROVEN |
| Satellite | PCI/USB/serial/network correlation | satellite | RADIO.SATELLITE | NOT_PROVEN |
| NFC/RFID | USB/PCI/ACPI | nfc_rfid | RADIO.NFC/RFID | NOT_PROVEN |
| IR | USB/PCI/ACPI | infrared | RADIO.IR | NOT_PROVEN |
| Keyboard | HID/PS2/USB | input | INPUT.KEYBOARD | DISCOVERY REQUIRED |
| Display | PCI/ACPI/display stack | display | DISPLAY.OUTPUT | DISCOVERY REQUIRED |

NOT_PROVEN is intentionally retained until an evidence-producing probe confirms or rejects the capability.
