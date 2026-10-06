use crate::operator_lang::{Program,Step};

#[derive(Clone,Debug,PartialEq,Eq)]
pub struct Trace { pub state:String, pub observation:String, pub action:String }

#[derive(Clone,Debug)]
pub struct DiscoveredOperator { pub name:String, pub rule:String, pub uses:usize, pub wins:usize }

pub fn discover(traces:&[Trace])->Vec<DiscoveredOperator> {
    let mut out=Vec::new();
    for a in traces {
        for b in traces {
            if a.observation != b.observation && a.action != b.action {
                let name=format!("transition_{}_{}", token(&a.action), token(&b.action));
                if !out.iter().any(|x:&DiscoveredOperator|x.name==name) {
                    out.push(DiscoveredOperator{name,rule:format!("when '{}' changes, compare '{}' with '{}'",a.observation,a.action,b.action),uses:0,wins:0});
                }
            }
        }
    }
    out
}

pub fn discover_programs(traces:&[Trace])->Vec<Program> {
    let mut programs=Vec::new();
    let seeds=[
        Step::RemoveNoise,Step::ReverseView,Step::SplitParts,Step::JoinParts,
        Step::ChangeFrame,Step::OrderActions,Step::ContrastCase,Step::ProbeFailure
    ];
    for a in traces {
        for b in traces {
            if a.observation != b.observation {
                let first=match token(&a.action).as_str() {
                    "split"=>Step::SplitParts, "contrast"=>Step::ContrastCase,
                    "probe"=>Step::ProbeFailure, "invert"=>Step::ReverseView,
                    _=>Step::ChangeFrame
                };
                let second=match token(&b.action).as_str() {
                    "split"=>Step::SplitParts, "contrast"=>Step::ContrastCase,
                    "probe"=>Step::ProbeFailure, "invert"=>Step::ReverseView,
                    _=>seeds[(a.state.len()+b.action.len())%seeds.len()].clone()
                };
                let p=Program{steps:vec![first,second]};
                if !programs.contains(&p) { programs.push(p); }
            }
        }
    }
    programs
}

fn token(s:&str)->String { s.split_whitespace().next().unwrap_or("empty").chars().filter(|c|c.is_alphanumeric()).take(12).collect() }

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn discovers_relations_from_trajectories() {
        let t=vec![
            Trace{state:"a".into(),observation:"blocked".into(),action:"split".into()},
            Trace{state:"b".into(),observation:"open".into(),action:"invert".into()},
        ];
        assert!(!discover(&t).is_empty());
        assert!(!discover_programs(&t).is_empty());
    }
}
