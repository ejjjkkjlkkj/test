$ErrorActionPreference = 'Stop'

function Assert-Equal([string]$Name, $Actual, $Expected) {
    if ($Actual -ne $Expected) {
        throw "$Name FAIL: expected [$Expected], got [$Actual]"
    }
    Write-Host "$Name PASS"
}

function Add-Exact([int]$A, [int]$B) {
    return $A + $B
}

function Evaluate-Rule([int]$A, [int]$B, [int]$Claimed) {
    $actual = Add-Exact $A $B
    if ($actual -ne $Claimed) { return 'INVALID' }
    return 'VALID'
}

function Check-Classification([string]$Source, [string]$Target) {
    if (($Source -eq 'INFERENCE' -and $Target -eq 'FACT') -or
        ($Source -eq 'PREDICTION' -and $Target -eq 'PROOF')) {
        return 'INVALID'
    }
    return 'VALID'
}

function Check-ProofPromotion([string]$Current, [string]$Requested) {
    if ($Current -eq 'SIMULATED' -and $Requested -eq 'HARDWARE') { return 'INVALID' }
    return 'VALID'
}

function Check-Unknown([string]$Evidence) {
    if ([string]::IsNullOrWhiteSpace($Evidence)) { return 'NOT_PROVEN' }
    return 'PROVEN'
}

Assert-Equal 'ARITHMETIC 1+1=2' (Evaluate-Rule 1 1 2) 'VALID'
Assert-Equal 'ARITHMETIC 1+1=3' (Evaluate-Rule 1 1 3) 'INVALID'
Assert-Equal 'INFERENCE->FACT' (Check-Classification 'INFERENCE' 'FACT') 'INVALID'
Assert-Equal 'PREDICTION->PROOF' (Check-Classification 'PREDICTION' 'PROOF') 'INVALID'
Assert-Equal 'SIMULATED->HARDWARE' (Check-ProofPromotion 'SIMULATED' 'HARDWARE') 'INVALID'
Assert-Equal 'UNKNOWN EVIDENCE' (Check-Unknown '') 'NOT_PROVEN'

# Contradiction: same formal input must not silently select two different results.
$r1 = Add-Exact 1 1
$r2 = Add-Exact 1 1
Assert-Equal 'DETERMINISM' "$r1/$r2" '2/2'

Write-Host 'ZERO_TRUTH_INVARIANTS_RESULT=PASS'
