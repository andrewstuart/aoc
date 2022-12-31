package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
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

type Range struct {
	From, To int
}

func (r Range) Contains(r2 Range) bool {
	return r2.From >= r.From && r2.To <= r.To
}

func (r Range) Overlaps(r2 Range) bool {
	return r.Has(r2.From) || r.Has(r2.To) || r2.Has(r.From) || r2.Has(r.To)
}

func (r Range) Has(p int) bool {
	return r.From <= p && p <= r.To
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([2]Range, error) {
		var out [2]Range
		if st == "" {
			return out, io.EOF
		}
		_, err := fmt.Sscanf(st, "%d-%d,%d-%d", &out[0].From, &out[0].To, &out[1].From, &out[1].To)
		return out, err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0
	for _, in := range inputs {
		if in[0].Overlaps(in[1]) {
			count++
		}
	}
	return count
}
