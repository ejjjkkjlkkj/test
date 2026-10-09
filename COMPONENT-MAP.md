# ZERO Component Map — source audit

Date: 2026-10-09
Target repository: `ejjjkkjlkkj/test` (ZERO), branch `main`.

## Evidence rules

- This file records source-level findings from repository files fetched from GitHub.
- A README claim is not equivalent to a successful build, runtime observation, or physical-hardware proof.
- No source is considered integrated into ZERO merely because it is listed here.
- Reuse is gated on exact source revision, dependency/license review, tests, and a defined adapter boundary.
- `QEMU/VM PASS != physical hardware PASS`.
- If a file or manifest was not found at the checked path, its license/build status remains `UNPROVEN`; do not infer it.

## Candidate component inventory

| Candidate | Exact source ref checked | Relevant source evidence | Candidate role in ZERO | Initial status | Main gate |
|---|---|---|---|---|---|
| ADMWS12 | `main`; capability blob `852c6e852918f71c804ed5a9f766900bcc79bb2f`; evidence blob `45f70957a19995685f85d9a55ee64487fc3cf0da`; HAL blob `dc4f853248baee4368c65e5861e1eedad2707f39` | `src/platform/capability.py`, `src/platform/evidence.py`, `src/platform/hal.py` | Semantic reference for capability states, evidence states, lifecycle and degraded operation; port behavior to ZERO-native implementation, not a parallel runtime | `PARTIAL / SOURCE REVIEWED` | Python implementation; no project license file found at root; test suite and all state transitions need exact-commit execution/review |
| omni-os | `main`; README blob `3dc0cc2ba7c20b83c041c8b13bf01817d00ad45a`; LICENSE blob `01506c469a30ac12c42a12e7fae25c5e33a3289f` | README describes Rust UEFI loader, x86-64 kernel, HII/IFR screen reader, native audio and QEMU boot claims | Primary candidate for boot/loader/kernel architecture and pre-OS accessibility research | `PARTIAL / CLAIMS RECORDED` | Root `Cargo.toml` not found at checked path; locate workspace manifests and per-crate tests; independently verify workflow results; QEMU evidence is not physical proof |
| accessible-windows | README on `main`; README identifies coherent line `repo-clean-consolidation-20260924`, consolidation commit `b53b62895ddb57094aa944cb1cdc549f0523668d`; cited baseline commit `8302a95b0b4d2fe4d57e2832bcc945819728b80d` | README describes UEFI screen reader, HII/IFR adapter, HDA integration and states parity issue #4 remains open; 24 kHz target not PASS | Candidate source for firmware semantic navigation, pre-OS screen reader and native audio integration | `PARTIAL / BRANCH SELECTION REQUIRED` | Audit the named consolidation branch directly; establish license and exact build/test evidence; parity is explicitly not closed |
| NVDA-RUST-UIA-STANDALONE | `main`; Cargo blob `cdc0d994e359eb09aa153b02ba103a383cded3f2`; README blob `515d064b8b1c6ef04799a3b6e8a99f54153f23cf` | Rust UIA/COM backend; portable event model; focus, TextChanged and TextSelectionChanged events; conditional Windows dependencies | Windows semantic UI observation adapter for ZERO's accessible projection; must not become a second core | `PARTIAL / SOURCE REVIEWED` | License file not found at root; execute locked tests and Windows x64 build; validate UIA event lifecycle and security boundaries |
| u (ST speech) | `main`; Cargo blob `e4624290a34c94373934ce3793c0def4664165c0`; README blob `07e1e759df03d67abdc1abd301671583c687af8a` | Rust library/binary; compact formant synthesis without external runtime dependency; neural backend described as private Python model folder; PCM output | Candidate speech adapter for local/offline announcements; compact backend is the lowest-dependency first candidate | `PARTIAL / SOURCE REVIEWED` | License file not found at root; do not import Python neural runtime into a Rust-only base; benchmark latency, intelligibility and interruptibility on the physical target |
| solution | `main`; README blob `b404db3a8f74697e8700ce739901416cd20e1934`; LICENSE blob `65dc5132432a0d66a459879ed20208c1c94d14bd` | README specifies software verification gates, IFR mutation fuzzing, replay and explicit hardware-only boundary; license states 0BSD | Candidate verification methodology and test-gate design, not a competing runtime core | `PARTIAL / METHODOLOGY REVIEWED` | Root Cargo manifest not found; README gives Python test commands; inspect source/tests and CI before adopting any tool |
| omni-security | `main`; README blob `6520fb1c175e5e60f157f7707d5ffc1d75af0d03` | Security and accessibility treated as co-equal requirements across hardware-to-OS chain; threat, protected-read-only and privilege model documents referenced | Cross-cutting security/accessibility requirements and review gates | `PARTIAL / REQUIREMENTS SOURCE` | No implementation claim from README alone; inspect individual requirement IDs, threat model and test coverage before mapping to ZERO |
| ZERO existing machine record | `main`; `go.mod` blob `b071d4b865b5ca89f52c6a88fac1495f3d570b11` | Go module name `zero`, Go 1.23; existing canonical record parser/tests modified in prior commits | Canonical wire/data model remains authoritative; adapters must map into it | `SOURCE OF TRUTH` | Run `go test ./...` on the exact commit; no successful test run is evidenced by this audit yet |

