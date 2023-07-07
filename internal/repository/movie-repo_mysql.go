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

const (
	// Do not use this query how it is, you have to concatenate it with another query.
	// where you want the id or anything like that
	qrySelectMovie = "SELECT id, title, date, cover_url FROM movies "
	qryInsertMovie = "INSERT INTO movies (title, date, cover_url) VALUES (?, ?, ?)"
	qryDeleteMovie = "DELETE FROM movies WHERE id = ?"
)

func (repo *mysqlMovieRepo) InsertMovie(ctx context.Context, movie models.Movie) error {
	_, err := repo.db.ExecContext(ctx, qryInsertMovie, movie.Title, movie.Date, movie.CoverUrl)
	return err
}

func (repo *mysqlMovieRepo) GetMovies(ctx context.Context, limit, offset uint) ([]models.Movie, error) {
	rows, err := repo.db.QueryContext(ctx, qrySelectMovie+"LIMIT ? OFFSET ?", limit, offset)
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

func (repo *mysqlMovieRepo) GetMovieById(ctx context.Context, id int) (models.Movie, error) {
	movie := models.Movie{}
	var movieDate string
	var err error
	if err := repo.db.QueryRowContext(ctx, qrySelectMovie+"WHERE id = ?", id).
		Scan(&movie.Id, &movie.Title, &movieDate, &movie.CoverUrl); err == sql.ErrNoRows {
		return movie, ErrNoRows
	} else if err != nil {
		return movie, err
	}
	movie.Date, err = time.Parse("2006-01-02", movieDate)
	return movie, err
}

func (repo mysqlMovieRepo) DeleteMovieById(ctx context.Context, id int) error {
	result, err := repo.db.ExecContext(ctx, qryDeleteMovie, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNoRows
	}
	return nil
}
