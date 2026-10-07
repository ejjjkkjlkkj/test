$ErrorActionPreference = 'Stop'

function Split-ZeroFields([string]$Line) {
    $fields = [System.Collections.Generic.List[string]]::new()
    $buffer = [System.Text.StringBuilder]::new()
    $escaped = $false
    foreach ($c in $Line.ToCharArray()) {
        if ($escaped) { [void]$buffer.Append('\'); [void]$buffer.Append($c); $escaped = $false; continue }
        if ($c -eq '\') { $escaped = $true; continue }
        if ($c -eq '|') { [void]$fields.Add($buffer.ToString()); [void]$buffer.Clear(); continue }
        [void]$buffer.Append($c)
    }
    if ($escaped) { throw 'FORMAT.ESCAPE' }
    [void]$fields.Add($buffer.ToString())
    return ,$fields.ToArray()
}

function Encode-ZeroField([string]$Value) {
    if ($null -eq $Value) { return '' }
    $Value -replace "\\", "\\\\" -replace "\|", "\\|" -replace "`r", "\\r" -replace "`n", "\\n" -replace "`t", "\\t"
}

function Decode-ZeroField([string]$Value) {
    $buffer = [System.Text.StringBuilder]::new()
    $escaped = $false
    foreach ($c in $Value.ToCharArray()) {
        if ($escaped) {
            switch ($c) {
                '\' { [void]$buffer.Append('\') }
                '|' { [void]$buffer.Append('|') }
                'n' { [void]$buffer.Append([char]10) }
                'r' { [void]$buffer.Append([char]13) }
                't' { [void]$buffer.Append([char]9) }
                default { throw 'FORMAT.ESCAPE' }
            }
            $escaped = $false
            continue
        }
        if ($c -eq '\') { $escaped = $true; continue }
        [void]$buffer.Append($c)
    }
    if ($escaped) { throw 'FORMAT.ESCAPE' }
    $buffer.ToString()
}

function ConvertTo-ZeroRecordLine($Record) {
    $values = @('RECORD',$Record.Version,$Record.Type,$Record.Identity,[string]$Record.Sequence,$Record.Time,$Record.Source,$Record.Target,$Record.Payload,$Record.Proof)
    ($values | ForEach-Object { Encode-ZeroField ([string]$_) }) -join '|'
}

function ConvertFrom-ZeroRecordLine([string]$Line) {
    $raw = Split-ZeroFields $Line
    if ($raw.Count -ne 10 -or $raw[0] -ne 'RECORD') { throw 'FORMAT.RECORD' }
    [ordered]@{
        Version=Decode-ZeroField $raw[1]; Type=Decode-ZeroField $raw[2]; Identity=Decode-ZeroField $raw[3]
        Sequence=[int64](Decode-ZeroField $raw[4]); Time=Decode-ZeroField $raw[5]
        Source=Decode-ZeroField $raw[6]; Target=Decode-ZeroField $raw[7]
        Payload=Decode-ZeroField $raw[8]; Proof=Decode-ZeroField $raw[9]
    }
}

function Convert-ZeroProjection($Record,[string]$Modality) {
    if ([string]::IsNullOrWhiteSpace($Modality)) { throw 'ACCESSIBILITY.MODALITY.EMPTY' }
    [ordered]@{ Modality=$Modality; Identity=$Record.Identity; Type=$Record.Type; Payload=$Record.Payload; Proof=$Record.Proof }
}

$record=[ordered]@{
    Version='ZERO-1'; Type='OBJECT'; Identity='TEST:FORMAT|1'; Sequence=1
    Time='2026-10-07T08:00:00.000Z'; Source='TEST:SOURCE'; Target='TEST:TARGET'
    Payload='meaning=accessible|machine;state=TESTED'; Proof='TESTED'
}

$encoded=ConvertTo-ZeroRecordLine $record
$decoded=ConvertFrom-ZeroRecordLine $encoded
foreach ($key in $record.Keys) { if ([string]$decoded[$key] -cne [string]$record[$key]) { throw "ROUNDTRIP.FAIL:$key" } }
if ((ConvertTo-ZeroRecordLine $decoded) -cne $encoded) { throw 'CANONICAL.FAIL' }

foreach ($m in @('VOICE','BRAILLE','KEYBOARD','DISPLAY','TOUCH','POINTER','NETWORK','AUTOMATION')) {
    $projection=Convert-ZeroProjection $decoded $m
    if ($projection.Identity -cne $decoded.Identity -or $projection.Type -cne $decoded.Type -or $projection.Payload -cne $decoded.Payload -or $projection.Proof -cne $decoded.Proof) { throw "ACCESSIBILITY.FAIL:$m" }
}

$bad='RECORD|ZERO-1|OBJECT|BAD|1|2026-10-07T08:00:00.000Z|S|T|broken\q|TESTED'
$rejected=$false
try { [void](ConvertFrom-ZeroRecordLine $bad) } catch { $rejected=$true }
if (-not $rejected) { throw 'INVALID_ESCAPE_ACCEPTED' }

Write-Host 'ZERO_MACHINE_SERIALIZER_RESULT=PASS'
Write-Host 'ZERO_MACHINE_DESERIALIZER_RESULT=PASS'
Write-Host 'ZERO_MACHINE_ROUNDTRIP_RESULT=PASS'
Write-Host 'ZERO_MACHINE_ACCESSIBILITY=PASS'
