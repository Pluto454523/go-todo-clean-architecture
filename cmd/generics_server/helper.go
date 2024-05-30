package main

import (
	"context"
	"fmt"
	"github.com/pluto454523/go-todo-list/cmd/generics_server/config"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server"
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

func initRepositories(cfg config.Config) (
	repository.TaskRepository,
	repository.UserRepository,
) {

	// Connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.Port,
		cfg.DB.SSLMode,
		cfg.DB.TimeZone,
	)

	return task_repository.NewTaskRepository(setupGorm(postgres.Open(dsn))),
		user_repository.NewUserRepository(setupGorm(postgres.Open(dsn)))
}

func initInterface(cfg config.Config, uc *usecases.UsecaseDependency) {

	wg := new(sync.WaitGroup)

	server := fiber_server.New(uc, &cfg)
	log.Info().Msg("Fiber server initialized")

	server.Start(wg)
	log.Info().Msg("Fiber server started")

	wg.Wait()
	log.Info().Msg("Application stopped")
}

func initLogger(cfg config.Config) {
	zerolog.SetGlobalLevel(cfg.Server.LogLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	multi := io.MultiWriter(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339Nano,
		},
	)

	log.Logger = zerolog.
		New(multi).
		With().
		Timestamp().
		Logger()

	log.Info().Msg("Logger initialized.")
}

func initTracer(cfg config.Config) {

	client := otlptracegrpc.NewClient(
		//otlptracegrpc.WithInsecure(), //เพื่อใช้การเชื่อมต่อที่ไม่เข้ารหัส (ไม่แนะนำใน production)
		otlptracegrpc.WithEndpoint(cfg.Services.OtelGrpcEndpoint),
	)

	//exporter, err := otlptrace.New(context.Background(), client)
	ctx := func() context.Context {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //** ไม่ควรทำ **
		defer cancel()                                                          // ให้ยกเลิก context เมื่อ main function จบการทำงาน
		return ctx
	}

	//traceเป็นoptional
	exporter, err := otlptrace.New(ctx(), client)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error init Otlp exporter.")
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("todo-list-backend-go"),
			semconv.ServiceVersion("v1.0.0"),
			attribute.String("environment", "development"),
		)),
	)

	otel.SetTracerProvider(tp)

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
