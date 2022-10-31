package errors

import (
	"errors"
)

var (
	NotFound404 = errors.New("Record not found.")
	BadRequest400 = errors.New("Bad request.")
)
