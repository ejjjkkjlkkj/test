$ErrorActionPreference = 'Stop'

function New-ZeroState {
    [ordered]@{
        Sequence = 0
        Objects = @{}
        Events = @()
        Results = @()
    }
}

function New-ZeroObject([string]$Identity,[string]$Meaning) {
    [ordered]@{
        Identity = $Identity
        Meaning = $Meaning
        State = 'CREATED'
    }
}

function Invoke-ZeroCreate($Machine,[string]$Identity,[string]$Meaning) {
    if ($Machine.Objects.ContainsKey($Identity)) { throw "OBJECT_EXISTS:$Identity" }
    $obj = New-ZeroObject $Identity $Meaning
    $Machine.Objects[$Identity] = $obj
    $Machine.Sequence++
    $Machine.Events += [ordered]@{
        Type='OBJECT.CREATED'; Identity=$Identity; Sequence=$Machine.Sequence
    }
    $result = [ordered]@{
        Identity=$Identity; Operation='CREATE'; State='COMPLETED'
        Result=$Identity; Proof='TESTED'
    }
    $Machine.Results += $result
    return $result
}

function Invoke-ZeroObserve($Machine,[string]$Identity) {
    if (-not $Machine.Objects.ContainsKey($Identity)) {
        return [ordered]@{Identity=$Identity; Operation='OBSERVE'; State='UNKNOWN'; Result=$null; Proof='TESTED'}
    }
    $obj=$Machine.Objects[$Identity]
    return [ordered]@{
        Identity=$obj.Identity; Operation='OBSERVE'; State=$obj.State
        Result=$obj.Meaning; Proof='TESTED'
    }
}

function Convert-ZeroProjection($Result,[string]$Modality) {
    if ([string]::IsNullOrWhiteSpace($Modality)) { throw 'ACCESSIBILITY.MODALITY.EMPTY' }
    [ordered]@{
        Modality=$Modality
        Identity=$Result.Identity
        Operation=$Result.Operation
        State=$Result.State
        Result=$Result.Result
        Proof=$Result.Proof
    }
}

$machine=New-ZeroState
$created=Invoke-ZeroCreate $machine 'CALCULATION:1+1' '2'
$observed=Invoke-ZeroObserve $machine 'CALCULATION:1+1'

if ($observed.Result -ne '2' -or $observed.State -ne 'CREATED') { throw 'ENGINE.OBSERVE.FAIL' }
if ($machine.Sequence -ne 1) { throw 'ENGINE.SEQUENCE.FAIL' }

$modalities=@('VOICE','BRAILLE','KEYBOARD','DISPLAY','TOUCH','POINTER','NETWORK','AUTOMATION')
foreach ($m in $modalities) {
    $p=Convert-ZeroProjection $created $m
    if ($p.Identity -ne $created.Identity -or $p.Operation -ne $created.Operation -or
        $p.Result -ne $created.Result -or $p.State -ne $created.State -or $p.Proof -ne $created.Proof) {
        throw "ACCESSIBILITY.PROJECTION.FAIL:$m"
    }
}

Write-Host 'ZERO_REFERENCE_ENGINE_RESULT=PASS'
Write-Host 'ZERO_REFERENCE_ENGINE_ACCESSIBILITY=PASS'
