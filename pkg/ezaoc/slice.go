package ezaoc

// Reslice takes a slice of comparable types and turns it into a 2d slice
func Reslice[T comparable, Ts ~[]T](ts Ts, delim T) []Ts {
	var out []Ts
	var curr Ts
	for _, t := range ts {
		if t == delim {
			out = append(out, curr)
			curr = curr[:0]
			continue
		}
		curr = append(curr, t)
	}
	return out
}
