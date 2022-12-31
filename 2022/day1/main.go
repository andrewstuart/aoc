package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
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
	inputs, err := ezaoc.ReadAOC(r, func(st string) (int, error) {
		if st == "" && last == "" {
			return -1, io.EOF
		}
		last = st

		n, err := strconv.Atoi(st)
		if err != nil {
			return -1, nil
		}

		return n, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// 	fmt.Printf("inputs = %+v\n", inputs)

	elves := ezaoc.Reslice(inputs, ezaoc.ResliceDelim(-1))

	sort.Slice(elves, func(i, j int) bool {
		return ezaoc.Sum(elves[i]) < ezaoc.Sum(elves[j])
	})
	fmt.Printf("elves = %+v\n", ezaoc.LastN(elves, 3))

	// mm := ezaoc.MaxOf(max, func(i int) int { return i })
	// // Add challenge logic here probably
	// count := 0
	// spew.Dump(next)
	// count = len(next)
	return ezaoc.Sum(ezaoc.FMap(ezaoc.LastN(elves, 3), ezaoc.Sum))
}
