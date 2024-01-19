package main

import (
	"bufio"
	"io"
	"log"
	"os"
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

type Game struct {
	ID    int
	Draws []Draw
}

type Draw struct {
	Green, Red, Blue int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (*Game, error) {
		if st == "" {
			return nil, io.EOF
		}

		//yuuck
		id, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(st, ":")[0], "Game "))
		st = strings.Split(st, ":")[1]

		var draws []Draw
		for _, draw := range strings.Split(st, ";") {
			var d Draw
			draw = strings.TrimSpace(draw)
			dice := strings.Split(draw, ",")
			for _, die := range dice {
				flds := strings.Fields(die)
				n, _ := strconv.Atoi(flds[0])
				color := flds[1]
				switch color {
				case "green":
					d.Green = n
				case "red":
					d.Red = n
				case "blue":
					d.Blue = n
				}
			}
			draws = append(draws, d)
		}

		return &Game{
			Draws: draws,
			ID:    id,
		}, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var count int
	for _, game := range inputs {
		var r, b, g int
		for _, d := range game.Draws {
			if d.Red > r {
				r = d.Red
			}
			if d.Blue > b {
				b = d.Blue
			}
			if d.Green > g {
				g = d.Green
			}
		}
		count += r * b * g
	}

	return count
}
