package main

import (
  "fmt"
  "os"
  "slices"
  "strconv"
  "strings"
)

type Range struct {
  DestRange int
  SourceRange int
  Length int
}

type Converter struct {
  Ranges [][]Range
}

func main() {
  file, _ := os.ReadFile("input1.txt")
  input1 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part1:", part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := strings.TrimRight(string(file), "\n")
  fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
  var ret []int
  parts := strings.Split(input, "\n\n")
  seeds := strings.Fields(parts[0])[1:] 

  c := NewConverter(parts[1:])
  for _, s := range seeds {
    seed, _ := strconv.Atoi(s)
    ret = append(ret, c.From(seed))
  }
  return slices.Min(ret)
}

func NewConverter(parts []string) *Converter {
  c := &Converter{
    Ranges: make([][]Range, 7),
  }
  for i, p := range parts {
    values := strings.Split(p, "\n")[1:]
    for _, value := range values {
      v := StringArrayToInt(strings.Fields(value))
      r := Range{
        DestRange: v[0],
        SourceRange: v[1],
        Length: v[2],
      }
      c.Ranges[i] = append(c.Ranges[i], r)
    }
  }
  return c
}

func (c *Converter) From(seed int) int {
  value := seed
  for _, k := range c.Ranges {
    for _, r := range k {
      if value < r.SourceRange {
        continue
      }
      if value > r.SourceRange + r.Length {
        continue
      }

      value = value - r.SourceRange + r.DestRange
      break
    }
  }
  return value
}

func StringArrayToInt(input []string) []int {
  var ret []int
  for _, s := range input {
    value, _ := strconv.Atoi(s)
    ret = append(ret, value)
  }
  return ret
}


func part2(input string) int {
  return 1
}
