package main

import (
	"fmt"
	"os"
	"strings"
)

type box struct {
	stacks [][]string
}

type move struct {
	quant int
	to    int
	from  int
}

func main() {
	file, _ := os.ReadFile("input.txt")
	input := string(file)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) string {
	box := &box{}
	box.stacks = make([][]string, 10)
	moves := box.parse(input)
	box.move1(moves)
	res := box.result()
	return res
}

func part2(input string) string {
	box := &box{}
	box.stacks = make([][]string, 10)
	moves := box.parse(input)
	box.move2(moves)
	res := box.result()
	return res
}

func (b *box) parse(input string) (moves []move) {
	sum := 0
	pos := 0
	splitted := strings.Split(input, "\n\n")

	stacks := splitted[0]

	for _, line := range strings.Split(stacks, "\n") {
		sum = 0
		pos = 0
		for i := range line {
			char := line[i]
			if char == '[' {
				pos = sum/4 + 1 + pos
				box := line[i+1]
				b.insertBottom(pos, string(box))
				sum = 0
			}
			if char == ' ' {
				sum++
			}
		}
	}
	steps := splitted[1]
	for _, line := range strings.Split(steps, "\n") {
		move := move{}
		fmt.Sscanf(line, "move %d from %d to %d", &move.quant, &move.from, &move.to)
		moves = append(moves, move)
	}

	return
}

func (b *box) move1(moves []move) {
	for _, move := range moves {
		for i := 0; i < move.quant; i++ {
			item := b.remove(move.from, 1)
			b.insertTop(move.to, item)
		}
	}
}

func (b *box) move2(moves []move) {
	for _, move := range moves {
		item := b.remove(move.from, move.quant)
		b.insertTop(move.to, item)
	}
}

func (b *box) result() (res string) {
	for _, stack := range b.stacks {
		if len(stack) == 0 {
			continue
		}
		res = res + stack[len(stack)-1]
	}
	return
}

func (b *box) getByPos(pos int) []string {
	return b.stacks[pos]
}

func (b *box) insertBottom(pos int, crate string) {
	stack := b.getByPos(pos)
	b.stacks[pos] = append([]string{crate}, stack...)
}

func (b *box) insertTop(pos int, crate []string) {
	stack := b.getByPos(pos)
	b.stacks[pos] = append(stack, crate...)
}

func (b *box) remove(pos, quantity int) []string {
	stack := b.getByPos(pos)
	ret := stack[len(stack)-quantity:]
	b.stacks[pos] = stack[:len(stack)-quantity]

	return ret

}
