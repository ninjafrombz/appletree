// Filename = internals/data/filters.go

package data

import "appletree.desireamagwula.net/internals/validator"

type Filters struct {
	Page int
	PageSize int
	Sort string
	SortList []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	// Check page and pagesize parameters
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 1000, "page", "must be a maximum of 1000")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
	// Check that the sort parameter matches a value in the sort list
	v.Check(validator.In(f.Sort, f.SortList...), "sort", "invalid sort value")
}