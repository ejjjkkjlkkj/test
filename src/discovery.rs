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
                    out.push(DiscoveredOperator{name,rule:format!("when '{}' is observed, contrast '{}' with '{}'",a.observation,a.action,b.action),uses:0,wins:0});
                }
            }
        }
    }
    out
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
    }
}
