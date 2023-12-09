package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
  "unicode"
)

func main() {
  file, _ := os.ReadFile("input1.txt")
  input1 := string(file)
  fmt.Println("Part1:", part1(input1))

  file, _ = os.ReadFile("input2.txt")
  input2 := string(file)
  fmt.Println("Part2:", part2(input2))
}

func part1(input string) int {
  sum := 0
  s := ""
  colides := false
  matrix := generateMatrix(input)

  for i, line := range matrix {
    for j := range line {
      item := matrix[i][j]

      if unicode.IsDigit(item) {
        s = s + string(item)
        // check each char of a number
        if !colides {
          colides = checkColision(i, j, matrix)
        }
      } else { // If not a digit, cast the number as int
        if colides {
          num, _ := strconv.Atoi(s)
          sum += num
        }
        s = ""
        colides = false
      }
    }
  }
  return sum
}

type Position struct {
  X int
  Y int
}
var positions = []Position {
  {X: -1, Y: -1},
  {X: -1, Y: 0},
  {X: -1, Y: 1},
  {X: 0, Y: -1},
  {X: 0, Y: 1},
  {X: 1, Y: -1},
  {X: 1, Y: 0},
  {X: 1, Y: 1},
}

func checkColision(x int, y int, matrix [][]rune) bool {
  for _, pos := range positions {
    dx := x + pos.X
    dy := y + pos.Y
    if dx < 0 || dy < 0 || dx >= len(matrix) || dy >= len(matrix) {
      continue 
    }
    if !unicode.IsDigit(matrix[dx][dy]) && matrix[dx][dy] != '.' {
      return true
    }
  }
  return false
}

func generateMatrix(input string) [][]rune {
  input = strings.TrimSpace(input)
  lines := strings.Split(input, "\n")

  matrix := make([][]rune, len(lines)) 
  for m := range matrix {
    matrix[m] = make([]rune, len(lines))
  }

  for x, line := range lines {
    for y, column := range line {
      matrix[x][y] = column
    }
  }

  return matrix
}

func part2(input string) int {
  sum := 0

  lineSize := len(strings.Split(input, "\n")) - 1
  lines := strings.TrimSpace(input)

  for i, item := range lines {
    if item == '*' {
      p := calculatePositions(i, lineSize)
      nums := calculateColisions(lines, p)
      if len(nums) == 2 {
        sum = sum + (nums[0] * nums[1])
      }
    }

  }
  return sum
}

func calculatePositions(i int, lineSize int) []int {
  var pos = []int {
    i - lineSize - 1 - 1,
    i - lineSize - 1,
    i - lineSize + 1 - 1,
    i - 1,
    i + 1,
    i + lineSize - 1 + 1,
    i + lineSize + 1,
    i + lineSize + 1 + 1,
  }
  return pos
}

func calculateColisions(input string, pos []int) []int {
  var nums []int
  lastAdded := 0

  for _, p := range pos {
    if !unicode.IsDigit(rune(input[p])) {
      continue
    }
  
    n := createNum(input, p)
    if n == lastAdded {
      continue
    }
   
    lastAdded = n
    nums = append(nums, n)
  }
  return nums
}

func createNum(input string, p int) int {
  s := string(input[p])
  i := 0

  for i = p - 1; i >= 0; i-- {
    if  !unicode.IsDigit(rune(input[i])) {
      break
    }
    s = string(input[i]) + s
  }
  
  for i = p + 1; i < len(input); i++ {
    if !unicode.IsDigit(rune(input[i])) {
      break
    }
    s = s + string(input[i]) 
  }

  num, _ := strconv.Atoi(s)
  return num
}
