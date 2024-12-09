package main

import (
	"fmt"
	"strconv"

	"github.com/RodrigoVieira938/aoc24/utils"
)

func generate_blockmap(diskmap string) []int {
	filenum := 0
	blockmap := []int{}
	for i, number_str := range diskmap {
		number, error := strconv.Atoi(string(number_str))
		if error != nil {
			continue
		}
		for x := 0; x < number; x++ {
			if i%2 == 0 {
				blockmap = append(blockmap, filenum)
			} else {
				blockmap = append(blockmap, -1)
			}
		}
		if i%2 == 0 {
			filenum++
		}

	}
	return blockmap
}

func puzzle1(blockmap []int) int {
	for i := len(blockmap) - 1; i >= 0; i-- {
		//find first blank space
		blank_space := -1
		for x, c := range blockmap {
			//if x is bigger than i, it would fragment instead of compress
			if c == -1 && i > x {
				blank_space = x
				break
			}
		}
		if blank_space == -1 {
			break
		}
		blockmap[blank_space] = blockmap[i]
		blockmap[i] = -1
	}
	checksum := 0

	for filepos, fileid := range blockmap {
		if fileid == -1 {
			break
		}
		checksum += filepos * fileid
	}

	return checksum
}

func main() {
	diskmap, error := utils.ReadFileStr("./data/day9.txt")
	if error == nil {
		//fmt.Println(diskmap)
		blockmap := generate_blockmap(diskmap)
		fmt.Printf("Day 9 - Puzzle 1 answer is %v\n", puzzle1(blockmap))
	}
}
