#[derive(Clone, Debug)]
pub struct Episode {
    pub problem: String,
    pub attempt: String,
    pub outcome: Outcome,
    pub lesson: String,
}

#[derive(Clone, Debug)]
pub enum Outcome {
    Unknown,
    Useful,
    Rejected,
}

#[derive(Default)]
pub struct Memory {
    episodes: Vec<Episode>,
}

impl Memory {
    pub fn remember(&mut self, episode: Episode) {
        self.episodes.push(episode);
        if self.episodes.len() > 256 {
            self.episodes.remove(0);
        }
    }

    pub fn relevant(&self, problem: &str) -> Vec<&Episode> {
        let words: Vec<&str> = problem.split_whitespace().collect();
        self.episodes
            .iter()
            .filter(|e| words.iter().any(|w| e.problem.contains(w)))
            .collect()
    }

    pub fn lessons(&self) -> impl Iterator<Item = &str> {
        self.episodes.iter().map(|e| e.lesson.as_str())
    }
}
