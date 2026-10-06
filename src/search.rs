use crate::operator::Operator;

#[derive(Clone,Debug)]
pub struct Candidate { pub text:String, pub history:Vec<Operator>, pub score:i64 }

pub fn evolve(problem:&str,seed:&str,generations:usize)->Candidate {
    let mut population=vec![Candidate{text:seed.into(),history:vec![],score:0}];
    let mut seen=std::collections::HashSet::new();
    for _ in 0..generations {
        let mut next=Vec::new();
        for c in &population {
            for op in Operator::all() {
                let text=op.apply(&c.text,problem);
                if seen.insert(text.clone()) {
                    let score=novelty(&text,problem) as i64+contradiction(&text) as i64-repetition(&c.history,op) as i64*3;
                    let mut history=c.history.clone(); history.push(op.clone());
                    next.push(Candidate{text,history,score});
                }
            }
        }
        next.sort_by_key(|c|-c.score); next.truncate(16);
        if next.is_empty(){break} population=next;
    }
    population.into_iter().max_by_key(|c|c.score).unwrap_or(Candidate{text:seed.into(),history:vec![],score:0})
}
fn novelty(t:&str,p:&str)->usize { let x:std::collections::HashSet<_>=p.split_whitespace().collect(); t.split_whitespace().filter(|w|!x.contains(w)).count() }
fn contradiction(t:&str)->usize { ["fail","contrast","different","invert","remove","probe"].iter().filter(|w|t.to_lowercase().contains(**w)).count() }
fn repetition(h:&[Operator],o:&Operator)->usize { h.iter().filter(|x|*x==o).count() }

#[cfg(test)]
mod tests { use super::*; #[test] fn search_is_deterministic(){let a=evolve("unknown","seed",3);let b=evolve("unknown","seed",3);assert_eq!(a.text,b.text);} }
