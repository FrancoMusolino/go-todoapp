package db

import (
	"github.com/FrancoMusolino/go-todoapp/utils/pagination"
	"gorm.io/gorm"
)

func Paginate(params *pagination.PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.PageNumber - 1) * params.PageSize
		return db.Offset(int(offset)).Limit(int(params.PageSize))
	}
}
