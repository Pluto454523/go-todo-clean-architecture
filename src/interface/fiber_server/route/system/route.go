package system

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.sellsuki.com/sellsuki/share/backend/boilerplate-backend-go/src/interface/fiber_server/config"
	"gitlab.sellsuki.com/sellsuki/share/backend/boilerplate-backend-go/src/interface/fiber_server/helper"
	"gitlab.sellsuki.com/sellsuki/share/backend/boilerplate-backend-go/src/use_case"
)

type routeSystem struct {
	config  *config.ServerConfig
	useCase *use_case.UseCase
}

func (r routeSystem) GetLiveliness(c *fiber.Ctx) error {
	return c.Send([]byte(helper.OK))
}

func (r routeSystem) GetLiveness(c *fiber.Ctx) error {
	return c.Send([]byte(helper.OK))
}

func (r routeSystem) GetReadiness(c *fiber.Ctx) error {
	err := r.useCase.HealthCheck(c.Context())
	if err != nil {
		return helper.ErrorHandler(c, err)
	}

	return c.Send([]byte(helper.OK))
}

func (r routeSystem) GetVersion(c *fiber.Ctx) error {
	return c.Send([]byte(r.config.AppVersion))
}

func NewRouteSystem(config *config.ServerConfig, useCase *use_case.UseCase) ServerInterface {
	return &routeSystem{
		config:  config,
		useCase: useCase,
	}
}
