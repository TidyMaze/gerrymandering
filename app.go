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

	fmt.Println(searchMemo(District{w, h}, voters, make(map[District]int)))
}

func searchMemo(district District, voters [][]int, memo map[District]int) int {
	// shortcut if already calculated
	if memo[district] != 0 {
		return memo[district]
	}
	maxScore := voters[district.height-1][district.width-1]
	for _, split := range getAllSplits(district.width, district.height) {
		first := searchMemo(split[0], voters, memo)
		snd := searchMemo(split[1], voters, memo)
		if first+snd > maxScore {
			maxScore = first + snd
		}
	}
	memo[district] = maxScore
	return maxScore
}

func getAllSplits(w int, h int) [][2]District {
	splits := make([][2]District, 0)

	// split horizontally
	for i := 0; i < h-1; i++ {
		splits = append(splits, [2]District{
			District{w, i + 1},
			District{w, h - i - 1},
		})
	}

	// split vertically
	for j := 0; j < w-1; j++ {
		splits = append(splits, [2]District{
			District{j + 1, h},
			District{w - j - 1, h},
		})
	}

	return splits
}
