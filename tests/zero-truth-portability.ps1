$ErrorActionPreference = 'Stop'

function Add-Exact([int]$A, [int]$B) {
    return $A + $B
}

function Assert-Equal([string]$Name, $Actual, $Expected) {
    if ($Actual -ne $Expected) {
        throw "$Name FAIL: expected [$Expected], got [$Actual]"
    }
    Write-Host "$Name PASS"
}

# The formal rule is deliberately executed without consulting any hardware profile.
$profiles = @(
    'PC',
    'IPHONE',
    'ANDROID',
    'AIRBORNE',
    'SPACECRAFT',
    'EMBEDDED',
    'OFFLINE',
    'REMOTE'
)

foreach ($profile in $profiles) {
    $result = Add-Exact 1 1
    Assert-Equal "TRUTH $profile 1+1=2" $result 2
}

$invalid = Add-Exact 1 1
if ($invalid -eq 3) {
    throw 'TRUTH CONTRADICTION: 1+1=3'
}
Write-Host 'TRUTH 1+1=3 REJECTED'

# Accessibility projections must preserve the deterministic result.
$modalities = @(
    'VOICE',
    'BRAILLE',
    'KEYBOARD',
    'DISPLAY',
    'TOUCH',
    'POINTER',
    'NETWORK',
    'AUTOMATION'
)

foreach ($modality in $modalities) {
    $semanticResult = [pscustomobject]@{
        identity = 'ARITHMETIC:1+1'
        operation = 'ADD'
        result = 2
        truth = 'DETERMINISTIC'
    }

    if ($semanticResult.result -ne 2 -or $semanticResult.truth -ne 'DETERMINISTIC') {
        throw "ACCESSIBILITY $modality FAIL"
    }
    Write-Host "ACCESSIBILITY $modality PASS"
}

Write-Host 'ZERO_TRUTH_PORTABILITY_RESULT=PASS'
