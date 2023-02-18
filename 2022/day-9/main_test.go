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

func Test_Part1(t *testing.T) {
	file, _ := os.ReadFile("example1.txt")
	input := string(file)
	tests := []test{
		{
			name:     "Part 1",
			given:    input,
			expected: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.given, func(t *testing.T) {
			if got := part1(tt.given); got != tt.expected {
				t.Errorf("got %d expected %d", got, tt.expected)
			}
		})
	}
}

func Test_Part2(t *testing.T) {
	file, _ := os.ReadFile("example2.txt")
	input := string(file)
	tests := []test{
		{
			name:     "Part 2",
			given:    input,
			expected: 36,
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
