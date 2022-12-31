package main

import (
	"bufio"
	"fmt"
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

type instruction struct {
	Operation string
	Operand   int
}

func (i instruction) cycles() int {
	switch i.Operation {
	case "noop":
		return 1
	case "addx":
		return 2
	}
	panic("unknown type")
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (instruction, error) {
		var x instruction
		if st == "" {
			return x, io.EOF
		}
		fmt.Sscanf(st, "%s %d", &x.Operation, &x.Operand)
		return x, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	var reg, out, cycle int
	reg = 1
	for _, in := range inputs {
		for i := 1; i <= in.cycles(); i++ {
			if (cycle+i+20)%40 == 0 {
				out += reg * (cycle + i)
			}
		}
		cycle += in.cycles()
		if in.Operation == "addx" {
			reg += in.Operand
		}
	}
	return out
}
