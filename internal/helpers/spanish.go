package helpers

import (
    "fmt"
    "math/rand"
)

// generateFakeSpanishDNI generates a fake (but valid) Spanish DNI (Documento Nacional de Identidad).
func generateFakeSpanishDNI() string {
    const letters = "TRWAGMYFPDXBNJZSQVHLCKE"
    number := rand.Intn(100000000)
    letter := letters[number%23]
    return fmt.Sprintf("%08d%c", number, letter)
}