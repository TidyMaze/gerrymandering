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
	firstDistrict  District
	secondDistrict District
}

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
			if votersByDimension[i] == nil {
				votersByDimension[i] = make([]int, w)
			}
			votersByDimension[i][j] = int(voters)
		}
	}

	fmt.Println(search(makeDistrict(w, h), votersByDimension, make(map[District]int)))
}

func search(district District, votersByDimension [][]int, memo map[District]int) int {
	if memo[district] != 0 {
		return memo[district]
	}

	maxScore := votersByDimension[district.height-1][district.width-1]
	splits := getAllWaysToSplit(district.width, district.height)
	for iSplit := 0; iSplit < len(splits); iSplit++ {
		fstScore := search(splits[iSplit].firstDistrict, votersByDimension, memo)
		sndScore := search(splits[iSplit].secondDistrict, votersByDimension, memo)
		splitScore := fstScore + sndScore

		if splitScore > maxScore {
			maxScore = splitScore
		}
	}

	// debug("memoizedFindMaxSplitScore:", district, maxScore)
	memo[district] = maxScore
	return maxScore
}

func makeDistrict(w int, h int) District {
	return District{
		width:  w,
		height: h,
	}
}

func getAllWaysToSplit(w int, h int) []Split {
	splits := make([]Split, 0)
	if w == 1 && h == 1 {
		return splits
	}

	for i := 0; i < h-1; i++ {
		splits = append(splits, Split{
			firstDistrict:  makeDistrict(w, i+1),
			secondDistrict: makeDistrict(w, h-i-1),
		})
	}

	for j := 0; j < w-1; j++ {
		splits = append(splits, Split{
			firstDistrict:  makeDistrict(j+1, h),
			secondDistrict: makeDistrict(w-j-1, h),
		})
	}

	return splits
}
