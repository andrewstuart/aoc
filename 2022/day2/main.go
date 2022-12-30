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
	Rock RPS = iota + 1
	Paper
	Scissors

	Win  = 6
	Draw = 3
	Loss = 0
)

type Strat struct {
	Them, Me RPS
	Result   int
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
		s.Me = Rock
	case "Y":
		s.Me = Paper
	case "Z":
		s.Me = Scissors
	}
	return &s
}

func (s Strat) outcome() int {
	if s.Them == s.Me {
		return Draw
	}
	if s.Me == Rock && s.Them == Scissors || int(s.Me)-2 == int(s.Them) {
		return Win
	}
	return Loss
}

func (s Strat) score() int {
	return int(s.Me) + s.outcome()
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
