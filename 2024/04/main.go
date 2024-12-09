package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

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

const word = "XMAS"

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]string, error) {
		if st == "" {
			return nil, io.EOF
		}
		return strings.Split(st, ""), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	ezaoc.VisitCells(inputs, func(c ezaoc.Cell[string]) error {
		// fmt.Println(c)
		for _, dir := range ezaoc.AllDirections {
			cs := ezaoc.GetCellsInDirection(inputs, dir, c.I, c.J, len(word))
			content := ""
			for _, c := range cs {
				content += c.Value
			}
			if content == word {
				count++
			}
			// fmt.Printf("x: %d, y: %d, dir: %s, content: %s\n", c.I, c.J, dir, content)
		}
		return nil
	})

	// 	ezaoc.Print2dGridWithNumbers(inputs)

	return count
}
