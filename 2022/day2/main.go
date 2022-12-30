package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/samber/lo"
)

type RPS int

const (
	Rock RPS = iota
	Paper
	Scissors

	Win  = 6
	Draw = 3
	Loss = 0
)

type Strat struct {
	Them   RPS
	Result int
}

func parse(st string) *Strat {
	var s Strat
	fs := strings.Fields(st)
	switch fs[0] {
	case "A":
		s.Them = Rock
	case "B":
		s.Them = Paper
	case "C":
		s.Them = Scissors
	}
	switch fs[1] {
	case "X":
		s.Result = Loss
	case "Y":
		s.Result = Draw
	case "Z":
		s.Result = Win
	}
	return &s
}

func (s Strat) outcome() int {
	return s.Result
}

func (s Strat) me() RPS {
	switch s.outcome() {
	case Win:
		return RPS((s.Them + 1) % 3)
	case Loss:
		o := s.Them - 1
		if o < 0 {
			return Scissors
		}
		return RPS(o)
	}
	return s.Them
}

func (s Strat) score() int {
	return int(s.me()+1) + s.outcome()
}

func main() {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (*Strat, error) {
		if st == "" {
			return nil, io.EOF
		}
		return parse(st), nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return ezaoc.Sum(lo.Map(inputs, func(s *Strat, _ int) int {
		return s.score()
	}))
}
