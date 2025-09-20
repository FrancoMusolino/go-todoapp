package pagination

import (
	"math"
)

type PaginationMetadata struct {
	PageSize     uint `json:"pageSize"`
	PageNumber   uint `json:"pageNumber"`
	TotalPages   uint `json:"totalPages"`
	TotalRecords uint `json:"totalRecords"`
}

func NewPaginationMetadata(pageSize, pageNumber, totalRecords uint) *PaginationMetadata {
	return &PaginationMetadata{
		PageSize:     pageSize,
		PageNumber:   pageNumber,
		TotalRecords: totalRecords,
		TotalPages:   uint(math.Ceil(float64(totalRecords / pageSize))),
	}
}

type PaginationParams struct {
	PageSize   uint
	PageNumber uint
}

func NewPaginationParams(pageSize, pageNumber uint) *PaginationParams {
	if pageNumber <= 0 {
		pageNumber = 1
	}

	switch {
	case pageSize > 50:
		pageSize = 50
	case pageSize <= 0:
		pageSize = 10
	}

	return &PaginationParams{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}
}
