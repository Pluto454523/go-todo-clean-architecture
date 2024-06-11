// main.go
package main

import (
	"github.com/pluto454523/go-todo-list/src/usecases"
)

type config struct {
	AppName     string `env:"APP_NAME" envDefault:"boilerplate-backend-go"`
	AppVersion  string `env:"APP_VERSION" envDefault:"v0.0.0"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
	Port        uint   `env:"PORT" envDefault:"8080"`
	GrpcPort    uint   `env:"GRPC_PORT" envDefault:"50051"`
	Debuglog    bool   `env:"DEBUG_LOG" envDefault:"true"`
	Services    struct {
		TaskPostgresqlUri     string `env:"DATABASE_TASK_POSTGRESQL_URI" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
		UserPostgresqlUri     string `env:"DATABASE_USER_POSTGRESQL_URI" envDefault:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
		OtelGrpcEndpoint      string `env:"OTEL_GRPC_ENDPOINT" envDefault:"localhost:4317"`
		AddressServiceBaseUrl string `env:"SERVICE_ADDRESS_BASE_URL" envDefault:"http://localhost:8080"`
	}
}

func main() {

	// Init Config envoriment
	conf := initEnvironment()

	// Init Logger config
	initLogger(conf)

	// Init Tracer Config
	initTracer(conf)

	// Initialize usecases
	usecase := usecases.New(initRepositories(conf))

	initInterface(conf, usecase)
}
