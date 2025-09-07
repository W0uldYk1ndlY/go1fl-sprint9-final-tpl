package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElement(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"negative size", -1, 0},
		{"zero size", 0, 0},
		{"positive size", 100_000, 100_000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := generateRandomElements(test.input)
			assert.Len(t, res, test.want)
		})
	}

}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"empty", []int{}, 0},
		{"nil", nil, 0},
		{"1 element", []int{128}, 128},
		{"5 elements", []int{0, -5, 28, 9, 130}, 130},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := maximum(test.input)
			assert.Equal(t, test.want, res)
		})

	}

}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"empty slice", make([]int, 0), 0},
		{"nil", nil, 0},
		{"1 element", []int{128}, 128},
		{"8 elements", []int{20, 19, 19, 11, 0, 17, 89, 12}, 89},
		{"15 elements", []int{20, 15, 13, 11, 0, 17, 89, 12, 150, 77, 89, 60, 54, 65, 163}, 163},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := maxChunks(test.input)
			assert.Equal(t, test.want, res)
		})
	}

}
