package fiber_server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pluto454523/go-todo-list/cmd/generics_server/config"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server/middleware"
	"github.com/pluto454523/go-todo-list/internal/interface/fiber_server/route"
	"github.com/pluto454523/go-todo-list/internal/usecases"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type FiberServer struct {
	useCase *usecases.UsecaseDependency
	server  *fiber.App
	config  *config.Config
}

func New(useCase *usecases.UsecaseDependency, cfg *config.Config) *FiberServer {

	server := fiber.New(fiber.Config{
		CaseSensitive:         false,
		StrictRouting:         false,
		DisableStartupMessage: true,
		ReadTimeout:           30 * time.Second,
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	f := &FiberServer{
		useCase: useCase,
		server:  server,
		config:  cfg,
	}
	// Create a new Fiber instance
	f.server.Use(middleware.LoggerMiddleware)
	f.server.Use(middleware.TracerMiddleware)

	// Initialize handler
	todoTaskHandler := route.NewTaskHandler(f.useCase)

	// Routes
	f.server.Post("/tasks", todoTaskHandler.CreateTask)
	f.server.Get("/tasks/:id", todoTaskHandler.GetTaskByID)
	f.server.Put("/tasks/:id", todoTaskHandler.UpdateTask)
	f.server.Patch("/tasks/:id", todoTaskHandler.PatchTask)
	f.server.Delete("/tasks/:id", todoTaskHandler.DeleteTask)
	f.server.Get("/tasks", todoTaskHandler.GetAllTask)
	f.server.Post("/changestatus/:id", todoTaskHandler.ChangeStatus)

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
		log.Info().Msg("Server is starting...")
		err := f.server.Listen(":" + strconv.Itoa(f.config.Server.Port))
		if err != nil {
			log.Error().Err(err).Msg("Server error")
		}
		log.Info().Msg("Server has been shutdown")
	}()
}
