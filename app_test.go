package main

import "testing"

func TestAllPossibleSplits(t *testing.T) {

	initialDistrict := makeDistrict(6, 6)
	initialDistricts := []District{initialDistrict}
	allPossibleSplits := findAllSplits(initialDistricts, 0)

	// print all possible splits
	for _, split := range allPossibleSplits {
		debug("split:", split)

		assert(districtsSize(split) == districtSize(initialDistrict))
	}
}
