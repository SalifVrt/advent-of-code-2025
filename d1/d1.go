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

func solve(dirs [][]string, part int) int {
	var ct int = 0
	var cur int = 50
	for _, v := range dirs {
		//Counting the times we reach 0 (Part 1 only)
		if part == 1 && cur == 0 {
			ct += 1
		}

		//To choose direction and distance for this move
		n, _ := strconv.Atoi(v[1])

		//Counting the times we reach 0 (Part 2 only)
		if part == 2 {
			// target step k0 such that (cur + step*k) % 100 == 0
			// for R (step = +1): k0 = (100 - cur) % 100
			// for L (step = -1): k0 = cur % 100
			k0 := 0
			if v[0] == "R" {
				k0 = (100 - cur) % 100
			} else {
				k0 = cur % 100
			}
			if k0 == 0 {
				k0 = 100 // first occurrence is at step 100,200,...
			}
			if n >= k0 {
				ct += 1 + (n-k0)/100
			}
		}

		switch v[0] {
		case "R":
			cur = (cur + n) % 100
		case "L":
			cur = (cur - (n % 100)) % 100
		}
		if cur < 0 {
			cur += 100
		}
	}
	return ct
}

func main() {
	dirs := format(input)
	//Part 1
	ans := solve(dirs, 1)
	fmt.Println("Answer Part 1:", ans)

	//Part 2
	ans = solve(dirs, 2)
	fmt.Println("Answer Part 2:", ans)
}
