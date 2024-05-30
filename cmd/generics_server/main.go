// main.go
package main

import (
	"github.com/pluto454523/go-todo-list/cmd/generics_server/config"
	"github.com/pluto454523/go-todo-list/internal/usecases"
)

var (
	conf config.Config
)

func main() {

	// Init Config envoriment
	conf = config.GetConfig()

	// Init Logger config
	initLogger(conf)

	// Init Tracer Config
	initTracer(conf)

	// Initialize usecases
	usecase := usecases.New(initRepositories(conf))

	initInterface(conf, usecase)
}
