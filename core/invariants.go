package core

// TruthInvariant is a read-only description of the normative mathematical
// invariant. The actual value is held in an unexported constant so callers
// cannot mutate the source of truth at runtime.
type TruthInvariant struct {
	Expression string
	Result     int64
	Immutable  bool
}

const canonicalOnePlusOne int64 = 2

// CanonicalTruth returns a value copy. No exported mutable global state is
// used as the source of truth.
func CanonicalTruth() TruthInvariant {
	return TruthInvariant{
		Expression: "1+1",
		Result:     canonicalOnePlusOne,
		Immutable:  true,
	}
}

// OnePlusOne returns the canonical mathematical invariant.
// Hardware, OS, runtime identity, clock, network, GPU and external input
// cannot participate in its definition.
func OnePlusOne() int64 {
	return canonicalOnePlusOne
}

// VerifyCanonicalTruth rejects any claim that contradicts the invariant.
func VerifyCanonicalTruth(claimed int64) bool {
	return claimed == canonicalOnePlusOne
}

// TruthCannotBePromotedOrOverridden verifies that execution context cannot
// replace the normative invariant.
func TruthCannotBePromotedOrOverridden() bool {
	return canonicalOnePlusOne == 2
}
