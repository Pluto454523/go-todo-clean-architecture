package config

import (
	"github.com/rs/zerolog"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DB       DB
		Server   Server
		Services Services
	}

	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
	}

	Server struct {
		Port     int
		LogLevel zerolog.Level
	}

	Services struct {
		OtelGrpcEndpoint string `env:"OTEL_GRPC_ENDPOINT" envDefault:"localhost:4317"`
	}
)

func GetConfig() Config {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load environment file")
		// log.Fatal().
		// 	Err(err).
		// 	Msg("Failed to load environment file")
	}

	ict, err := time.LoadLocation(os.Getenv("POSTGRES_TIMEZONE"))
	if err != nil {
		panic(err)
	}
	time.Local = ict

	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic(err)
	}

	svPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	configInstance := Config{
		DB: DB{
			Host:     os.Getenv("POSTGRES_HOSTNAME"),
			Port:     dbPort,
			DBName:   os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
			TimeZone: os.Getenv("POSTGRES_TIMEZONE"),
		},
		Server: Server{
			Port:     svPort,
			LogLevel: zerolog.InfoLevel,
		},
		Services: Services{
			OtelGrpcEndpoint: os.Getenv("OTEL_GRPC_ENDPOINT"),
		},
	}

	return configInstance
}
