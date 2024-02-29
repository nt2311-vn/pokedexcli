package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas")

	for _, area := range res.Results {
		fmt.Printf("- %s\n", area.Name)
	}
	fmt.Println("")

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}
