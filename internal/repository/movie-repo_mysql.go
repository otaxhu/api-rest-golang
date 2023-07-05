package repository

import (
	"context"
	"database/sql"

	"github.com/otaxhu/api-rest-golang/internal/models"
)

type mysqlMovieRepo struct {
	db *sql.DB
}

func newMysqlMovieRepo(db *sql.DB) MovieRepository {
	return &mysqlMovieRepo{
		db,
	}
}

func (repo *mysqlMovieRepo) InsertMovie(ctx context.Context, movie models.Movie) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO movies (title, date, cover_url) VALUES (?, ?, ?)", movie.Title, movie.Date, movie.CoverUrl)
	return err
}
