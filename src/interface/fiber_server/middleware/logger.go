package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"strings"
)

func LoggerMiddleware(c *fiber.Ctx) error {

	bodyReq := strings.ReplaceAll(string(c.Body()), "\n", "")
	bodyReq = strings.ReplaceAll(bodyReq, " ", "")

	log.Info().
		Str("ip_address", c.IP()).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("handler", c.Route().Path).
		//Str("query", c.Queries()).
		Str("body", bodyReq).
		Msg("request_payload")

	err := c.Next()

	log.Info().
		Int("status", c.Response().StatusCode()).
		Str("body", string(c.Response().Body())).
		Msg("response_payload")

	return err
}
