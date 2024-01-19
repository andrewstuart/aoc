package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"

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

var num = regexp.MustCompile(`^\d+`)

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) (string, error) {
		return st, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i, line := range inputs {
		for j := 0; j < len(line); j++ {
			if inputs[i][j] != '*' {
				continue
			}

		}
	}
	return count
}

func check(s []string, row, i, l int) bool {
	prev, next := row-1, row+1
	if prev >= 0 {
		if checkRow(s, prev, i, l) {
			return true
		}
	}
	if next < len(s) {
		if checkRow(s, next, i, l) {
			return true
		}
	}
	if i > 0 && checkChar(s[row][i-1]) {
		return true
	}
	if i+l < len(s[row]) && checkChar(s[row][i+l]) {
		return true
	}
	return false
}

func checkRow(s []string, row, J, l int) bool {
	for i := J - 1; i < J+l+1; i++ {
		if i < 0 || i >= len(s[row]) {
			continue
		}
		chr := s[row][i]
		// fmt.Print(row, i, string([]byte{chr}), ",")
		if checkChar(chr) {
			return true
		}
	}
	// fmt.Println()
	return false
}

func checkChar(chr byte) bool {
	isNum := chr >= '0' && chr <= '9'
	return !isNum && chr != '.'
}
