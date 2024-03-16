package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestContainsCyrillicOrLatin(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Hello", true},
		{"Привет", true},
		{"123", false},
		{"", false},
		{"$%*", false},
	}

	for _, tc := range testCases {
		result := containsCyrillicOrLatin(tc.input)
		if result != tc.expected {
			t.Errorf("containsCyrillicOrLatin(%s) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}

func TestRandomWord(t *testing.T) {
	testMap := map[string]string{
		"apple":  "яблоко",
		"banana": "банан",
	}

	var buf bytes.Buffer
	oldStdout := os.Stdout
	os.Stdout = &buf

	randomWord(testMap)

	os.Stdout = oldStdout

	output := buf.String()
	expectedOutput := "Переведи apple:\n"
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("randomWord did not print the expected output: %s", expectedOutput)
	}
}

func TestMainFunc(t *testing.T) {
	// Redirect Stdin for testing
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	input := "apple=яблоко\nbanana=банан\ndone\nbanana\nbanana\ndone\n"
	os.Stdin = strings.NewReader(input)

	var buf bytes.Buffer
	oldStdout := os.Stdout
	os.Stdout = &buf

	main()

	os.Stdout = oldStdout

	output := buf.String()

	expectedOutput := "Correct! Try next word\nMap contents:\napple: яблоко\nbanana: банан\n"
	if output != expectedOutput {
		t.Errorf("main function did not produce the expected output: %s", expectedOutput)
	}
}
