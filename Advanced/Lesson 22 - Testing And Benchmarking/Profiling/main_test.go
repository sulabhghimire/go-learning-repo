package main

import (
	"math/rand"
	"testing"
)

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {

	for b.Loop() {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {

	slice := GenerateRandomSlice(1000)
	b.ResetTimer()
	for b.Loop() {
		SumSlice(slice)
	}
}
