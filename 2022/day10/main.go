package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
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

	fmt.Println(aoc(br))
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

func aoc(r io.Reader) string {
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

	var reg, cycle int
	reg = 1
	out := &bytes.Buffer{}
	for _, in := range inputs {
		for i := 0; i < in.cycles(); i++ {
			pix := (cycle + i) % 40
			if math.Abs(float64(reg-pix)) < 2 {
				fmt.Fprint(out, "#")
			} else {
				fmt.Fprint(out, ".")
			}
			if (cycle+i+1)%40 == 0 {
				fmt.Fprintln(out)
			}
		}
		cycle += in.cycles()
		if in.Operation == "addx" {
			reg += in.Operand
		}
	}
	return strings.TrimSpace(out.String())
}
