package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
  RED =  12
  GREEN = 13
  BLUE = 14
)

var colorValues = map[string]int {
  "RED": RED,
  "GREEN": GREEN,
  "BLUE": BLUE,
}

func main() {
	file, _ := os.ReadFile("input1.txt")
	input1 := string(file)
  fmt.Println(part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := string(file)
  fmt.Println(part2(input2))
}

func part1(input string) int {
  sum := 0

  for id, line := range strings.Split(input, "\n") {
    if line == "" {
      continue
    }
    line = strings.ReplaceAll(line, ";", ",")

    if isGamePossible(line) {
      sum = sum + id + 1
    }
  }
  return sum
}

func isGamePossible(line string) bool {
  line = strings.Split(line, ":")[1]
  splitted := strings.Split(line, ",")

  for _, item := range splitted {
    item = strings.TrimSpace(item)
    s := strings.Split(item, " ")
    
    color := strings.ToUpper(s[1])
    num, _ := strconv.Atoi(s[0])

    if colorValues[color] < num {
      return false
    }
  }
  return true
}

func part2(input string) int {
  sum := 0

  for _, line := range strings.Split(input, "\n") {
    if line == "" {
      continue
    }
    line = strings.ReplaceAll(line, ";", ",")

    sum += getPowerSet(line)
  }
  return sum
}

func getPowerSet(line string) int {
  var colors = map[string]int{
    "RED": -1,
    "GREEN": -1,
    "BLUE": -1,
  }
  sum := 1

  line = strings.Split(line, ":")[1]
  splitted := strings.Split(line, ",")

  for _, item := range splitted {
    item = strings.TrimSpace(item)
    s := strings.Split(item, " ")
    
    color := strings.ToUpper(s[1])
    num, _ := strconv.Atoi(s[0])

    if colors[color] < num {
      colors[color] = num
    }
  }

  for _, val := range colors {
    sum *= val
  }
  return sum

}
