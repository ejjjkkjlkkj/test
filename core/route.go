package core

type RouteCapability struct {
    Transport Transport
    Available bool
    Reachable bool
    StoreAndForward bool
    CostClass string
}

func SelectRoute(capabilities []RouteCapability) (Transport, bool) {
    best := TransportNone
    bestRank := int(^uint(0) >> 1)
    for _, capability := range capabilities {
        if !capability.Available || !capability.Reachable { continue }
        rank := transportRank(capability.Transport)
        if rank < 0 || rank >= bestRank { continue }
        best, bestRank = capability.Transport, rank
    }
    if best == TransportNone { return TransportNone, false }
    return best, true
}

func transportRank(transport Transport) int {
    switch transport {
    case TransportSoftware: return 0
    case TransportMesh: return 1
    case TransportDTN: return 2
    case TransportI2P: return 3
    case TransportRadio: return 4
    case TransportSatellite: return 5
    case TransportD2D: return 6
    default: return -1
    }
}
