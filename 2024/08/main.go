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

	ans := ezaoc.Set[ezaoc.Cell[string]]{}
	byFreq := map[string][]ezaoc.Cell[string]{}

	ezaoc.VisitCells(inputs, func(c ezaoc.Cell[string]) error {
		if c.Value == "." {
			return nil
		}
		byFreq[c.Value] = append(byFreq[c.Value], c)
		return nil
	})

	for _, nodes := range byFreq {
		for i, node := range nodes {
			for _, node2 := range nodes[i+1:] {
				iDiff, jDiff := node.I-node2.I, node.J-node2.J
				c1 := ezaoc.Cell[string]{I: node.I + iDiff, J: node.J + jDiff}
				c2 := ezaoc.Cell[string]{I: node2.I - iDiff, J: node2.J - jDiff}
				if ezaoc.IsInBounds(inputs, c1.I, c1.J) {
					c1.Set(inputs, "#")
					ans.Add(c1)
				}
				if ezaoc.IsInBounds(inputs, c2.I, c2.J) {
					ans.Add(c2)
					c2.Set(inputs, "#")
				}
			}
		}
	}

	// Add challenge logic here probably

	return len(ans)
}
