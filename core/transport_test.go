package core

import "testing"

func TestTransportDoesNotChangeMeaning(t *testing.T) {
    base,err:=NewEnvelope("ARITHMETIC:1+1","ZERO","DEVICE","2",ProofTested)
    if err!=nil { t.Fatal(err) }
    for _,transport:=range []Transport{TransportSoftware,TransportMesh,TransportDTN,TransportI2P,TransportRadio,TransportSatellite,TransportD2D} {
        candidate:=base.WithTransport(transport).WithDelivery(DeliveryDeferred)
        if !SameMeaning(base,candidate) { t.Fatalf("%s changed semantic meaning",transport) }
        if candidate.Payload!="2" || candidate.Proof!=ProofTested { t.Fatalf("%s changed semantic payload or proof",transport) }
    }
}

func TestDeliveryFailureDoesNotInvalidateTruth(t *testing.T) {
    envelope,err:=NewEnvelope("ARITHMETIC:1+1","ZERO","DEVICE","2",ProofTested)
    if err!=nil { t.Fatal(err) }
    failed:=envelope.WithTransport(TransportSatellite).WithDelivery(DeliveryFailed)
    if !SameMeaning(envelope,failed) || !VerifyCanonicalTruth(2) { t.Fatal("delivery failure modified semantic truth") }
}

func TestEnvelopeRejectsMissingSemanticIdentity(t *testing.T) {
    if _,err:=NewEnvelope("","A","B","2",ProofTested); err==nil { t.Fatal("empty identity accepted") }
    if _,err:=NewEnvelope("X","A","","2",ProofTested); err==nil { t.Fatal("empty target accepted") }
}
