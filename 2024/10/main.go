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
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0
	ezaoc.Print2dGrid(inputs)

	var trailheads []ezaoc.Cell[int]
	ezaoc.VisitCells(inputs, func(c ezaoc.Cell[int]) error {
		if c.Value == 0 {
			// fmt.Println("starting at ", c)
			visited := ezaoc.Set[ezaoc.Cell[int]]{}
			count += Search(inputs, c, visited)
			trailheads = append(trailheads, c)
		}
		return nil
	})

	return count
}

func Search(inputs [][]int, c ezaoc.Cell[int], visited ezaoc.Set[ezaoc.Cell[int]]) int {
	if visited.Contains(c) {
		return 0
	}
	// fmt.Println(c)
	// fmt.Println(path)
	visited.Add(c)
	defer visited.Remove(c)
	if c.Value == 9 {
		return 1
	}

	found := 0

	ns := ezaoc.NonDiagSliceNeighbors(inputs, c.I, c.J)
	// fmt.Println(ns)
	// time.Sleep(1 * time.Second)
	for _, n := range ns {
		if n.Value == c.Value+1 {
			found += Search(inputs, n, visited)
		}
	}
	return found
}
