package core

import (
	"errors"
	"fmt"
	"strings"
)

// Transport identifies a possible carrier. It is metadata, never part of
// semantic truth and never required for a valid semantic result.
type Transport string

const (
	TransportNone      Transport = "NONE"
	TransportSoftware  Transport = "SOFTWARE"
	TransportMesh      Transport = "MESH"
	TransportDTN       Transport = "DTN"
	TransportI2P       Transport = "I2P"
	TransportRadio     Transport = "RADIO"
	TransportSatellite Transport = "SATELLITE"
	TransportD2D       Transport = "D2D"
)

type DeliveryState string

const (
	DeliveryUnknown   DeliveryState = "UNKNOWN"
	DeliveryPending   DeliveryState = "PENDING"
	DeliveryDelivered DeliveryState = "DELIVERED"
	DeliveryDeferred  DeliveryState = "DEFERRED"
	DeliveryRejected  DeliveryState = "REJECTED"
	DeliveryFailed    DeliveryState = "FAILED"
)

// Envelope separates the meaning of a message from the carrier used to move it.
// A carrier may disappear without changing the envelope's semantic payload.
type Envelope struct {
	Identity  string
	Sequence  int64
	Source    string
	Target    string
	Payload   string
	Proof     Proof
	Transport Transport
	Delivery  DeliveryState
}

func NewEnvelope(identity, source, target, payload string, proof Proof) (Envelope, error) {
	if strings.TrimSpace(identity) == "" {
		return Envelope{}, errors.New("ENVELOPE.IDENTITY.EMPTY")
	}
	if strings.TrimSpace(target) == "" {
		return Envelope{}, errors.New("ENVELOPE.TARGET.EMPTY")
	}
	if proof == "" {
		return Envelope{}, errors.New("ENVELOPE.PROOF.EMPTY")
	}
	return Envelope{
		Identity: identity,
		Source: source,
		Target: target,
		Payload: payload,
		Proof: proof,
		Transport: TransportNone,
		Delivery: DeliveryUnknown,
	}, nil
}

// WithTransport changes only delivery metadata. It cannot change identity,
// payload, proof, or semantic meaning.
func (e Envelope) WithTransport(transport Transport) Envelope {
	e.Transport = transport
	return e
}

// WithDelivery changes only delivery state.
func (e Envelope) WithDelivery(state DeliveryState) Envelope {
	e.Delivery = state
	return e
}

// SameMeaning proves that two envelopes carry identical semantic content even
// if their transport or delivery state differs.
func SameMeaning(a, b Envelope) bool {
	return a.Identity == b.Identity &&
		a.Sequence == b.Sequence &&
		a.Source == b.Source &&
		a.Target == b.Target &&
		a.Payload == b.Payload &&
		a.Proof == b.Proof
}

// Validate rejects malformed delivery metadata but never requires a specific
// physical carrier.
func (e Envelope) Validate() error {
	if strings.TrimSpace(e.Identity) == "" {
		return errors.New("ENVELOPE.IDENTITY.EMPTY")
	}
	if strings.TrimSpace(e.Target) == "" {
		return errors.New("ENVELOPE.TARGET.EMPTY")
	}
	if e.Sequence < 0 {
		return errors.New("ENVELOPE.SEQUENCE.INVALID")
	}
	if e.Proof == "" {
		return errors.New("ENVELOPE.PROOF.EMPTY")
	}
	return nil
}

func (e Envelope) String() string {
	return fmt.Sprintf("%s:%d:%s->%s [%s/%s]", e.Identity, e.Sequence, e.Source, e.Target, e.Transport, e.Delivery)
}
