package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
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

type Move struct {
	Dir   string
	Count int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (Move, error) {
		var mv Move
		if st == "" {
			return mv, io.EOF
		}
		_, err := fmt.Sscanf(st, "%s %d", &mv.Dir, &mv.Count)
		return mv, err
	})
	if err != nil {
		log.Fatal(err)
	}

	visited := ezaoc.Set[cell]{}
	var head, tail cell
	visited.Add(tail)

	// Add challenge logic here probably
	for _, in := range inputs {
		switch in.Dir {
		case "L":
			head.I -= in.Count
		case "R":
			head.I += in.Count
		case "U":
			head.J += in.Count
		case "D":
			head.J -= in.Count
		}

		visited.Add(tail.movesTo(head)...)
	}

	return len(visited)
}

type cell struct {
	I, J int
}

func (c *cell) movesTo(c2 cell) []cell {
	var out []cell
	for math.Abs(float64(c.I-c2.I)) > 1 || math.Abs(float64(c.J-c2.J)) > 1 {
		next := *c
		if c2.I > c.I {
			next.I++
		}
		if c2.I < c.I {
			next.I--
		}
		if c2.J > c.J {
			next.J++
		}
		if c2.J < c.J {
			next.J--
		}
		out = append(out, next)
		*c = next
	}
	return out
}
