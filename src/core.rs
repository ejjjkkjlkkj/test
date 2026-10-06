use crate::memory::{Episode, Memory, Outcome};
use crate::world;

pub struct Mind {
    memory: Memory,
    generation: u64,
}

pub struct ResultReport {
    pub answer: String,
    pub path: Vec<String>,
    pub revisions: usize,
}

impl ResultReport {
    pub fn render(&self) -> String {
        let mut out = format!("ANSWER\n{}\n\n", self.answer);
        out.push_str(&format!("REVISIONS: {}\n", self.revisions));
        out.push_str("PATH:\n");
        for (i, step) in self.path.iter().enumerate() {
            out.push_str(&format!("{}. {}\n", i + 1, step));
        }
        out
    }
}

impl Mind {
    pub fn new() -> Self {
        Self { memory: Memory::default(), generation: 0 }
    }

    pub fn solve(&mut self, problem: &str) -> ResultReport {
        self.generation += 1;
        let mut path = Vec::new();
        let mut candidate = self.construct(problem);
        let mut revisions = 0;

        for round in 0..16 {
            path.push(format!("cycle {round}: confrontation"));
            let observation = world::observe(problem, &candidate);

            if let Some(tension) = observation.tension {
                path.push(format!("tension: {tension}"));
                candidate = self.transform(problem, &candidate, &tension, round);
                revisions += 1;
                continue;
            }

            if self.repeated(&candidate, problem) {
                path.push("collision with previous experience".into());
                candidate = self.escape(problem, &candidate, round);
                revisions += 1;
                continue;
            }

            let novelty = self.novelty(&candidate, problem);
            path.push(format!("novelty signal: {novelty}"));

            self.memory.remember(Episode {
                problem: problem.to_string(),
                attempt: candidate.clone(),
                outcome: Outcome::Useful,
                lesson: format!("generation {}: candidate survived; novelty={novelty}", self.generation),
            });

            return ResultReport { answer: candidate, path, revisions };
        }

        self.memory.remember(Episode {
            problem: problem.to_string(),
            attempt: candidate.clone(),
            outcome: Outcome::Unknown,
            lesson: "bounded search exhausted without sufficient evidence".into(),
        });

        ResultReport { answer: candidate, path, revisions }
    }

    fn construct(&self, problem: &str) -> String {
        format!(
            "Build an answer from observable consequences rather than inherited procedure. Problem: {problem}"
        )
    }

    fn transform(&self, problem: &str, candidate: &str, tension: &str, round: usize) -> String {
        match round % 8 {
            0 => format!("Remove the fragile assumption ({tension}); rebuild from what can be observed in: {problem}"),
            1 => format!("Try to destroy this candidate with a counter-case: {candidate}"),
            2 => format!("Invert the goal of '{problem}' and inspect what becomes possible"),
            3 => format!("Split '{problem}' into independent effects; solve the effects before naming a method"),
            4 => format!("Treat the contradiction '{tension}' as information, not failure; redesign the candidate"),
            5 => format!("Change representation of '{problem}': describe transitions instead of objects"),
            6 => format!("Ask what evidence would make the current idea useless, then rebuild around that evidence"),
            _ => format!("Discard the current strategy and synthesize a different experiment for: {problem}"),
        }
    }

    fn repeated(&self, candidate: &str, problem: &str) -> bool {
        self.memory.relevant(problem).iter().any(|e| e.attempt == candidate)
    }

    fn escape(&self, problem: &str, candidate: &str, round: usize) -> String {
        format!(
            "Novel escape {round}: preserve only consequences that survived from '{candidate}', then create a different transformation for '{problem}'."
        )
    }

    fn novelty(&self, candidate: &str, problem: &str) -> usize {
        let p: std::collections::HashSet<_> = problem.split_whitespace().collect();
        candidate.split_whitespace().filter(|w| !p.contains(w)).count()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn produces_a_report() {
        let mut mind = Mind::new();
        let result = mind.solve("organize unknown information");
        assert!(!result.answer.is_empty());
        assert!(!result.path.is_empty());
    }

    #[test]
    fn repeated_experience_forces_escape() {
        let mut mind = Mind::new();
        let a = mind.solve("repair an unknown process");
        let b = mind.solve("repair an unknown process");
        assert!(b.revisions >= a.revisions);
    }
}
