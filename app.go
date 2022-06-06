package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type SplitDirection int

// const (
// 	Horizontal SplitDirection = iota
// 	Vertical
// )

type District struct {
	width  int
	height int
}

type Split struct {
	districts []District
}

// debug any value to stderr
func debug(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	var w, h int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &w, &h)

	votersByDimension := make(map[int]map[int]int)

	for i := 0; i < h; i++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		for j := 0; j < w; j++ {
			voters, _ := strconv.ParseInt(inputs[j], 10, 32)
			_ = voters
			if votersByDimension[i+1] == nil {
				votersByDimension[i+1] = make(map[int]int)
			}
			votersByDimension[i+1][j+1] = int(voters)
		}
	}

	// debug("votersByDimension:", votersByDimension)

	max := findMaxDistrictScore(w, h, votersByDimension)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(max) // Write answer to stdout
}

func getVotersByDimension(votersByDimension map[int]map[int]int, w int, h int) int {
	return votersByDimension[h][w]
}

func findMaxSplitScore(districts []District, depth int, votersByDimension map[int]map[int]int) int {
	// debug("depth:", depth, "districts length:", len(districts), "districts:", districts)

	maxScore := computeDistrictsScore(districts, votersByDimension)

	// for each district, try to split it, for each possible split, store the resulting districts in results
	for iDistrict := 0; iDistrict < len(districts); iDistrict++ {
		// all but the iDistrict
		otherDistricts := make([]District, 0)
		for iOtherDistrict := 0; iOtherDistrict < len(districts); iOtherDistrict++ {
			if iOtherDistrict != iDistrict {
				otherDistricts = append(otherDistricts, districts[iOtherDistrict])
			}
		}

		for _, split := range getAllWaysToSplit(districts[iDistrict].width, districts[iDistrict].height) {
			maxSubSplitsScore := findMaxSplitScore(split.districts, depth+1, votersByDimension)

			splitScore := computeDistrictsScore(otherDistricts, votersByDimension) + maxSubSplitsScore

			if splitScore > maxScore {
				maxScore = splitScore
				// debug("Found new max score:", maxScore, "for split:", split)
			}
		}
	}

	return maxScore
}

func assert(condition bool) {
	if !condition {
		panic("assertion failed")
	}
}

func makeDistrict(w int, h int) District {
	assert(w > 0 && h > 0)

	return District{
		width:  w,
		height: h,
	}
}

func getAllWaysToSplit(w int, h int) []Split {
	results := make([]Split, 0)

	if w == 1 && h == 1 {
		return results
	}

	// debug("Splitting:", w, "x", h)

	for i := 0; i < h-1; i++ {
		newSplit := Split{
			districts: []District{
				makeDistrict(w, i+1),
				makeDistrict(w, h-i-1),
			},
		}
		results = append(results, newSplit)
	}

	for j := 0; j < w-1; j++ {
		newSplit := Split{
			districts: []District{
				makeDistrict(j+1, h),
				makeDistrict(w-j-1, h),
			},
		}
		results = append(results, newSplit)
	}

	return results
}

func districtSize(district District) int {
	return district.width * district.height
}

func districtsSize(districts []District) int {
	size := 0
	for _, district := range districts {
		size += districtSize(district)
	}
	return size
}

func computeDistrictsScore(districts []District, votersByDimension map[int]map[int]int) int {
	score := 0
	for _, district := range districts {
		voters := getVotersByDimension(votersByDimension, district.width, district.height)
		// debug("voters:", voters, "for district:", district)
		score += voters
	}
	return score
}

func findMaxDistrictScore(w int, h int, votersByDimension map[int]map[int]int) int {
	return findMaxSplitScore([]District{makeDistrict(w, h)}, 0, votersByDimension)
}
