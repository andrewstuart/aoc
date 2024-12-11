package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

type Card struct {
	Num     int
	Winning []int
	Mine    []int
}

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
	i := 1
	inputs, err := ezaoc.ReadAOC(r, func(st string) (Card, error) {
		if st == "" {
			return Card{}, io.EOF
		}
		fs := strings.Fields(st)
		mine := false
		card := Card{Num: i}
		i++
		for _, f := range fs[2:] {
			if f == "|" {
				mine = true
				continue
			}
			n := ezaoc.MustAtoi(f)
			if !mine {
				card.Winning = append(card.Winning, n)
			} else {
				card.Mine = append(card.Mine, n)
			}
		}
		return card, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for _, card := range inputs {
		val := 0
		for _, n := range card.Winning {
			if slices.Contains(card.Mine, n) {
				if val == 0 {
					val = 1
					continue
				}
				val *= 2
			}
		}
		count += val
	}

	return count
}
