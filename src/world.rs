#[derive(Clone, Debug)]
pub struct Observation {
    pub statement: String,
    pub tension: Option<String>,
}

pub fn observe(problem: &str, candidate: &str) -> Observation {
    let lower = candidate.to_lowercase();
    let tension = if lower.contains("always") || lower.contains("never") {
        Some("absolute claim detected; seek a counterexample".into())
    } else if candidate.len() < problem.len() / 3 {
        Some("candidate may be under-specified".into())
    } else {
        None
    };

    Observation {
        statement: candidate.to_string(),
        tension,
    }
}
