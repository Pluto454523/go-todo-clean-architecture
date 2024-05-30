package repository

import (
	"errors"
	"time"
)

var (
	ErrorUnsupportedFilterOption = errors.New("unsupported filter option")
)

type FilterOption interface {
	FilterOption()
}

type CreatedBetweenFilterOption struct {
	After  time.Time
	Before time.Time
}

func (CreatedBetweenFilterOption) FilterOption() {}

type CustomFieldValueFilterOption struct {
	Field string
	Value string
}

func (CustomFieldValueFilterOption) FilterOption() {}
