use std::{collections::HashMap, env, fs, process, time::Instant};

#[derive(Debug)]
enum Op {
    Const(String, i64),
    Add(String, String, String),
    Sub(String, String, String),
    Mul(String, String, String),
    Print(String),
}

fn parse(source: &str) -> Result<Vec<Op>, String> {
    let mut ops = Vec::new();
    for (idx, raw) in source.lines().enumerate() {
        let line = raw.split('#').next().unwrap_or("").trim();
        if line.is_empty() { continue; }
        let t: Vec<&str> = line.split_whitespace().collect();
        let err = || format!("line {}: invalid instruction: {}", idx + 1, line);
        let op = match t.as_slice() {
            ["const", dst, value] => Op::Const((*dst).into(), value.parse().map_err(|_| err())?),
            ["add", dst, a, b] => Op::Add((*dst).into(), (*a).into(), (*b).into()),
            ["sub", dst, a, b] => Op::Sub((*dst).into(), (*a).into(), (*b).into()),
            ["mul", dst, a, b] => Op::Mul((*dst).into(), (*a).into(), (*b).into()),
            ["print", name] => Op::Print((*name).into()),
            _ => return Err(err()),
        };
        ops.push(op);
    }
    Ok(ops)
}

fn value(vars: &HashMap<String, i64>, name: &str, line: usize) -> Result<i64, String> {
    vars.get(name).copied().ok_or_else(|| format!("instruction {}: undefined value '{}'", line, name))
}

fn run(ops: &[Op], repeat: usize, print_output: bool) -> Result<i64, String> {
    let mut final_value = 0;
    for _ in 0..repeat {
        let mut vars: HashMap<String, i64> = HashMap::with_capacity(ops.len());
        for (i, op) in ops.iter().enumerate() {
            let n = i + 1;
            match op {
                Op::Const(dst, v) => { vars.insert(dst.clone(), *v); }
                Op::Add(dst, a, b) => {
                    let v = value(&vars, a, n)?.checked_add(value(&vars, b, n)?)
                        .ok_or_else(|| format!("instruction {}: integer overflow", n))?;
                    vars.insert(dst.clone(), v);
                }
                Op::Sub(dst, a, b) => {
                    let v = value(&vars, a, n)?.checked_sub(value(&vars, b, n)?)
                        .ok_or_else(|| format!("instruction {}: integer overflow", n))?;
                    vars.insert(dst.clone(), v);
                }
                Op::Mul(dst, a, b) => {
                    let v = value(&vars, a, n)?.checked_mul(value(&vars, b, n)?)
                        .ok_or_else(|| format!("instruction {}: integer overflow", n))?;
                    vars.insert(dst.clone(), v);
                }
                Op::Print(name) => {
                    let v = value(&vars, name, n)?;
                    final_value = v;
                    if print_output { println!("{}", v); }
                }
            }
        }
    }
    Ok(final_value)
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 || args.iter().any(|a| a == "--help" || a == "-h") {
        println!("ZERO-IR reference interpreter\nUsage: zero-language <program.zero> [--bench N]\nInstructions: const NAME INT | add DST A B | sub DST A B | mul DST A B | print NAME\nLines beginning with # are comments.");
        return;
    }
    let path = &args[1];
    let repeat = args.windows(2).find(|w| w[0] == "--bench")
        .and_then(|w| w[1].parse::<usize>().ok()).unwrap_or(1).max(1);
    let source = match fs::read_to_string(path) {
        Ok(s) => s,
        Err(e) => { eprintln!("FAIL: cannot read {}: {}", path, e); process::exit(2); }
    };
    let ops = match parse(&source) {
        Ok(p) => p,
        Err(e) => { eprintln!("FAIL: {}", e); process::exit(2); }
    };
    if ops.is_empty() { eprintln!("FAIL: program contains no instructions"); process::exit(2); }
    let start = Instant::now();
    match run(&ops, repeat, repeat == 1) {
        Ok(last) => {
            let elapsed = start.elapsed();
            println!("STATUS=PASS");
            println!("ENGINE=ZERO-IR-reference-interpreter");
            println!("INSTRUCTIONS={}", ops.len());
            println!("REPETITIONS={}", repeat);
            println!("ELAPSED_NS={}", elapsed.as_nanos());
            println!("LAST_VALUE={}", last);
            if repeat > 1 {
                println!("NS_PER_RUN={}", elapsed.as_nanos() / repeat as u128);
            }
        }
        Err(e) => { eprintln!("FAIL: {}", e); process::exit(2); }
    }
}
