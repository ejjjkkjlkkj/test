Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'
$KnownTypes = @('OBJECT','EVENT','OPERATION','RESULT','PROOF','CAPABILITY','OBSERVATION')
$Proofs = @('UNKNOWN','DEFINED','IMPLEMENTED','TESTED','SIMULATED','QEMU','HARDWARE','RF_PROVEN','SATELLITE_LINK_PROVEN')
$LastSequence = @{}

function Split-ZeroRecord([string]$Line) {
    $fields = [System.Collections.Generic.List[string]]::new()
    $buffer = [System.Text.StringBuilder]::new()
    $escaped = $false
    foreach ($c in $Line.ToCharArray()) {
        if ($escaped) { [void]$buffer.Append($c); $escaped = $false; continue }
        if ($c -eq '\') { $escaped = $true; continue }
        if ($c -eq '|') { $fields.Add($buffer.ToString()); $buffer.Clear(); continue }
        [void]$buffer.Append($c)
    }
    if ($escaped) { throw 'FORMAT.ESCAPE' }
    $fields.Add($buffer.ToString())
    return ,$fields.ToArray()
}

function Decode-Zero([string]$Value) {
    $out = [System.Text.StringBuilder]::new()
    $escaped = $false
    foreach ($c in $Value.ToCharArray()) {
        if ($escaped) {
            switch ($c) {
                '\' { [void]$out.Append('\') }
                'n' { [void]$out.Append([char]10) }
                'r' { [void]$out.Append([char]13) }
                't' { [void]$out.Append([char]9) }
                '|' { [void]$out.Append('|') }
                default { throw 'FORMAT.ESCAPE' }
            }
            $escaped = $false
        } elseif ($c -eq '\') { $escaped = $true }
        else { [void]$out.Append($c) }
    }
    if ($escaped) { throw 'FORMAT.ESCAPE' }
    return $out.ToString()
}

function Test-ZeroRecord([string]$Line, [string]$Stream='default') {
    $fields = Split-ZeroRecord $Line
    if ($fields.Count -ne 10) { return @{ Valid=$false; Code='FORMAT.FIELD_COUNT' } }
    if ($fields[0] -ne 'RECORD') { return @{ Valid=$false; Code='FORMAT.PREFIX' } }
    try {
        $version=Decode-Zero $fields[1]; $type=Decode-Zero $fields[2]; $identity=Decode-Zero $fields[3]
        $sequenceText=Decode-Zero $fields[4]; $time=Decode-Zero $fields[5]; $source=Decode-Zero $fields[6]
        $target=Decode-Zero $fields[7]; $payload=Decode-Zero $fields[8]; $proof=Decode-Zero $fields[9]
    } catch { return @{ Valid=$false; Code=$_.Exception.Message } }
    if ($version -ne '1') { return @{ Valid=$false; Code='FORMAT.VERSION' } }
    if ([string]::IsNullOrEmpty($type)) { return @{ Valid=$false; Code='FORMAT.TYPE' } }
    if ([string]::IsNullOrEmpty($identity)) { return @{ Valid=$false; Code='FORMAT.IDENTITY' } }
    if ($sequenceText -notmatch '^(0|[1-9][0-9]*)$') { return @{ Valid=$false; Code='FORMAT.SEQUENCE' } }
    if ($time -ne 'UNKNOWN' -and $time -notmatch '^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]{3}Z$') { return @{ Valid=$false; Code='FORMAT.TIME' } }
    if ([string]::IsNullOrEmpty($source)) { return @{ Valid=$false; Code='FORMAT.SOURCE' } }
    if ([string]::IsNullOrEmpty($target)) { return @{ Valid=$false; Code='FORMAT.TARGET' } }
    if ($proof -notin $Proofs) { return @{ Valid=$false; Code='FORMAT.PROOF' } }
    $sequence=[int64]$sequenceText
    if ($LastSequence.ContainsKey($Stream) -and $sequence -le $LastSequence[$Stream]) { return @{ Valid=$false; Code='FORMAT.SEQUENCE' } }
    $LastSequence[$Stream]=$sequence
    $normalizedType=if($type -in $KnownTypes){$type}else{'UNKNOWN_TYPE'}
    return @{ Valid=$true; Code='ACCEPTED'; Type=$normalizedType; RawType=$(if($normalizedType -eq 'UNKNOWN_TYPE'){$type}else{''}); Proof=$proof }
}

function Assert-Result([string]$Name,[string]$Line,[bool]$ExpectedValid,[string]$ExpectedCode,[string]$Stream) {
    $r=Test-ZeroRecord $Line $Stream
    if($r.Valid -ne $ExpectedValid){throw "$Name FAIL validity"}
    if($r.Code -ne $ExpectedCode){throw "$Name FAIL code expected=$ExpectedCode actual=$($r.Code)"}
    Write-Output "$Name PASS"
}

$root=Split-Path -Parent $PSScriptRoot
foreach($relative in @('TEST-VECTORS/object-create.zero','TEST-VECTORS/event-discover.zero','TEST-VECTORS/observe-result.zero','TEST-VECTORS/satellite-link.zero')){
    $line=(Get-Content -Raw (Join-Path $root $relative)).TrimEnd([char]13,[char]10)
    Assert-Result $relative $line $true 'ACCEPTED' $relative
}
Assert-Result 'V002' 'RECORD|1|OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=test' $false 'FORMAT.FIELD_COUNT' 'V002'
Assert-Result 'V004' 'RECORD|1|OBJECT|x|1|NOT-A-TIME|ZERO|WORLD|meaning=test|DEFINED' $false 'FORMAT.TIME' 'V004'
Assert-Result 'V003-first' 'RECORD|1|OBJECT|x|2|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=test|DEFINED' $true 'ACCEPTED' 'V003'
Assert-Result 'V003-second' 'RECORD|1|EVENT|x|1|2026-10-07T00:00:01.000Z|ZERO|x|kind=test|DEFINED' $false 'FORMAT.SEQUENCE' 'V003'
$v005=Test-ZeroRecord 'RECORD|1|FUTURE_OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|future=value|DEFINED' 'V005'
if(-not $v005.Valid -or $v005.Type -ne 'UNKNOWN_TYPE' -or $v005.RawType -ne 'FUTURE_OBJECT'){throw 'V005 FAIL'}; Write-Output 'V005 PASS'
$v006=Test-ZeroRecord 'RECORD|1|OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=test|SIMULATED' 'V006'
if(-not $v006.Valid -or $v006.Proof -ne 'SIMULATED'){throw 'V006 FAIL'}; Write-Output 'V006 PASS'
Assert-Result 'V007' 'RECORD|1|OBJECT|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|meaning=ligne\navec\|separateur\\reel|DEFINED' $true 'ACCEPTED' 'V007'
Assert-Result 'V008' 'RECORD|1|OPERATION|x|1|2026-10-07T00:00:00.000Z|ZERO|WORLD|intent=ACT|DEFINED' $true 'ACCEPTED' 'V008'
Write-Output 'ZERO_VALIDATOR_RESULT=PASS'
