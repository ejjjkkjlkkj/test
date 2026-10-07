package zero

import "testing"

func TestCanonicalEscaping(t *testing.T) {
	r := Record{
		Version: "1", Type: "EVENT", Identity: "escape", Sequence: 7,
		Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE",
		Payload: "a|b\\c\nd\te", Proof: "DEFINED",
	}
	wire, err := Encode(r)
	if err != nil { t.Fatal(err) }
	got, err := Decode(wire)
	if err != nil { t.Fatal(err) }
	if got != r { t.Fatalf("escaping changed record: %#v != %#v", got, r) }
}

func TestCanonicalInteger(t *testing.T) {
	r := Record{
		Version: "1", Type: "EVENT", Identity: "integer", Sequence: 1,
		Time: "UNKNOWN", Source: "LOCAL", Target: "REMOTE",
		Payload: "x", Proof: "DEFINED",
	}
	wire, err := Encode(r)
	if err != nil { t.Fatal(err) }
	if wire == "" { t.Fatal("empty encoding") }
	got, err := Decode(wire)
	if err != nil || got.Sequence != 1 { t.Fatalf("bad canonical sequence: %v %#v", err, got) }
}
