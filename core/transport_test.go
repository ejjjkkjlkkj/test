package core

import "testing"

func TestTransportDoesNotChangeMeaning(t *testing.T) {
	base, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", ProofTested)
	if err != nil {
		t.Fatal(err)
	}

	transports := []Transport{
		TransportSoftware, TransportMesh, TransportDTN, TransportI2P,
		TransportRadio, TransportSatellite, TransportD2D,
	}
	for _, transport := range transports {
		candidate := base.WithTransport(transport).WithDelivery(DeliveryDeferred)
		if !SameMeaning(base, candidate) {
			t.Fatalf("%s changed semantic meaning", transport)
		}
		if candidate.Payload != "2" || candidate.Proof != ProofTested {
			t.Fatalf("%s changed semantic payload or proof", transport)
		}
	}
}

func TestDeliveryFailureDoesNotInvalidateTruth(t *testing.T) {
	envelope, err := NewEnvelope("ARITHMETIC:1+1", "ZERO", "DEVICE", "2", ProofTested)
	if err != nil {
		t.Fatal(err)
	}
	failed := envelope.WithTransport(TransportSatellite).WithDelivery(DeliveryFailed)
	if !SameMeaning(envelope, failed) {
		t.Fatal("delivery failure modified semantic meaning")
	}
	if !VerifyCanonicalTruth(2) {
		t.Fatal("transport failure modified canonical truth")
	}
}

func TestEnvelopeRejectsMissingSemanticIdentity(t *testing.T) {
	if _, err := NewEnvelope("", "A", "B", "2", ProofTested); err == nil {
		t.Fatal("empty identity accepted")
	}
	if _, err := NewEnvelope("X", "A", "", "2", ProofTested); err == nil {
		t.Fatal("empty target accepted")
	}
}
