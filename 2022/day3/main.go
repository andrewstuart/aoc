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
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]rune, error) {
		if st == "" {
			return nil, io.EOF
		}
		return []rune(st), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0

	for _, input := range inputs {
		a, b := split(input)
		s := ezaoc.SetFrom(a)
		s2 := ezaoc.Set[rune]{}
		for _, ch := range b {
			if s.Contains(ch) {
				s2.Add(ch)
			}
		}
		count += ezaoc.Sum(lo.Map(s2.Items(), func(r rune, _ int) int { return pri(r) }))
	}

	return count
}

func pri(ch rune) int {
	if 'A' <= ch && ch <= 'Z' {
		return int(ch-'A') + 27
	}
	if 'a' <= ch && ch <= 'z' {
		return int(ch-'a') + 1
	}
	panic("out of range priority")
}

func split[T any, Ts ~[]T](in Ts) (Ts, Ts) {
	l := len(in) / 2
	return in[:l], in[l:]
}
