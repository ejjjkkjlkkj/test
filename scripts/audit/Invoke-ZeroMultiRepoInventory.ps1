# ZERO multi-repository inventory (read-only)
# Requires Git, GitHub CLI (gh), and an authenticated gh session.
# Inventories accessible repositories, every listed branch, and recursive Git trees.
# Does not clone, checkout, push, delete, or modify remote repositories.
# Status INVENTORIED means metadata captured; it does not mean source code was semantically audited.

[CmdletBinding()]
param(
    [string]$OutputDirectory = (Join-Path (Get-Location) ("audit-output\multi-repo-" + (Get-Date -Format "yyyyMMdd-HHmmss"))),
    [string]$Owner = "",
    [int]$MaxRepositories = 0,
    [switch]$IncludeArchived
)

$ErrorActionPreference = "Stop"
$script:Rows = [System.Collections.Generic.List[object]]::new()
$script:Failures = [System.Collections.Generic.List[object]]::new()

function Invoke-GhJson {
    param([Parameter(Mandatory)][string[]]$Arguments)
    $raw = & gh @Arguments 2>&1
    if ($LASTEXITCODE -ne 0) {
        throw ("gh failed (code {0}): {1}" -f $LASTEXITCODE, ($raw -join [Environment]::NewLine))
    }
    $text = ($raw -join [Environment]::NewLine).Trim()
    if ([string]::IsNullOrWhiteSpace($text)) { return $null }
    return ($text | ConvertFrom-Json -Depth 100)
}

function Get-PaginatedArray {
    param([Parameter(Mandatory)][string]$Endpoint)
    $raw = & gh api --paginate --slurp $Endpoint 2>&1
    if ($LASTEXITCODE -ne 0) {
        throw ("gh pagination failed (code {0}): {1}" -f $LASTEXITCODE, ($raw -join [Environment]::NewLine))
    }
    $text = ($raw -join [Environment]::NewLine).Trim()
    $pages = @()
    if (-not [string]::IsNullOrWhiteSpace($text)) {
        $pages = ConvertFrom-Json -InputObject $text -Depth 100 -NoEnumerate
    }
    $items = [System.Collections.Generic.List[object]]::new()
    foreach ($page in $pages) {
        if ($null -eq $page) { continue }
        foreach ($item in $page) { if ($null -ne $item) { $items.Add($item) } }
    }
    return $items.ToArray()
}

function Write-JsonLines {
    param([Parameter(Mandatory)][string]$Path, [Parameter(Mandatory)][object[]]$Items)
    $utf8NoBom = [System.Text.UTF8Encoding]::new($false)
    $writer = [System.IO.StreamWriter]::new($Path, $false, $utf8NoBom)
    try {
        foreach ($item in $Items) {
            $writer.WriteLine(($item | ConvertTo-Json -Depth 30 -Compress))
        }
    } finally { $writer.Dispose() }
}

