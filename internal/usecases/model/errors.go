package model

import "errors"

var (
	ErrtaskNotFound       = errors.New("task not found")
	ErrtaskNoAlreadyExist = errors.New("task no already exist")
)
