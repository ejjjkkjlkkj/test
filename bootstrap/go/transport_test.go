package zero

import "testing"

func TestTransportDoesNotChangeMeaning(t *testing.T) {
	base, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", "TESTED")
	if err != nil { t.Fatal(err) }
	for _, transport := range []Transport{TransportSoftware, TransportMesh, TransportDTN, TransportI2P, TransportRadio, TransportSatellite, TransportD2D} {
		candidate := base.WithTransport(transport).WithDelivery(DeliveryDeferred)
		if !SameMeaning(base, candidate) { t.Fatalf("%s changed semantic meaning", transport) }
	}
}

func TestTransportFailureDoesNotChangeTruthPayload(t *testing.T) {
	base, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", "TESTED")
	if err != nil { t.Fatal(err) }
	failed := base.WithTransport(TransportSatellite).WithDelivery(DeliveryFailed)
	if !SameMeaning(base, failed) || failed.Payload != "2" { t.Fatal("transport failure changed semantic content") }
}

func TestRouteSelectionIsDeterministic(t *testing.T) {
	capabilities := []RouteCapability{
		{Transport: TransportSatellite, Available: true, Reachable: true},
		{Transport: TransportDTN, Available: true, Reachable: true},
		{Transport: TransportMesh, Available: true, Reachable: true},
	}
	for i := 0; i < 100; i++ {
		selected, ok := SelectRoute(capabilities)
		if !ok || selected != TransportMesh { t.Fatalf("unexpected route: %s %v", selected, ok) }
	}
}

func TestRouteSelectionSkipsUnavailableCarriers(t *testing.T) {
	selected, ok := SelectRoute([]RouteCapability{
		{Transport: TransportSatellite, Available: true, Reachable: true},
		{Transport: TransportMesh, Available: false, Reachable: true},
		{Transport: TransportDTN, Available: true, Reachable: false},
	})
	if !ok || selected != TransportSatellite { t.Fatalf("unexpected route: %s %v", selected, ok) }
}

func TestNoRouteDoesNotChangeSemanticTruth(t *testing.T) {
	base, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", "TESTED")
	if err != nil { t.Fatal(err) }
	if _, ok := SelectRoute(nil); ok { t.Fatal("empty capability set unexpectedly selected a route") }
	if base.Payload != "2" { t.Fatal("missing route changed semantic payload") }
}

func TestEnvelopeRejectsMissingIdentityAndTarget(t *testing.T) {
	if _, err := NewEnvelope("", "ZERO", "DEVICE", "2", "TESTED"); err == nil { t.Fatal("empty identity accepted") }
	if _, err := NewEnvelope("X", "ZERO", "", "2", "TESTED"); err == nil { t.Fatal("empty target accepted") }
}
