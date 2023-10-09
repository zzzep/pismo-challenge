package httpserver

import (
	"github.com/zzzep/pismo-challenge/src/adapters/primary"
	"github.com/zzzep/pismo-challenge/src/adapters/secondary/migrations"
	"log"
)

const addr = "0.0.0.0:80"

func Run() {
	if err := migrations.Run(); err != nil {
		log.Fatal(err)
		return
	}

	c := primary.NewContainer()
	SetRoutes(c)

	if err := c.Router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
