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

type file struct {
	id     int
	blocks int
	start  int
}

func aoc(r io.Reader) int {
	inputs, err := ezaoc.ReadAOC(r, ezaoc.IntSlicer(""))
	if err != nil {
		log.Fatal(err)
	}
	ins := inputs[0]

	count := 0
	files := map[int]*file{}
	var blocks []int
	for i, in := range ins {
		if i%2 == 0 {
			id := i / 2
			files[id] = &file{id: id, blocks: in, start: len(blocks)}
			blk := make([]int, in)
			fill(blk, id)
			blocks = append(blocks, blk...)
			continue
		}
		blk := make([]int, in)
		fill(blk, -1)
		blocks = append(blocks, blk...)
	}

	print(blocks)
	blks := make([]int, len(blocks))
	copy(blks, blocks)

	i := 0
	maxID := len(files) - 1
	for {
		// don't go out of bounds
		if i >= len(blks) {
			break
		}
		// don't overwrite files
		if blks[i] != -1 {
			i++
			continue
		}

		// find the number of free blocks
		free := 0
		for i+free < len(blks) && blks[i+free] == -1 {
			free++
		}

		print(blks)
		fmt.Println(free)

		found := false
		// find the largest file ID with a small enough size to fit
		for jj := maxID; jj > 0; jj-- {
			f, ok := files[jj]
			if !ok || f.blocks > free || f.start <= i {
				continue
			}
			found = true
			for ii := 0; ii < f.blocks; ii++ {
				blks[i+ii] = f.id
				blks[f.start+ii] = -1
			}
			i += f.blocks
			delete(files, f.id)
			break
		}
		if !found {
			i++
		}
	}

	for i, id := range blks {
		if id == -1 {
			continue
		}
		count += i * id
	}

	return count
}

func print(is []int) {
	for _, i := range is {
		if i == -1 {
			fmt.Print(".")
			continue
		}
		fmt.Printf("%d", i)
	}
	fmt.Println()
}

func fill(bs []int, fill int) {
	for i := 0; i < len(bs); i++ {
		bs[i] = fill
	}
}
