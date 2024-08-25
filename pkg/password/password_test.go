package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		word1, word2 string
	}{
		{"hello", "world"},
		{"test", "password"},
		{"golang", "rocks"},
	}

	for _, tc := range testCases {
		password := Generate(tc.word1, tc.word2)

		if len(password) != passwordLength {
			t.Errorf("Expected password length %d, got %d", passwordLength, len(password))
		}

		if !ContainsRequiredChars(password) {
			t.Errorf("Password %s does not meet complexity requirements", password)
		}

		// Test idempotence
		password2 := Generate(tc.word1, tc.word2)
		if password != password2 {
			t.Errorf("Generate function is not idempotent for inputs %s and %s", tc.word1, tc.word2)
		}
	}
}

func TestContainsRequiredChars(t *testing.T) {
	validPassword := "aA1!bcdefghijklm"
	if !ContainsRequiredChars(validPassword) {
		t.Errorf("Expected %s to be valid, but it was not", validPassword)
	}

	invalidPasswords := []string{
		"abcdefghijklmnop", // No uppercase, digit, or symbol
		"ABCDEFGHIJKLMNOP", // No lowercase, digit, or symbol
		"1234567890123456", // No uppercase, lowercase, or symbol
		"!@#$%^&*()_+{}[]", // No uppercase, lowercase, or digit
	}

	for _, invalidPassword := range invalidPasswords {
		if ContainsRequiredChars(invalidPassword) {
			t.Errorf("Expected %s to be invalid, but it was not", invalidPassword)
		}
	}
}
