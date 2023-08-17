package main

import (
	"github.com/zzzep/pismo-challenge/config"
	"github.com/zzzep/pismo-challenge/migrations"
	"log"
)

const addr = "0.0.0.0:80"

func main() {
	if err := migrations.Run(); err != nil {
		log.Fatal(err)
		return
	}

	c := config.NewContainer()
	c.SetRoutes()

	if err := c.Router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
