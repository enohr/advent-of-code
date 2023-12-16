package main

import (
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

type Hand struct {
  Bid int
  Cards string
  Rank int64
  Type int
}

const (
  HIGH_CARD = 1 + iota
  ONE_PAIR
  TWOPAIR
  THREEKIND
  FULLHOUSE
  FOURKIND
  FIVEKIND
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
  var hands []*Hand
  matches := make([][]*Hand, 8)
  var orderedMatches []*Hand
  parts := strings.Split(input, "\n")

  for _, p := range parts {
    fields := strings.Fields(p)
    cards := fields[0]
    bid, _ := strconv.Atoi(fields[1])
    hands = append(hands, &Hand{
      Bid : bid,
      Cards: cards,
      Rank: 0,
    })
  }

  for _, h := range hands {
    h.SetType()
    h.CalculateRank()
    matches[h.Type] = append(matches[h.Type], h)
  }

  for _, m := range matches {
    sort.Slice(m, func(i, j int) bool {
      return m[i].Rank < m[j].Rank
    })
    orderedMatches = append(orderedMatches, m...)
  }

  // Something wrong. Answer too high

  sum := 0
  for i, om := range orderedMatches {
    sum += (i + 1) * om.Bid
  }
  return sum
}

func (h *Hand) SetType() {
  m := make(map[rune]int, 5)

  for _, card := range h.Cards {
    m[card]++
  }


  switch len(m) {
  case 1:
    h.Type = FIVEKIND
  case 2:
    first, second := rune(h.Cards[0]), rune(h.Cards[1])
    if m[first] == 4 || m[second] == 4 {
      h.Type = FOURKIND 
    } else {
      h.Type = FULLHOUSE
    }
  case 3:
    first, second, third := rune(h.Cards[0]), rune(h.Cards[1]), rune(h.Cards[2])
    if m[first] == 3 || m[second] == 3 || m[third] == 3 {
      h.Type = THREEKIND
    } else {
      h.Type = TWOPAIR
    }
  case 4:
    h.Type = ONE_PAIR
  case 5:
    h.Type = HIGH_CARD
  }
}

func (h *Hand) CalculateRank() {
  cardsHex := h.Cards
  
  // Transform cards in hex
  // A should goes first to not be replaced again
  cardsHex = strings.ReplaceAll(cardsHex, "A", "E")
  cardsHex = strings.ReplaceAll(cardsHex, "T", "A")
  cardsHex = strings.ReplaceAll(cardsHex, "J", "B")
  cardsHex = strings.ReplaceAll(cardsHex, "Q", "C")
  cardsHex = strings.ReplaceAll(cardsHex, "K", "D")

  h.Rank, _ = strconv.ParseInt(string(cardsHex), 16, 0)
}

func part2(input string) int {
  return 1
}
