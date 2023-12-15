package main

import (
  "fmt"
  "os"
  "slices"
  "sort"
  "strconv"
  "strings"
)

type SeedRange struct {
  Start int
  End int
}

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

func part2(input string) int {
  var seeds []SeedRange
  parts := strings.Split(input, "\n\n")
  list := StringArrayToInt(strings.Fields(parts[0])[1:])
  for i := 0; i < len(list); i+=2 {
    seeds = append(seeds, SeedRange{Start: list[i], End: list[i] + list[i+1]}) 
  }

  c := NewConverter(parts[1:])
  return c.FromRange(seeds)
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

func (c *Converter) FromRange(seeds []SeedRange) int {
  for _, k := range c.Ranges {
    var next []SeedRange
    index := 0
    // iterate over all seeds. In case we add more seeds on loop,
    // it will iterate until has seeds to calculate
    for {
      if index >= len(seeds) {
        break
      }
      sr := seeds[index]
      index++
      for i, r := range k {
        // Get where our seed range fits on the actual Range
        start := Max(sr.Start, r.SourceRange)
        end := Min(sr.End, r.SourceRange + r.Length)

        // At least one part of seed range inside of actual Range
        if start < end {
          // It means we have a range starting before SourceRange starts
          // Then, create a new range starting on previous start and ending on
          // calculated range start
          if start > sr.Start {
            srNew := SeedRange{Start: sr.Start, End: start}
            seeds = append(seeds, srNew)
          }

          // It means we have a range after SourceRange ends
          // Then, create a new range starting on calculated range ends
          // and ending on previous end
          if sr.End > end {
            srNew := SeedRange{Start: end, End: sr.End} 
            seeds = append(seeds, srNew)
          }

          // Calculate the new value of starting and ending range
          sr.Start = start - r.SourceRange + r.DestRange
          sr.End = end - r.SourceRange + r.DestRange
          next = append(next, sr)
          break
        }
        // If not match founded, add the current value again
        if i+1 == len(k) {
          next = append(next, sr)
        }
      }
    }
    seeds = next
  }
  sort.Slice(seeds, func(i, j int) bool {
    return seeds[i].Start < seeds[j].Start
  })
  return seeds[0].Start
}

func StringArrayToInt(input []string) []int {
  var ret []int
  for _, s := range input {
    value, _ := strconv.Atoi(s)
    ret = append(ret, value)
  }
  return ret
}

func Max(a, b int) int {
  if a > b { return a }
  return b
}


func Min(a, b int) int {
  if a < b { return a }
  return b
}

