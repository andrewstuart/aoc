package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

	fmt.Printf("inputs = %+v\n", inputs)

	elves := ezaoc.Reslice(inputs, -1)

	fmt.Printf("elves = %+v\n", elves)
	fmt.Printf("len(elves) = %+v\n", len(elves))

	i, max := ezaoc.MaxOf(elves, func(cals []int) int {
		ct := 0
		for _, i := range cals {
			ct += i
		}
		return ct
	})
	fmt.Printf("i = %+v\n", i)
	fmt.Printf("elves[i] = %+v\n", elves[i])

	// mm := ezaoc.MaxOf(max, func(i int) int { return i })
	// // Add challenge logic here probably
	// count := 0
	// spew.Dump(next)
	// count = len(next)

	return max
}
