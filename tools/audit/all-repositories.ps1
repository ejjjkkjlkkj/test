#requires -Version 7.0
[CmdletBinding()]
param(
    [string]$Owner = 'ejjjkkjlkkj',
    [string]$OutputRoot = (Join-Path (Get-Location) 'audit-output\all-repositories'),
    [int]$MaxRepositories = 0,
    [switch]$IncludeArchived
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

# Compatibility entry point. The canonical implementation is maintained at:
# scripts/audit/Invoke-ZeroMultiRepoInventory.ps1
$target = Join-Path $PSScriptRoot '..\..\scripts\audit\Invoke-ZeroMultiRepoInventory.ps1'
if (-not (Test-Path -LiteralPath $target -PathType Leaf)) {
    throw "BLOCKED: canonical inventory script not found: $target"
}

$params = @{
    Owner = $Owner
    OutputDirectory = [IO.Path]::GetFullPath($OutputRoot)
    MaxRepositories = $MaxRepositories
}
if ($IncludeArchived) {
    $params.IncludeArchived = $true
}
& $target @params
