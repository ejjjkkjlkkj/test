# ZERO language — first executable slice

This crate is a real, minimal Rust reference interpreter for the initial ZERO-IR instruction format. It is **not yet a universal machine language, compiler, or bare-metal runtime**.

## Run

```powershell
cargo run --release --manifest-path crates/zero-language/Cargo.toml -- crates/zero-language/examples/hello.zero
```

Expected program output includes `299792459` and `STATUS=PASS`. Timing values are machine-specific and must not be compared across machines without controlling the environment.

## Benchmark interpreter overhead

```powershell
cargo run --release --manifest-path crates/zero-language/Cargo.toml -- crates/zero-language/examples/hello.zero --bench 100000
```

This measures repeated interpreter execution, including its current per-run variable-map allocation. It is not a hardware propagation measurement and not a compiler-vs-compiler benchmark.

## Initial instructions

- `const NAME INTEGER`
- `add DEST LEFT RIGHT`
- `sub DEST LEFT RIGHT`
- `mul DEST LEFT RIGHT`
- `print NAME`
- `# comment`

Undefined names, malformed instructions, and checked integer overflow fail explicitly.

## Status

- Source committed: yes.
- Runtime tests executed by the author/CI: not yet confirmed for this revision.
- Native compiler/backend: not implemented.
- Bare-metal execution: not implemented.
- Universal hardware support: unproven.
- Sound/light comparison: this interpreter reports elapsed wall-clock time only. A separate experiment must define distance, medium, instrumentation, uncertainty, and comparison baseline before making a propagation claim.
