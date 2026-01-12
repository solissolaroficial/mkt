package errors

import (
	"errors"
)

var (
	ErrOfflineActionNotFound      = errors.New("offline action not found")
	ErrShowroomItemNotFound       = errors.New("showroom item not found")
	ErrRepMarketingActionNotFound = errors.New("rep marketing action not found")
)
