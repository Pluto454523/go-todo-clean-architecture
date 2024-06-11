package helper

import "github.com/pluto454523/go-todo-list/src/usecases/model"

type SetError struct {
	Code    int16
	Message string
}

var errorList = map[error]ErrorMapInfo{
	model.ErrtaskNoAlreadyExist: {500, "task_no_already_exist"},
	model.ErrtaskNotFound:       {404, "task_not_found"},
}
