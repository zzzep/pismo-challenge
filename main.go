package main

import (
	_ "github.com/zzzep/pismo-challenge/docs"
	"github.com/zzzep/pismo-challenge/src/config"
	"github.com/zzzep/pismo-challenge/src/data/migrations"
	"log"
)

const addr = "0.0.0.0:80"

// @title           Pismo Challenge Giuseppe
// @version         1.0
// @description     This is a Challenge made by Giuseppe to Pismo

// @contact.name   Giuseppe Fechio
// @contact.url    http://github.com/zzzep/pismo-challenge
// @contact.email  giuseppe.fechio@gmail.com

// @host      localhost:80
// @BasePath  /
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
