#[derive(Clone,Debug)]
pub struct Task { pub prompt:String, pub hidden:String }

pub fn held_out()->Vec<Task> {
    vec![
        Task{prompt:"A plan works until one condition changes.".into(),hidden:"changed condition".into()},
        Task{prompt:"Two observations disagree about the same event.".into(),hidden:"disagreement".into()},
        Task{prompt:"A useful category stops predicting what happens.".into(),hidden:"category failure".into()},
        Task{prompt:"The first approach produces the same dead end repeatedly.".into(),hidden:"dead end".into()},
        Task{prompt:"A new object does not fit any existing description.".into(),hidden:"novel object".into()},
    ]
}

pub fn judge(output:&str,task:&Task)->bool {
    let x=output.to_lowercase();
    x.contains("contrast") || x.contains("probe") || x.contains("reverse") || x.contains("change")
        || (task.hidden=="novel object" && x.contains("separate"))
}
