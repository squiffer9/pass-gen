package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Save original args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Test with correct number of arguments
	os.Args = []string{"cmd", "test1", "test2"}
	main()
	// Note: We can't easily test the output here, but we can ensure it doesn't panic

	// Test with incorrect number of arguments
	os.Args = []string{"cmd"}
	main()
	// Again, we can't easily test the output, but we ensure it doesn't panic
}
