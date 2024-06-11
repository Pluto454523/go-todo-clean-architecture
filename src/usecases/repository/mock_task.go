package repository

import (
	"context"
	"github.com/pluto454523/go-todo-list/src/entity/task"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, et task.TaskEntity) (uint, error) {
	args := m.Called(ctx, et)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(ctx context.Context, id uint) (task.TaskEntity, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(task.TaskEntity), args.Error(1)
}

func (m *MockTaskRepository) GetAllTask(ctx context.Context, fo FilterOption, so SortOption) ([]task.TaskEntity, error) {

	args := m.Called(ctx, fo, so)
	return args.Get(0).([]task.TaskEntity), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, et task.TaskEntity) error {

	args := m.Called(ctx, et)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id uint) error {

	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) HardDeleteTask(ctx context.Context, id uint) error {

	args := m.Called(ctx, id)
	return args.Error(0)
}
