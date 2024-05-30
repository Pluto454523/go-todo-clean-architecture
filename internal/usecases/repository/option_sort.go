package repository

import "errors"

var (
	ErrorUnsupportedSortOption = errors.New("unsupported sort option")
)

type SortOption interface {
	SortOption()
}

// This file is generic option for sorting
// If there are specific sorting option, you may want to put them in that use-case file instead

type IdSortOption struct {
	Desc bool
}

func (IdSortOption) SortOption() {}

type CreatedAtSortOption struct {
	Desc bool
}

func (CreatedAtSortOption) SortOption() {}

type UpdatedAtSortOption struct {
	Desc bool
}

func (UpdatedAtSortOption) SortOption() {}

type CustomFieldSortOption struct {
	Field string
	Desc  bool
}

func (CustomFieldSortOption) SortOption() {}
