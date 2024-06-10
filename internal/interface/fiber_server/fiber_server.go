package fiber_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server/config"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server/middleware"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server/route"
	"github.com/pluto454523/go-todo-list/internal/usecases"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type FiberServer struct {
	useCase *usecases.UsecaseDependency
	server  *fiber.App
	config  *config.ServerConfig
}

func New(uc *usecases.UsecaseDependency, cfg *config.ServerConfig) *FiberServer {

	server := fiber.New(fiber.Config{
		CaseSensitive:         false,
		StrictRouting:         false,
		DisableStartupMessage: true,
		ReadTimeout:           30 * time.Second,
	})

	if cfg.CorsAllowAll {
		server.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		}))
	}

	f := &FiberServer{
		useCase: uc,
		server:  server,
		config:  cfg,
	}

	//helper.AddSwaggerUI(server, task_spec.GetSwagger, "/v1")

	if cfg.RequestLog {
		server.Use(middleware.LoggerMiddleware)
	}

	server.Use(middleware.TracerMiddleware)

	// Initialize handler and route
	todoTaskHandler := route.NewTaskHandler(f.useCase)
	server.Post("/tasks", todoTaskHandler.CreateTask)
	server.Get("/tasks/:id", todoTaskHandler.GetTaskByID)
	server.Put("/tasks/:id", todoTaskHandler.UpdateTask)
	server.Patch("/tasks/:id", todoTaskHandler.PatchTask)
	server.Delete("/tasks/:id", todoTaskHandler.DeleteTask)
	server.Get("/tasks", todoTaskHandler.GetAllTask)
	server.Post("/changestatus/:id", todoTaskHandler.ChangeStatus)

	//task_spec.RegisterHandlersWithOptions(
	//	f.server,
	//	route.NewRouteTaskV1(uc),
	//	task_spec.FiberServerOptions{
	//		BaseURL: "/v1",
	//	},
	//)

	return f
}

func (f FiberServer) Start(wg *sync.WaitGroup) {
	wg.Add(2)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer wg.Done()
		<-exit
		log.Info().Msg("Shutting down server...")

		err := f.server.Shutdown()
		if err != nil {
			log.Error().Err(err).Msg("Server shutdown with error")
		} else {
			log.Info().Msg("Server gracefully shutdown")
		}
	}()

	go func() {
		defer wg.Done()
		log.Info().Msgf("Server is starting....%v", f.config.ListenAddress)
		err := f.server.Listen(f.config.ListenAddress)

		if err != nil {
			log.Error().Err(err).Msg("Server error")
		}

		log.Info().Msg("Server has been shutdown")
	}()
}
