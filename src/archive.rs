use crate::operator_lang::Program;

#[derive(Clone,Debug)]
pub struct Entry { pub program:Program, pub trials:usize, pub wins:usize, pub failures:usize }

#[derive(Default,Debug)]
pub struct Archive { entries:Vec<Entry> }

impl Archive {
    pub fn admit(&mut self, program:Program, win:bool) {
        if let Some(e)=self.entries.iter_mut().find(|e|e.program==program) {
            e.trials+=1; if win { e.wins+=1 } else { e.failures+=1 }
        } else {
            self.entries.push(Entry{program,trials:1,wins:win as usize,failures:(!win) as usize});
        }
        self.entries.retain(|e| e.trials<4 || e.wins>0);
        self.entries.sort_by_key(|e| e.failures*20 + e.program.steps.len());
        if self.entries.len()>64 { self.entries.truncate(64); }
    }
    pub fn usable(&self)->Vec<Program> { self.entries.iter().filter(|e|e.wins>0).map(|e|e.program.clone()).collect() }
    pub fn len(&self)->usize { self.entries.len() }
}
