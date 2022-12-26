package main

import (
	"os"
	"testing"
)

type test struct {
	name     string
	given    string
	expected int
}

var input string

func init() {
	file, _ := os.ReadFile("example.txt")
	input = string(file)
}

func Test_Part1(t *testing.T) {
	tests := []test{
		{
			name:     "Part 1",
			given:    input,
			expected: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.given); got != tt.expected {
				t.Errorf("got %d expected %d", got, tt.expected)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
	tests := []test{
		{
			name:     "Part 2",
			given:    input,
			expected: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.given, func(t *testing.T) {
			if got := part2(tt.given); got != tt.expected {
				t.Errorf("got %d expected %d", got, tt.expected)
			}
		})
	}
}
