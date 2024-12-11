package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
)

const ticks = 75

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br, ticks))
}

func aoc(r io.Reader, ticks int) int {
	input, _ := ezaoc.ReadAOC(r, ezaoc.IntSlicer(" "))

	stones := map[int]int{}

	for _, in := range input[0] {
		stones[in]++
	}
	// Add challenge logic here probably

	for i := 0; i < ticks; i++ {
		// spew.Dump(stones)
		nextStones := map[int]int{}
		for val, ct := range stones {
			switch {
			case val == 0:
				nextStones[1] += ct
			case len(fmt.Sprint(val))%2 == 0:
				// fmt.Println("val", val)
				digits := fmt.Sprint(val)
				left, right := digits[:len(digits)/2], digits[len(digits)/2:]
				// fmt.Println("lr", left, right)
				lNum, rNum := ezaoc.MustAtoi(left), ezaoc.MustAtoi(right)
				// fmt.Println("lrn", lNum, rNum)
				nextStones[lNum] += ct
				nextStones[rNum] += ct
			default:
				nextStones[val*2024] += ct
			}
		}
		stones = nextStones
	}

	count := 0
	for _, ct := range stones {
		count += ct
	}

	return count
}
