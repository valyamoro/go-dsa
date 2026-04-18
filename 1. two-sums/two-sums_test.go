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
			name:     "Базовый случай с положительными числами",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Индексы в конце массива",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "С отрицательными числами",
			nums:     []int{-1, -2, -3, 5, 10},
			target:   7,
			expected: []int{2, 4},
		},
		{
			name:     "Оба числа отрицательные",
			nums:     []int{-10, 5, 2, -5},
			target:   -15,
			expected: []int{0, 3},
		},
		{
			name:     "С нулём в массиве",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "Ноль + положительное число",
			nums:     []int{0, 1, 2, 3},
			target:   3,
			expected: []int{1, 2},
		},
		{
			name:     "Одинаковые элементы (разные индексы)",
			nums:     []int{5, 5, 5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "Несколько дубликатов",
			nums:     []int{1, 1, 1, 2, 2, 3},
			target:   4,
			expected: []int{3, 4},
		},
		{
			name:     "Массив из двух элементов",
			nums:     []int{1, 2},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "Два элемента с нулём",
			nums:     []int{0, 0},
			target:   0,
			expected: []int{0, 1},
		},
		{
			name:     "Большие положительные числа",
			nums:     []int{1000000, 2000000, 3000000},
			target:   3000000,
			expected: []int{0, 1},
		},
		{
			name:     "Большой диапазон значений",
			nums:     []int{-1000000, 1, 1000000},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Решение в начале массива",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: []int{0, 1},
		},
		{
			name:     "Решение в конце массива",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4},
		},
		{
			name:     "Нет решения - сумма слишком большая",
			nums:     []int{1, 2, 3},
			target:   100,
			expected: nil,
		},
		{
			name:     "Нет решения - сумма слишком маленькая",
			nums:     []int{1, 2, 3},
			target:   0,
			expected: nil,
		},
		{
			name:     "Пустой массив",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
		{
			name:     "Массив с одним элементом",
			nums:     []int{5},
			target:   10,
			expected: nil,
		},
		{
			name:     "Смешанные положительные и отрицательные",
			nums:     []int{-5, -2, 0, 3, 8, 10},
			target:   5,
			expected: []int{0, 5},
		},
		{
			name:     "Последовательные числа",
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
						t.Logf("twoSum(%v, %d) = %v (альтернативный порядок, но OK)", tt.nums, tt.target, result)
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
