use crate::discovery::Trace;

#[derive(Clone,Debug,Default)]
pub struct Evaluation { pub progress:i64, pub novelty:i64, pub recovery:i64, pub cost:i64 }

pub fn compare(before:&str,after:&str,traces:&[Trace])->Evaluation {
    let b=before.split_whitespace().count() as i64;
    let a=after.split_whitespace().count() as i64;
    let novelty=after.split_whitespace().filter(|w|!before.split_whitespace().any(|x|x==*w)).count() as i64;
    let recovery=traces.windows(2).filter(|w|w[0].observation!=w[1].observation && w[0].action!=w[1].action).count() as i64;
    Evaluation{progress:a-b,novelty,recovery,cost:a.max(1)}
}
