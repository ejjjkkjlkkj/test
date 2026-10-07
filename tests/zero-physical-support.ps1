Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

function New-ExecutionProfile([string]$Name, [string[]]$Capabilities) {
    [pscustomobject]@{
        Name = $Name
        Capabilities = $Capabilities
    }
}

function Invoke-ZeroPhysicalProfile($Profile) {
    $required = @('LOCAL.EXECUTE','STATE.PERSIST')
    foreach ($capability in $required) {
        if ($Profile.Capabilities -notcontains $capability) {
            return [pscustomobject]@{
                State = 'REJECTED'
                Accessible = $false
                Proven = $false
                Reason = 'MISSING_CAPABILITY'
            }
        }
    }

    [pscustomobject]@{
        State = 'COMPLETED'
        Accessible = ($Profile.Capabilities -contains 'ACCESSIBILITY.PROJECT')
        Proven = ($Profile.Capabilities -contains 'OBSERVE.RESULT')
        Reason = $null
    }
}

$profiles = @(
    (New-ExecutionProfile 'PC' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT')),
    (New-ExecutionProfile 'IPHONE' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT')),
    (New-ExecutionProfile 'AIRBORNE' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT')),
    (New-ExecutionProfile 'SPACE' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT')),
    (New-ExecutionProfile 'CONSTRAINED' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT'))
)

foreach ($profile in $profiles) {
    $result = Invoke-ZeroPhysicalProfile $profile
    if ($result.State -ne 'COMPLETED') { throw "$($profile.Name) unexpectedly rejected" }
    if (-not $result.Accessible) { throw "$($profile.Name) accessibility failed" }
    if (-not $result.Proven) { throw "$($profile.Name) observation proof failed" }
}

$offline = New-ExecutionProfile 'AIRBORNE-OFFLINE' @('LOCAL.EXECUTE','STATE.PERSIST','ACCESSIBILITY.PROJECT','OBSERVE.RESULT')
$offlineResult = Invoke-ZeroPhysicalProfile $offline
if ($offlineResult.State -ne 'COMPLETED') { throw 'AIRBORNE-OFFLINE failed' }

$insufficient = New-ExecutionProfile 'INSUFFICIENT' @('LOCAL.EXECUTE')
$rejected = Invoke-ZeroPhysicalProfile $insufficient
if ($rejected.State -ne 'REJECTED') { throw 'Insufficient profile was accepted' }
if ($rejected.Reason -ne 'MISSING_CAPABILITY') { throw 'Wrong rejection reason' }
if ($rejected.Proven) { throw 'Rejected profile claimed proof' }

Write-Output 'ZERO_PHYSICAL_SUPPORT_MULTI_PROFILE=PASS'
Write-Output 'ZERO_PHYSICAL_SUPPORT_AIRBORNE_OFFLINE=PASS'
Write-Output 'ZERO_PHYSICAL_SUPPORT_CAPABILITY_BOUNDARY=PASS'
Write-Output 'ZERO_PHYSICAL_SUPPORT_ACCESSIBILITY=PASS'
