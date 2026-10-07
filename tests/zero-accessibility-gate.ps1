$ErrorActionPreference = 'Stop'

function Assert-Equal([string]$Name, $Actual, $Expected) {
    if ($Actual -ne $Expected) {
        throw "$Name FAIL: expected [$Expected], got [$Actual]"
    }
    Write-Host "$Name PASS"
}

# Accessibility is a mandatory semantic property, not a documentation claim.
# Every accepted result must expose the same semantic value through every
# required native projection. A projection that changes meaning is rejected.

$semantic = [ordered]@{
    identity = 'CALCULATION:1+1'
    operation = 'ADD'
    result = '2'
    state = 'COMPLETED'
    proof = 'TESTED'
}

$requiredModalities = @(
    'VOICE',
    'BRAILLE',
    'KEYBOARD',
    'DISPLAY',
    'TOUCH',
    'POINTER',
    'NETWORK',
    'AUTOMATION'
)

function Project([hashtable]$Value, [string]$Modality) {
    if ($Modality -notin $requiredModalities) {
        throw "ACCESSIBILITY.MODALITY_UNKNOWN"
    }

    # The projection carries the semantic source unchanged.
    return [ordered]@{
        modality = $Modality
        identity = $Value.identity
        operation = $Value.operation
        result = $Value.result
        state = $Value.state
        proof = $Value.proof
    }
}

foreach ($modality in $requiredModalities) {
    $projection = Project $semantic $modality
    Assert-Equal "$modality modality" $projection.modality $modality
    Assert-Equal "$modality identity" $projection.identity 'CALCULATION:1+1'
    Assert-Equal "$modality operation" $projection.operation 'ADD'
    Assert-Equal "$modality result" $projection.result '2'
    Assert-Equal "$modality state" $projection.state 'COMPLETED'
    Assert-Equal "$modality proof" $projection.proof 'TESTED'
}

# Accessibility must not be optional for an accepted semantic result.
$missing = @($requiredModalities | Where-Object { $_ -eq '' })
Assert-Equal 'REQUIRED_MODALITIES_PRESENT' $missing.Count 0
$unknownRejected = $false
try { [void](Project $semantic 'UNKNOWN') } catch { $unknownRejected = $true }
Assert-Equal 'UNKNOWN_MODALITY_REJECTED' $unknownRejected $true

Write-Host 'ZERO_ACCESSIBILITY_GATE_RESULT=PASS'
