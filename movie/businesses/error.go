package businesses

import "errors"

var (
	ErrDuplicateData  = errors.New("duplicate data found")
	ErrInternalServer = errors.New("internal Server Error")
	ErrUserIdNotFound = errors.New("user id not found")
)
