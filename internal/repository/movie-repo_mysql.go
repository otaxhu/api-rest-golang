package repository

import (
	"context"
	"database/sql"
	"time"

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

func (repo *mysqlMovieRepo) GetMovies(ctx context.Context, limit, offset uint) ([]models.Movie, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, title, date, cover_url FROM movies LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	movies := []models.Movie{}
	for rows.Next() {
		movie := models.Movie{}
		var movieDate string
		if err := rows.Scan(&movie.Id, &movie.Title, &movieDate, &movie.CoverUrl); err != nil {
			continue
		}
		movie.Date, err = time.Parse("2006-01-02", movieDate)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if len(movies) == 0 {
		return nil, ErrNoRows
	}
	return movies, nil
}
