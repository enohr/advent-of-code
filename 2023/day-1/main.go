package main

import (
	"fmt"
	"os"
	"strings"
  "unicode"
  "strconv"
)

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

	for _, line := range strings.Split(input, "\n") {
    first, second := "", ""
     
    for _, char := range line {
      if !unicode.IsDigit(char) {
        continue
      }
      
      if first == "" {
        first = string(char)
      }
      second = string(char)
    }

    num, _ := strconv.Atoi(first + second)
    sum = sum + num
  }
  return sum
}

func part2(input string) int {
  sum := 0
  
	for _, line := range strings.Split(input, "\n") {
    first, second := "", ""
    
    line = transformData(line)

    for _, char := range line {
      if !unicode.IsDigit(char) {
        continue 
      }
      if first == "" {
        first = string(char)
      }
      second = string(char)
    }

    num, _ := strconv.Atoi(first + second)
    sum = sum + num
  }

  return sum
}

// Needs to keep the letters for word overlapping 
// like: oneight, which is 18
var validNumbers = map[string]string {
  "one": "o1e",
  "two": "t2o",
  "three": "t3ee",
  "four": "f4r",
  "five": "f5e",
  "six": "s6",
  "seven": "7n",
  "eight": "e8t",
  "nine": "n9e",
}

// If we replace the words if its respective number
// we only need to repeat part 1.
func transformData(line string) string{
  newString := line

  for key, value := range(validNumbers) {
    newString = strings.ReplaceAll(newString, key, value)
  }
  return newString
}
