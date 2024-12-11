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

	wins := map[int]int{}

	for i, card := range inputs {
		matches := 0
		for _, n := range card.Winning {
			if slices.Contains(card.Mine, n) {
				matches++
			}
		}
		for j := 0; j < matches; j++ {
			wins[i+j+1] += wins[i] + 1
		}
	}

	count := 0
	for _, w := range wins {
		count += w
	}
	return count + len(inputs)
}
