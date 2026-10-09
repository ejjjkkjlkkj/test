use std::{env, hint::black_box, process, time::Instant};

const C_VACUUM_M_S: f64 = 299_792_458.0;

fn arg_value(args: &[String], name: &str, default: &str) -> String {
    args.windows(2).find(|w| w[0] == name).map(|w| w[1].clone()).unwrap_or_else(|| default.to_string())
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.iter().any(|a| a == "--help" || a == "-h") {
        println!("ZERO sound/light reference comparison (not a propagation measurement)\nUsage: zero-propagation --distance-m 100 --temperature-c 20 --iterations 10000000\nOptions: --distance-m <positive metres> --temperature-c <Celsius> --iterations <positive integer>");
        return;
    }
    let distance: f64 = match arg_value(&args, "--distance-m", "100").parse() {
        Ok(v) if v.is_finite() && v > 0.0 => v,
        _ => { eprintln!("STATUS=INVALID_INPUT\nERROR=distance must be a finite positive number"); process::exit(2); }
    };
    let temp: f64 = match arg_value(&args, "--temperature-c", "20").parse() {
        Ok(v) if v.is_finite() && (-100.0..=100.0).contains(&v) => v,
        _ => { eprintln!("STATUS=INVALID_INPUT\nERROR=temperature must be between -100 and 100 C"); process::exit(2); }
    };
    let iterations: u64 = match arg_value(&args, "--iterations", "10000000").parse() {
        Ok(v) if v > 0 => v,
        _ => { eprintln!("STATUS=INVALID_INPUT\nERROR=iterations must be a positive integer"); process::exit(2); }
    };

    // Approximation for dry air near ordinary conditions; not an environmental measurement.
    let sound_speed = 331.3 + 0.606 * temp;
    let light_time_s = distance / C_VACUUM_M_S;
    let sound_time_s = distance / sound_speed;

    // Deterministic integer workload; black_box discourages dead-code elimination.
    let start = Instant::now();
    let mut x: u64 = 0x9E3779B97F4A7C15;
    for i in 0..iterations {
        x = black_box(x.rotate_left(7).wrapping_add(i ^ 0xD1B54A32D192ED03));
        x = black_box(x ^ x.wrapping_mul(0x94D049BB133111EB));
    }
    black_box(x);
    let elapsed = start.elapsed();
    let measured_s = elapsed.as_secs_f64();

    println!("FORMAT=ZERO-PROPAGATION-REFERENCE-v1");
    println!("STATUS=PASS");
    println!("MEASUREMENT=local_compute_wall_time");
    println!("PHYSICAL_PROPAGATION_MEASURED=false");
    println!("DISTANCE_M={:.9}", distance);
    println!("AIR_TEMPERATURE_C={:.3}", temp);
    println!("SOUND_SPEED_APPROX_M_S={:.6}", sound_speed);
    println!("LIGHT_SPEED_VACUUM_M_S={:.0}", C_VACUUM_M_S);
    println!("ITERATIONS={}", iterations);
    println!("COMPUTE_TIME_NS={}", elapsed.as_nanos());
    println!("COMPUTE_TIME_S={:.12}", measured_s);
    println!("SOUND_REFERENCE_TIME_MS={:.9}", sound_time_s * 1e3);
    println!("LIGHT_VACUUM_REFERENCE_TIME_NS={:.9}", light_time_s * 1e9);
    println!("SOUND_REFERENCE_RATIO_COMPUTE_OVER={:.12}", measured_s / sound_time_s);
    println!("LIGHT_REFERENCE_RATIO_COMPUTE_OVER={:.12}", measured_s / light_time_s);
    println!("SOUND_RESULT={}", if measured_s < sound_time_s { "COMPUTE_BEFORE_REFERENCE" } else { "COMPUTE_NOT_BEFORE_REFERENCE" });
    println!("LIGHT_RESULT={}", if measured_s < light_time_s { "COMPUTE_BEFORE_REFERENCE" } else { "COMPUTE_NOT_BEFORE_REFERENCE" });
    println!("CHECKSUM={}", x);
    println!("LIMITATION=Comparison to calculated reference times only; not proof of physical signal propagation speed.");
}
