use crate::archive::Archive;
use crate::discovery::{discover_programs,Trace};
use crate::evaluator::compare;
use crate::generator::{held_out,judge};

pub fn run() -> String {
    let traces=vec![
        Trace{state:"unknown".into(),observation:"blocked".into(),action:"split".into()},
        Trace{state:"unknown".into(),observation:"conflict".into(),action:"contrast".into()},
        Trace{state:"unknown".into(),observation:"blocked".into(),action:"invert".into()},
        Trace{state:"unknown".into(),observation:"open".into(),action:"probe".into()},
    ];
    let programs=discover_programs(&traces);
    let mut archive=Archive::default();
    let tasks=held_out();
    let mut hits=0usize;
    for p in &programs {
        let mut win=false;
        for task in &tasks {
            if judge(&p.run("initial state",&task.prompt),task) { win=true; }
        }
        archive.admit(p.clone(),win);
        hits+=win as usize;
    }
    let e=compare("blocked conflict","split contrast invert probe",&traces);
    format!("PROGRAMS={} ARCHIVE={} HELDOUT_HITS={} PROGRESS={} NOVELTY={} RECOVERY={} COST={}\n",
        programs.len(),archive.len(),hits,e.progress,e.novelty,e.recovery,e.cost)
}
