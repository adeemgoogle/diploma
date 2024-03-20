package main

import (
	"diploma/internal/app"
	"diploma/pkg/config"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Println("Cannot load configs")
	}

	err = app.Start(conf)
	if err != nil {
		return
	}

}
