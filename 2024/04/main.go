package main

import (
	"bufio"
	"fmt"
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
		if checkX(inputs, "MAS", c.I, c.J) {
			// fmt.Println(c.I, c.J, c.Value)
			count++
		}
		return nil
	})

	// ezaoc.Print2dGrid(inputs)
	// ezaoc.Print2dGridWithNumbers(inputs)

	return count
}

func checkX(inputs [][]string, s string, i, j int) bool {
	return checkDirEitherWay(inputs, s, i, j, ezaoc.UpLeft) && checkDirEitherWay(inputs, s, i, j, ezaoc.DownLeft)
}

func checkDirEitherWay(inputs [][]string, s string, i, j int, d ezaoc.Direction) bool {
	return checkDir(inputs, s, i, j, d) || checkDir(inputs, s, i, j, d.Opposite())
}

func checkDir(inputs [][]string, s string, i, j int, d ezaoc.Direction) bool {
	c := ezaoc.GetCellsInDirection(inputs, d, i, j, 2)
	if len(c) < 2 {
		return false
	}

	cs := ezaoc.GetCellsInDirection(inputs, d.Opposite(), c[1].I, c[1].J, len(s))

	chk := ""
	for _, v := range cs {
		chk += v.Value
	}
	return chk == s
}
