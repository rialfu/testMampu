package helpers

import (
	"fmt"
	"rialfu/wallet/pkg/constants"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Pagination struct {
	Page  int
	Limit int
	Sort  string
	Order string
}

func (p *Pagination) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "asc"
	}
}

func GetModelColumns(db *gorm.DB, model any) map[string]bool {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model)

	cols := make(map[string]bool)
	for _, field := range stmt.Schema.Fields {
		cols[field.DBName] = true
	}
	return cols
}

func ApplyPagination(
	db *gorm.DB,
	model any,
	queryParams map[string][]string,
) (*gorm.DB, Pagination) {

	p := Pagination{
		Page:  parseInt(queryParams["page"], constants.ENUM_PAGINATION_PAGE),
		Limit: parseInt(queryParams["limit"], constants.ENUM_PAGINATION_PER_PAGE),
		Sort:  parseString(queryParams["sort"]),
		Order: parseString(queryParams["order"]),
	}
	p.Normalize()

	validColumns := GetModelColumns(db, model)

	// ===== Sorting =====
	if p.Sort != "" && validColumns[p.Sort] {
		db = db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: p.Sort},
			Desc:   p.Order == "desc",
		})
	}

	// ===== Search / Filter =====
	for key, values := range queryParams {
		if !validColumns[key] {
			continue
		}
		if len(values) == 0 || values[0] == "" {
			continue
		}

		db = db.Where(
			fmt.Sprintf("%s ILIKE ?", key),
			"%"+values[0]+"%",
		)
	}

	offset := (p.Page - 1) * p.Limit
	db = db.Offset(offset).Limit(p.Limit)

	return db, p
}

type PaginateData[T any] struct {
	Data  []T   `json:"data"`
	Limit int   `json:"limit"`
	Page  int   `json:"page"`
	Total int64 `json:"total"`
}
