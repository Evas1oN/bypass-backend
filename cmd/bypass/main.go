package main

import (
	"log"
	"pavrat/bypass/internal/config"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(config.Database.Url)
}
