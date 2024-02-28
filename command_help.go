package main

import "fmt"

func callbackHelp() {
	fmt.Println("Welcome to Pokedex help menu!")
	fmt.Println("Here are your available commands:")
	availableCommands := getCommands()

	for _, command := range availableCommands {
		fmt.Printf("- %v  %v\n", command.name, command.description)
	}
	fmt.Println("")
}
