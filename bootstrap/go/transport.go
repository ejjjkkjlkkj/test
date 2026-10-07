package zero

import "errors"

// Transport identifies a possible carrier. It is metadata and is never the
// source of semantic truth.
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

type Envelope struct {
	Identity  string
	Sequence  uint64
	Source    string
	Target    string
	Payload   string
	Proof     string
	Transport Transport
	Delivery  DeliveryState
}

func NewEnvelope(identity, source, target, payload, proof string) (Envelope, error) {
	if identity == "" {
		return Envelope{}, errors.New("ENVELOPE.IDENTITY.EMPTY")
	}
	if target == "" {
		return Envelope{}, errors.New("ENVELOPE.TARGET.EMPTY")
	}
	if proof == "" {
		return Envelope{}, errors.New("ENVELOPE.PROOF.EMPTY")
	}
	return Envelope{Identity: identity, Source: source, Target: target, Payload: payload, Proof: proof, Transport: TransportNone, Delivery: DeliveryUnknown}, nil
}

func (e Envelope) WithTransport(transport Transport) Envelope {
	e.Transport = transport
	return e
}

func (e Envelope) WithDelivery(delivery DeliveryState) Envelope {
	e.Delivery = delivery
	return e
}

// SameMeaning ignores only transport and delivery metadata.
func SameMeaning(a, b Envelope) bool {
	return a.Identity == b.Identity && a.Sequence == b.Sequence && a.Source == b.Source &&
		a.Target == b.Target && a.Payload == b.Payload && a.Proof == b.Proof
}

func (e Envelope) Validate() error {
	if e.Identity == "" {
		return errors.New("ENVELOPE.IDENTITY.EMPTY")
	}
	if e.Target == "" {
		return errors.New("ENVELOPE.TARGET.EMPTY")
	}
	if e.Proof == "" {
		return errors.New("ENVELOPE.PROOF.EMPTY")
	}
	return nil
}

type RouteCapability struct {
	Transport       Transport
	Available       bool
	Reachable       bool
	StoreAndForward bool
	CostClass       string
}

func SelectRoute(capabilities []RouteCapability) (Transport, bool) {
	best := TransportNone
	bestRank := int(^uint(0) >> 1)
	for _, capability := range capabilities {
		if !capability.Available || !capability.Reachable {
			continue
		}
		rank := transportRank(capability.Transport)
		if rank < 0 || rank >= bestRank {
			continue
		}
		best = capability.Transport
		bestRank = rank
	}
	return best, best != TransportNone
}

func transportRank(transport Transport) int {
	switch transport {
	case TransportSoftware:
		return 0
	case TransportMesh:
		return 1
	case TransportDTN:
		return 2
	case TransportI2P:
		return 3
	case TransportRadio:
		return 4
	case TransportSatellite:
		return 5
	case TransportD2D:
		return 6
	default:
		return -1
	}
}
