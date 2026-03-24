package main

import (
	"strings"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func TestAdd(t *testing.T) {
	result := Add(2 , 3)

	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

}

func TestAddTableDriven(t *testing.T)  {
	test := []struct {a , b , expected int}{
		{2, 3, 5},
		{1, 1, 2},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _ , tc := range test {
		result := Add(tc.a , tc.b)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, result)
		}
	}
}

func TestAddSubtest(t *testing.T)  {
	tests := []struct {a , b , expected int}{
		{2, 3, 5},
		{1, 1, 2},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _ , tc := range tests {
		t.Run("Testing Add function", func(t *testing.T) {
			result := Add(tc.a , tc.b)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})

		t.Run("Testing Subtract function", func(t *testing.T) {
			result := Subtract(tc.a , tc.b)
			expected := tc.a - tc.b
			if result != expected {
				t.Errorf("Expected %d, but got %d", expected, result)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(5, 3)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.WriteString("Hello")
		builder.WriteString(" ")
		builder.WriteString("World")
		_ = builder.String()
	}
}
