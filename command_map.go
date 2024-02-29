package main

import (
	"fmt"
	"log"

	"github.com/nt2311-vn/pokedexcli/internal/pokeapi"
)

func callbackMap() error {
	pokeClient := pokeapi.NewClient()
	res, err := pokeClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas")

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Println("")

	return nil
}
