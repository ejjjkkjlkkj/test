#[derive(Clone,Debug,PartialEq,Eq,Hash)]
pub enum Step { RemoveNoise, ReverseView, SplitParts, JoinParts, ChangeFrame, OrderActions, ContrastCase, ProbeFailure }

#[derive(Clone,Debug,PartialEq,Eq,Hash)]
pub struct Program { pub steps:Vec<Step> }

impl Program {
    pub fn single(step:Step)->Self { Self{steps:vec![step]} }
    pub fn run(&self,state:&str,problem:&str)->String {
        let mut out=state.to_string();
        for step in &self.steps {
            out=match step {
                Step::RemoveNoise=>format!("strip unsupported detail from: {out}"),
                Step::ReverseView=>format!("reverse the current view of: {out}"),
                Step::SplitParts=>format!("separate the parts of: {problem}"),
                Step::JoinParts=>format!("join compatible consequences from: {out}"),
                Step::ChangeFrame=>format!("change the frame used to inspect: {problem}"),
                Step::OrderActions=>format!("put the possible actions into an experiment sequence: {out}"),
                Step::ContrastCase=>format!("construct a contrasting case against: {out}"),
                Step::ProbeFailure=>format!("probe what observation would make this fail: {out}"),
            };
        }
        out
    }
    pub fn label(&self)->String {
        self.steps.iter().map(|s|match s {
            Step::RemoveNoise=>"remove", Step::ReverseView=>"reverse", Step::SplitParts=>"split",
            Step::JoinParts=>"join", Step::ChangeFrame=>"reframe", Step::OrderActions=>"sequence",
            Step::ContrastCase=>"contrast", Step::ProbeFailure=>"probe",
        }).collect::<Vec<_>>().join(" -> ")
    }
}
