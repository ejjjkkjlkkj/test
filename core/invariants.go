package core

type TruthInvariant struct {
    Expression string
    Result int64
    Immutable bool
}

const canonicalOnePlusOne int64 = 2

func CanonicalTruth() TruthInvariant {
    return TruthInvariant{Expression:"1+1", Result:canonicalOnePlusOne, Immutable:true}
}

func OnePlusOne() int64 { return canonicalOnePlusOne }
func VerifyCanonicalTruth(claimed int64) bool { return claimed == canonicalOnePlusOne }
func TruthCannotBePromotedOrOverridden() bool { return canonicalOnePlusOne == 2 }
