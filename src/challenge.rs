use crate::core::Mind;

#[derive(Debug)]
pub struct Challenge {
    pub name: &'static str,
    pub prompt: &'static str,
    pub signals: &'static [&'static str],
}

pub fn suite() -> Vec<Challenge> {
    vec![
        Challenge { name: "hidden-assumption", prompt: "A process fails when one hidden assumption changes. Recover without knowing the original procedure.", signals: &["assumption", "evidence"] },
        Challenge { name: "contradiction", prompt: "Two observations conflict. Find a way to progress without deleting either observation.", signals: &["contradiction", "information"] },
        Challenge { name: "unknown-objects", prompt: "Unlabeled objects must be organized. Invent categories only if experiments justify them.", signals: &["experiment", "category"] },
        Challenge { name: "dead-end", prompt: "Your first strategy repeatedly fails. Create a genuinely different route instead of polishing it.", signals: &["different", "strategy"] },
    ]
}

pub fn run(mind: &mut Mind) -> String {
    let mut out = String::new();
    for c in suite() {
        let r = mind.solve(c.prompt);
        let s = r.answer.to_lowercase();
        let hits = c.signals.iter().filter(|x| s.contains(**x)).count();
        out.push_str(&format!("{} | signals={}/{} | revisions={} | steps={}\n", c.name, hits, c.signals.len(), r.revisions, r.path.len()));
    }
    out
}
