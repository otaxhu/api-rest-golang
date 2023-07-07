package service

import "fmt"

var (
	ErrInvalidMovieObject = fmt.Errorf("invalid movie object")
	ErrSavingMovie        = fmt.Errorf("internal server error during saving movie")
	ErrInvalidParams      = fmt.Errorf("invalid params")
	ErrNotFound           = fmt.Errorf("not found")
	ErrInternalServer     = fmt.Errorf("internal server error")
	ErrDeletingMovie      = fmt.Errorf("internal server error during deleting movie")
)
