package main

import (
	"final-project/config"
	"final-project/router"
	"log"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
