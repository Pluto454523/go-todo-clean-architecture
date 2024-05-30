package user

import (
	"fmt"
	"unicode/utf8"
)

type (
	UserEnity struct {
		ID   uint
		Name string
	}
)

var (
	ErrNameIsEmpty    string = "ชื่อใหม่ของคุณเป็นค่าว่าง"
	ErrNameOverLenght string = "ชื่อใหม่ของคุณเป็นค่าว่าง"
)

func (ue UserEnity) ValidateNewName(name string) (newName string, err error) {

	if name == "" {
		return name, fmt.Errorf(ErrNameIsEmpty)
	}

	if utf8.RuneCountInString(name) >= 16 {
		return name, fmt.Errorf(ErrNameOverLenght)
	}

	return name, nil
}
