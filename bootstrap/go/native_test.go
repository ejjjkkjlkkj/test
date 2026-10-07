package zero

import (
	"bytes"
	"testing"
)

func TestNativeRoundTrip(t *testing.T) {
	r := Record{
		Version:"1", Type:"FUTURE_OBJECT", Identity:"SATELLITE:test",
		Sequence:7, Time:"2026-10-07T00:00:00.000Z",
		Source:"ZERO", Target:"WORLD",
		Payload:"meaning=machine spatiale de test|state=SIMULATED",
		Proof:"SIMULATED",
	}
	got, err := DecodeNative(mustNativeEncode(r))
	if err != nil { t.Fatal(err) }
	if got != r { t.Fatalf("native round-trip changed record: %#v != %#v", got, r) }
}

func TestNativeIsDeterministic(t *testing.T) {
	r := Record{Version:"1", Type:"OBJECT", Identity:"x", Sequence:1, Time:"UNKNOWN", Source:"ZERO", Target:"WORLD", Payload:"UNKNOWN", Proof:"DEFINED"}
	a := mustNativeEncode(r)
	b := mustNativeEncode(r)
	if !bytes.Equal(a,b) { t.Fatal("native encoding is not deterministic") }
}

func TestNativeRejectsTruncation(t *testing.T) {
	r := Record{Version:"1", Type:"OBJECT", Identity:"x", Sequence:1, Time:"UNKNOWN", Source:"ZERO", Target:"WORLD", Payload:"x", Proof:"DEFINED"}
	data := mustNativeEncode(r)
	for i := 0; i < len(data); i++ {
		if _, err := DecodeNative(data[:i]); err == nil { t.Fatalf("accepted truncated native record at %d", i) }
	}
}

func mustNativeEncode(r Record) []byte {
	b, err := EncodeNative(r)
	if err != nil { panic(err) }
	return b
}


func FuzzDecodeNativeNeverPanics(f *testing.F) {
	seed := mustNativeEncode(Record{
		Version:"1", Type:"OBJECT", Identity:"seed", Sequence:1,
		Time:"UNKNOWN", Source:"ZERO", Target:"WORLD",
		Payload:"payload", Proof:"DEFINED",
	})
	f.Add(seed)
	f.Add([]byte("ZER0"))
	f.Add([]byte{})
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = DecodeNative(data)
	})
}
