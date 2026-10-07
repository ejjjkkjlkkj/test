package core

import "testing"

func TestSelectRouteSkipsUnavailableCapabilities(t *testing.T) {
	selected, ok := SelectRoute([]RouteCapability{
		{Transport: TransportSatellite, Available: true, Reachable: true},
		{Transport: TransportSoftware, Available: false, Reachable: true},
		{Transport: TransportDTN, Available: true, Reachable: false},
	})
	if !ok || selected != TransportSatellite {
		t.Fatalf("unexpected route: %s %v", selected, ok)
	}
}

func TestSelectRouteIsDeterministic(t *testing.T) {
	caps := []RouteCapability{
		{Transport: TransportI2P, Available: true, Reachable: true},
		{Transport: TransportMesh, Available: true, Reachable: true},
		{Transport: TransportDTN, Available: true, Reachable: true},
	}
	first, ok := SelectRoute(caps)
	if !ok {
		t.Fatal("no route selected")
	}
	for i := 0; i < 100; i++ {
		got, ok := SelectRoute(caps)
		if !ok || got != first {
			t.Fatalf("selection changed: %s %v != %s %v", got, ok, first)
		}
	}
	if first != TransportMesh {
		t.Fatalf("expected deterministic preferred route, got %s", first)
	}
}

func TestNoRouteIsNotSemanticFailure(t *testing.T) {
	envelope, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", ProofTested)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := SelectRoute(nil); ok {
		t.Fatal("empty capabilities unexpectedly produced a route")
	}
	if !VerifyCanonicalTruth(2) {
		t.Fatal("route absence changed truth")
	}
	if envelope.Payload != "2" {
		t.Fatal("route absence changed payload")
	}
}
