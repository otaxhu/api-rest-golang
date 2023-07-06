package service

import "fmt"

var (
	ErrInvalidMovieObject = fmt.Errorf("invalid movie object")
	ErrSavingMovie        = fmt.Errorf("internal server error during saving movie")
	ErrInvalidPageParam   = fmt.Errorf("invalid page param")
	ErrNotFound           = fmt.Errorf("not found")
	ErrInternalServer     = fmt.Errorf("internal server error")
)
