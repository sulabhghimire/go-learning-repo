package main

import "testing"

func Add(i, j int) int {
	return i + j
}

func BenchmarkAddSmallInput(b *testing.B) {
	for b.Loop() {
		Add(2, 3)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	for b.Loop() {
		Add(200, 300)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	for b.Loop() {
		Add(2000, 3000)
	}
}
