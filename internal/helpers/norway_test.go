package helpers

import (
	"testing"

	"syreclabs.com/go/faker"
)

func TestSSN(t *testing.T) {
	ssn := generateFakeNorwegianSSN(faker.Date().Birthday(18, 90))

	if !isValidNorwegianSSN(ssn) {
		t.Errorf("Generated SSN is invalid: %s", ssn)
	}
}

// isValidNorwegianSSN checks if a given Norwegian SSN is valid (assumes the DOB is correct, as we generate that ourselves).
func isValidNorwegianSSN(ssn string) bool {
	if len(ssn) != 11 {
		return false
	}

	// Assuming generateControlDigits is implemented and returns control digits correctly
	control1, control2 := generateControlDigits(ssn[:9])
	expectedControl1 := int(ssn[9] - '0')
	expectedControl2 := int(ssn[10] - '0')

	return control1 == expectedControl1 && control2 == expectedControl2
}
