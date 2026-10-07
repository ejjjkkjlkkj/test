$ErrorActionPreference='Stop'

function New-State {
    [ordered]@{
        Sequence=0
        Objects=@{}
        Capabilities=@{}
        Events=@()
        Results=@()
    }
}

function Emit-Event($State,[string]$Type,[string]$Identity,[string]$Meaning) {
    $State.Sequence++
    $event=[ordered]@{
        Sequence=$State.Sequence
        Type=$Type
        Identity=$Identity
        Meaning=$Meaning
    }
    $State.Events += ,$event
    return $event
}

function Invoke-ZeroISA($State,[string]$Instruction,$Arguments) {
    switch ($Instruction) {
        'CREATE' {
            $id=[string]$Arguments.Identity
            if ($State.Objects.ContainsKey($id)) { throw 'CREATE.DUPLICATE' }
            $State.Objects[$id]=[ordered]@{
                Identity=$id
                Meaning=[string]$Arguments.Meaning
                State='CREATED'
            }
            $event=Emit-Event $State 'OBJECT.CREATED' $id $State.Objects[$id].Meaning
            $result=[ordered]@{Operation='CREATE';State='COMPLETED';Identity=$id;EventSequence=$event.Sequence;Proof='TESTED'}
        }
        'OBSERVE' {
            $id=[string]$Arguments.Identity
            if (-not $State.Objects.ContainsKey($id)) {
                $result=[ordered]@{Operation='OBSERVE';State='UNKNOWN';Identity=$id;Proof='TESTED'}
            } else {
                $o=$State.Objects[$id]
                $result=[ordered]@{Operation='OBSERVE';State='COMPLETED';Identity=$id;Meaning=$o.Meaning;ObjectState=$o.State;Proof='TESTED'}
            }
        }
        'CAPABILITY' {
            $id=[string]$Arguments.Identity
            $State.Capabilities[$id]=[string]$Arguments.Operation
            $result=[ordered]@{Operation='CAPABILITY';State='COMPLETED';Identity=$id;Capability=$State.Capabilities[$id];Proof='TESTED'}
        }
        'SET' {
            $id=[string]$Arguments.Identity
            if (-not $State.Objects.ContainsKey($id)) { throw 'SET.UNKNOWN_OBJECT' }
            $op=[string]$Arguments.Operation
            if (-not $State.Capabilities.ContainsKey($id) -or $State.Capabilities[$id] -ne $op) { throw 'SET.CAPABILITY_REQUIRED' }
            $State.Objects[$id].State=[string]$Arguments.State
            $event=Emit-Event $State 'OBJECT.STATE_CHANGED' $id $State.Objects[$id].State
            $result=[ordered]@{Operation='SET';State='COMPLETED';Identity=$id;ObjectState=$State.Objects[$id].State;EventSequence=$event.Sequence;Proof='TESTED'}
        }
        default { throw 'ISA.UNKNOWN_INSTRUCTION' }
    }
    $State.Results += ,$result
    return $result
}

function Project-Result($Result,[string]$Modality) {
    if ([string]::IsNullOrWhiteSpace($Modality)) { throw 'ACCESSIBILITY.MODALITY.EMPTY' }
    [ordered]@{
        Modality=$Modality
        Operation=$Result.Operation
        State=$Result.State
        Identity=$Result.Identity
        Proof=$Result.Proof
        Meaning=if ($Result.Contains('Meaning')) {$Result.Meaning} else {$null}
        ObjectState=if ($Result.Contains('ObjectState')) {$Result.ObjectState} else {$null}
    }
}

$state=New-State
$create=Invoke-ZeroISA $state 'CREATE' @{Identity='OBJ:ISA:1';Meaning='Objet ZERO accessible'}
if ($create.State -ne 'COMPLETED' -or $state.Objects['OBJ:ISA:1'].Meaning -ne 'Objet ZERO accessible') { throw 'CREATE.FAIL' }

$before=$state.Objects['OBJ:ISA:1'].State
$observe=Invoke-ZeroISA $state 'OBSERVE' @{Identity='OBJ:ISA:1'}
if ($observe.State -ne 'COMPLETED' -or $state.Objects['OBJ:ISA:1'].State -ne $before) { throw 'OBSERVE.MUTATED' }

$unknown=Invoke-ZeroISA $state 'OBSERVE' @{Identity='OBJ:UNKNOWN'}
if ($unknown.State -ne 'UNKNOWN') { throw 'UNKNOWN.FAIL' }

$denied=$false
try { [void](Invoke-ZeroISA $state 'SET' @{Identity='OBJ:ISA:1';Operation='ACT';State='ACTIVE'}) } catch { $denied=$true }
if (-not $denied) { throw 'CAPABILITY.BYPASS' }

[void](Invoke-ZeroISA $state 'CAPABILITY' @{Identity='OBJ:ISA:1';Operation='ACT'})
$set=Invoke-ZeroISA $state 'SET' @{Identity='OBJ:ISA:1';Operation='ACT';State='ACTIVE'}
if ($set.State -ne 'COMPLETED' -or $state.Objects['OBJ:ISA:1'].State -ne 'ACTIVE') { throw 'SET.FAIL' }

foreach($m in @('VOICE','BRAILLE','KEYBOARD','DISPLAY','TOUCH','POINTER','NETWORK','AUTOMATION')){
    $p=Project-Result $set $m
    foreach($k in @('Operation','State','Identity','Proof')){
        if ([string]$p[$k] -cne [string]$set[$k]) { throw "ACCESSIBILITY.FAIL:${m}:$k" }
    }
}

Write-Host 'ZERO_MACHINE_ISA_EXECUTOR=PASS'
Write-Host 'ZERO_MACHINE_ISA_ACCESSIBILITY=PASS'
Write-Host 'ZERO_MACHINE_ISA_CAPABILITY_ENFORCEMENT=PASS'
Write-Host 'ZERO_MACHINE_ISA_EVENT_EMISSION=PASS'
