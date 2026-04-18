package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Simple palindrome",
			input:    121,
			expected: true,
		},
		{
			name:     "Large palindrome",
			input:    12321,
			expected: true,
		},
		{
			name:     "Very large palindrome",
			input:    123454321,
			expected: true,
		},
		{
			name:     "Single digit",
			input:    1,
			expected: true,
		},
		{
			name:     "Zero",
			input:    0,
			expected: true,
		},
		{
			name:     "Nine",
			input:    9,
			expected: true,
		},
		{
			name:     "Negative number",
			input:    -121,
			expected: false,
		},
		{
			name:     "Negative palindrome",
			input:    -12321,
			expected: false,
		},
		{
			name:     "Minus one",
			input:    -1,
			expected: false,
		},
		{
			name:     "Not palindrome",
			input:    123,
			expected: false,
		},
		{
			name:     "Ten",
			input:    10,
			expected: false,
		},
		{
			name:     "One hundred",
			input:    100,
			expected: false,
		},
		{
			name:     "Two hundred thirty one",
			input:    231,
			expected: false,
		},
		{
			name:     "Two digits palindrome",
			input:    11,
			expected: true,
		},
		{
			name:     "Two digits not palindrome",
			input:    12,
			expected: false,
		},
		{
			name:     "Three digits palindrome",
			input:    121,
			expected: true,
		},
		{
			name:     "Three digits not palindrome",
			input:    123,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	testCases := []int{121, 12321, 123454321, 123456789, 0, -121}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, num := range testCases {
			isPalindrome(num)
		}
	}
}
