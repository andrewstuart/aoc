package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br, false))
}

var ErrFoundLoop = errors.New("found loop")

func aoc(r io.Reader, print bool) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]string, error) {
		if st == "" {
			return nil, io.EOF
		}
		return strings.Split(st, ""), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// if print {
	// ezaoc.Print2dGridWithNumbers(inputs)
	// }

	var position ezaoc.Cell[string]
	var direction ezaoc.Direction
	ezaoc.VisitCells(inputs, func(c ezaoc.Cell[string]) error {
		switch c.Value {
		case "^", "v", "<", ">":
			position = c
			direction = getDir(c.Value)
		}
		return nil
	})

	start := position
	startDir := direction

	log.Printf("Position: %v, Direction: %v\n", position, direction)

	type step struct {
		C ezaoc.Cell[string]
		D ezaoc.Direction
	}

	simulate := func(inputs [][]string, position ezaoc.Cell[string], direction ezaoc.Direction, print bool) ([]step, error) {
		path := []step{}

		for {
			path = append(path, step{C: position, D: direction})
			next := ezaoc.GetCellsInDirection(inputs, direction, position.I, position.J, 2)
			if len(next) < 2 {
				break
			}

			if next[1].Value == "#" || next[1].Value == "O" {
				direction = direction.Turn(ezaoc.TurnRight)
				// If we've been here and turned before
				if slices.Contains(path, step{C: position, D: direction}) {
					return path, ErrFoundLoop
				}
			} else {
				inputs[position.I][position.J] = "."
				position = next[1]
			}
			inputs[position.I][position.J] = getDirString(direction)
			if print {
				ezaoc.Print2dGrid(inputs)
				fmt.Println()
				time.Sleep(time.Second / 20)
			}
		}
		return path, nil
	}

	t := ezaoc.Copy2dSlice(inputs)
	visited, _ := simulate(t, position, direction, false)

	// part1
	// 	return len(lo.UniqBy(visited, func(s step) ezaoc.Cell[string] { return s.C }))

	// for part 2 we want to check all the visited cells.
	// if we can add an obstacle in front that creates a loop then increment count and try again

	reported := ezaoc.Set[ezaoc.Cell[string]]{}
	for _, c := range visited {
		if c.C == start || reported.Contains(c.C) {
			continue
		}
		inCopy := ezaoc.Copy2dSlice(inputs)
		next := ezaoc.GetCellsInDirection(inCopy, c.D, c.C.I, c.C.J, 2)
		if len(next) < 2 {
			continue
		}
		n := next[1]
		n.Set(inCopy, "O")
		v, err := simulate(inCopy, start, startDir, print)
		if err == ErrFoundLoop {
			reported.Add(n)
			fmt.Println("Loop found by adding obstacle at", n)
			for _, c := range v {
				fmt.Printf("%d,%d ", c.C.I, c.C.J)
			}
			fmt.Println()
		}
	}

	return len(reported)
}

func getDirString(d ezaoc.Direction) string {
	switch d {
	case ezaoc.Up:
		return "^"
	case ezaoc.Down:
		return "v"
	case ezaoc.Left:
		return "<"
	case ezaoc.Right:
		return ">"
	default:
		return "?"
	}
}

func getDir(s string) ezaoc.Direction {
	switch s {
	case "^":
		return ezaoc.Up
	case "v":
		return ezaoc.Down
	case "<":
		return ezaoc.Left
	case ">":
		return ezaoc.Right
	default:
		return ezaoc.Unknown
	}
}
