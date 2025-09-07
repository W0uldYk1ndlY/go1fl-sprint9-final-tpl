package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("Size must be greater than 0")
		return nil
	}

	numbers := make([]int, size)

	for i := range numbers {
		numbers[i] = r.Int()
	}

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Println("empty data")
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	max := data[0]
	for i := range data {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		fmt.Println("empty data")
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	var wg sync.WaitGroup
	wg.Add(CHUNKS)
	fin := make([]int, CHUNKS)

	slice_size := len(data) / CHUNKS
	remainder := len(data) % CHUNKS
	first_index := 0
	for i := range CHUNKS {
		var chunk_size int

		if i < remainder {
			chunk_size = slice_size + 1
		} else {
			chunk_size = slice_size
		}

		last_index := first_index + chunk_size
		chunk := data[first_index:last_index]
		first_index = last_index

		go func(chunk []int, index int) {
			res := maximum(chunk)
			fin[index] = res
			wg.Done()
		}(chunk, i)
	}
	wg.Wait()
	return maximum(fin)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
