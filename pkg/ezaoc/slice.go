package ezaoc

// Reslice takes a slice of comparable types and turns it into a 2d slice
func Reslice[T comparable, Ts ~[]T](inputs Ts, delim T) []Ts {
	var out []Ts
	var curr Ts
	for _, in := range inputs {
		if in == delim {
			out = append(out, curr)
			curr = *new(Ts)
			continue
		}
		curr = append(curr, in)
	}
	return out
}
