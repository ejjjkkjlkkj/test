#requires -Version 7.0
[CmdletBinding()]
param(
    [string]$Owner = 'ejjjkkjlkkj',
    [string]$OutputRoot = (Join-Path $PSScriptRoot '..\..\audit-output\all-repositories'),
    [int]$MaxRepositories = 0
)

Set-StrictMode -Version Latest
$ErrorActionPreference = 'Stop'

# Read-only remote audit. This script creates local reports only; it never
# creates, updates, deletes, merges, or rewrites any GitHub branch or file.

$gh = Get-Command gh -ErrorAction SilentlyContinue
if (-not $gh) {
    throw 'BLOCKED: GitHub CLI (gh) is required.'
}
& gh auth status *> $null
if ($LASTEXITCODE -ne 0) {
    throw 'BLOCKED: run gh auth login with an account that can read all target repositories.'
}

$OutputRoot = [IO.Path]::GetFullPath($OutputRoot)
$manifestRoot = Join-Path $OutputRoot 'manifests'
New-Item -ItemType Directory -Path $manifestRoot -Force | Out-Null

$tab = [string][char]9
$repoLines = @(& gh api --paginate "user/repos?per_page=100&affiliation=owner,collaborator,organization_member" --jq '.[] | [.full_name, .default_branch, (.private|tostring)] | @tsv')
if ($LASTEXITCODE -ne 0) {
    throw "BLOCKED: could not enumerate repositories for owner $Owner."
}

$repositories = foreach ($line in $repoLines) {
    if ([string]::IsNullOrWhiteSpace($line)) { continue }
    $parts = $line -split $tab, 3
    if ($parts.Count -ge 3 -and $parts[0] -like "$Owner/*") {
        [pscustomobject]@{
            Name = $parts[0]
            DefaultBranch = $parts[1]
            Private = $parts[2]
        }
    }
}
if ($MaxRepositories -gt 0) {
    $repositories = @($repositories | Select-Object -First $MaxRepositories)
}

$branchRows = [Collections.Generic.List[object]]::new()
$treeCache = @{}
$repoIndex = 0

foreach ($repository in $repositories) {
    $repoIndex++
    Write-Progress -Activity 'ZERO cross-repository inventory' -Status "$repoIndex / $($repositories.Count): $($repository.Name)" -PercentComplete (($repoIndex / [Math]::Max(1, $repositories.Count)) * 100)

    $branchLines = @(& gh api --paginate "repos/$($repository.Name)/branches?per_page=100" --jq '.[] | [.name, .commit.sha] | @tsv')
    if ($LASTEXITCODE -ne 0) {
        $branchRows.Add([pscustomobject]@{
            Repository = $repository.Name; Branch = ''; CommitSHA = ''; TreeSHA = ''
            Files = 0; Bytes = 0; Truncated = ''; Status = 'BRANCH_ENUMERATION_FAILED'
            Manifest = ''
        })
        continue
    }

    foreach ($line in $branchLines) {
        if ([string]::IsNullOrWhiteSpace($line)) { continue }
        $parts = $line -split $tab, 2
        if ($parts.Count -ne 2) { continue }
        $branch = $parts[0]
        $commit = $parts[1]
        $cacheKey = "$($repository.Name)|$commit"

        if (-not $treeCache.ContainsKey($cacheKey)) {
            $treeJson = (& gh api "repos/$($repository.Name)/git/trees/$commit?recursive=1" | Out-String)
            if ($LASTEXITCODE -ne 0) {
                $treeCache[$cacheKey] = [pscustomobject]@{
                    TreeSHA = ''; Files = 0; Bytes = 0; Truncated = ''
                    Status = 'TREE_FETCH_FAILED'; Manifest = ''
                }
            }
            else {
                try {
                    $tree = $treeJson | ConvertFrom-Json -Depth 100
                    $blobs = @($tree.tree | Where-Object { $_.type -eq 'blob' })
                    $bytes = [long]0
                    foreach ($blob in $blobs) {
                        if ($null -ne $blob.size) { $bytes += [long]$blob.size }
                    }

                    $repoSlug = $repository.Name -replace '[^A-Za-z0-9_.-]', '_'
                    $manifestName = "$repoSlug-$commit.jsonl"
                    $manifestPath = Join-Path $manifestRoot $manifestName
                    $writer = [IO.StreamWriter]::new($manifestPath, $false, [Text.UTF8Encoding]::new($false))
                    try {
                        foreach ($blob in $blobs) {
                            $row = [ordered]@{
                                repository = $repository.Name
                                commit_sha = $commit
                                tree_sha = [string]$tree.sha
                                path = [string]$blob.path
                                mode = [string]$blob.mode
                                blob_sha = [string]$blob.sha
                                size_bytes = $blob.size
                            }
                            $writer.WriteLine(($row | ConvertTo-Json -Compress -Depth 10))
                        }
                    }
                    finally {
                        $writer.Dispose()
                    }

                    $treeCache[$cacheKey] = [pscustomobject]@{
                        TreeSHA = [string]$tree.sha
                        Files = $blobs.Count
                        Bytes = $bytes
                        Truncated = [bool]$tree.truncated
                        Status = $(if ($tree.truncated) { 'TREE_TRUNCATED' } else { 'PASS' })
                        Manifest = $manifestName
                    }
                }
                catch {
                    $treeCache[$cacheKey] = [pscustomobject]@{
                        TreeSHA = ''; Files = 0; Bytes = 0; Truncated = ''
                        Status = 'TREE_PARSE_FAILED'; Manifest = ''
                    }
                }
            }
        }

        $cached = $treeCache[$cacheKey]
        $branchRows.Add([pscustomobject]@{
            Repository = $repository.Name
            Branch = $branch
            CommitSHA = $commit
            TreeSHA = $cached.TreeSHA
            Files = $cached.Files
            Bytes = $cached.Bytes
            Truncated = $cached.Truncated
            Status = $cached.Status
            Manifest = $cached.Manifest
        })
    }
}

Write-Progress -Activity 'ZERO cross-repository inventory' -Completed
$summaryPath = Join-Path $OutputRoot 'summary.tsv'
$branchRows | Export-Csv -LiteralPath $summaryPath -Delimiter $tab -NoTypeInformation -Encoding utf8

$report = [pscustomobject]@{
    GeneratedUTC = [DateTime]::UtcNow.ToString('o')
    Owner = $Owner
    RepositoryCount = @($repositories).Count
    BranchRows = $branchRows.Count
    UniqueRepositoryCommitTrees = $treeCache.Count
    FilesInventoried = [long](($branchRows | Measure-Object -Property Files -Sum).Sum)
    FailedRows = @($branchRows | Where-Object { $_.Status -ne 'PASS' }).Count
    SummaryPath = $summaryPath
    ManifestDirectory = $manifestRoot
    RemoteMutations = 'NONE'
}
$report | ConvertTo-Json -Depth 5 | Set-Content -LiteralPath (Join-Path $OutputRoot 'report.json') -Encoding utf8
$report
