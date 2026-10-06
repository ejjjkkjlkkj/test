#[derive(Clone, Debug, PartialEq, Eq, Hash)]
pub enum Operator { Remove, Invert, Split, Merge, Reframe, Sequence, Contrast, Probe }

impl Operator {
    pub fn all() -> &'static [Operator] {
        &[Operator::Remove,Operator::Invert,Operator::Split,Operator::Merge,Operator::Reframe,Operator::Sequence,Operator::Contrast,Operator::Probe]
    }
    pub fn apply(&self, state:&str, problem:&str)->String {
        match self {
            Operator::Remove=>format!("Remove unsupported parts from: {state}"),
            Operator::Invert=>format!("Invert the current relation and inspect consequences for: {problem}"),
            Operator::Split=>format!("Separate {problem} into independent transitions"),
            Operator::Merge=>format!("Combine compatible consequences from: {state}"),
            Operator::Reframe=>format!("Represent {problem} from a different viewpoint"),
            Operator::Sequence=>format!("Turn {state} into an ordered experiment"),
            Operator::Contrast=>format!("Construct the strongest contrasting case against: {state}"),
            Operator::Probe=>format!("Design an observation that could make this idea fail: {state}"),
        }
    }
}
