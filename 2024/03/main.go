package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/rs/zerolog"
)

var lg = zerolog.New(os.Stderr).With().Timestamp().Logger()

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
	bs, _ := io.ReadAll(r)
	input := string(bs)

	re := regexp.MustCompile(`(mul\(\d+\,\d+\)|do()|don't())`)

	muls := re.FindAllString(input, -1)
	lg.Info().Msgf("muls: %v", muls)
	tot := 0
	for _, m := range muls {
		var a, b int
		fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
		tot += a * b
	}
	// Add challenge logic here probably
	return tot
}
