package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/samber/lo"
)

// efficient solution
// for every map, we can calculate the final position of the seed by generating new maps based on the combined operations of the previous maps.

type Range struct {
	Start, Len int
}

func (r Range) FT() (int, int) {
	return r.Start, r.Start + r.Len - 1
}

func (r Range) Has(i int) bool {
	return i >= r.Start && i < r.Start+r.Len
}

// 5-6 1-10 true
// 1-10 5-6 true
// 5-6 6-7 true
// 5-6 7-8 false
func (r Range) Overlaps(r2 Range) bool {
	min1, max1 := r.FT()
	min2, max2 := r2.FT()
	// fmt.Printf("%d-%d, %d-%d\n", min1, max1, min2, max2)
	return r2.Has(min1) || r2.Has(max1) || r.Has(min2) || r.Has(max2)
}

type Puzzle struct {
	Seeds []Range

	Maps [][]Map
}

type Map struct {
	Range
	Dest int
}

func (m Map) DestR() Range {
	return Range{m.Dest, m.Len}
}

func (m Map) MapRange(r Range) []Range {
	if !m.Overlaps(r) {
		return []Range{r}
	}
	offset := m.Dest - m.Start + 1
	mapMin, mapMax := m.FT()
	rMin, rMax := r.FT()
	// Contained
	if mapMin <= rMin && mapMax >= rMax {
		return []Range{{Start: r.Start + offset, Len: r.Len}}
	}
	var out []Range
	if mapMin > rMin && mapMax < rMax {
		out = append(out, Range{Start: rMin + offset, Len: m.Len})
	}
	if mapMax < rMax {
		out = append(out, Range{Start: mapMax + 1, Len: rMax - mapMax})
	}
	if mapMin > rMin {
		out = append(out, Range{Start: rMin, Len: mapMin - rMin})
	}
	slices.SortFunc(out, func(r1, r2 Range) int {
		return r1.Start - r2.Start
	})
	return out
}

// func (m Map) Split(nextSrc Map) []Map {
// 	dst := m.DestR()
// 	if !dst.Overlaps(nextSrc.Range) {
// 		return []Map{m, nextSrc}
// 	}
// 	return []Map{}
// }

// Gets a slice of slices of maps and expands it out to one slice
// eg
// <dest> <start> <len>
// map:
// 50 90 10 90-99 -> 50->59

// map:
// 5 50 5 50-54 -> 5->9
// 10 90 5 90-94 -> 10->14
//
// should become
// 90-94 -> 10-14
// 50-54 -> 5-9
// func FlattenMaps(maps [][]Map) []Map {
// 	var out []Map
// 	for i, m := range maps {
// 	}
// }

func GetMap(in int, maps []Map) int {
	for _, m := range maps {
		if in >= m.Start && in < m.Start+m.Len {
			return m.Dest + in - m.Start
		}
	}
	return in
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
	var p Puzzle
	var currentMaplist []Map

	ezaoc.ReadAOC(r, func(st string) (string, error) {
		if strings.HasPrefix(st, "seeds:") {
			ranges := lo.Map(strings.Fields(st)[1:], ezaoc.MapNoI(ezaoc.MustAtoi))
			for i := 0; i < len(ranges); i += 2 {
				p.Seeds = append(p.Seeds, Range{Start: ranges[i], Len: ranges[i+1]})
			}
			return "", nil
		}
		if st == "" {
			if len(currentMaplist) > 0 {
				p.Maps = append(p.Maps, currentMaplist)
			}
			return "", nil
		}
		if strings.HasSuffix(st, "map:") {
			currentMaplist = []Map{}
			return "", nil
		}

		var m Map
		vals := lo.Map(strings.Fields(st), ezaoc.MapNoI(ezaoc.MustAtoi))
		m.Dest = vals[0]
		m.Start = vals[1]
		m.Len = vals[2]
		currentMaplist = append(currentMaplist, m)

		return "", nil
	})
	p.Maps = append(p.Maps, currentMaplist)

	fmt.Println(p.Maps)

	// Add challenge logic here probably
	minim := -1

	seeds := make([]Range, len(p.Seeds))
	copy(seeds, p.Seeds)
	for _, ms := range p.Maps {
		// spew.Dump(seeds)
		var next []Range
		fmt.Println("maps", ms)
		for _, rng := range seeds {
			for _, m := range ms {
				next = append(next, m.MapRange(rng)...)
			}
		}
		seeds = next
	}

	for _, s := range seeds {
		if minim == -1 || s.Start < minim {
			minim = s.Start
		}
	}

	return minim
}
