package ports

import (
	"errors"
)

var (
	ErrDuplicatedKey  = errors.New("duplicated key not allowed")
	ErrRecordNotFound = errors.New("record not found")
)
