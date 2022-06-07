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

	fmt.Println(search(District{w, h}, voters, make(map[District]int)))
}

func search(district District, voters [][]int, memo map[District]int) int {
	if memo[district] != 0 {
		return memo[district]
	}

	maxScore := voters[district.height-1][district.width-1]
	for _, s := range getAllSplits(district.width, district.height) {
		first := search(s.first, voters, memo)
		snd := search(s.second, voters, memo)
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
	if w == 1 && h == 1 {
		return splits
	}

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
