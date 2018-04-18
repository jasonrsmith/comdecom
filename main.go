package main

import (
	"bytes"
	"strconv"
)

func main() {
}

func IsNumber(b rune) bool {
	return b >= 48 && b <= 57
}

func ScanNumber(s string) (int, int) {
	var (
		i      int
		length int
		v      rune
	)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
	}
	for i, v = range s {
		if !IsNumber(rune(v)) {
			break
		}
		length++
	}

	number, _ := strconv.Atoi(s[:i])
	return number, length
}

func Decompress(compressed string) string {
	result, _ := decompressHelp(compressed)
	return result
}

func decompressHelp(compressed string) (string, int) {
	var numberToRepeat int
	var buf bytes.Buffer

	i := 0

	compressedRuneList := []rune(compressed)
	inBrackets := false
	for i < len(compressedRuneList) {
		v := compressedRuneList[i]
		if IsNumber(v) {
			var length int
			numberToRepeat, length = ScanNumber(compressed[i:])
			i += length
		} else if v == '[' {
			i++
			recursiveResult, readLen := decompressHelp(string(compressedRuneList[i:]))
			for j := 0; j < numberToRepeat; j++ {
				buf.WriteString(recursiveResult)
			}
			i += readLen
			inBrackets = true
		} else if v == ']' {
			if inBrackets == false {
				break
			}
			inBrackets = false
			i++
		} else {
			buf.WriteRune(v)
			i++
		}
	}

	return buf.String(), i
}
