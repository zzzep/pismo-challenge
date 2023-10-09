package main

import (
	_ "github.com/zzzep/pismo-challenge/docs"
	_ "github.com/zzzep/pismo-challenge/src/adapters/primary"
	"github.com/zzzep/pismo-challenge/src/adapters/primary/httpserver"
)

// @title           Pismo Challenge Giuseppe
// @version         1.0
// @description     This is a Challenge made by Giuseppe Fechio to Pismo

// @contact.name   Giuseppe Fechio
// @contact.url    http://github.com/zzzep/pismo-challenge
// @contact.email  giuseppe.fechio@gmail.com

// @host      localhost:80
// @BasePath  /
func main() {
	httpserver.Run()
}
