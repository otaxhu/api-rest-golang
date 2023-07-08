package service

import "fmt"

var (
	ErrInvalidMovieObject = fmt.Errorf("invalid movie object")
	ErrInvalidParams      = fmt.Errorf("invalid params")
	ErrNotFound           = fmt.Errorf("not found")
	ErrInternalServer     = fmt.Errorf("internal server error")
	ErrNoEntries          = fmt.Errorf("no entries in the resource")
)
