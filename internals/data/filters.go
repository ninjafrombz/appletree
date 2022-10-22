// Filename = internals/data/filters.go

package data

import (
	"strings"

	"appletree.desireamagwula.net/internals/validator"
)

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

// The sort column method safely extracts the sort field query parameter 
func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// Get the sort order method determines whether we should sort by descending or ascending. 
func (f Filters) sortOrder() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"

	}
	return "ASC"
}

// The limit method determines the limit 
func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offSet() int {
	return (f.Page - 1) * f.PageSize
}