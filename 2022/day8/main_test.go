package main

import (
	"strings"
	"testing"

	"github.com/andrewstuart/aoc2022/pkg/ezaoc"
	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	asrt := assert.New(t)
	asrt.Equal(21, aoc(strings.NewReader(`30373
25512
65332
33549
35390`)))
}

func TestSplit(t *testing.T) {
	asrt := assert.New(t)

	asrt.Equal(
		[][]byte{[]byte("abc"), []byte("efg")},
		ezaoc.Reslice([]byte("abcdefg"), splitExcept[byte](3)),
	)

	asrt.Equal(
		[][]byte{[]byte("abcdef"), {}},
		split([]byte("abcdefg"), 6),
	)
}
