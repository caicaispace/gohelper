package utils

import (
	"fmt"
	"testing"
)

func TestNumberLen(t *testing.T) {
	num := 9999999999999
	fmt.Println(NumberLen(num))
	fmt.Println(NumberLen(int64(num)))
	fmt.Println(NumberLen(int32(num)))
	fmt.Println(NumberLen(int16(num)))
	fmt.Println(NumberLen(int8(num)))
	fmt.Println(NumberLen(int64(num)))
	fmt.Println(NumberLen(uint(num)))
	fmt.Println(NumberLen(uint64(num)))
	fmt.Println(NumberLen(uint32(num)))
	fmt.Println(NumberLen(uint16(num)))
	fmt.Println(NumberLen(uint8(num)))
}

func TestRandInt(t *testing.T) {
	fmt.Println("RandInt(0, 100):", RandInt(0, 1))
}

func Test_UnicodeDecode(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Hello, 世界", "Hello, 世界"},
		{"こんにちは", "こんにちは"},
		{"Привет, мир", "Привет, мир"},
		{"你好，世界！", "你好，世界！"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := UnicodeDecode(testCase.input)
			if result != testCase.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", testCase.input, testCase.expected, result)
			}
		})
	}
}

func Test_UnicodeDecodeV2(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Hello, 世界", "Hello, 世界"},
		{"こんにちは", "こんにちは"},
		{"Привет, мир", "Привет, мир"},
		{"你好，世界！", "你好，世界！"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := UnicodeDecodeV2(testCase.input)
			if result != testCase.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", testCase.input, testCase.expected, result)
			}
		})
	}
}
