package password

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"

	"password-generator/pkg/utils"
)

const passwordLength = 16
const symbols = "!@#$%^*-_=+[]:,.?/"
const upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const digits = "0123456789"

func Generate(word1, word2 string) string {
	combinedWords := word1 + word2
	hash := sha256.Sum256([]byte(combinedWords))
	hashHex := hex.EncodeToString(hash[:])

	if len(hashHex) < 16 {
		panic("hashHex length is less than 16")
	}

	seed := utils.HexToInt64(hashHex[:16])
	r := rand.New(rand.NewSource(seed))

	// Generate base password
	password := make([]byte, passwordLength-4) // Reserve 4 spots for required characters
	allChars := symbols + upperLetters + lowerLetters + digits
	for i := range password {
		password[i] = allChars[r.Intn(len(allChars))]
	}

	// Ensure at least one of each required character type
	password = append(password, symbols[r.Intn(len(symbols))])
	password = append(password, upperLetters[r.Intn(len(upperLetters))])
	password = append(password, lowerLetters[r.Intn(len(lowerLetters))])
	password = append(password, digits[r.Intn(len(digits))])

	// Shuffle the password
	r.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func ContainsRequiredChars(password string) bool {
	return strings.ContainsAny(password, symbols) &&
		strings.ContainsAny(password, upperLetters) &&
		strings.ContainsAny(password, lowerLetters) &&
		strings.ContainsAny(password, digits)
}
