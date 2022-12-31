package main

import (
	"bufio"
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

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
		if st == "" {
			return st, io.EOF
		}
		return st, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	const ct = 14
	// Add challenge logic here probably
	count := 0
	for i := range inputs[0] {
		s := ezaoc.SetFrom([]rune(inputs[0][i : i+ct]))
		if len(s) == ct {
			count = i + ct
			break
		}
	}

	return count
}
