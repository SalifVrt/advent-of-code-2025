package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func format(txt string) [][]string {
	lines := strings.Split(strings.TrimSpace(txt), ",")
	var ranges [][]string = make([][]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(strings.TrimSpace(line), "-")
		ranges[i] = []string{parts[0], parts[1]}
	}
	return ranges
}

func solve(ranges [][]string) int {
	var sum int = 0
	var rad int = 0

	for _, v := range ranges {
		var len_rad int = len(v[0]) / 2
		rad1, _ := strconv.Atoi(v[0][:len_rad])
		rad2, _ := strconv.Atoi(v[0][len_rad:])

		//Choose the radical to use (radical is the number split in half, which will then be doubled)
		rad = min(rad1, rad2)

		//Calculate the full number by doubling the chosen radical
		full_num, _ := strconv.Atoi(strconv.Itoa(rad) + strconv.Itoa(rad))
		inf, _ := strconv.Atoi(v[0]) //Lower limit
		sup, _ := strconv.Atoi(v[1]) //Upper limit
		for full_num <= sup {
			if full_num >= inf {
				sum += full_num
			}
			rad++
			full_num, _ = strconv.Atoi(strconv.Itoa(rad) + strconv.Itoa(rad))
		}
	}
	return sum
}

func main() {
	//fmt.Println(format(input))
	//Part 1
	ans := solve(format(input))
	fmt.Println("Answer Part 1:", ans)
}
