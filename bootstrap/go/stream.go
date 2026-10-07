package zero

import "fmt"

func ValidateSequence(records []Record) error {
	if len(records) == 0 {
		return nil
	}
	seen := make(map[string]struct{}, len(records))
	var previous uint64
	for i, r := range records {
		if err := Validate(r); err != nil {
			return fmt.Errorf("record %d: %w", i, err)
		}
		if i > 0 && r.Sequence <= previous {
			return fmt.Errorf("record %d: non-increasing sequence", i)
		}
		if _, ok := seen[r.Identity]; ok {
			return fmt.Errorf("record %d: duplicate identity", i)
		}
		seen[r.Identity] = struct{}{}
		previous = r.Sequence
	}
	return nil
}
