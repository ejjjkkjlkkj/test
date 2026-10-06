use crate::core::Mind;

#[derive(Debug)]
pub struct Challenge {
    pub name: &'static str,
    pub prompt: &'static str,
    pub required_signals: &'static [&'static str],
}

pub fn suite() -> Vec<Challenge> {
    vec![
        Challenge {
            name: "unknown-process",
            prompt: "A process suddenly stops working after an assumption changes. Find a recovery strategy.",
            required_signals: &["assumption", "test"],
        },
        Challenge {
            name: "conflicting-observations",
            prompt: "Two observations appear incompatible. Find a way to proceed without pretending one is false.",
            required_signals: &["observation", "context"],
        },
        Challenge {
            name: "novel-category",
            prompt: "You receive objects with no labels. Invent a useful organization without assuming the correct categories are known.",
            required_signals: &["organize", "experiment"],
        },
    ]
}

pub fn run(mind: &mut Mind) -> String {
    let mut out = String::new();
    for challenge in suite() {
        let result = mind.solve(challenge.prompt);
        let lower = result.answer.to_lowercase();
        let hits = challenge
            .required_signals
            .iter()
            .filter(|signal| lower.contains(**signal))
            .count();

        out.push_str(&format!(
            "{} | signals={}/{} | revisions={}\n",
            challenge.name,
            hits,
            challenge.required_signals.len(),
            result.revisions
        ));
    }
    out
}
