package ezaoc

// FMap is map operation across a slice using a mapping func
func FMap[T, T2 any, U ~[]T, F func(T) T2](u U, f F) []T2 {
	var out []T2
	for _, n := range u {
		out = append(out, f(n))
	}
	return out
}

// Sum returns the sum of a slice of integers
func Sum(is []int) int {
	out := 0
	for _, i := range is {
		out += i
	}
	return out
}

// LastN returns a slice of the last N items in an input slice. Reuses the
// backing array of the input slice, so beware.
func LastN[T any, U ~[]T](t U, ct uint) U {
	if ct > uint(len(t)) {
		panic("Out of range")
	}

	return t[len(t)-int(ct):]
}