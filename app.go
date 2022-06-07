package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	votersByDimension := make([][]int, h)

	for i := 0; i < h; i++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		for j := 0; j < w; j++ {
			voters, _ := strconv.ParseInt(inputs[j], 10, 32)
			_ = voters
			if votersByDimension[i] == nil {
				votersByDimension[i] = make([]int, w)
			}
			votersByDimension[i][j] = int(voters)
		}
	}

	// debug("votersByDimension:", votersByDimension)

	max := findMaxDistrictScore(w, h, votersByDimension)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(max) // Write answer to stdout
}

func getVotersByDimension(votersByDimension [][]int, w int, h int) int {
	return votersByDimension[h-1][w-1]
}

func findMaxSplitScore(district District, depth int, votersByDimension [][]int) int {
	// debug("depth:", depth, "districts:", district)
	maxScore := computeDistrictsScore([]District{district}, votersByDimension)
	splits := make([]Split, 0)
	getAllWaysToSplit(district.width, district.height, &splits)
	for iSplit := 0; iSplit < len(splits); iSplit++ {
		firstDistrictMaxScore := findMaxSplitScore(splits[iSplit].districts[0], depth+1, votersByDimension)
		secondDistrictMaxScore := findMaxSplitScore(splits[iSplit].districts[1], depth+1, votersByDimension)
		splitScore := firstDistrictMaxScore + secondDistrictMaxScore

		if splitScore > maxScore {
			maxScore = splitScore
			// debug("Found new max score:", maxScore, "for split:", splits[iSplit])
		}
	}

	return maxScore
}

func memoizedFindMaxSplitScore(district District, depth int, votersByDimension [][]int, memo map[District]int) int {
	if memo[district] != 0 {
		debug("memoizedFindMaxSplitScore:", district, "=", memo[district])
		return memo[district]
	}

	maxScore := computeDistrictsScore([]District{district}, votersByDimension)
	splits := make([]Split, 0)
	getAllWaysToSplit(district.width, district.height, &splits)
	for iSplit := 0; iSplit < len(splits); iSplit++ {
		firstDistrictMaxScore := memoizedFindMaxSplitScore(splits[iSplit].districts[0], depth+1, votersByDimension, memo)
		secondDistrictMaxScore := memoizedFindMaxSplitScore(splits[iSplit].districts[1], depth+1, votersByDimension, memo)
		splitScore := firstDistrictMaxScore + secondDistrictMaxScore

		if splitScore > maxScore {
			maxScore = splitScore
			// debug("Found new max score:", maxScore, "for split:", splits[iSplit])
		}
	}

	memo[district] = maxScore
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

func getAllWaysToSplit(w int, h int, splits *[]Split) {
	*splits = (*splits)[:0]
	if w == 1 && h == 1 {
		return
	}

	// debug("Splitting:", w, "x", h)

	for i := 0; i < h-1; i++ {
		*splits = append(*splits, Split{
			districts: []District{
				makeDistrict(w, i+1),
				makeDistrict(w, h-i-1),
			},
		})
	}

	for j := 0; j < w-1; j++ {
		*splits = append(*splits, Split{
			districts: []District{
				makeDistrict(j+1, h),
				makeDistrict(w-j-1, h),
			},
		})
	}
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

func computeDistrictsScore(districts []District, votersByDimension [][]int) int {
	score := 0

	for iDistrict := 0; iDistrict < len(districts); iDistrict++ {
		voters := getVotersByDimension(votersByDimension, districts[iDistrict].width, districts[iDistrict].height)
		// debug("voters:", voters, "for district:", district)
		score += voters
	}
	return score
}

func findMaxDistrictScore(w int, h int, votersByDimension [][]int) int {
	cache := make(map[District]int)
	return memoizedFindMaxSplitScore(makeDistrict(w, h), 0, votersByDimension, cache)
}
