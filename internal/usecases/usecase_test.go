package usecases

import (
	"github.com/pluto454523/go-todo-list/internal/entity/task"
	"github.com/pluto454523/go-todo-list/internal/usecases/repository"
	"time"
)

func newMocksRepository() (
	*repository.MockTaskRepository,
	*repository.MockUserRepository,
	*UsecaseDependency,
) {
	tr := new(repository.MockTaskRepository)
	ur := new(repository.MockUserRepository)

	uc := New(tr, ur)

	return tr, ur, uc
}

func newValidTask() task.TaskEntity {
	return task.TaskEntity{
		ID:          1,
		Title:       "New Task",
		Status:      "todo",
		DueDate:     time.Now(),
		Description: "Description of new task",
	}
}
