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
        Self {
            memory: Memory::default(),
            generation: 0,
        }
    }

    pub fn solve(&mut self, problem: &str) -> ResultReport {
        self.generation += 1;
        let mut path = Vec::new();
        let mut candidate = self.first_construction(problem);
        let mut revisions = 0;

        path.push("construction initiale".into());

        for round in 0..8 {
            let observation = world::observe(problem, &candidate);

            if let Some(tension) = observation.tension {
                path.push(format!("tension: {tension}"));
                candidate = self.revise(problem, &candidate, &tension, round);
                revisions += 1;
                continue;
            }

            if let Some(previous) = self.memory.relevant(problem).first() {
                if previous.attempt == candidate {
                    candidate = self.escape(problem, &candidate);
                    path.push("échappement d'une solution déjà rencontrée".into());
                    revisions += 1;
                    continue;
                }
            }

            path.push("confrontation réussie".into());
            self.memory.remember(Episode {
                problem: problem.to_string(),
                attempt: candidate.clone(),
                outcome: Outcome::Useful,
                lesson: format!("la construction '{}' résiste à la confrontation", candidate),
            });

            return ResultReport {
                answer: candidate,
                path,
                revisions,
            };
        }

        self.memory.remember(Episode {
            problem: problem.to_string(),
            attempt: candidate.clone(),
            outcome: Outcome::Unknown,
            lesson: "la recherche a atteint sa limite sans preuve suffisante".into(),
        });

        ResultReport {
            answer: candidate,
            path,
            revisions,
        }
    }

    fn first_construction(&self, problem: &str) -> String {
        let lessons: Vec<&str> = self.memory.lessons().take(3).collect();
        if lessons.is_empty() {
            format!(
                "Ne pas supposer une procédure connue. Décrire le problème, isoler ce qui est certain, puis produire une première transformation testable: {problem}"
            )
        } else {
            format!(
                "Repartir des observations précédentes sans les considérer comme vraies: {problem}. Leçons disponibles: {}",
                lessons.join(" | ")
            )
        }
    }

    fn revise(&self, problem: &str, candidate: &str, tension: &str, round: usize) -> String {
        match round % 4 {
            0 => format!(
                "Retirer l'affirmation fragile ({tension}) et reconstruire autour de l'observable: {problem}"
            ),
            1 => format!(
                "Chercher le cas qui invaliderait l'approche précédente, puis conserver seulement ce qui survit: {candidate}"
            ),
            2 => format!(
                "Changer de point de vue: traiter '{problem}' comme un système à explorer plutôt qu'une question à répondre"
            ),
            _ => format!(
                "Composer une nouvelle stratégie à partir de la contradiction détectée: {tension}; contexte: {problem}"
            ),
        }
    }

    fn escape(&self, problem: &str, candidate: &str) -> String {
        format!(
            "Transformation alternative de '{problem}': inverser l'hypothèse centrale de '{candidate}', observer ce que cela rend possible, puis tester cette nouvelle voie."
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn mind_produces_a_report() {
        let mut mind = Mind::new();
        let result = mind.solve("organize unknown information");
        assert!(!result.answer.is_empty());
        assert!(!result.path.is_empty());
    }

    #[test]
    fn memory_changes_future_search() {
        let mut mind = Mind::new();
        let first = mind.solve("repair an unknown process");
        let second = mind.solve("repair an unknown process");
        assert!(second.revisions >= first.revisions);
    }
}
