package usecases

import (
	"github.com/pluto454523/go-todo-list/src/usecases/repository"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("UsecaseDependency")

type (
	UsecaseDependency struct {
		TaskRepository repository.TaskRepository
		UserRepository repository.UserRepository
	}
)

func New(
	TaskRepository repository.TaskRepository,
	UserRepository repository.UserRepository,
) *UsecaseDependency {

	return &UsecaseDependency{
		TaskRepository: TaskRepository,
		UserRepository: UserRepository,
	}
}
