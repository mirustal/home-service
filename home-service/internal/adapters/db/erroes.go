package dbErr

import "errors"

var (
	ErrHouseNotFound = errors.New("house not found")
	ErrFlatNotFound  = errors.New("flat not found")
)