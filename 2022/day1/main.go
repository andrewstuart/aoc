package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/davecgh/go-spew/spew"
)

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
	last := ""
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
		if st == "" && last == "" {
			return "", io.EOF
		}
		last = st

		return st, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	next := ezaoc.Slice(inputs, "")

	// Add challenge logic here probably
	count := 0
	spew.Dump(next)
	count = len(next)

	return count
}
