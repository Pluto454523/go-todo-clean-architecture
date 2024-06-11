package model

import (
	"time"
)

type (
	TaskPayload struct {
		Title       string    `json:"title" validate:"required,min=5,max=20"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
		Status      string    `json:"status" validate:"required, lowercase"`
	}

	TaskResponse struct {
		ID          uint      `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"due_date"`
		Status      string    `json:"status"`
	}

	TaskOptional struct {
		Order  string `query:"order" validate:"omitempty,oneof=id title description status due_date"`
		Sort   string `query:"sort" validate:"omitempty,oneof=desc asc"`
		Filter string `query:"filter" validate:"omitempty,oneof='title description status"`
		Value  string `query:"value" validate:"omitempty"`
	}
)
