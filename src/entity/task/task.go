package task

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"time"
)

type TaskEntity struct {
	ID          uint
	Title       string    `validate:"required,min=2,max=64"`
	Description string    `validate:"required,min=2,max=64"`
	Status      string    `validate:"required,status"`
	DueDate     time.Time `validate:"required"`
}

func (task TaskEntity) ChangeStatus(status string) (TaskEntity, error) {

	if status != "todo" && status != "doing" && status != "done" {
		return TaskEntity{}, errors.New("can't change status")
	}

	if task.Status == "done" {
		return TaskEntity{}, errors.New("task is done")
	}

	if task.Status == status {
		return TaskEntity{}, errors.New("task already have a status")
	}

	if status == "done" {
		if task.Status != "doing" {
			return TaskEntity{}, errors.New("can't change status")
		}
	}

	task.Status = status
	return task, nil
}

// isZero checks if a reflect.Value is zero value.
func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

// PatchTask updates the current TaskEntity with non-zero values from newTask.
func (task TaskEntity) PatchTask(newTask TaskEntity) TaskEntity {
	newValue := reflect.ValueOf(newTask)
	currentValue := reflect.ValueOf(&task).Elem()

	for i := 0; i < newValue.NumField(); i++ {
		field := newValue.Field(i)
		if !isZero(field) {
			currentValue.Field(i).Set(field)
		}
	}

	return task
}

func ValidateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	switch status {
	case "todo", "doing", "done":
		return true
	default:
		return false
	}
}

func ValidateSlug(fl validator.FieldLevel) bool {
	regX, _ := regexp.Compile("^[a-z0-9]+(?:-[a-z0-9]+)*$")
	return regX.MatchString(fl.Field().String())
}

func (task TaskEntity) Validate() error {
	v := validator.New()
	var err error

	err = v.RegisterValidation("slug", ValidateSlug)
	if err != nil {
		return err
	}

	err = v.RegisterValidation("status", ValidateStatus)
	if err != nil {
		return err
	}

	return v.Struct(task)
}
