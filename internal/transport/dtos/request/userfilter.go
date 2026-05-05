package request

import "time"

type UserFilter struct {
	Search    string
	FromDate  *time.Time
	ToDate    *time.Time
	Page      int
	PageSize  int
	SortBy    string
	SortOrder string
}
