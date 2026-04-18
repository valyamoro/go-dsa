package main

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic case with positive numbers",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Indices at end of array",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "With negative numbers",
			nums:     []int{-1, -2, -3, 5, 10},
			target:   7,
			expected: []int{2, 4},
		},
		{
			name:     "Both numbers negative",
			nums:     []int{-10, 5, 2, -5},
			target:   -15,
			expected: []int{0, 3},
		},
		{
			name:     "With zero in array",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "Zero plus positive number",
			nums:     []int{0, 1, 2, 3},
			target:   3,
			expected: []int{1, 2},
		},
		{
			name:     "Same elements different indices",
			nums:     []int{5, 5, 5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "Multiple duplicates",
			nums:     []int{1, 1, 1, 2, 2, 3},
			target:   4,
			expected: []int{3, 4},
		},
		{
			name:     "Array of two elements",
			nums:     []int{1, 2},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "Two elements with zero",
			nums:     []int{0, 0},
			target:   0,
			expected: []int{0, 1},
		},
		{
			name:     "Large positive numbers",
			nums:     []int{1000000, 2000000, 3000000},
			target:   3000000,
			expected: []int{0, 1},
		},
		{
			name:     "Large range of values",
			nums:     []int{-1000000, 1, 1000000},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Solution at beginning of array",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "Solution at end of array",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "No solution sum too large",
			nums:     []int{1, 2, 3},
			target:   100,
			expected: nil,
		},
		{
			name:     "No solution sum too small",
			nums:     []int{1, 2, 3},
			target:   0,
			expected: nil,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
		{
			name:     "Array with single element",
			nums:     []int{5},
			target:   10,
			expected: nil,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-5, -2, 0, 3, 8, 10},
			target:   5,
			expected: []int{0, 5},
		},
		{
			name:     "Consecutive numbers",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   11,
			expected: []int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)

			if result == nil && tt.expected == nil {
				return
			}

			if (result == nil) != (tt.expected == nil) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
				return
			}

			if !slices.Equal(result, tt.expected) {
				if len(result) == 2 && len(tt.expected) == 2 {
					if result[0] == tt.expected[1] && result[1] == tt.expected[0] {
						t.Logf("twoSum(%v, %d) = %v (alternative order but OK)", tt.nums, tt.target, result)
						return
					}
				}
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	target := 19998

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
