package gematriacalculator

import "testing"

func TestOrdinal(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 1},
		{'b', 2},
		{'z', 26},
		{'A', 1},
		{'B', 2},
		{'Z', 26},
		{'#', 0}, // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterOrdinal(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected position %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestReverseOrdinal(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 26},
		{'b', 25},
		{'z', 1},
		{'A', 26},
		{'B', 25},
		{'Z', 1},
		{'#', 0}, // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterOrdinalReverse(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected reverse position %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestReduction(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 1},
		{'b', 2},
		{'z', 8},
		{'A', 1}, // Test uppercase, should also map to 1
		{'B', 2}, // Test uppercase, should also map to 2
		{'Z', 8}, // Test uppercase, should also map to 8
		{'#', 0}, // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterReduction(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestReductionReverse(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 8},
		{'b', 7},
		{'z', 1},
		{'A', 8}, // Test uppercase, should also map to 8
		{'B', 7}, // Test uppercase, should also map to 7
		{'Z', 1}, // Test uppercase, should also map to 1
		{'#', 0}, // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterReductionReverse(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected reverse value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestSumerian(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 6},
		{'b', 12},
		{'c', 18},
		{'y', 150},
		{'z', 156},
		{'A', 6},   // Test uppercase, should also map to 6
		{'B', 12},  // Test uppercase, should also map to 12
		{'C', 18},  // Test uppercase, should also map to 18
		{'Y', 150}, // Test uppercase, should also map to 150
		{'Z', 156}, // Test uppercase, should also map to 156
		{'#', 0},   // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterSumerian(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected custom value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestSumerianReverse(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 156},
		{'b', 150},
		{'c', 144},
		{'y', 12},
		{'z', 6},
		{'A', 156}, // Test uppercase, should also map to 6
		{'B', 150}, // Test uppercase, should also map to 12
		{'C', 144}, // Test uppercase, should also map to 18
		{'Y', 12},  // Test uppercase, should also map to 150
		{'Z', 6},   // Test uppercase, should also map to 156
		{'#', 0},   // Non-alphabet character, should return 0
	}

	for _, test := range tests {
		result := resolveLetterSumerianReverse(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected custom value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestHebrewLatinCharacters(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'a', 1},
		{'b', 2},
		{'k', 10},
		{'q', 70},
		{'r', 80},
		{'s', 90},
		{'t', 100},
		{'z', 500},
		{'j', 600},
		{'w', 900},
		{'#', 0},
	}
	for _, test := range tests {
		result := resolveLetterHebrewLatinCharacters(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected custom value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestHebrew(t *testing.T) {
	tests := []struct {
		char     rune
		expected int
	}{
		{'א', 1},
		{'ח', 8},
		{'צ', 90},
		{'k', 0},
		{'q', 0},
		{'#', 0},
	}
	for _, test := range tests {
		result := resolveLetterHebrew(test.char)
		if result != test.expected {
			t.Errorf("For character %c, expected custom value %d but got %d", test.char, test.expected, result)
		}
	}
}

func TestGematria(t *testing.T) {
	tests := []struct {
		text   string
		cipher Cipher
		value  int
	}{
		{
			text:   "car",
			cipher: Ordinal,
			value:  22,
		},
		{
			text:   "car",
			cipher: OrdinalReverse,
			value:  59,
		},
		{
			text:   "car",
			cipher: Reduction,
			value:  13,
		},
		{
			text:   "car",
			cipher: ReductionReverse,
			value:  23,
		},
		{
			text:   "car",
			cipher: Sumerian,
			value:  132,
		},
		{
			text:   "car",
			cipher: SumerianReverse,
			value:  354,
		},
		{
			text:   "מְכוֹנִית",
			cipher: Hebrew,
			value:  526,
		},
		{
			text:   "בַּיִת",
			cipher: Hebrew,
			value:  412,
		},
		{
			text:   "דָּוִד",
			cipher: Hebrew,
			value:  14,
		},
	}
	for _, test := range tests {
		result := ComputeGematria(test.text, test.cipher)
		if result != test.value {
			t.Errorf("unexpected value for text %s with cipher %v: expected %d but got %d", test.text, test.cipher, test.value, result)
		}
	}
}
