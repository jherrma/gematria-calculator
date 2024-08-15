package gematriacalculator

import (
	"unicode"
)

type Cipher int

const (
	Ordinal Cipher = iota
	OrdinalReverse
	Reduction
	ReductionReverse
	Sumerian
	SumerianReverse
	HebrewLatinCharacters
	Hebrew
)

func ComputeGematria(text string, method Cipher) int {
	sum := 0

	for _, letter := range text {
		lowerLetter := unicode.ToLower(letter)

		if int(lowerLetter) >= '0' && int(lowerLetter) <= '9' {
			sum += resolveNumber(lowerLetter)
		}

		switch method {
		case Ordinal:
			value := resolveLetterOrdinal(lowerLetter)
			sum += value
		case OrdinalReverse:
			value := resolveLetterOrdinalReverse(lowerLetter)
			sum += value
		case Reduction:
			value := resolveLetterReduction(lowerLetter)
			sum += value

		case ReductionReverse:
			value := resolveLetterReductionReverse(lowerLetter)
			sum += value
		case Sumerian:
			value := resolveLetterSumerian(lowerLetter)
			sum += value
		case SumerianReverse:
			value := resolveLetterSumerianReverse(lowerLetter)
			sum += value
		case HebrewLatinCharacters:
			value := resolveLetterHebrewLatinCharacters(lowerLetter)
			sum += value
		case Hebrew:
			value := resolveLetterHebrew(lowerLetter)
			sum += value
		}
	}

	return sum
}

func resolveNumber(char rune) int {
	switch char {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		return 0
	}
}

func resolveLetterOrdinal(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	value := int(char - 'a' + 1)
	return value
}

func resolveLetterOrdinalReverse(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	value := 26 - (int(char - 'a'))
	return value
}

func resolveLetterReduction(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	mapping := "12345678912345678912345678"
	value := int(mapping[int(char-'a')] - '0')
	return value
}

func resolveLetterReductionReverse(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	reverseMapping := "87654321987654321987654321"
	value := int(reverseMapping[int(char-'a')] - '0')
	return value
}

func resolveLetterSumerian(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	value := int(char - 'a' + 1)
	return value * 6
}

func resolveLetterSumerianReverse(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
		return 0
	}
	value := 26 - (int(char - 'a'))
	return value * 6
}

func resolveLetterHebrewLatinCharacters(char rune) int {
	if int(char) < int('a') || int(char) > int('z') {
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

func resolveLetterHebrew(char rune) int {
	if int(char) < int('א') || int(char) > int('ת') {
		return 0
	}

	charMap := map[rune]int{
		'א': 1,
		'ב': 2,
		'ג': 3,
		'ד': 4,
		'ה': 5,
		'ו': 6,
		'ז': 7,
		'ח': 8,
		'ט': 9,
		'י': 10,
		'כ': 20,
		'ל': 30,
		'מ': 40,
		'נ': 50,
		'ס': 60,
		'ע': 70,
		'פ': 80,
		'צ': 90,
		'ק': 100,
		'ר': 200,
		'ש': 300,
		'ת': 400,
	}

	value := charMap[char]
	return value
}
