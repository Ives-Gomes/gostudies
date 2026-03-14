package main

import (
	"fmt"

	"github.com/Ives-Gomes/gostudies/config"
	"github.com/Ives-Gomes/gostudies/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Printf("Error initializing config: %v\n", err)
		return
	}

	router.Initialize()
}
