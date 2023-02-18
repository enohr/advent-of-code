package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

var moves = map[string]position{
	"U": {x: 1, y: 0},
	"R": {x: 0, y: 1},
	"D": {x: -1, y: 0},
	"L": {x: 0, y: -1},
}

// Rope is an array like:
// [T, X, X, X, X, H]
// Could be an LinkedList too (H.next, etc), but array works as well
var rope []position

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	// Use map to mark visited cause don't allow dups
	visited := map[position]bool{}
	head := &position{
		x: 0, y: 0,
	}

	tail := &position{
		x: 0, y: 0,
	}

	visited[*tail] = true

	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, " ")
		dir := splitted[0]
		qty, _ := strconv.Atoi(splitted[1])
		move := moves[dir]
		for i := 0; i < qty; i++ {
			head.x += move.x
			head.y += move.y
			updateTail(head, tail)
			visited[*tail] = true
		}
	}
	return len(visited)
}

func part2(input string) int {
	// Use map to mark visited cause don't allow dups
	visited := map[position]bool{}
	var head, tail *position
	rope = make([]position, 10)

	for _, line := range strings.Split(input, "\n") {
		splitted := strings.Split(line, " ")
		dir := splitted[0]
		qty, _ := strconv.Atoi(splitted[1])
		move := moves[dir]
		for i := 0; i < qty; i++ {
			head = &rope[len(rope)-1]
			head.x += move.x
			head.y += move.y
			rope[len(rope)-1] = *head

			// Update all rope parts, starting from head to tail
			for i := len(rope) - 1; i > 0; i-- {
				head = &rope[i]
				tail = &rope[i-1]
				updateTail(head, tail)
				rope[i-1] = *tail
			}
			visited[*tail] = true
		}
	}

	return len(visited)
}

func updateTail(head, tail *position) {
	diffX := head.x - tail.x
	diffY := head.y - tail.y
	absX := abs(diffX)
	absY := abs(diffY)

	x := 1
	y := 1
	// Check if should increase or decrease
	if diffX < 0 {
		x = -1
	}
	if diffY < 0 {
		y = -1
	}

	if absX > 1 || absY > 1 {
		if absX > 0 && absY > 0 {
			tail.x += int(x)
			tail.y += int(y)
		} else if absY > 1 {
			tail.y += int(y)
		} else if absX > 1 {
			tail.x += int(x)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
