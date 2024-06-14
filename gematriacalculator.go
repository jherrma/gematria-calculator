package gematriacalculator

import "unicode"

type Method int

const (
	Ordinal Method = iota
	OrdinalReverse
	Reduction
	ReductionReverse
	Sumerian
	SumerianReverse
	Jewish
)

func ComputeGematria(text string, method Method) int {
	sum := 0
	switch method {
	case Ordinal:
		for _, letter := range text {
			value := resolveLetterOrdinal(letter)
			sum += value
		}
	case OrdinalReverse:
		for _, letter := range text {
			value := resolveLetterOrdinalReverse(letter)
			sum += value
		}
	case Reduction:
		for _, letter := range text {
			value := resolveLetterReduction(letter)
			sum += value
		}
	case ReductionReverse:
		for _, letter := range text {
			value := resolveLetterReductionReverse(letter)
			sum += value
		}
	case Sumerian:
		for _, letter := range text {
			value := resolveLetterSumerian(letter)
			sum += value
		}
	case SumerianReverse:
		for _, letter := range text {
			value := resolveLetterSumerianReverse(letter)
			sum += value
		}
	case Jewish:
		for _, letter := range text {
			value := resolveLetterJewish(letter)
			sum += value
		}
	}

	return sum
}

func resolveLetterOrdinal(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	value := int(lowerChar - 'a' + 1)
	return value
}

func resolveLetterOrdinalReverse(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	value := 26 - (int(lowerChar - 'a'))
	return value
}

func resolveLetterReduction(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	mapping := "12345678912345678912345678"
	value := int(mapping[int(lowerChar-'a')] - '0')
	return value
}

func resolveLetterReductionReverse(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	reverseMapping := "87654321987654321987654321"
	value := int(reverseMapping[int(lowerChar-'a')] - '0')
	return value
}

func resolveLetterSumerian(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	value := int(lowerChar - 'a' + 1)
	return value * 6
}

func resolveLetterSumerianReverse(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}
	value := 26 - (int(lowerChar - 'a'))
	return value * 6
}

func resolveLetterJewish(char rune) int {
	lowerChar := unicode.ToLower(char)
	if int(lowerChar) < int('a') || int(lowerChar) > int('z') {
		return 0
	}

	charMap := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'k': 10,
		'l': 20,
		'm': 30,
		'n': 40,
		'o': 50,
		'p': 60,
		'q': 70,
		'r': 80,
		's': 90,
		't': 100,
		'u': 200,
		'x': 300,
		'y': 400,
		'z': 500,
		'j': 600,
		'v': 700,
		'w': 900,
	}

	value := charMap[char]
	return value
}
