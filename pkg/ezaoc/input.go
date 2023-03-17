package ezaoc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var ErrIgnore = fmt.Errorf("Ignore this output")

func readMultiString(br *bufio.Reader, delim string) (string, error) {
	out := ""
	for i := 0; i < len(delim); i++ {
		st, err := br.ReadString(delim[i])
		out += st
		if err != nil {
			return out, err
		}
		if !strings.HasSuffix(out, delim[:i+1]) {
			idx := -1
			// if we've read XYX and delim is XX, reset to length of the longest suffix that still matches delim
			for j := i; j >= -1; j-- {
				if strings.HasSuffix(out, delim[:j-1]) {
					idx = j - 1
					break
				}
			}
			i = idx
		}
	}
	return out, nil
}

func Read[T any](r io.Reader, delim string, f func(string) (T, error)) ([]T, error) {
	br := bufio.NewReader(r)
	var ts []T

accum:
	for {
		st, err := readMultiString(br, delim)
		if err != nil && err != io.EOF {
			return nil, err
		}
		st = strings.TrimSuffix(st, delim)
		if st == "" && err == io.EOF {
			return ts, nil
		}

		next, err := f(st)
		if err != nil {
			if err != nil {
				switch err {
				case io.EOF: // Callees may return io.EOF to end our use of this reader.
					return ts, nil
				case ErrIgnore:
					continue accum
				default:
					return nil, err
				}
			}
			return ts, err
		}
		ts = append(ts, next)
	}
}

// ReadAOC takes any io.Reader (suggest using bufio.Reader to prevent loss of
// bytes in io.EOF cases) and calls the provided func on every space-trimmed
// input line, returning a slice of that item and any errors encountered. The
// callee should return io.EOF to cease use of the reader, e.g. in the case of
// header or footer. If a callee returns ezaoc.ErrIgnore, the returned item
// will be ignored.
func ReadAOC[T any](r io.Reader, f func(string) (T, error)) ([]T, error) {
	return Read(r, "\n", f)
}

// Raw is the same as above but doesn't trim space, for when that matters in input files.
func RawReadAOC[T any](r io.Reader, f func(string) (T, error)) ([]T, error) {
	br := bufio.NewReader(r)
	var ts []T
accum:
	for {
		st, err := br.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}
		if strings.TrimSpace(st) == "" && err == io.EOF {
			return ts, nil
		}

		next, err := f(st)
		if err != nil {
			if err != nil {
				switch err {
				case io.EOF: // Callees may return io.EOF to end our use of this reader.
					return ts, nil
				case ErrIgnore:
					continue accum
				default:
					return nil, err
				}
			}
			return ts, err
		}
		ts = append(ts, next)
	}
}
