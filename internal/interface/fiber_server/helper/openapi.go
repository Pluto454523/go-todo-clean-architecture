package helper

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/swaggest/swgui/v5emb"
)

type spec interface {
	GetSwagger() (*openapi3.T, error)
}

func AddSwaggerUI(router fiber.Router, s func() (*openapi3.T, error), baseUrl string) error {
	oapi, err := s()
	if err != nil {
		return err
	}

	oapi.AddServer(&openapi3.Server{
		URL: fmt.Sprintf("%s", baseUrl),
	})

	swagger, err := oapi.MarshalJSON()
	if err != nil {
		return err
	}

	p := fmt.Sprintf("%s/swagger", baseUrl)

	g := router.Group(p)
	g.Get("swagger.json", func(c *fiber.Ctx) error {
		return c.Send(swagger)
	})
	g.Get("*", adaptor.HTTPHandler(v5emb.New(
		oapi.Info.Title,
		fmt.Sprintf("%s/swagger.json", p),
		p,
	)))

	return nil
}
