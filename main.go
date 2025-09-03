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
		numbers[i] = r.Intn(SIZE)
	}

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 || data == nil {
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
	if len(data) == 0 || data == nil {
		fmt.Println("empty data")
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	var wg sync.WaitGroup
	wg.Add(8)
	fin := make([]int, CHUNKS+7)
	remainder := len(data) % CHUNKS
	if remainder != 0 {
		remainder_slice := data[(len(data) - remainder):]
		fin = append(fin, remainder_slice...)
	}
	slice_size := len(data) / CHUNKS
	for i := range CHUNKS {
		go func(data []int) {
			first_index := i * slice_size
			res := maximum(data[first_index:(first_index + slice_size)])
			fin = append(fin, res)
			wg.Done()
		}(data)
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
