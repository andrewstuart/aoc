package ezaoc

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadString(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)
	br := bufio.NewReader(strings.NewReader("fooXbarXYZfooXYbarXYZbaz"))
	st, err := readMultiString(br, "XYZ")
	rq.NoError(err)
	asrt.Equal(st, "fooXbarXYZ")

	st, err = readMultiString(br, "XYZ")
	rq.NoError(err)
	asrt.Equal(st, "fooXYbarXYZ")
}

func TestReadMultiNewline(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)
	br := bufio.NewReader(strings.NewReader("foo\nbar\n\n\nfoo\n\nbar\n\n\nbaz"))
	st, err := readMultiString(br, "\n\n\n")
	rq.NoError(err)
	asrt.Equal("foo\nbar\n\n\n", st)

	st, err = readMultiString(br, "\n\n\n")
	rq.NoError(err)
	asrt.Equal("foo\n\nbar\n\n\n", st)

	st, err = readMultiString(br, "\n\n\n")
	asrt.Error(err)
	asrt.Equal("baz", st)
}
