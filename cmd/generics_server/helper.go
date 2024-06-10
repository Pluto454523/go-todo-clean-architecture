package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server"
	sv_config "github.com/pluto454523/go-todo-list/internal/interface/fiber_server/config"
	"github.com/pluto454523/go-todo-list/internal/repository/task_repository"
	"github.com/pluto454523/go-todo-list/internal/repository/user_repository"
	"github.com/pluto454523/go-todo-list/internal/usecases"
	"github.com/pluto454523/go-todo-list/internal/usecases/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"os"
	"sync"
	"time"
)

func initEnvironment() config {

	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msgf("Failed loading .env file: %s", err)
	}

	var cfg config
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error parse env to struct: %s", err)
	}

	return cfg

}

func initRepositories(cfg config) (
	repository.TaskRepository,
	repository.UserRepository,
) {

	taskRepo := task_repository.NewTaskRepository(setupGorm(postgres.Open(cfg.Services.TaskPostgresqlUri)))
	userRepo := user_repository.NewUserRepository(setupGorm(postgres.Open(cfg.Services.UserPostgresqlUri)))

	return taskRepo, userRepo
}

func initInterface(cfg config, uc *usecases.UsecaseDependency) {

	wg := new(sync.WaitGroup)

	server := fiber_server.New(uc, &sv_config.ServerConfig{
		AppVersion:    cfg.AppVersion,
		ListenAddress: fmt.Sprintf(":%d", cfg.Port),
		RequestLog:    true,
	})

	log.Info().Msg("Fiber server initialized")

	server.Start(wg)
	log.Info().Msg("Fiber server started")

	wg.Wait()
	log.Info().Msg("Application stopped")
}

func initLogger(cfg config) {

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	multi := io.MultiWriter(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339Nano,
		},
	)

	if cfg.Debuglog {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.
		New(multi).
		With().
		Timestamp().
		Logger()

	log.Info().Msg("Logger initialized.")
}

func initTracer(cfg config) {
	// tracing is optional
	// คือหาก trace ใช้งานไม่ได้ระบบควรจะทำงานได้โดยไม่มี ผลกระทบกับระบบหลัก

	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(), // เพื่อใช้การเชื่อมต่อที่ไม่เข้ารหัส not recommemded on production
		otlptracegrpc.WithEndpoint(cfg.Services.OtelGrpcEndpoint),
	)

	// use otlptrace exporter
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		log.Fatal().Err(err).Msg("Error init Otlp exporter")
	}

	// create new resource with attributes (global)
	resources := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(cfg.AppName),
		semconv.ServiceVersion(cfg.AppVersion),
		attribute.String("environment", cfg.Environment),
	)

	// Create new tracer provider
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resources),
	)

	// Set tracer provider
	otel.SetTracerProvider(provider)

	log.Info().Msg("Tracer initialized.")
}

func setupGorm(d gorm.Dialector) *gorm.DB {

	// New logger for detailed SQL logging
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // Slow SQL threshold
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // Enable color
	//	},
	//)

	db, err := gorm.Open(d, &gorm.Config{
		//Logger: newLogger,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error open postgres url")
	}

	//sqlDB, err := db.DB()
	//if err != nil {
	//	slog.Fatal("Error get sql db").WithError(err).Write()
	//}
	//
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(100)
	//sqlDB.SetConnMaxLifetime(0)
	////sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
