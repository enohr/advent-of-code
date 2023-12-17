package main

import (
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

type Hand struct {
  Cards string
  Bid int
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
  return playCamelCards(input, false)
}

func part2(input string) int {
  return playCamelCards(input, true)
}

func playCamelCards(input string, replaceJoker bool) int {
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
    h.CalculateRank(replaceJoker)
    h.SetType(replaceJoker)
    matches[h.Type] = append(matches[h.Type], h)
  }

  // Order the hands based on type and rank
  for _, m := range matches {
    sort.Slice(m, func(i, j int) bool {
      return m[i].Rank < m[j].Rank
    })
    orderedMatches = append(orderedMatches, m...)
  }

  sum := 0
  for i, om := range orderedMatches {
    sum += (i + 1) * om.Bid
  }
  return sum
}

func (h *Hand) SetType(replaceJoker bool) {
  m := make(map[rune]int, 5)

  if replaceJoker {
    h.ReplaceJoker()
  }

  // Calculate the card frequency
  for _, card := range h.Cards {
    m[card]++
  }

  switch len(m) {
  case 1:
    h.Type = FIVEKIND
  case 2:
    // If first or second item count is 4, then the other one will be 1.
    // So, its a FOURKIND. Otherwhise, it will be a FULLHOUSE
    first, second := rune(h.Cards[0]), rune(h.Cards[1])
    if m[first] == 4 || m[second] == 4 {
      h.Type = FOURKIND 
    } else {
      h.Type = FULLHOUSE
    }
  case 3:
    // If one of the item count is 3, then the other two will be 1.
    // So its a THREEKIND. Otherwhise, it will be a TWOPAIR.
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

// Calculate the rank by transforming the hand in
// hexadecimal number
func (h *Hand) CalculateRank(replaceJoker bool) {
  cardsHex := h.Cards

  // Transform cards in hex
  if replaceJoker {
    cardsHex = strings.ReplaceAll(cardsHex, "J", "1")
  }
  // A should goes first to not be replaced again
  cardsHex = strings.ReplaceAll(cardsHex, "A", "E")
  cardsHex = strings.ReplaceAll(cardsHex, "T", "A")
  cardsHex = strings.ReplaceAll(cardsHex, "J", "B")
  cardsHex = strings.ReplaceAll(cardsHex, "Q", "C")
  cardsHex = strings.ReplaceAll(cardsHex, "K", "D")

  h.Cards = cardsHex
  h.Rank, _ = strconv.ParseInt(string(cardsHex), 16, 0)
}

// This function will replace Joker (char 1)
// by the greatest and most frequent Card.
// e.g:
// A2221 -> A2222
// 22AA1 -> 22AAA
func (h *Hand) ReplaceJoker() {
  m := make(map[rune]int, 5)
  for _, card := range h.Cards {
    m[card]++
  }

  count := 0
  value := rune(0)
  for k, v := range m {
    if k == '1' {
      continue
    }
    if v > count {
      count = v
      value = k
      continue
    }
    if v == count && k > value {
      count = v
      value = k
    }
  }
  h.Cards = strings.ReplaceAll(h.Cards, "1", string(value))
}

