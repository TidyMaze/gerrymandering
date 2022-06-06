package main

import "testing"

func TestAllPossibleSplits(t *testing.T) {

	width := 7
	height := 7

	initialDistrict := makeDistrict(width, height)

	votersByDimension := make([][]int, height)

	for i := 0; i < height; i++ {
		votersByDimension[i] = make([]int, width)
		for j := 0; j < width; j++ {
			votersByDimension[i][j] = 1
		}
	}

	initialDistricts := []District{initialDistrict}
	best := findMaxSplitScore(initialDistricts, 0, votersByDimension)

	debug("best:", best)
}
