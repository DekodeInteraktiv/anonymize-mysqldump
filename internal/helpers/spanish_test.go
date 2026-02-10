package helpers

import (
	"testing"
)

// TestSpanishDNI checks that the generated DNI is valid: 8 digits followed by a control letter
func TestSpanishDNI(t *testing.T) {
	dni := string(generateSpanishDNI(nil).Val)

	if !isValidSpanishDNI(dni) {
		t.Errorf("Generated DNI is invalid: %s", dni)
	}
}

// isValidSpanishDNI checks if a given Spanish DNI is valid
func isValidSpanishDNI(dni string) bool {
	const letters = "TRWAGMYFPDXBNJZSQVHLCKE"

	if len(dni) != 9 {
		return false
	}

	// Check that the first 8 characters are digits
	for i := 0; i < 8; i++ {
		if dni[i] < '0' || dni[i] > '9' {
			return false
		}
	}

	// Check that the last character matches the expected letter
	number := 0
	for i := 0; i < 8; i++ {
		number = number*10 + int(dni[i]-'0')
	}
	expectedLetter := letters[number%23]

	return dni[8] == expectedLetter
}
