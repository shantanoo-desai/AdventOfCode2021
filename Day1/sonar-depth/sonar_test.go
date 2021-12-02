package main

import (
	"fmt"
	"testing"
)

func TestDepthReport(t *testing.T) {
	t.Run("Test for Sonar Test Report", func(t *testing.T) {
		got := DepthReport([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
		want := 7

		if got != want {
			t.Errorf("Expected: '%d, Got: '%d'", want, got)
		}
	})

	t.Run("Faulty Single Measurement", func(t *testing.T) {
		got := DepthReport([]int{199})
		want := 0

		if got != want {
			t.Errorf("Expected: '%d', Want: '%d", want, got)
		}
	})
}

func TestSlidingWindow(t *testing.T) {
	got := SlidingWindow([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	want := 5

	if got != want {
		t.Errorf("Expected: '%d', Got: '%d'", want, got)
	}
}

func BenchmarkDepthReport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DepthReport([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	}
}

func BenchmarkSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlidingWindow([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	}
}

func ExampleDepthReport() {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	result := DepthReport(data)
	fmt.Println(result)
	// Output: 7
}

func ExampleSlidingWindow() {
	data := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	result := SlidingWindow(data)
	fmt.Println(result)
	// Output: 5
}
