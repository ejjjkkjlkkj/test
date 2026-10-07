package core

// TruthInvariant is deliberately independent of hardware, operating system,
// runtime identity, clock, network, GPU and external services.
//
// Mathematical semantics are not configurable runtime state.
type TruthInvariant struct {
	Expression string
	Result     int64
	Immutable  bool
}

var CanonicalTruth = TruthInvariant{
	Expression: "1+1",
	Result:     2,
	Immutable:  true,
}

// OnePlusOne is the canonical semantic invariant. No hardware capability,
// environment value or external input participates in its definition.
func OnePlusOne() int64 {
	return CanonicalTruth.Result
}

// VerifyCanonicalTruth rejects any claim that contradicts the invariant.
func VerifyCanonicalTruth(claimed int64) bool {
	return claimed == CanonicalTruth.Result
}

// TruthCannotBePromotedOrOverridden prevents treating execution context as a
// source of mathematical truth.
func TruthCannotBePromotedOrOverridden() bool {
	return CanonicalTruth.Immutable && OnePlusOne() == 2
}
