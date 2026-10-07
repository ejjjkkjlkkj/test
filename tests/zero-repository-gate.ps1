$ErrorActionPreference = 'Stop'

# Global acceptance gate.
# This gate is intentionally strict: specification-only components are not
# accepted as functional components.

$required = @(
    [pscustomobject]@{ Name='ZERO reference validator'; Functional=$true; Accessible=$true },
    [pscustomobject]@{ Name='ZERO deterministic truth'; Functional=$true; Accessible=$true },
    [pscustomobject]@{ Name='ZERO accessibility projection'; Functional=$true; Accessible=$true },

    # These are declared required but have no executable implementation yet.
    [pscustomobject]@{ Name='ZERO reference engine'; Functional=$true; Accessible=$true },
    [pscustomobject]@{ Name='ZERO machine ISA executor'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO machine serializer'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO machine deserializer'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO bootstrap runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO native object runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO event runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO capability runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO network runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO browser'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO radio runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO satellite runtime'; Functional=$false; Accessible=$false },
    [pscustomobject]@{ Name='ZERO AI-native runtime'; Functional=$false; Accessible=$false }
)

$failed = @()

foreach ($unit in $required) {
    if (-not ($unit.Functional -and $unit.Accessible)) {
        $failed += $unit.Name
        Write-Host "REJECTED: $($unit.Name) FUNCTIONAL=$($unit.Functional) ACCESSIBLE=$($unit.Accessible)"
    } else {
        Write-Host "ACCEPTED: $($unit.Name)"
    }
}

if ($failed.Count -gt 0) {
    Write-Host ''
    Write-Host "ZERO_REPOSITORY_GATE=FAIL"
    Write-Host "REJECTED_UNITS=$($failed.Count)"
    exit 1
}

Write-Host "ZERO_REPOSITORY_GATE=PASS"
