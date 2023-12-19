package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	input1 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part1:", part1(input1))

	file, _ = os.ReadFile("input2.txt")
	input2 := strings.TrimRight(string(file), "\n")
	fmt.Println("Part2:", part2(input2))
}

type Pos struct {
	X int
	Y int
}

func part1(input string) int {
	matrix, starting := generateMatrix(input)
	moves := 0

	goal := starting
	previous := starting

	current := FindFirstMove(starting, matrix)
	for current != goal {
		next := GetNextMove(current, previous, matrix)
		previous = current
		current = next
		moves++
	}
	return (moves / 2) + 1
}

func part2(input string) int {
	return 1
}

func GetNextMove(current, previous Pos, matrix [][]rune) Pos {
	currentValue := matrix[current.X][current.Y]
	var possible []Pos

	switch currentValue {
	case '|': // up or down
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X + 1, Y: current.Y}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case '-': // left or right
		pos1 := Pos{X: current.X, Y: current.Y - 1}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'J': // up or left
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y - 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'L': // up or right
		pos1 := Pos{X: current.X - 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case '7': // down or left
		pos1 := Pos{X: current.X + 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y - 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	case 'F': // down or right
		pos1 := Pos{X: current.X + 1, Y: current.Y}
		pos2 := Pos{X: current.X, Y: current.Y + 1}
		possible = append(possible, pos1)
		possible = append(possible, pos2)
	}

	// We cant go back, so get the possible move
	// different than the previous one
	if possible[0] == previous {
		return possible[1]
	} else {
		return possible[0]
	}
}

func FindFirstMove(p Pos, board [][]rune) Pos {
	positions := []Pos{
		{X: -1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: 1, Y: 0},
	}

	for _, position := range positions {
		x := position.X + p.X
		y := position.Y + p.Y

		if x >= len(board) || y >= len(board) || x < 0 || y < 0 {
			continue
		}
		if board[x][y] == '.' {
			continue
		}
		return Pos{X: x, Y: y}
	}
	return Pos{X: 0, Y: 0}
}

func generateMatrix(input string) ([][]rune, Pos) {
	starting := Pos{X: 0, Y: 0}
	lines := strings.Split(input, "\n")

	matrix := make([][]rune, len(lines))
	for m := range matrix {
		matrix[m] = make([]rune, len(lines))
	}

	for x, line := range lines {
		for y, column := range line {
			if column == 'S' {
				starting = Pos{
					X: x,
					Y: y,
				}
			}
			matrix[x][y] = column
		}
	}
	return matrix, starting
}
