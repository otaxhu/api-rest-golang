package models

import "time"

type Movie struct {
	Id       int
	Title    string
	Date     time.Time
	CoverUrl string
}
