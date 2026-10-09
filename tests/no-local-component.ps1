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

    # Audit reports, inventories, and historical state snapshots are retained
    # as evidence, not runtime dependencies. Keep them intact and out of this
    # source-component gate; scan implementation and active configuration files.
    if ($path.Name -like 'RAPPEL-ETAT-*.md' -or
        $path.Name -like '*AUDIT*.md' -or
        $path.Name -like '*INVENTORY*.md' -or
        $path.Name -like '*FILE-INVENTORY-*.md' -or
        $path.Name -eq 'COMPONENT-MAP.md') { continue }

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
