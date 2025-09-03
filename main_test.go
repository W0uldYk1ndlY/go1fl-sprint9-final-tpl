package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле

func TestGenerateRandomElement(t *testing.T) {
	//negative size
	res := generateRandomElements(-1)
	assert.Nil(t, res)

	//zero size
	res = generateRandomElements(0)
	assert.Nil(t, res)

	//positive size
	res = generateRandomElements(100000)
	assert.Len(t, res, 100000)
}

func TestMaximum(t *testing.T) {
	//empty slice
	data := make([]int, 0)
	res := maximum(data)
	assert.Zero(t, res)

	//nil
	res = maximum(nil)
	assert.Zero(t, res)

	//1 element
	data = make([]int, 1)
	data[0] = 128
	res = maximum(data)
	assert.Equal(t, res, 128)

	//5 elements
	data = []int{0, -5, 28, 9, 130}
	res = maximum(data)
	assert.Equal(t, res, 130)

}

func TestMaxChunks(t *testing.T) {
	//empty slice
	data := make([]int, 0)
	res := maxChunks(data)
	assert.Zero(t, res)

	//nil
	res = maxChunks(nil)
	assert.Zero(t, res)

	//1 element
	data = make([]int, 1)
	data[0] = 128
	res = maxChunks(data)
	assert.Equal(t, res, 128)

	//8 elements
	data = []int{20, 19, 19, 11, 0, 17, 89, 12}
	res = maxChunks(data)
	assert.Equal(t, res, 89)

	//15 elements
	data = []int{20, 19, 19, 11, 0, 17, 89, 12, 150, 77, 89, 60, 54, 65, 63}
	res = maxChunks(data)
	assert.Equal(t, res, 150)
}
