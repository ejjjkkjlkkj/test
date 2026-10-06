use crate::discovery::{discover,Trace};
use crate::evaluator::compare;

pub fn run() -> String {
    let traces=vec![
        Trace{state:"unknown".into(),observation:"blocked".into(),action:"split".into()},
        Trace{state:"unknown".into(),observation:"conflict".into(),action:"contrast".into()},
        Trace{state:"unknown".into(),observation:"blocked".into(),action:"invert".into()},
        Trace{state:"unknown".into(),observation:"open".into(),action:"probe".into()},
    ];
    let ops=discover(&traces);
    let e=compare("blocked conflict","split contrast invert probe",&traces);
    format!("DISCOVERED_OPERATORS={} PROGRESS={} NOVELTY={} RECOVERY={} COST={}\n",ops.len(),e.progress,e.novelty,e.recovery,e.cost)
}
