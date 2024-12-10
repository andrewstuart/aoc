package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

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

type Op struct {
	Result   int
	Operands []int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (Op, error) {
		if st == "" {
			return Op{}, io.EOF
		}
		const f = "%d: "
		var op Op
		_, err := fmt.Sscanf(st, f, &op.Result)
		if err != nil {
			return op, err
		}
		op.Operands, err = ezaoc.IntSlicer(" ")(st[len(fmt.Sprintf(f, op.Result)):])
		if err != nil {
			return op, fmt.Errorf("error slicing ints: %w", err)
		}
		return op, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0

ins:
	for _, op := range inputs {
		ops := gen(len(op.Operands)-1, []string{"+", "*", "||"})
		fmt.Println(op, len(ops))
		for _, ops := range ops {
			if op.eval(ops...) == op.Result {
				count += op.Result
				continue ins
			}
		}
	}

	return count
}

func (o Op) eval(ops ...string) int {
	if len(ops) != len(o.Operands)-1 {
		panic("invalid number of operators")
	}
	accum := o.Operands[0]
	for i, op := range ops {
		switch op {
		case "+":
			accum += o.Operands[i+1]
		case "*":
			accum *= o.Operands[i+1]
		case "||":
			accum, _ = strconv.Atoi(strconv.Itoa(accum) + strconv.Itoa(o.Operands[i+1]))
		}
	}
	return accum
}

func gen(n int, ops []string) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	var out [][]string
	for _, op := range ops {
		for _, sub := range gen(n-1, ops) {
			out = append(out, append([]string{op}, sub...))
		}
	}
	return out
}
