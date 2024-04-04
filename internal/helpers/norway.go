package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// generateControlDigits generates the control digits for a given Norwegian SSN
func generateControlDigits(ssn string) (int, int) {
	// Weights for calculating control digits
	var weights1 = []int{3, 7, 6, 1, 8, 9, 4, 5, 2}
	var weights2 = []int{5, 4, 3, 2, 7, 6, 5, 4, 3, 2}

	var sum1, sum2 int
	for i, r := range ssn {
		digit := int(r - '0')
		if i < 9 { // First control digit calculation
			sum1 += digit * weights1[i]
		}
		if i < 10 { // Second control digit calculation includes first control digit
			sum2 += digit * weights2[i]
		}
	}

	// Calculate control digits
	control1 := 11 - (sum1 % 11)
	control2 := 11 - (sum2 % 11)

	// Adjust for special cases
	if control1 == 11 {
		control1 = 0
	}
	if control2 == 11 {
		control2 = 0
	}

	return control1, control2
}

// generateFakeNorwegianSSN generates a fake (but valid) Norwegian SSN.
func generateFakeNorwegianSSN(dob time.Time) string {
	individualNumber := rand.Intn(500)
	year, month, day := dob.Date()
	ssn := fmt.Sprintf("%02d%02d%02d%03d", day, month, year%100, individualNumber)

	control1, control2 := generateControlDigits(ssn)
	ssn += fmt.Sprintf("%d%d", control1, control2)

	return ssn
}
