package main

import "testing"

// func TestAllPossibleSplits(t *testing.T) {
// 	oneMeasure()
// }

func BenchmarkOneMeasure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oneMeasure()
	}
}

func oneMeasure() {
	width := 50
	height := 50

	initialDistrict := makeDistrict(width, height)

	votersByDimension := make([][]int, height)

	for i := 0; i < height; i++ {
		votersByDimension[i] = make([]int, width)
		for j := 0; j < width; j++ {
			votersByDimension[i][j] = 1
		}
	}

	cache := make(map[District]int)
	memoizedFindMaxSplitScore(initialDistrict, votersByDimension, cache)
}
