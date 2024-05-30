package task_repository

import (
	"fmt"
	"github.com/pluto454523/go-todo-list/internal/usecases/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyPgFilterOption(q *gorm.DB, fo repository.FilterOption) (*gorm.DB, error) {
	switch x := (fo).(type) {
	case nil:
		return q, nil

	case *repository.CreatedBetweenFilterOption:
		return q.Where("created_at BETWEEN ? AND ?", x.After, x.Before), nil
	case *repository.CustomFieldValueFilterOption:
		return q.Where(fmt.Sprintf("%v like '%v%%'", x.Field, x.Value)), nil

	//case *repository.InProviderFilterOption:
	//	return q.Where("provider_code = ?", x.ProviderCode), nil
	//
	//case *repository.CreatedBetweenAndInStoreFilterOption:
	//	return q.Where("store_id = ? AND \"order\".created_at BETWEEN ? AND ?", x.StoreId, x.After, x.Before), nil
	//
	////	never try if filter yet, hope it will work, XD
	//case *repository.InStoreAndHaveProductAlsoCreatedInDateRangeFilterOption:
	//	return q.Where("store_id = ? AND EXISTS (SELECT 1 FROM item WHERE item.order_id = order.order_id AND product_id = ?) AND \"order\".created_at BETWEEN ? AND ?", x.StoreId, x.ProductId, x.After, x.Before), nil

	default:
		return nil, repository.ErrorUnsupportedFilterOption
	}
}

func applyPgSortOption(q *gorm.DB, so repository.SortOption) (*gorm.DB, error) {
	switch x := (so).(type) {
	case nil:
		return q, nil

	case *repository.CustomFieldSortOption:
		return q.Order(clause.OrderByColumn{
			Column: clause.Column{Name: x.Field},
			Desc:   x.Desc,
		}), nil

	case *repository.IdSortOption:
		return q.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "id"},
			Desc:   x.Desc,
		}), nil

	case *repository.CreatedAtSortOption:
		return q.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "created_at"},
			Desc:   x.Desc,
		}), nil

	case *repository.UpdatedAtSortOption:
		return q.Order(clause.OrderByColumn{
			Column: clause.Column{Name: "updated_at"},
			Desc:   x.Desc,
		}), nil

	default:
		return nil, repository.ErrorUnsupportedSortOption
	}
}
