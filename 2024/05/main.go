package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

type Page struct {
	ID, ComesBefore int
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
	bs, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	in := string(bs)

	parts := strings.Split(in, "\n\n")

	inputs, err := ezaoc.ReadAOC(strings.NewReader(parts[0]), func(st string) (Page, error) {
		if st == "" {
			return Page{}, io.EOF
		}
		p := Page{}
		_, err := fmt.Sscanf(st, "%d|%d", &p.ID, &p.ComesBefore)
		if err != nil {
			return p, fmt.Errorf("error scanning: %w", err)
		}
		return p, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// as we encouter each page, we add it to a set. Then for the next page, we
	// check its "must come before" list, and if any of those pages are in the
	// set it's out of order

	order := make(map[int][]int)
	for _, p := range inputs {
		order[p.ID] = append(order[p.ID], p.ComesBefore)
	}

	updates, err := ezaoc.ReadAOC(strings.NewReader(parts[1]), ezaoc.IntSlicer(","))
	if err != nil {
		log.Fatal(err)
	}
	count := 0

	good := ezaoc.Set[int]{}
allUpdates:
	for i, up := range updates {
		// fmt.Println(up)
		s := ezaoc.Set[int]{}
		for _, page := range up {
			for _, seen := range order[page] {
				if s.Contains(seen) {
					continue allUpdates
				}
			}
			s.Add(page)
		}

		good.Add(i)
	}

	for i, up := range updates {
		if good.Contains(i) {
			continue
		}

		sort.Slice(up, func(i, j int) bool {
			ord := order[up[i]]
			return slices.Contains(ord, up[j])
		})

		count += up[len(up)/2]

	}

	return count
}
