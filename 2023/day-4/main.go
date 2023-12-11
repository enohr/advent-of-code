package main

import (
  "fmt"
  "os"
  "strings"
  "slices"
  "math"
)

func main() {
  file, _ := os.ReadFile("input2.txt")
  input1 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part1:", part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
  sum := 0 

  for _, line := range strings.Split(input, "\n") {
    w := 0

    values := strings.FieldsFunc(line, func(r rune) bool {
      return r == ':' || r == '|'
    })

    winnings := strings.Fields(values[1])
    ours := strings.Fields(values[2])

    for _, our := range ours {
      if slices.Contains(winnings, our) {
        w = w + 1
      }
    }
    sum = sum + IntPow(2, w-1) 
  }
  return sum
}

func IntPow(n, m int) int {
  return int(math.Pow(float64(n), float64(m)))
}

func part2(input string) int {
  sum := 0
  lines := strings.Split(input, "\n")
  cards := make([]int, len(lines))

  for i, line := range lines {
    w := 0

    values := strings.FieldsFunc(line, func(r rune) bool {
      return r == ':' || r == '|'
    })

    winnings := strings.Fields(values[1])
    ours := strings.Fields(values[2])

    for _, our := range ours {
      if slices.Contains(winnings, our) {
        w = w + 1
      }
    }

    // Add 1 to current
    cards[i]++
    copies := cards[i]

    // Iterate over i+1 .. w+i
    for j := i+1; j <= w + i; j++ {
      if j >= len(cards) {
        continue
      }
      cards[j] = cards[j] + copies
    }

    sum += cards[i]
  }
  return sum
}
