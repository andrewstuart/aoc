package main

import (
	"bufio"
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

type file struct {
	id     int
	blocks int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		log.Fatal(err)
	}
	ins := inputs[0]

	count := 0
	files := map[int]file{}
	var blocks []int
	for i, in := range ins {
		if i%2 == 0 {
			id := i / 2
			files[in] = file{id: id, blocks: in}
			blk := make([]int, in)
			fill(blk, id)
			blocks = append(blocks, blk...)
			continue
		}
		blk := make([]int, in)
		fill(blk, -1)
		blocks = append(blocks, blk...)
	}

	blks := make([]int, len(blocks))
	copy(blks, blocks)

	i, j := 0, len(blks)-1
	for {
		if blks[i] != -1 {
			i++
			continue
		}
		if blks[j] == -1 {
			j--
			continue
		}
		if i > j {
			break
		}
		blks[i], blks[j] = blks[j], blks[i]
	}

	for i, id := range blks {
		if id == -1 {
			continue
		}
		count += i * id
	}

	return count
}

func fill(bs []int, fill int) {
	for i := 0; i < len(bs); i++ {
		bs[i] = fill
	}
}
