mod archive;
mod challenge;
mod core;
mod discovery;
mod emergence;
mod evaluator;
mod experiment;
mod generator;
mod memory;
mod operator;
mod operator_lang;
mod search;
mod world;

use core::Mind;

fn main() {
    let mut mind=Mind::new();
    let args:Vec<String>=std::env::args().collect();
    match args.get(1).map(String::as_str) {
        Some("solve")=>{
            let problem=args.get(2..).unwrap_or_default().join(" ");
            if problem.trim().is_empty(){eprintln!("usage: cargo run -- solve \"problem\"");std::process::exit(2);}
            println!("{}",mind.solve(&problem).render());
        }
        Some("experiment")=>println!("{}",experiment::run()),
        Some("discover")=>println!("{}",experiment::run()),
        Some("emerge")=>println!("{}",emergence::report()),
        Some("benchmark")=>println!("{}",challenge::run(&mut mind)),
        _=>println!("TEST experimental intelligence core\n  solve <problem>\n  experiment\n  emerge\n  benchmark"),
    }
}
