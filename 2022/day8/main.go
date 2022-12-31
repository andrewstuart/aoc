package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/samber/lo"
)

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]int, error) {
		if st == "" {
			return nil, io.EOF
		}
		return ezaoc.IntSlicer("")(st)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	// count := 2*len(inputs) + 2*len(inputs[0]) - 4 // 4 corners would be counted twice
	count := 0
	ezaoc.VisitCells(inputs, func(c ezaoc.Cell[int]) error {
		scans := append(split(inputs[c.I], c.J), split(ezaoc.RawCols(inputs, c.J), c.I)...)

		tot := 1
		for _, scan := range scans {
			_, idx, _ := lo.FindIndexOf(scan, func(ht int) bool {
				return ht >= c.Value
			})
			if idx < 0 {
				idx = len(scan) - 1
			}
			tot *= (idx + 1)
		}
		if tot > count {
			count = tot
		}
		return nil
	})
	return count
}

func split[T any](ts []T, n int) [][]T {
	return [][]T{reverse(ts[:n]), ts[n+1:]}
}

func reverse[T any](input []T) []T {
	var output []T

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func splitExcept[T any](n int) func(_ T, idx int) (bool, bool) {
	return func(_ T, idx int) (bool, bool) {
		return idx == n, false
	}
}
