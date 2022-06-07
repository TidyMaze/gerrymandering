package main

import "testing"

func TestAllPossibleSplits(t *testing.T) {

	width := 70
	height := 70

	initialDistrict := makeDistrict(width, height)

	votersByDimension := make([][]int, height)

	for i := 0; i < height; i++ {
		votersByDimension[i] = make([]int, width)
		for j := 0; j < width; j++ {
			votersByDimension[i][j] = 1
		}
	}

	cache := make(map[District]int)
	best := memoizedFindMaxSplitScore(initialDistrict, votersByDimension, cache)

	debug("best:", best)
}
