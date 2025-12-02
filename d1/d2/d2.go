package main

import (
	_ "embed"
	"fmt"
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
	sum := 0
	for _, v := range ranges {
		var len_rad int = len(v[0][0])/2
		var rad int = v[0][0]
	return sum
}

func main() {
	fmt.Println(format(input))
}
