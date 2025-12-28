package valueobject

// Pagination represents pagination parameters for queries
type Pagination struct {
	Page     int
	PageSize int
}

// NewPagination creates a new Pagination instance with validation
// Page minimum is 1, PageSize must be between 1 and 100 (default 10)
func NewPagination(page, pageSize int) Pagination {
	// Validate and set default values
	if page < 1 {
		page = 1
	}

	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	return Pagination{
		Page:     page,
		PageSize: pageSize,
	}
}

// Offset returns the database offset for pagination (Page - 1) * PageSize
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit returns the page size for database queries
func (p Pagination) Limit() int {
	return p.PageSize
}
