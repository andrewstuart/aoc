package main

import (
	"bufio"
	"bytes"
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

type stacks [][]string

func (st stacks) String() string {
	_, maxLen := ezaoc.MaxOf(st, func(sts []string) int {
		return len(sts)
	})
	buf := &bytes.Buffer{}
	for i := maxLen - 1; i >= 0; i-- {
		for j := 0; j < len(st); j++ {
			if len(st[j]) > i {
				fmt.Fprintf(buf, "[%s] ", st[j][i])
				continue
			}
			buf.WriteString("    ")
		}
		buf.WriteString("\n")
	}
	for n := range st {
		fmt.Fprintf(buf, " %d  ", n+1)
	}
	buf.WriteString("\n\n")
	return buf.String()
}

func (s stacks) move(n, from, to int) {
	tmp := make([]string, n)
	copy(tmp, ezaoc.LastN(s[from], n))
	s[from] = s[from][:len(s[from])-n]
	for i := len(tmp) - 1; i >= 0; i-- {
		s[to] = append(s[to], tmp[i])
	}
}

func readStacks(r io.Reader) [][]string {
	stacks, err := ezaoc.RawReadAOC(r, func(st string) ([]string, error) {
		if strings.TrimSpace(st) == "" {
			return nil, io.EOF
		}
		rs := []rune(st)
		stax := ezaoc.Reslice(rs, ezaoc.ResliceGroupN[rune](4))
		out := make([]string, len(stax))
		for i, st := range stax {
			if st[0] == '[' {
				out[i] = string([]rune{st[1]})
			}
		}
		if strings.TrimSpace(strings.Join(out, "")) == "" {
			return out, ezaoc.ErrIgnore
		}

		return out, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	outStacks := make([][]string, len(stacks[0]))
	for i := len(stacks) - 1; i >= 0; i-- {
		for j := 0; j < len(stacks[i]); j++ {
			if stacks[i][j] != "" {
				outStacks[j] = append(outStacks[j], stacks[i][j])
			}
		}
	}
	return outStacks
}

func aoc(r io.Reader) string {
	br := bufio.NewReader(r)
	st := stacks(readStacks(br))

	moves, err := ezaoc.ReadAOC(br, func(st string) ([3]int, error) {
		var out [3]int
		if strings.TrimSpace(st) == "" {
			return out, io.EOF
		}
		_, err := fmt.Sscanf(st, "move %d from %d to %d", &out[0], &out[1], &out[2])
		return out, err
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range moves {
		fmt.Println(st)
		fmt.Println(m)
		st.move(m[0], m[1]-1, m[2]-1)
		fmt.Println(st)
	}

	// Add challenge logic here probably

	sts := ezaoc.FMap(st, ezaoc.Last[[]string])
	return strings.Join(sts, "")
}
