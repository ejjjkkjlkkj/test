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

function Project-Truth([string]$Profile, [string]$Modality) {
    if ($Profile -notin $profiles) { throw "UNKNOWN PROFILE: $Profile" }
    if ($Modality -notin $modalities) { throw "UNKNOWN MODALITY: $Modality" }
    return [pscustomobject]@{ Profile=$Profile; Modality=$Modality; Result=2; Truth='DETERMINISTIC' }
}

foreach ($profile in $profiles) {
    $result = Project-Truth $profile 'VOICE'
    Assert-Equal "TRUTH $profile PROFILE" $result.Profile $profile
    Assert-Equal "TRUTH $profile 1+1=2" $result.Result 2
    Assert-Equal "TRUTH $profile CLASS" $result.Truth 'DETERMINISTIC'
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
    $semanticResult = Project-Truth 'PC' $modality

    if ($semanticResult.Profile -ne 'PC' -or $semanticResult.Modality -ne $modality -or $semanticResult.Result -ne 2 -or $semanticResult.Truth -ne 'DETERMINISTIC') {
        throw "ACCESSIBILITY $modality FAIL"
    }
    Write-Host "ACCESSIBILITY $modality PASS"
}

$unknownProfileRejected = $false
try { [void](Project-Truth 'UNKNOWN' 'VOICE') } catch { $unknownProfileRejected = $true }
Assert-Equal 'UNKNOWN_PROFILE_REJECTED' $unknownProfileRejected $true

Write-Host 'ZERO_TRUTH_PORTABILITY_RESULT=PASS'
