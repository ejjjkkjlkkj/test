package zero

import "testing"

func TestCanonicalRoundTrip(t *testing.T) {
	in := "RECORD|1|EVENT|0001|1|UNKNOWN|LOCAL|REMOTE|RULE:ARITHMETIC.ADD\\|INPUT:1,1\\|EXPECTED:2|DEFINED"
	out, err := RoundTrip(in)
	if err != nil {
		t.Fatal(err)
	}
	if out != in {
		t.Fatalf("round-trip changed record: %q != %q", out, in)
	}
}

func TestPayloadPreserved(t *testing.T) {
	in := "RECORD|1|EVENT|0002|2|UNKNOWN|LOCAL|REMOTE|UNKNOWN|DEFINED"
	r, err := Decode(in)
	if err != nil {
		t.Fatal(err)
	}
	if r.Payload != "UNKNOWN" {
		t.Fatalf("payload changed: %q", r.Payload)
	}
}

func TestMissingFieldRejected(t *testing.T) {
	in := "RECORD|1|EVENT|0003|3|UNKNOWN|LOCAL|REMOTE||DEFINED"
	if _, err := Decode(in); err == nil {
		t.Fatal("expected missing payload rejection")
	}
}

func TestContradictionDoesNotDeliver(t *testing.T) {
	r := Record{Version: "1", Type: "EVENT", Identity: "0004", Sequence: 4, Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE", Payload: "x", Proof: "DEFINED"}
	got := Execute(r, "SET_STATE:DELIVERED")
	if got.Status != "CONTRADICTION" || got.Next == "DELIVERED" {
		t.Fatalf("invalid transition accepted: %+v", got)
	}
}

func TestDeferredIsNotDelivered(t *testing.T) {
	r := Record{Version: "1", Type: "EVENT", Identity: "0005", Sequence: 5, Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE", Payload: "x", Proof: "DEFINED"}
	got := Execute(r, "DELIVERY:DEFERRED")
	if got.Status != "DEFERRED" || got.Next == "DELIVERED" {
		t.Fatalf("deferred delivery became delivered: %+v", got)
	}
}

func TestDeterministicExecution(t *testing.T) {
	r := Record{Version: "1", Type: "EVENT", Identity: "0006", Sequence: 6, Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE", Payload: "x", Proof: "DEFINED"}
	want := Execute(r, "OBSERVE")
	for i := 0; i < 100; i++ {
		if got := Execute(r, "OBSERVE"); got != want {
			t.Fatalf("non-deterministic result at run %d: %+v != %+v", i, got, want)
		}
	}
}


func FuzzDecodeNeverPanics(f *testing.F) {
	f.Add("RECORD|1|EVENT|seed|1|UNKNOWN|LOCAL|REMOTE|payload|DEFINED")
	f.Add("")
	f.Add("RECORD")
	f.Fuzz(func(t *testing.T, input string) {
		_, _ = Decode(input)
	})
}
