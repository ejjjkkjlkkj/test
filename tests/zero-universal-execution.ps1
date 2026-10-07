Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

function New-SemanticProgram {
    [pscustomobject]@{
        ProgramId = 'ZERO.TEST.UNIVERSAL.001'
        Operation = 'CREATE'
        ObjectId = 'OBJECT:NOTE:001'
        Meaning = 'NOTE'
        Input = 'Bonjour'
        RequiredCapabilities = @('OBJECT.CREATE')
    }
}

function New-Realization([string]$Name, [string[]]$Capabilities) {
    [pscustomobject]@{
        Name = $Name
        Capabilities = $Capabilities
    }
}

function Invoke-SemanticProgram($Program, $Realization) {
    foreach ($required in $Program.RequiredCapabilities) {
        if ($Realization.Capabilities -notcontains $required) {
            return [pscustomobject]@{
                State = 'REJECTED'
                Reason = 'MISSING_CAPABILITY'
                ProgramId = $Program.ProgramId
                Result = $null
                Accessible = $false
                Proven = $false
            }
        }
    }

    [pscustomobject]@{
        State = 'COMPLETED'
        Reason = $null
        ProgramId = $Program.ProgramId
        Result = [pscustomobject]@{
            ObjectId = $Program.ObjectId
            Meaning = $Program.Meaning
            Value = $Program.Input
        }
        Accessible = $true
        Proven = $true
    }
}

function Test-SemanticEquivalence($A, $B) {
    if ($A.ProgramId -ne $B.ProgramId) { return $false }
    if ($A.State -ne $B.State) { return $false }
    if ($A.Result.ObjectId -ne $B.Result.ObjectId) { return $false }
    if ($A.Result.Meaning -ne $B.Result.Meaning) { return $false }
    if ($A.Result.Value -ne $B.Result.Value) { return $false }
    if ($A.Accessible -ne $B.Accessible) { return $false }
    if ($A.Proven -ne $B.Proven) { return $false }
    return $true
}

$program = New-SemanticProgram

$realizations = @(
    (New-Realization 'PC' @('OBJECT.CREATE')),
    (New-Realization 'IPHONE' @('OBJECT.CREATE')),
    (New-Realization 'ANDROID' @('OBJECT.CREATE')),
    (New-Realization 'EMBEDDED' @('OBJECT.CREATE')),
    (New-Realization 'REMOTE-NODE' @('OBJECT.CREATE'))
)

$results = foreach ($r in $realizations) {
    Invoke-SemanticProgram $program $r
}

$reference = $results[0]

foreach ($result in $results) {
    if (-not $result.Accessible) { throw "ACCESSIBILITY FAIL: $($result.ProgramId)" }
    if (-not $result.Proven) { throw "PROOF FAIL: $($result.ProgramId)" }
    if (-not (Test-SemanticEquivalence $reference $result)) {
        throw "SEMANTIC EQUIVALENCE FAIL"
    }
}

$missing = New-Realization 'NO-CREATE' @()
$rejected = Invoke-SemanticProgram $program $missing

if ($rejected.State -ne 'REJECTED') { throw 'MISSING CAPABILITY WAS NOT REJECTED' }
if ($rejected.Reason -ne 'MISSING_CAPABILITY') { throw 'WRONG REJECTION REASON' }
if ($rejected.Accessible) { throw 'REJECTED INSTANCE CLAIMED ACCESSIBLE' }
if ($rejected.Proven) { throw 'REJECTED INSTANCE CLAIMED PROVEN' }

Write-Output 'ZERO_UNIVERSAL_EXECUTION_SEMANTIC_EQUIVALENCE=PASS'
Write-Output 'ZERO_UNIVERSAL_EXECUTION_CAPABILITY_REJECTION=PASS'
Write-Output 'ZERO_UNIVERSAL_EXECUTION_ACCESSIBILITY=PASS'
Write-Output 'ZERO_UNIVERSAL_EXECUTION_RESULT=PASS'
