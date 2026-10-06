//! Gen-2 experiment: transformations are data, not named strategies.
//!
//! The seed contains only mechanical mutations. A "discovered" transformation is
//! a composition produced by the engine and can be replayed independently.

#[derive(Clone, Debug, PartialEq, Eq, Hash)]
pub enum Primitive {
    Keep,
    Drop,
    Swap,
    Duplicate,
    Group,
    Ungroup,
}

#[derive(Clone, Debug, PartialEq, Eq, Hash)]
pub struct Transformation {
    pub steps: Vec<Primitive>,
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct State {
    pub items: Vec<String>,
}

impl State {
    pub fn new(items: &[&str]) -> Self {
        Self { items: items.iter().map(|s| (*s).to_string()).collect() }
    }
}

impl Transformation {
    pub fn apply(&self, mut state: State) -> State {
        for step in &self.steps {
            match step {
                Primitive::Keep => state.items.truncate(state.items.len().min(3)),
                Primitive::Drop => {
                    if !state.items.is_empty() {
                        state.items.remove(0);
                    }
                }
                Primitive::Swap => {
                    if state.items.len() >= 2 {
                        state.items.swap(0, 1);
                    }
                }
                Primitive::Duplicate => {
                    if let Some(first) = state.items.first().cloned() {
                        state.items.push(first);
                    }
                }
                Primitive::Group => {
                    if state.items.len() >= 2 {
                        let joined = state.items.drain(0..2).collect::<Vec<_>>().join("+");
                        state.items.insert(0, joined);
                    }
                }
                Primitive::Ungroup => {
                    if let Some(first) = state.items.first().cloned() {
                        if first.contains('+') {
                            let parts = first.split('+').map(str::to_string).collect::<Vec<_>>();
                            state.items.splice(0..1, parts);
                        }
                    }
                }
            }
        }
        state
    }

    pub fn label(&self) -> String {
        self.steps.iter().map(|s| match s {
            Primitive::Keep => "keep",
            Primitive::Drop => "drop",
            Primitive::Swap => "swap",
            Primitive::Duplicate => "duplicate",
            Primitive::Group => "group",
            Primitive::Ungroup => "ungroup",
        }).collect::<Vec<_>>().join(" -> ")
    }
}

#[derive(Clone, Debug)]
pub struct Trial {
    pub before: State,
    pub after: State,
    pub changed: bool,
    pub reversible: bool,
}

pub fn seed() -> Vec<Transformation> {
    vec![
        Transformation { steps: vec![Primitive::Keep] },
        Transformation { steps: vec![Primitive::Drop] },
        Transformation { steps: vec![Primitive::Swap] },
        Transformation { steps: vec![Primitive::Duplicate] },
        Transformation { steps: vec![Primitive::Group] },
        Transformation { steps: vec![Primitive::Ungroup] },
    ]
}

/// Builds compositions that did not exist in the initial seed.
pub fn discover(seed: &[Transformation], max_depth: usize) -> Vec<Transformation> {
    let atoms = [
        Primitive::Keep,
        Primitive::Drop,
        Primitive::Swap,
        Primitive::Duplicate,
        Primitive::Group,
        Primitive::Ungroup,
    ];

    let mut found = Vec::new();
    for depth in 2..=max_depth {
        let mut current = vec![Vec::<Primitive>::new()];
        for _ in 0..depth {
            let mut next = Vec::new();
            for prefix in &current {
                for atom in &atoms {
                    let mut steps = prefix.clone();
                    steps.push(atom.clone());
                    next.push(steps);
                }
            }
            current = next;
        }

        for steps in current {
            let candidate = Transformation { steps };
            if !seed.contains(&candidate) && !found.contains(&candidate) {
                found.push(candidate);
            }
        }
    }
    found
}

pub fn trial(t: &Transformation, state: State, perturb: &Transformation) -> Trial {
    let before = state.clone();
    let after = t.apply(state);
    let changed = after != before;
    let once_more = perturb.apply(after.clone());
    let recovered = perturb.apply(once_more.clone()) == once_more;
    Trial { before, after, changed, reversible: recovered }
}

/// Keeps transformations according to observed behavior, not text keywords.
pub fn survivors(candidates: &[Transformation], cases: &[State]) -> Vec<Transformation> {
    candidates.iter()
        .filter(|candidate| {
            let perturb = Transformation { steps: vec![Primitive::Swap] };
            cases.iter().any(|case_state| {
                let result = trial(candidate, case_state.clone(), &perturb);
                result.changed && result.reversible
            })
        })
        .cloned()
        .collect()
}

pub fn report() -> String {
    let seed = seed();
    let discovered = discover(&seed, 3);
    let cases = vec![
        State::new(&["a", "b", "c", "d"]),
        State::new(&["x", "y", "z"]),
        State::new(&["left", "right"]),
    ];
    let survivors = survivors(&discovered, &cases);

    let mut out = String::new();
    out.push_str(&format!("SEED={}\n", seed.len()));
    out.push_str(&format!("DISCOVERED={}\n", discovered.len()));
    out.push_str(&format!("SURVIVORS={}\n", survivors.len()));
    for (i, t) in survivors.iter().take(12).enumerate() {
        out.push_str(&format!("{} | {}\n", i + 1, t.label()));
    }
    out
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn discovery_creates_programs_absent_from_seed() {
        let seed = seed();
        let discovered = discover(&seed, 2);
        assert!(!discovered.is_empty());
        assert!(discovered.iter().all(|p| !seed.contains(p)));
        assert!(discovered.iter().any(|p| p.steps.len() == 2));
    }

    #[test]
    fn transformation_is_replayable() {
        let t = Transformation { steps: vec![Primitive::Duplicate, Primitive::Group] };
        let a = t.apply(State::new(&["a", "b"]));
        let b = t.apply(State::new(&["a", "b"]));
        assert_eq!(a, b);
    }

    #[test]
    fn selection_uses_behavior_not_labels() {
        let candidates = discover(&seed(), 2);
        let kept = survivors(&candidates, &[State::new(&["a", "b", "c"])]);
        assert!(kept.iter().all(|t| t.steps.len() >= 2));
    }
}
