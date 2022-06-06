package main

import "testing"

func TestAllPossibleSplits(t *testing.T) {

	width := 7
	height := 6

	initialDistrict := makeDistrict(7, 7)

	votersByDimension := make(map[int]map[int]int)

	for i := 0; i < height; i++ {
		votersByDimension[i] = make(map[int]int)
		for j := 0; j < width; j++ {
			votersByDimension[i][j] = 1
		}
	}

	initialDistricts := []District{initialDistrict}
	best := findMaxSplitScore(initialDistricts, 0, votersByDimension)

	debug("best:", best)
}
