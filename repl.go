package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")
		scanner.Scan()

		text := scanner.Text()

		cleaned := cleanInput(text)

		fmt.Println("echoing: ", cleaned)

	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
