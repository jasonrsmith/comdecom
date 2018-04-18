package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestIsNumber(t *testing.T) {
	var tests = []struct {
		el             rune
		expectedResult bool
	}{
		{'0', true},
		{'1', true},
		{'a', false},
		{'9', true},
	}
	for _, tt := range tests {
		actualResult := IsNumber(tt.el)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}

func TestScanNumber(t *testing.T) {
	var tests = []struct {
		s              string
		expectedInt    int
		expectedLength int
	}{
		{"a", 0, 0},
		{"19", 1, 2},
		{"a19", 0, 0},
		{"0", 0, 1},
		{"19a", 19, 2},
	}
	for _, tt := range tests {
		actualInt, actualLength := ScanNumber(tt.s)
		assert.Equal(t, tt.expectedInt, actualInt)
		assert.Equal(t, tt.expectedLength, actualLength)
	}
}

func TestSimpleStringWithoutCompressionReturnsItself(t *testing.T) {
	compressed := "abcde"
	result := Decompress(compressed)
	assert.Equal(t, compressed, result)
}

func TestSingleDigitIntCompression(t *testing.T) {
	compressed := "3[abc]"
	expectedResult := "abcabcabc"
	result := Decompress(compressed)
	assert.Equal(t, expectedResult, result)
}

func TestMoreComplex(t *testing.T) {
	var tests = []struct {
		compressed string
		expected   string
	}{
		{"ab3[abc5[xyz]]", "ababcxyzxyzxyzxyzxyzabcxyzxyzxyzxyzxyzabcxyzxyzxyzxyzxyz"},
		{"19", ""},
		{"a19", "a"},
		{"19a", "a"},
		{"2[3[a]b]", "aaabaaab"},
		{"2[a]b", "aab"},
		{"4[a]", "aaaa"},
		{"4[🌞]", "🌞🌞🌞🌞"},
		{"⁣    ☁☁☁☁☁　 　    ☁　  🌞　  ☁　   ☁  　 　 　   ☁ ☁🌴🌴🌴🌴🌴🌴🌴☁ ⁣ ☁　 🐬　 🍹　 ☁   ☁🌊🌊🌊🌊🌊☁　      ☁☁☁☁☁　 　      💭　 　 　    🙎", "⁣    ☁☁☁☁☁　 　    ☁　  🌞　  ☁　   ☁  　 　 　   ☁ ☁🌴🌴🌴🌴🌴🌴🌴☁ ⁣ ☁　 🐬　 🍹　 ☁   ☁🌊🌊🌊🌊🌊☁　      ☁☁☁☁☁　 　      💭　 　 　    🙎"},
		{"3[💩]", "💩💩💩"},
		// TODO why won't this work?
		//{"3[💩4[a]]", ""},
	}
	for _, tt := range tests {
		actual := Decompress(tt.compressed)
		assert.Equal(t, tt.expected, actual)
	}
}
