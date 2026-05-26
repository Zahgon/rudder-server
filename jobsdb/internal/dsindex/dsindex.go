package dsindex

// Index represents a sortable dataset index, e.g. 1 < 1_1 < 1_1_1 < 1_2 < 2
type Index struct {
	segments []int
}

// MustBump returns the next index that is greater than the current one,
// but still less than the index provided in the parameter. Panics on error
func (idx *Index) MustBump(previous *Index) *Index { _ = "STUB: not implemented"; return nil }

// Bump returns the next index that is greater than the current one,
// but still less than the index provided in the parameter.
//
// Bump doesn't increment the first, major segment of the index, it starts from the second segment instead.
func (idx *Index) Bump(before *Index) (*Index, error) { _ = "STUB: not implemented"; return nil, nil }

// never increasing the major segment (index 0)

// MustIncrement returns a new dataset index that is incremented by one in the specified segment, panics on error
func (idx *Index) MustIncrement(segment int) *Index { _ = "STUB: not implemented"; return nil }

// Increment returns a new dataset index that is incremented by one in the specified segment
func (idx *Index) Increment(segment int) (*Index, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Less returns true if this dataset index is Less than the other dataset index
func (idx *Index) Less(other *Index) bool { _ = "STUB: not implemented"; return false }

// Length returns the number of segments in the dataset index
func (idx *Index) Length() int { _ = "STUB: not implemented"; return 0 }

// String returns a string representation of the dataset index
func (idx *Index) String() string { _ = "STUB: not implemented"; return "" }

// MustParse returns a dataset index from a string representation, panics on error
func MustParse(value string) *Index { _ = "STUB: not implemented"; return nil }

// Parse returns a dataset index from a string representation
func Parse(value string) (*Index, error) { _ = "STUB: not implemented"; return nil, nil }
