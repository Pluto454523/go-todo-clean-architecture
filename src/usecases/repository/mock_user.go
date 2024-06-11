package repository

import (
	"github.com/pluto454523/go-todo-list/src/entity/user"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(id uint) (user.UserEnity, error) {
	args := m.Called(id)
	return args.Get(0).(user.UserEnity), args.Error(1)
}
