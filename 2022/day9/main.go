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
	rope := make([]cell, 10)
	head, tail := rope[0], rope[9]
	visited.Add(tail)

	// Add challenge logic here probably
	for _, in := range inputs {
		for i := 0; i < in.Count; i++ {
			switch in.Dir {
			case "L":
				head.I--
			case "R":
				head.I++
			case "U":
				head.J++
			case "D":
				head.J--
			}

			next := head
			for i := range rope[1:9] {
				rope[i].movesTo(next)
				next = rope[i]
			}
			visited.Add(tail.movesTo(next)...)
		}
	}

	return len(visited)
}

type cell struct {
	I, J int
}

const dist = 1

func (c *cell) movesTo(c2 cell) []cell {
	var out []cell
	for math.Abs(float64(c.I-c2.I)) > dist || math.Abs(float64(c.J-c2.J)) > dist {
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
