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
	lines := strings.Split(strings.TrimSpace(txt), "\n")
	var dirs [][]string = make([][]string, len(lines))
	for i, line := range lines {
		dirs[i] = make([]string, 2)
		for j, v := range line {
			if j >= 1 {
				if j == 1 {
					dirs[i][1] = ""
				}
				dirs[i][1] += string(v)
			} else if j == 0 {
				dirs[i][0] = string(v)
			}
		}
	}
	return dirs
}

func solve1(dirs [][]string) int {
	var ct int = 0
	var cur int = 50
	for _, v := range dirs {
		//Counting the times we reach 0
		if cur == 0 {
			ct += 1
		}

		//To chose in which direction we turn the dial
		n, _ := strconv.Atoi(v[1])
		switch v[0] {
		case "R":
			cur += n
		case "L":
			cur -= n
		}

		//If we go above the limits
		if cur < 0 {
			for cur < 0 {
				cur = 100 + cur
			}
		} else if cur > 99 {
			for cur > 99 {
				cur = cur - 100
			}

		}
	}
	return ct
}

func main() {
	dirs := format(input)
	//Part 1
	ans := solve1(dirs)
	fmt.Println("Answer Part 1:", ans)
}
