package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/davecgh/go-spew/spew"
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

type Monkey struct {
	N         int
	Items     []int
	Operation string
	Test      int
	True      int
	False     int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.Read(r, "\n\n", func(st string) (Monkey, error) {
		var m Monkey
		if st == "" {
			return m, io.EOF
		}
		br := bufio.NewReader(strings.NewReader(st))
		_, err := fmt.Fscanf(br, "Monkey %d:\n", &m.N)
		if err != nil {
			return m, err
		}
		itemsRaw, err := br.ReadString('\n')
		if err != nil {
			return m, err
		}
		m.Items, err = ezaoc.IntSlicer(", ")(strings.TrimPrefix(strings.TrimSpace(itemsRaw), "Starting items: "))
		if err != nil {
			return m, err
		}
		raw, err := br.ReadString('\n')
		if err != nil {
			return m, err
		}
		m.Operation = strings.Split(strings.TrimSpace(raw), ": ")[1]
		_, err = fmt.Fscanf(br, "  Test: divisible by %d\n", &m.Test)
		if err != nil {
			return m, err
		}
		_, err = fmt.Fscanf(br, "    If true: throw to monkey %d\n", &m.True)
		if err != nil {
			return m, err
		}
		_, err = fmt.Fscanf(br, "    If false: throw to monkey %d\n", &m.False)
		if err != nil {
			return m, err
		}
		return m, err
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0
	spew.Dump(inputs)
	count = len(inputs)

	return count
}
