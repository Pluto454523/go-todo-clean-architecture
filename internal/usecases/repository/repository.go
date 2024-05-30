package repository

import (
	"github.com/pluto454523/go-todo-list/internal/entity/task"
	"github.com/pluto454523/go-todo-list/internal/entity/user"
	"golang.org/x/net/context"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, et task.TaskEntity) (uint, error)
	DeleteTask(ctx context.Context, id uint) error
	GetAllTask(ctx context.Context, fo FilterOption, so SortOption) ([]task.TaskEntity, error)
	GetTaskByID(ctx context.Context, id uint) (task.TaskEntity, error)
	HardDeleteTask(ctx context.Context, id uint) error
	UpdateTask(ctx context.Context, et task.TaskEntity) (err error)
}

type UserRepository interface {
	GetUserByID(id uint) (user user.UserEnity, err error)
}
