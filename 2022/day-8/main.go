package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	arr := parse(input)

	// Start with all external trees
	sum := len(arr)*2 + (len(arr)-2)*2

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-1; j++ {
			val := arr[i][j]

			// We'll assume it's visible for all sides
			top, left, down, right := true, true, true, true

			//check left
			for k := j - 1; k >= 0; k-- {
				n := arr[i][k]
				if n >= val {
					left = false
					break
				}
			}

			// top
			for k := i - 1; k >= 0; k-- {
				n := arr[k][j]
				if n >= val {
					top = false
					break
				}
			}

			//right
			for k := j + 1; k <= len(arr)-1; k++ {
				n := arr[i][k]

				if n >= val {
					right = false
					break
				}
			}

			//down
			for k := i + 1; k <= len(arr)-1; k++ {
				n := arr[k][j]
				if n >= val {
					down = false
					break
				}
			}

			// If it's visible for a side, sum+1
			if left || top || right || down {
				sum++
			}
		}

	}
	return sum

}

func part2(input string) int {
	arr := parse(input)
	max := 0

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr)-1; j++ {
			val := arr[i][j]

			top, left, down, right := 0, 0, 0, 0
			//check left
			for k := j - 1; k >= 0; k-- {
				left++
				tree := arr[i][k]
				if tree >= val {
					break
				}
			}

			// top
			for k := i - 1; k >= 0; k-- {
				top++
				tree := arr[k][j]
				if tree >= val {
					break
				}
			}

			//right
			for k := j + 1; k <= len(arr)-1; k++ {
				right++
				tree := arr[i][k]
				if tree >= val {
					break
				}
			}

			//down
			for k := i + 1; k <= len(arr)-1; k++ {
				down++
				tree := arr[k][j]
				if tree >= val {
					break
				}
			}
			visible := top * left * down * right

			if visible > max {
				max = visible
			}
		}

	}
	return max

}

func parse(input string) [][]int {
	var arr [][]int
	for _, line := range strings.Split(input, "\n") {
		r := make([]int, 0)
		for _, char := range strings.Split(line, "") {
			n, _ := strconv.Atoi(char)
			r = append(r, n)
		}
		arr = append(arr, r)
	}
	return arr
}
