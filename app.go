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
	first  District
	second District
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

	voters := make([][]int, h)

	for i := 0; i < h; i++ {
		scanner.Scan()
		inputs = strings.Split(scanner.Text(), " ")
		for j := 0; j < w; j++ {
			cost, _ := strconv.ParseInt(inputs[j], 10, 32)
			if voters[i] == nil {
				voters[i] = make([]int, w)
			}
			voters[i][j] = int(cost)
		}
	}

	fmt.Println(search(District{w, h}, voters))
}

// find the maximum score after splitting a district
func search(district District, voters [][]int) int {
	return searchMemo(district, voters, make(map[District]int))
}

func searchMemo(district District, voters [][]int, memo map[District]int) int {
	// shortcut if already calculated
	if memo[district] != 0 {
		return memo[district]
	}

	maxScore := voters[district.height-1][district.width-1]
	for _, s := range getAllSplits(district.width, district.height) {
		first := searchMemo(s.first, voters, memo)
		snd := searchMemo(s.second, voters, memo)
		maxScore = max(maxScore, first+snd)
	}

	memo[district] = maxScore
	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getAllSplits(w int, h int) []Split {
	splits := make([]Split, 0)

	for i := 0; i < h-1; i++ {
		splits = append(splits, Split{
			first:  District{w, i + 1},
			second: District{w, h - i - 1},
		})
	}

	for j := 0; j < w-1; j++ {
		splits = append(splits, Split{
			first:  District{j + 1, h},
			second: District{w - j - 1, h},
		})
	}

	return splits
}
