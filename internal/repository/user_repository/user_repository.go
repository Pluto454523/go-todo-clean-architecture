package user_repository

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("user_repository")
