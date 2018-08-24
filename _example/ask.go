package main

import (
	"log"

	"github.com/wulfheart/go-input"
)

func main() {

	ui := &input.UI{}

	query := "What is your name?"
	name, err := ui.Ask(query, &input.Options{
		Required:  true,
		HideOrder: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer is %s\n", name)
}
