package movie_repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/otaxhu/api-rest-golang/internal/models"
	repo_errs "github.com/otaxhu/api-rest-golang/internal/repository/errors"
)

type mysqlMovieRepo struct {
	db *sql.DB
}

func newMysqlMovieRepo(db *sql.DB) MovieRepository {
	return &mysqlMovieRepo{
		db: db,
	}
}

const (
	// Do not use this query how it is, you have to concatenate it with another query.
	// where you want the id or anything like that
	qrySelectMovie = "SELECT id, title, date, cover_url FROM movies "
	qryInsertMovie = "INSERT INTO movies (title, date, cover_url) VALUES (?, ?, ?)"
	qryDeleteMovie = "DELETE FROM movies WHERE id = ?"
	qryUpdateMovie = "UPDATE movies SET title = ?, date = ?, cover_url = ? WHERE id = ?"
)

func (repo *mysqlMovieRepo) InsertMovie(ctx context.Context, movie models.Movie) (tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	_, err = tx.ExecContext(ctx, qryInsertMovie, movie.Title, movie.Date, movie.CoverUrl)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
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
		return nil, repo_errs.ErrNoRows
	}
	return movies, nil
}

func (repo *mysqlMovieRepo) GetMovieById(ctx context.Context, id int) (models.Movie, error) {
	movie := models.Movie{}
	var movieDate string
	var err error
	if err := repo.db.QueryRowContext(ctx, qrySelectMovie+"WHERE id = ?", id).
		Scan(&movie.Id, &movie.Title, &movieDate, &movie.CoverUrl); err == sql.ErrNoRows {
		return movie, repo_errs.ErrNoRows
	} else if err != nil {
		return movie, err
	}
	movie.Date, err = time.Parse("2006-01-02", movieDate)
	return movie, err
}

func (repo *mysqlMovieRepo) DeleteMovie(ctx context.Context, id int) (tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, qryDeleteMovie, id); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}

func (repo *mysqlMovieRepo) UpdateMovie(ctx context.Context, movie models.Movie) (tx, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if _, err := tx.ExecContext(ctx, qryUpdateMovie, movie.Title, movie.Date, movie.CoverUrl, movie.Id); err != nil {
		tx.Rollback()
		return nil, err
	}
	return tx, nil
}
