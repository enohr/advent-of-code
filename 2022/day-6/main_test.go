package main

import (
	"testing"
)

type test struct {
	given    string
	expected int
}

func Test_Part1(t *testing.T) {
	tests := []test{
		{
			given:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 5,
		},
		{
			given:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 6,
		},
		{
			given:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 10,
		},
		{
			given:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 11,
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
	tests := []test{
		{
			given:    "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			expected: 19,
		},
		{
			given:    "bvwbjplbgvbhsrlpgdmjqwftvncz",
			expected: 23,
		},
		{
			given:    "nppdvjthqldpwncqszvftbrmjlhg",
			expected: 23,
		},
		{
			given:    "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			expected: 29,
		},
		{
			given:    "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			expected: 26,
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
