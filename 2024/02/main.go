package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/rs/zerolog"
)

var lg = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.DebugLevel)

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)

	log.Println(aoc(br))
}

const limit = 3

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, func(st string) ([]int, error) {
		if st == "" {
			return nil, io.EOF
		}
		ints, err := ezaoc.IntSlicer(" ")(st)
		return ints, err
	})
	if err != nil {
		log.Fatal(err)
	}
	safe := 0

	for _, in := range inputs {
		if check(in) {
			safe++
		}
	}
	return safe
}

func check(in []int) bool {
	last := in[0]
	dir := in[1] - last // used to test that direction doesn't change
	bad := 0
	for j, each := range in[1:] {
		if j == 1 && bad == 1 {
			if check(in[1:]) {
				return true
			}
		}
		if bad > 1 {
			break
		}

		diff := each - last
		if diff*dir <= 0 { // sign is different or zero
			lg.Debug().Int("last", last).Int("each", each).Int("dir", dir).Int("diff", diff).Int("bad", bad).Msg("unsafe dir change")
			bad++
			continue
		}
		if diff > limit || diff < -limit { // abs of diff is too large
			lg.Debug().Int("last", last).Int("each", each).Int("dir", dir).Int("diff", diff).Int("bad", bad).Msg("unsafe dist")
			bad++
			continue
		}
		last = each
	}
	if bad > 1 {
		lg.Warn().Any("level", in).Int("bad", bad).Msg("unsafe")
	}
	return bad <= 1
}
