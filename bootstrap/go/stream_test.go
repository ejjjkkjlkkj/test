package zero

import "testing"

func TestSequenceMustIncrease(t *testing.T) {
	base := Record{Version: "1", Type: "EVENT", Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE", Payload: "x", Proof: "DEFINED"}
	records := []Record{
		{Version: base.Version, Type: base.Type, Identity: "a", Sequence: 1, Time: base.Time, Source: base.Source, Target: base.Target, Payload: base.Payload, Proof: base.Proof},
		{Version: base.Version, Type: base.Type, Identity: "b", Sequence: 2, Time: base.Time, Source: base.Source, Target: base.Target, Payload: base.Payload, Proof: base.Proof},
	}
	if err := ValidateSequence(records); err != nil { t.Fatal(err) }
	records[1].Sequence = 1
	if err := ValidateSequence(records); err == nil { t.Fatal("expected sequence rejection") }
}

func TestDuplicateIdentityRejected(t *testing.T) {
	r := Record{Version: "1", Type: "EVENT", Identity: "same", Sequence: 1, Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE", Payload: "x", Proof: "DEFINED"}
	if err := ValidateSequence([]Record{r, r}); err == nil { t.Fatal("expected duplicate identity rejection") }
}
