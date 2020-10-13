package algos

import (
	"strconv"
	"strings"
)

// UniqueChars - Unique character detection based on ASCII or Extended ASCII.
func UniqueChars(input string) bool {
	chars := make([]bool, 256, 256)

	for _, character := range input {
		charvalue := int(character)

		if chars[charvalue] {
			return false
		}
		chars[charvalue] = true

	}
	return true
}

// UniqueCharsHashset - Uses a hashset to determine if string is made up of unique characters. Really works better for Unicode.
func UniqueCharsHashset(input string) bool {
	chars := make(map[rune]bool)

	// Luckily golang's range over string implicitly does unicode decoding so we don't have to worry about interpreting bytes
	for _, character := range input {
		if chars[character] {
			return false
		}

		chars[character] = true
	}
	return true
}

// UniqueCharsBrute - Brute force implementation of checking for unique characters in a string
func UniqueCharsBrute(input string) bool {
	for i, leftchar := range input[:len(input)-1] {
		for _, rightchar := range input[i+1:] {
			if leftchar == rightchar {
				return false
			}
		}
	}
	return true
}

// IsPermutation - Algorithm to determine if a string is a permutation of a second string
func IsPermutation(a string, b string) bool {
	// permutations have the multiset of characters
	// The must be the same length, and all the characters in them must be the same
	if len(a) != len(b) {
		return false
	}

	var characters [256]int

	for _, character := range a {
		characters[int(character)]++
	}
	for _, character := range b {
		characters[int(character)]--
	}

	for _, count := range characters {
		if count != 0 {
			return false
		}
	}

	return true
}

// ReplaceSpacesBuffer - A serializer of spaces to %20, implemented using string buffer
func ReplaceSpacesBuffer(input string, truelength int) string {
	var builder strings.Builder

	for _, character := range input {
		if truelength == 0 {
			break
		}
		if character == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteRune(character)
		}
		truelength--
	}

	return builder.String()
}

// ReplaceSpacesInPlace - A serializer of spaces to %20, implemented using in place array.
func ReplaceSpacesInPlace(input []rune, truelength int) string {
	// For each rune. We are working with rune array so each one is already a unicode codepoint
	// Check if it's a space
	// If it's a space, shift everything over 2x
	// replace current and next 2 with %20

	if len(input) == 0 {
		return ""
	}

	for i := 0; i < truelength; i++ {
		if truelength == len(input) {
			break
		}

		if input[i] == ' ' {
			shiftRight(&input, i, truelength-1, 2)

			input[i] = '%'

			i++
			input[i] = '2'
			truelength++

			i++
			input[i] = '0'
			truelength++
		}
	}

	return string(input)
}

// Shift everything to the right of start index by n. You can assume the array is the correct size
func shiftRight(array *[]rune, startIndex int, endIndex int, n int) {
	for i := endIndex; i > startIndex; i-- {
		(*array)[i+n] = (*array)[i]
	}
}

// StringCompression - Simple compression for strings by counting duplicate chars together. i.e. aaabbb becomes a3b3
func StringCompression(input string) string {
	if len(input) == 0 {
		return input
	}

	characters := []rune(input)

	currChar := characters[0]
	currCharCount := 0
	isCompressed := false

	var buffer strings.Builder

	for _, char := range characters {
		if char != currChar {
			buffer.WriteString(string(currChar))
			buffer.WriteString(strconv.Itoa(currCharCount))

			currChar = char
			currCharCount = 1
		} else {
			currCharCount++

			if currCharCount > 1 {
				isCompressed = true
			}
		}
	}

	// Write the last character
	buffer.WriteString(string(currChar))
	buffer.WriteString(strconv.Itoa(currCharCount))

	if isCompressed {
		return buffer.String()
	}

	return input
}

// ZeroMatrixRowColumn - When the value in a matrix is 0, set column and row to zero
func ZeroMatrixRowColumn(matrix [][]int) [][]int {
	// My assumption is that you are masking
	// Row: i, Column j

	if len(matrix) == 0 {
		return matrix
	}

	if len(matrix[0]) == 0 {
		return matrix
	}

	rows := make([]bool, len(matrix))
	columns := make([]bool, len(matrix[0]))

	for i := 0; i < len(matrix); i++ {
		if rows[i] {
			continue
		}

		for j := 0; j < len(matrix[i]); j++ {
			if columns[j] {
				continue
			}

			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for row, set := range rows {
		for j := 0; j < len(matrix[row]); j++ {
			if set {
				matrix[row][j] = 0
			}
		}
	}

	for column, set := range rows {
		for i := 0; i < len(matrix[column]); i++ {
			if set {
				matrix[i][column] = 0
			}
		}
	}

	return matrix
}
