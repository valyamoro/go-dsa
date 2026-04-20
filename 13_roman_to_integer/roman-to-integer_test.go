package main

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "III equals 3",
			input:    "III",
			expected: 3,
		},
		{
			name:     "LVIII equals 58",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "MCMXCIV equals 1994",
			input:    "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "IV equals 4",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "IX equals 9",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "XL equals 40",
			input:    "XL",
			expected: 40,
		},
		{
			name:     "XC equals 90",
			input:    "XC",
			expected: 90,
		},
		{
			name:     "CD equals 400",
			input:    "CD",
			expected: 400,
		},
		{
			name:     "CM equals 900",
			input:    "CM",
			expected: 900,
		},
		{
			name:     "I equals 1",
			input:    "I",
			expected: 1,
		},
		{
			name:     "V equals 5",
			input:    "V",
			expected: 5,
		},
		{
			name:     "X equals 10",
			input:    "X",
			expected: 10,
		},
		{
			name:     "L equals 50",
			input:    "L",
			expected: 50,
		},
		{
			name:     "C equals 100",
			input:    "C",
			expected: 100,
		},
		{
			name:     "D equals 500",
			input:    "D",
			expected: 500,
		},
		{
			name:     "M equals 1000",
			input:    "M",
			expected: 1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := romanToInt(tt.input)
			if result != tt.expected {
				t.Errorf("romanToInt(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
