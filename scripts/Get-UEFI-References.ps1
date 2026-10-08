#requires -Version 7.0
[CmdletBinding()]
param(
  [string]$Root = (Join-Path $PWD 'UEFI-Reference')
)

$ErrorActionPreference = 'Stop'
New-Item -ItemType Directory -Force -Path $Root | Out-Null

$repos = @(
  @{ Name='EDK2'; Url='https://github.com/tianocore/edk2.git' },
  @{ Name='TianoCore-Docs'; Url='https://github.com/tianocore-docs/Docs.git' },
  @{ Name='Driver-Writer'; Url='https://github.com/tianocore-docs/edk2-UefiDriverWritersGuide.git' },
  @{ Name='VFR'; Url='https://github.com/tianocore-docs/edk2-VfrSpecification.git' },
  @{ Name='INF'; Url='https://github.com/tianocore-docs/edk2-InfSpecification.git' },
  @{ Name='DSC'; Url='https://github.com/tianocore-docs/edk2-DscSpecification.git' },
  @{ Name='DEC'; Url='https://github.com/tianocore-docs/edk2-DecSpecification.git' },
  @{ Name='UNI'; Url='https://github.com/tianocore-docs/edk2-UniSpecification.git' },
  @{ Name='Secure-Coding'; Url='https://github.com/tianocore-docs/EDK_II_Secure_Coding_Guide.git' },
  @{ Name='Secure-Boot-Chain'; Url='https://github.com/tianocore-docs/Understanding_UEFI_Secure_Boot_Chain.git' }
)

foreach ($r in $repos) {
  $dst = Join-Path $Root $r.Name
  if (Test-Path (Join-Path $dst '.git')) {
    Write-Host "PRESENT $($r.Name)"
    continue
  }
  if (Test-Path $dst) {
    throw "Destination exists but is not a Git repository: $dst"
  }
  git clone --filter=blob:none --no-checkout $r.Url $dst
  if ($LASTEXITCODE -ne 0) { throw "FAIL clone $($r.Name)" }
  Write-Host "OK $($r.Name)"
}

# Official UEFI specification source is reference-only; keep it outside the public
# collection unless its redistribution terms have been explicitly cleared.
$spec = Join-Path $Root 'OFFICIAL-REFERENCE-UEFI-Specification-Release'
if (-not (Test-Path $spec)) {
  git clone --filter=blob:none --no-checkout https://github.com/UEFI/UEFI-Specification-Release.git $spec
  if ($LASTEXITCODE -ne 0) { throw 'FAIL clone UEFI Specification Release' }
  Write-Host 'OK OFFICIAL-REFERENCE-UEFI-Specification-Release'
}

Write-Host ''
Write-Host 'PASS: références UEFI/TianoCore récupérées sans suppression ni modification BIOS.'
