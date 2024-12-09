package ezaoc

func ResliceDelim[T comparable](t T) func(T, int) (bool, bool) {
	return func(comp T, _ int) (bool, bool) {
		return t == comp, false
	}
}

func ResliceGroupN[T any](n int) func(T, int) (bool, bool) {
	return func(_ T, i int) (bool, bool) {
		return (i+1)%n == 0, true
	}
}

// Reslice takes a slice and a function that returns a boolean and a boolean.
// It returns a slice of slices where the input slice is split at the points
// where the function returns true. If the second boolean is true, the element
// that caused the split is included in the new slice.
func Reslice[T any, Ts ~[]T](inputs Ts, f func(T, int) (bool, bool)) []Ts {
	var out []Ts
	var curr Ts
	for i, in := range inputs {
		if split, keep := f(in, i); split {
			if keep {
				curr = append(curr, in)
			}
			out = append(out, curr)
			curr = *new(Ts)
			continue
		}
		curr = append(curr, in)
	}
	if len(curr) > 0 {
		out = append(out, curr)
	}
	return out
}

func ResliceIncludeLastEmpty[T any, Ts ~[]T](inputs Ts, f func(T, int) (bool, bool)) []Ts {
	var out []Ts
	var curr Ts
	for i, in := range inputs {
		if split, keep := f(in, i); split {
			if keep {
				curr = append(curr, in)
			}
			out = append(out, curr)
			curr = *new(Ts)
			continue
		}
		curr = append(curr, in)
	}
	out = append(out, curr)
	return out
}
