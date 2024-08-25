package main

import (
	"fmt"
	"os"

	"password-generator/pkg/password"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: pass-gen <word1> <word2>")
		return
	}

	word1 := os.Args[1]
	word2 := os.Args[2]

	password := password.Generate(word1, word2)
	fmt.Println(password)
}
