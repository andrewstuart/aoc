package main

import (
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSplit(t *testing.T) {
	asrt := assert.New(t)
	a, b := split([]rune("vJrwpWtwJgWrhcsFMMfFFhFp"))
	asrt.Equal("vJrwpWtwJgWr", string(a))
	asrt.Equal("hcsFMMfFFhFp", string(b))

}

func TestPri(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(38, pri('L'))
	asrt.Equal(16, pri('p'))
}

func TestAOC(t *testing.T) {
	asrt := assert.New(t)
	in := strings.NewReader(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

	asrt.Equal(157, aoc(in))
}
