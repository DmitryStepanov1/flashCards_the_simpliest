package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func containsCyrillicOrLatin(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Cyrillic, r) || unicode.Is(unicode.Latin, r) {
			return true
		}
	}
	return false
}

func randomWord(m map[string]string) string {
	//k := rand.Intn(len(m))

	for i, v := range m {
		s := fmt.Sprintf("Переведи %s:", i)
		fmt.Println(s)
		return v
	}

	return ""

}

func main() {
	// Create a map to store input values
	inputMap := make(map[string]string)

	// Create a scanner to read input from terminal
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter key-value pairs (key=value) separated by newline. Enter 'done' to finish:")

	// Keep reading input until the user enters 'done'
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the user wants to finish entering input
		if input == "done" {
			break
		}

		// Split input by "=" to separate key and value
		parts := strings.Split(input, "=")
		if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 || containsCyrillicOrLatin(parts[0]) == false || containsCyrillicOrLatin(parts[1]) == false {
			fmt.Println("Invalid input format. Please use 'key=value' format.")
			continue
		}

		// Store key-value pair in the map
		key := parts[0]
		value := parts[1]
		inputMap[key] = value
	}

	for {
		v := randomWord(inputMap)

		scanner := bufio.NewScanner(os.Stdin)

	jumpTo:

		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the user wants to finish entering input
		if input == "done" {
			break
		} else if input == v {
			fmt.Println("Correct! Try next word")
			continue
		} else {
			fmt.Println("Wrong, try again")
			goto jumpTo
		}

	}

	// Print the map
	fmt.Println("Map contents:")
	for key, value := range inputMap {
		fmt.Printf("%s: %s\n", key, value)
	}
}
