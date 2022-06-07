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

	fmt.Println(memoizedFindMaxSplitScore(makeDistrict(w, h), 0, votersByDimension, make(map[District]int)))
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
		fstScore := memoizedFindMaxSplitScore(splits[iSplit].districts[0], depth+1, votersByDimension, memo)
		sndScore := memoizedFindMaxSplitScore(splits[iSplit].districts[1], depth+1, votersByDimension, memo)
		splitScore := fstScore + sndScore

		if splitScore > maxScore {
			maxScore = splitScore
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

func computeDistrictsScore(districts []District, votersByDimension [][]int) int {
	score := 0
	for i := 0; i < len(districts); i++ {
		score += votersByDimension[districts[i].height-1][districts[i].width-1]
	}
	return score
}
