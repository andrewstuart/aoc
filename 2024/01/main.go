package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
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
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([2]int, error) {
		if st == "" {
			return [2]int{0, 0}, io.EOF
		}
		fs := strings.Fields(st)
		a, _ := strconv.Atoi(fs[0])
		b, _ := strconv.Atoi(fs[1])
		return [2]int{a, b}, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	as, bs := make([]int, len(inputs)), make([]int, len(inputs))
	for i, v := range inputs {
		as[i] = v[0]
		bs[i] = v[1]
	}

	slices.Sort(as)
	slices.Sort(bs)
	// Add challenge logic here probably
	dist := 0
	for i := range as {
		d := as[i] - bs[i]
		if d < 0 {
			d = -d
		}
		dist += d
	}

	return dist
}
