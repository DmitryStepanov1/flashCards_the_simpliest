package main

import (
	"os"
	"strings"
	"testing"
)

func TestContainsCyrillicOrLatin(t *testing.T) {
	testData := map[string]bool{
		"Hello":  true,
		"Привет": true,
		"123":    false,
		"":       false,
	}

	for input, expected := range testData {
		result := containsCyrillicOrLatin(input)
		if result != expected {
			t.Errorf("containsCyrillicOrLatin(%s) = %t; want %t", input, result, expected)
		}
	}
}

func TestRandomWord(t *testing.T) {
	testMap := map[string]string{
		"apple":  "яблоко",
		"banana": "банан",
	}

	// Test that randomWord returns a value from the map
	result := randomWord(testMap)
	if _, ok := testMap[result]; !ok {
		t.Errorf("randomWord did not return a value from the provided map")
	}
}

func TestCheckWord(t *testing.T) {
	// Mock user input using a strings.Reader
	userInput := "banana\n"
	inputReader := strings.NewReader(userInput)

	// Redirect Stdin for testing
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = inputReader

	// Capture output for testing
	output := captureOutput(func() {
		v := "банан"
		checkWord(v)
	})

	expectedOutput := "Wrong, try again\n"
	if output != expectedOutput {
		t.Errorf("checkWord output: %s; want %s", output, expectedOutput)
	}
}

// Helper function to capture stdout
func captureOutput(f func()) string {
	res := ""
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = os.Stdout

	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		res += string(buf[:n])
	}
	return res
}
