Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

$root = Split-Path -Parent $PSScriptRoot
$forbidden = @(
    ('oll' + 'ama'),
    ('zero-' + 'ai'),
    ('local-' + 'ai'),
    ('LOCAL-' + 'AI.md')
)

$extensions = @('.go','.md','.ps1','.zero','.txt','.json','.yml','.yaml','.c','.h','.sh')
$violations = @()

foreach ($path in Get-ChildItem -LiteralPath $root -Recurse -File -Force) {
    if ($path.FullName -like "$root\.git\*") { continue }
    if ($extensions -notcontains $path.Extension.ToLowerInvariant()) { continue }

    $text = Get-Content -LiteralPath $path.FullName -Raw -ErrorAction Stop
    foreach ($needle in $forbidden) {
        if ($text.IndexOf($needle, [System.StringComparison]::OrdinalIgnoreCase) -ge 0) {
            $violations += "$($path.FullName):$needle"
        }
    }
}

if ($violations.Count -gt 0) {
    $violations | ForEach-Object { Write-Error "FORBIDDEN_COMPONENT: $_" }
    exit 1
}

Write-Output 'ZERO_COMPONENT_CLEAN_RESULT=PASS'
