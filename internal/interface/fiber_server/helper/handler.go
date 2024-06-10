package helper

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
)

const OK string = "OK"

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
	IssueId   string `json:"issue_id"`
}

type ErrorMapInfo struct {
	StatusCode int
	ErrorCode  string
}

func SendError(c *fiber.Ctx, status int, err error, errCode, issueId string) error {

	if issueId == "" {
		span := trace.SpanFromContext(c.Context())
		issueId = span.SpanContext().TraceID().String()
	}

	c.Locals("error", err)

	return c.Status(status).JSON(ErrorResponse{
		Error:     err.Error(),
		ErrorCode: errCode,
		IssueId:   issueId,
	})
}

func ErrorHandler(c *fiber.Ctx, err error, issueIds ...string) error {
	unwrapErr := errors.Unwrap(err)
	if unwrapErr == nil {
		unwrapErr = err
	}

	var issueId string
	if len(issueIds) > 0 {
		issueId = issueIds[0]
	}

	for iErr, code := range errorList {
		if errors.Is(err, iErr) {
			return SendError(c, code.StatusCode, err, code.ErrorCode, issueId)
		}
	}

	return SendError(c, 500, err, "unexpected_error", issueId)
}