Write-Host "[INFO] ZERO multi-repository inventory — READ ONLY" -ForegroundColor Cyan
Write-Host "[INFO] Output: $OutputDirectory"
Write-Host "[INFO] API requests are sequential to protect rate limits."
if (-not (Get-Command gh -ErrorAction SilentlyContinue)) {
    Write-Host "[FAIL] GitHub CLI (gh) is not installed or not on PATH." -ForegroundColor Red
} else {
    try {
        $null = Invoke-GhJson -Arguments @("api", "user")
        New-Item -ItemType Directory -Path $OutputDirectory -Force | Out-Null
        $reposEndpoint = "user/repos?per_page=100&affiliation=owner,collaborator,organization_member&sort=full_name&direction=asc"
        $repositories = @(Get-PaginatedArray -Endpoint $reposEndpoint)
        if ($Owner) {
            $repositories = @($repositories | Where-Object { $_.owner.login -ieq $Owner })
        }
        if (-not $IncludeArchived) {
            $repositories = @($repositories | Where-Object { -not $_.archived })
        }
        if ($MaxRepositories -gt 0) {
            $repositories = @($repositories | Select-Object -First $MaxRepositories)
        }
        Write-Host ("[INFO] Repositories selected: {0}" -f $repositories.Count)

        $repoIndex = 0
        foreach ($repo in $repositories) {
            $repoIndex++
            $fullName = [string]$repo.full_name
            $repoHash = [Convert]::ToHexString([System.Security.Cryptography.SHA256]::HashData([System.Text.Encoding]::UTF8.GetBytes($fullName))).Substring(0, 16).ToLowerInvariant()
            $safeRepo = (($fullName -replace '[^A-Za-z0-9._-]', '_') + "-" + $repoHash)
            $repoDir = Join-Path $OutputDirectory $safeRepo
            New-Item -ItemType Directory -Path $repoDir -Force | Out-Null
            Write-Host ("[{0}/{1}] {2}" -f $repoIndex, $repositories.Count, $fullName)

            try {
                $branches = @(Get-PaginatedArray -Endpoint ("repos/{0}/branches?per_page=100" -f $fullName))
                $branchManifest = foreach ($branch in $branches) {
                    [pscustomobject]@{
                        repository = $fullName
                        branch = [string]$branch.name
                        commit_sha = [string]$branch.commit.sha
                        protected = [bool]$branch.protected
                    }
                }
                $branchManifest | ConvertTo-Json -Depth 20 | Set-Content -LiteralPath (Join-Path $repoDir "branches.json") -Encoding utf8

                foreach ($branch in $branches) {
                    $branchName = [string]$branch.name
                    $commitSha = [string]$branch.commit.sha
                    $branchHash = [Convert]::ToHexString([System.Security.Cryptography.SHA256]::HashData([System.Text.Encoding]::UTF8.GetBytes($branchName))).Substring(0, 16).ToLowerInvariant()
                    $branchSafe = (($branchName -replace '[^A-Za-z0-9._-]', '_') + "-" + $branchHash)
                    $status = "INVENTORIED"
                    $treeSha = ""
                    $fileCount = 0
                    $totalBytes = [long]0
                    $truncated = $false
                    $manifestPath = Join-Path $repoDir ("tree-" + $branchSafe + ".jsonl")
                    try {
                        $commit = Invoke-GhJson -Arguments @("api", ("repos/{0}/git/commits/{1}" -f $fullName, $commitSha))
                        $treeSha = [string]$commit.tree.sha
                        $tree = Invoke-GhJson -Arguments @("api", ("repos/{0}/git/trees/{1}?recursive=1" -f $fullName, $treeSha))
                        $truncated = [bool]$tree.truncated
                        $entries = @($tree.tree)
                        $fileEntries = @($entries | Where-Object { $_.type -eq "blob" })
                        $fileCount = $fileEntries.Count
                        foreach ($entry in $fileEntries) {
                            if ($null -ne $entry.size) { $totalBytes += [long]$entry.size }
                        }
                        $manifestItems = foreach ($entry in $entries) {
                            [pscustomobject]@{
                                repository = $fullName
                                branch = $branchName
                                commit_sha = $commitSha
                                tree_sha = $treeSha
                                path = [string]$entry.path
                                type = [string]$entry.type
                                mode = [string]$entry.mode
                                object_sha = [string]$entry.sha
                                size_bytes = $entry.size
                            }
                        }
                        Write-JsonLines -Path $manifestPath -Items @($manifestItems)
                        if ($truncated) { $status = "PARTIAL_TREE_TRUNCATED" }
                    } catch {
                        $status = "FAILED"
                        $script:Failures.Add([pscustomobject]@{
                            repository = $fullName
                            branch = $branchName
                            stage = "tree"
                            error = $_.Exception.Message
                        })
                        Set-Content -LiteralPath $manifestPath -Value ("# FAILED: " + $_.Exception.Message) -Encoding utf8
                    }
                    $script:Rows.Add([pscustomobject]@{
                        repository = $fullName
                        branch = $branchName
                        commit_sha = $commitSha
                        tree_sha = $treeSha
                        file_count = $fileCount
                        total_blob_bytes = $totalBytes
                        tree_truncated = $truncated
                        status = $status
                        manifest = (Resolve-Path -LiteralPath $manifestPath -ErrorAction SilentlyContinue).Path
                    })
                    Start-Sleep -Milliseconds 150
                }
                if ($branches.Count -eq 0) {
                    $script:Rows.Add([pscustomobject]@{
                        repository = $fullName; branch = ""; commit_sha = ""; tree_sha = ""
                        file_count = 0; total_blob_bytes = 0; tree_truncated = $false
                        status = "NO_BRANCHES_RETURNED"; manifest = ""
                    })
                }
            } catch {
                $script:Failures.Add([pscustomobject]@{
                    repository = $fullName
                    branch = ""
                    stage = "branches"
                    error = $_.Exception.Message
                })
                $script:Rows.Add([pscustomobject]@{
                    repository = $fullName; branch = ""; commit_sha = ""; tree_sha = ""
                    file_count = 0; total_blob_bytes = 0; tree_truncated = $false
                    status = "FAILED_BRANCH_ENUMERATION"; manifest = ""
                })
                Write-Host ("[FAIL] {0}: {1}" -f $fullName, $_.Exception.Message) -ForegroundColor Red
            }
        }

        $script:Rows | Export-Csv -LiteralPath (Join-Path $OutputDirectory "branch-summary.csv") -NoTypeInformation -Encoding utf8
        $script:Failures | ConvertTo-Json -Depth 20 | Set-Content -LiteralPath (Join-Path $OutputDirectory "failures.json") -Encoding utf8
        $runInfo = [pscustomobject]@{
            created_utc = [DateTime]::UtcNow.ToString("o")
            mode = "READ_ONLY"
            account = (Invoke-GhJson -Arguments @("api", "user")).login
            repositories_selected = $repositories.Count
            branch_rows = $script:Rows.Count
            failures = $script:Failures.Count
            complete_semantic_audit = $false
            note = "Inventory only. No source-level semantic audit, build, or tests were run."
        }
        $runInfo | ConvertTo-Json -Depth 10 | Set-Content -LiteralPath (Join-Path $OutputDirectory "run-summary.json") -Encoding utf8
        Write-Host ("[PASS] Inventaire écrit: {0}" -f $OutputDirectory) -ForegroundColor Green
        Write-Host ("[INFO] Repositories={0}; branches={1}; failures={2}" -f $repositories.Count, $script:Rows.Count, $script:Failures.Count)
        if (@($script:Rows | Where-Object { $_.status -eq "PARTIAL_TREE_TRUNCATED" }).Count -gt 0) {
            Write-Host "[PARTIAL] Au moins un arbre GitHub est tronqué; ces dépôts nécessitent une analyse locale Git complémentaire." -ForegroundColor Yellow
        }
        if ($script:Failures.Count -gt 0) {
            Write-Host "[PARTIAL] Consulte failures.json; les erreurs ne sont pas masquées." -ForegroundColor Yellow
        }
    } catch {
        Write-Host ("[FAIL] {0}" -f $_.Exception.Message) -ForegroundColor Red
        Write-Host "[INFO] Aucun changement distant effectué. Vérifier gh auth status et relancer." -ForegroundColor Yellow
    }
}
Write-Host "[INFO] Fin du script. La session PowerShell reste ouverte." -ForegroundColor Cyan
