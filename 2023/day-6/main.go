package main

import (
  "fmt"
  "os"
  "strconv"
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

func part1(input string) int {
  parts := strings.Split(input, "\n")
  times := strings.Fields(parts[0])[1:] 
  distances := strings.Fields(parts[1])[1:]
  result := 1

  for i, time := range times {
    sum := 0
    record, _ := strconv.Atoi(distances[i])
    t, _ := strconv.Atoi(time)

    for i := 0; i <= t; i++ {
      ms := i
      ml := t - i
      traveled := ms * ml
      if traveled > record {
        sum++
      }
    }
    result *= sum
  }
  return result
}

func part2(input string) int {
  parts := strings.Split(input, "\n")
  times := strings.Fields(parts[0])[1:]
  distances := strings.Fields(parts[1])[1:]

  time, _ := strconv.Atoi(strings.Join(times, ""))
  record, _ := strconv.Atoi(strings.Join(distances, ""))

  sum := 0
  for i := 0; i <= time; i++ {
    ms := i
    ml := time - i
    traveled := ms * ml
    if traveled > record {
      sum++
    }
  }
  return sum
}
