# ZERO propagation-reference experiment

This executable measures a local CPU workload and compares its measured wall-clock duration with calculated sound/light travel-time references for a user-specified distance. It does **not** measure physical signal propagation.

Run from the repository root:

```powershell
cargo run --release --manifest-path crates/zero-propagation/Cargo.toml -- --distance-m 100 --temperature-c 20 --iterations 10000000
```

Use a distance appropriate to the question. The light reference is for vacuum at exactly 299,792,458 m/s. The sound reference uses the approximation 331.3 + 0.606*T m/s for air and the temperature entered by the user; it does not measure local air conditions.

The program emits labelled key/value output including measured compute duration, reference travel times, ratios, and checksum. Repeat runs and report variance; CPU scheduling, thermal state, power policy, and timer resolution affect results. For physical propagation proof, use a calibrated physical emitter/receiver setup and report uncertainty.

See `docs/physics/zero-real-sound-light-references.md` for sources and limitations.
