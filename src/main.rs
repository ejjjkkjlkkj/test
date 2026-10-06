mod challenge;
mod core;
mod memory;
mod world;

use core::Mind;

fn main() {
    let mut mind = Mind::new();

    let args: Vec<String> = std::env::args().collect();
    match args.get(1).map(String::as_str) {
        Some("solve") => {
            let problem = args.get(2..).unwrap_or_default().join(" ");
            if problem.trim().is_empty() {
                eprintln!("usage: cargo run -- solve \"problem\"");
                std::process::exit(2);
            }
            println!("{}", mind.solve(&problem).render());
        }
        Some("experiment") => {
            for problem in [
                "find a useful way to organize an unknown collection",
                "repair a process when one assumption becomes false",
                "create a better explanation when the first explanation fails",
            ] {
                println!("\n=== {} ===", problem);
                println!("{}", mind.solve(problem).render());
            }
        }
        Some("benchmark") => println!("{}", challenge::run(&mut mind)),
        _ => {
            println!("TEST experimental intelligence core");
            println!("  solve <problem>");
            println!("  experiment");
            println!("  benchmark");
        }
    }
}
