package main

import (
	"fmt"
	"log"

	"github.com/nt2311-vn/pokedexcli/internal/pokeapi"
)

func main() {
	// startRepl()

	pokeClient := pokeapi.NewClient()

	res, err := pokeClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
