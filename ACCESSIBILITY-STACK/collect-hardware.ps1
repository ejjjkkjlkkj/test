#requires -Version 7.0
[CmdletBinding()]
param([string]$OutputDirectory=(Join-Path $PWD "accessibility-audit"))
# READ ONLY: no device/driver/firmware/NVRAM/network modification.
Set-StrictMode -Version Latest
$ErrorActionPreference="Continue"
New-Item -ItemType Directory -Force -Path $OutputDirectory | Out-Null
$stamp=Get-Date -Format "yyyyMMdd-HHmmss"
$records=[System.Collections.Generic.List[object]]::new()
function Add-Record([object]$x){$records.Add($x)}
function Safe([scriptblock]$s){try{&$s}catch{$null}}
$computer=Safe {Get-CimInstance Win32_ComputerSystem | Select Manufacturer,Model,SystemType,TotalPhysicalMemory}
$os=Safe {Get-CimInstance Win32_OperatingSystem | Select Caption,Version,BuildNumber,OSArchitecture}
$pnp=Safe {Get-PnpDevice}
foreach($d in @($pnp)){
  $keys="DEVPKEY_Device_HardwareIds","DEVPKEY_Device_CompatibleIds","DEVPKEY_Device_Service","DEVPKEY_Device_Driver","DEVPKEY_Device_DriverVersion","DEVPKEY_Device_DriverDate","DEVPKEY_Device_BusReportedDeviceDesc","DEVPKEY_Device_LocationInfo","DEVPKEY_Device_Parent"
  $props=Safe {Get-PnpDeviceProperty -InstanceId $d.InstanceId -KeyName $keys}
  $ids=@($props|Where KeyName -eq "DEVPKEY_Device_HardwareIds"|Select -Expand Data -ErrorAction SilentlyContinue)
  $compat=@($props|Where KeyName -eq "DEVPKEY_Device_CompatibleIds"|Select -Expand Data -ErrorAction SilentlyContinue)
  $service=($props|Where KeyName -eq "DEVPKEY_Device_Service"|Select -Expand Data -First 1)
  $driver=($props|Where KeyName -eq "DEVPKEY_Device_Driver"|Select -Expand Data -First 1)
  $driverVersion=($props|Where KeyName -eq "DEVPKEY_Device_DriverVersion"|Select -Expand Data -First 1)
  $driverDate=($props|Where KeyName -eq "DEVPKEY_Device_DriverDate"|Select -Expand Data -First 1)
  $parent=($props|Where KeyName -eq "DEVPKEY_Device_Parent"|Select -Expand Data -First 1)
  $bus=if($d.InstanceId -match "^PCI\\"){"PCI"}elseif($d.InstanceId -match "^USB\\"){"USB"}elseif($d.InstanceId -match "^ACPI\\"){"ACPI"}elseif($d.InstanceId -match "^HID\\"){"HID"}else{"OTHER"}
  $radio=[bool]($d.FriendlyName -match "Wi-Fi|Wireless|Bluetooth|WWAN|LTE|5G|Cell|GNSS|GPS|NFC|RFID|Infrared|IR|Satellite|Radio")
  $state=if($d.ConfigManagerErrorCode -eq 24){"PHANTOM"}elseif($d.Status -eq "OK"){"PRESENT"}else{"ERROR"}
  Add-Record ([pscustomobject]@{type="device";id=$d.InstanceId;name=$d.FriendlyName;class=$d.Class;status=$d.Status;problem_code=$d.ConfigManagerErrorCode;bus=$bus;hardware_ids=$ids;compatible_ids=$compat;service=$service;driver=$driver;driver_version=$driverVersion;driver_date=$driverDate;parent=$parent;radio_hint=$radio;state=$state;evidence=if($d.InstanceId -match "^(PCI|USB)\\"){"P2"}else{"P0"}})
}
$adapters=Safe {Get-NetAdapter -IncludeHidden | Select Name,InterfaceDescription,ifIndex,Status,MacAddress,LinkSpeed,Virtual,HardwareInterface}
foreach($a in @($adapters)){Add-Record ([pscustomobject]@{type="network_interface";name=$a.Name;description=$a.InterfaceDescription;status=$a.Status;mac=$a.MacAddress;virtual=$a.Virtual;hardware_interface=$a.HardwareInterface})}
$wifi=@($adapters|Where {$_.Name -match "Wi-Fi|Wireless" -or $_.InterfaceDescription -match "Wi-Fi|Wireless|802.11|WLAN"})
Add-Record ([pscustomobject]@{type="capability";capability="RADIO.WIFI";state=if($wifi.Count){"PROVEN"}else{"NOT_PROVEN"};evidence="network interface; PNP records"})
$bluetooth=@($pnp|Where {$_.FriendlyName -match "Bluetooth" -or $_.InstanceId -match "^USB\\.*BTH"})
Add-Record ([pscustomobject]@{type="capability";capability="RADIO.BLUETOOTH";state=if($bluetooth.Count){"PROVEN"}else{"NOT_PROVEN"};evidence="PNP records"})
$patterns=@(@{cap="RADIO.WWAN";rx="WWAN|LTE|5G|Cellular|Mobile Broadband|Modem"},@{cap="RADIO.GNSS";rx="GNSS|GPS|Galileo|GLONASS|BeiDou|NavIC"},@{cap="RADIO.NFC";rx="NFC|Near Field"},@{cap="RADIO.RFID";rx="RFID"},@{cap="RADIO.IR";rx="Infrared|IR Blaster|Consumer IR"},@{cap="RADIO.SATELLITE";rx="Satellite|SATCOM|Iridium|Globalstar|Thuraya|Inmarsat"})
foreach($p in $patterns){$m=@($pnp|Where {$_.FriendlyName -match $p.rx -or $_.InstanceId -match $p.rx});Add-Record ([pscustomobject]@{type="capability";capability=$p.cap;state=if($m.Count){"EVIDENCE_FOUND"}else{"NOT_PROVEN"};matches=@($m|Select InstanceId,FriendlyName,Class,Status)})}
$report=[pscustomobject]@{protocol_version="0.1";timestamp=(Get-Date).ToString("o");safety="READ_ONLY";computer=$computer;os=$os;records=$records}
$json=Join-Path $OutputDirectory "accessibility-hardware-$stamp.json"
$txt=Join-Path $OutputDirectory "accessibility-hardware-$stamp.txt"
$report|ConvertTo-Json -Depth 12|Set-Content -Encoding UTF8 $json
@("ACCESSIBILITY HARDWARE DISCOVERY — READ ONLY","DATE: $($report.timestamp)","MACHINE: $($computer.Model)","","CAPABILITIES",($records|Where type -eq "capability"|Format-Table -AutoSize|Out-String),"","RECORDS: $($records.Count)")|Set-Content -Encoding UTF8 $txt
Write-Output "PASS — audit complete"
Write-Output "JSON: $json"
Write-Output "TEXT: $txt"
