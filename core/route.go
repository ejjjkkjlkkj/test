package core

// RouteCapability describes what a carrier can currently do. It is data,
// not a hardware handle. The semantic core can therefore select a route
// without importing an operating system, radio, network or device API.
type RouteCapability struct {
	Transport Transport
	Available bool
	Reachable bool
	StoreAndForward bool
	CostClass string
}

// SelectRoute chooses a deterministic available carrier. It never changes
// the envelope and never treats carrier availability as semantic truth.
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
	if best == TransportNone {
		return TransportNone, false
	}
	return best, true
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
