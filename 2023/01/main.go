package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

var digitNames = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine",
}

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

var digits *regexp.Regexp

func init() {
	digRE := "^(\\d|"
	for _, d := range digitNames {
		digRE += d + "|"
	}
	digRE = digRE[:len(digRE)-1] + ")"
	digits = regexp.MustCompile(digRE)
}

func toNum(s string) int {
	if len(s) == 1 {
		n, _ := strconv.Atoi(s)
		return n
	}

	return slices.Index(digitNames, s)
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (int, error) {
		var first, last int
		for i := 0; i < len(st); i++ {
			sub := st[i:]
			st := digits.FindString(sub)
			if st == "" {
				continue
			}
			last = toNum(st)
			if first == 0 {
				first = last
			}
		}
		fmt.Println(first, last)
		return first*10 + last, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	tot := 0
	for _, i := range inputs {
		tot += i
	}

	return tot
}
