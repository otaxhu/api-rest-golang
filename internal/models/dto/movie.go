package dto

import (
	"mime/multipart"
	"time"
)

// This is a representation of the movie in the service layer when you want to save it.
// The reason of this is because I want the gin framework decoupled from my business logic.
// So I want the Cover file decoupled of it and save it by myself.
type SaveMovie struct {
	Title string                `validate:"required"`
	Date  time.Time             `validate:"required"`
	Cover *multipart.FileHeader `validate:"required"`
}

type GetMovie struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	CoverUrl string    `json:"cover_url"`
}

type UpdateMovie struct {
	Id    int
	Title string
	Date  time.Time
	Cover *multipart.FileHeader
}
