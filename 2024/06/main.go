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

	// Add challenge logic here probably
	count := 0
	ezaoc.Print2dGrid(inputs)

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

	log.Printf("Position: %v, Direction: %v\n", position, direction)

	visited := ezaoc.Set[ezaoc.Cell[string]]{}

	for {
		next := ezaoc.GetCellsInDirection(inputs, direction, position.I, position.J, 2)
		if len(next) < 2 {
			break
		}

		inputs[position.I][position.J] = "."

		// 		fmt.Println(next[1])

		if next[1].Value == "#" {
			direction = direction.Turn(ezaoc.TurnRight)
		} else {
			position = next[1]
			if !visited.Contains(position) {
				count++
				visited.Add(position)
			}
		}
		inputs[position.I][position.J] = getDirString(direction)
		// ezaoc.Print2dGrid(inputs)
		// fmt.Println()
		// time.Sleep(time.Second / 2)
	}

	return count
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
