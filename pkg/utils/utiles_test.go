package utils

import (
	"testing"
)

func TestHexToInt64(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"0000000000000000", 0},
		{"0000000000000001", 1},
		{"00000000000000ff", 255},
		{"0000000000000100", 256},
		{"ffffffffffffffff", -1},
		{"8000000000000000", -9223372036854775808},
		{"7fffffffffffffff", 9223372036854775807},
	}

	for _, tc := range testCases {
		result := HexToInt64(tc.input)
		if result != tc.expected {
			t.Errorf("HexToInt64(%s) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func TestHexToInt64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("HexToInt64 did not panic on invalid input")
		}
	}()

	HexToInt64("invalid hex string")
}
