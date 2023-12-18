package main

import (
	"fmt"
	"os"
	"strings"
  "math"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
  input1 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part1:", part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part2:", part2(input2))
}

type Node struct {
  Left string
  Right string
}

func part1(input string) int {
  parts := strings.Split(input, "\n")
  order := parts[0]
  mapping := parts[2:]
  
  nodes := make(map[string]Node, len(mapping))
  for _, m := range mapping {
    split := strings.FieldsFunc(m, func(r rune) bool {
      return r == '=' || r == ',' || r == '(' || r == ')'
    })
    value := strings.TrimSpace(split[0])
    left := strings.TrimSpace(split[2])
    right := strings.TrimSpace(split[3])
    nodes[value] = Node{Left: left, Right: right}
  }
  
  current := "AAA"
  movements := 0
  for current != "ZZZ" {
    movement := order[movements % len(order)]
    node := nodes[current]
    if movement == 'L' {
      current = node.Left
    } else {
      current = node.Right
    }
    movements++
  }
  return movements
}

func part2(input string) int {
  parts := strings.Split(input, "\n")
  order := parts[0]
  mapping := parts[2:]
  var starting []string

  nodes := make(map[string]Node, len(mapping))
  for _, m := range mapping {
    split := strings.FieldsFunc(m, func(r rune) bool {
      return r == '=' || r == ',' || r == '(' || r == ')'
    })
    value := strings.TrimSpace(split[0])
    left := strings.TrimSpace(split[2])
    right := strings.TrimSpace(split[3])
    nodes[value] = Node{Left: left, Right: right}
    if value[2] == 'A' {
      starting = append(starting, value)
    }
  }
  
  var movements []int
  for _, start := range starting {
    movements = append(movements, walk(start, order, nodes))
  }

  // If we reach the first xxZ in 2 steps, it will hit again
  // on all 2 multiples (4, 6, 8 and so on), so the cycles has
  // the same amount of steps to complete.
  // e.g.: If the steps are 2 and 3, this will happen on 2, 4, 6... and 3, 6, 9...
  // then the result will be 6, when the two cycles ends at the sime time.
  // To find this result, just need to calculate the least common multiple.
  return leastCommonMultiple(movements)
}

func walk(current, order string, nodes map[string]Node) int {
  movements := 0
  for current[2] != 'Z' {
    movement := order[movements % len(order)]
    node := nodes[current]
    if movement == 'L' {
      current = node.Left
    } else {
      current = node.Right
    }
    movements++
  }
  return movements
}

// lcm of a list of numbers is the same as 
// lcm(a, lcm(b, lcm(c,d)))
func leastCommonMultiple(values []int) int {
  value := 1
  for _, v := range values {
    value = lcm(value, v) 
  }
  return value
}

// Source: https://github.com/TheAlgorithms/Go/blob/master/math/lcm/lcm.go
func lcm(a, b int) int {
  return int(math.Abs(float64(a*b))) / gcd(a,b)
}

// Source: https://github.com/TheAlgorithms/Go/blob/master/math/gcd/gcd.go
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
