package task_repository

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("task_repository")