## Integration order

1. **Semantic contract first:** map ADMWS12 capability/evidence states into ZERO's existing capability and proof model. Keep source state and provenance; never silently map unknown/failed to available.
2. **Verification gates second:** extract only reusable, license-cleared test methodology from `solution`; represent NOT_RUN and UNPROVEN as blockers.
3. **Pre-OS runtime research:** inspect `omni-os` and the exact `accessible-windows` consolidation branch at crate/file level before selecting a single source path for loader, HII/IFR, audio and kernel.
4. **Windows semantic observation:** evaluate UIA standalone as an adapter behind ZERO's input/observation interfaces.
5. **Speech:** prototype the compact Rust ST backend behind a replaceable speech interface; keep neural assets/runtime optional and local.
6. **Security gates:** map `omni-security` requirements to explicit ZERO capability authorization, proof and recovery checks.

## Not yet done

- No code from these candidate repositories has been copied or merged by this audit.
- No license conclusion is made for repositories whose root license file was not found.
- No tests or builds from the candidate repositories were executed during this source review.
- No physical hardware, firmware write, NVRAM write, USB format or reboot validation was performed.
- No performance, accessibility parity, or 99.99% reliability claim is established.


## Follow-up: exact accessible-windows status evidence (2026-10-09)

Directly fetched `docs/STATUS.md` from `accessible-windows@repo-clean-consolidation-20260924`, blob `780acbc2111171535f12277065af809b3fbdbd67`.

The source document reports:
- successful GitHub Actions run `35846636674` dated 2026-09-23 on a Windows self-hosted runner, using project commit `8302a95b0b4d2fe4d57e2832bcc945819728b80d`;
- reported markers for LocalSystem identity, native voice build, codec round-trip, NAVIGATION.EFI build, VMware UEFI boot and screen-reader discovery;
- 16 kHz speech output for that baseline;
- a later failure at run `35848464855` in the physical PsExec stage, after `STAGE=SYSTEM_VOICE_READY`;
- the 24 kHz target and screen-reader parity closure remain open.

**Evidence classification:** `SOURCE-REPORTED`, not independently re-run in this session. These historical markers are scoped to the exact commits and workflow runs above; they do not establish current HEAD success or final parity.

## Follow-up: adapter implementation gate

Created `ADMWS12-ADAPTER-IMPLEMENTATION.md` at commit `5babd9006d847a322a369dbed6ec0d7d99b17d48`. It defines loss-minimizing state mapping and tests. It is a specification only; adapter code has not been added and tests have not been run.
