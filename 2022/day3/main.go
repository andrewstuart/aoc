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

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]rune, error) {
		if st == "" {
			return nil, io.EOF
		}
		return []rune(st), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Add challenge logic here probably
	count := 0

	groups := ezaoc.Reslice(inputs, ezaoc.ResliceGroupN[[]rune](3))
	fmt.Printf("groups = %+v\n", groups)

	for _, group := range groups {
		s := ezaoc.SetFrom(group[0])
		for _, g2 := range group[0:] {
			s2 := ezaoc.Set[rune]{}
			for _, ch := range g2 {
				if s.Contains(ch) {
					s2.Add(ch)
				}
			}
			s = s2
		}
		count += ezaoc.Sum(ezaoc.FMap(s.Items(), pri))
		fmt.Printf("string(s.Items()) = %+v\n", string(s.Items()))
	}

	return count
}

func pri(ch rune) int {
	if 'A' <= ch && ch <= 'Z' {
		return int(ch-'A') + 27
	}
	if 'a' <= ch && ch <= 'z' {
		return int(ch-'a') + 1
	}
	panic("out of range priority")
}

func split[T any, Ts ~[]T](in Ts) (Ts, Ts) {
	l := len(in) / 2
	return in[:l], in[l:]
}
