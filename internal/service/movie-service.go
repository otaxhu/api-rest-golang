package service

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/otaxhu/api-rest-golang/internal/models"
	"github.com/otaxhu/api-rest-golang/internal/models/dto"
	"github.com/otaxhu/api-rest-golang/internal/repository"
)

type MovieService interface {
	SaveMovie(ctx context.Context, movie dto.SaveMovie) error
	GetMovies(ctx context.Context, page int) ([]dto.GetMovies, error)
	GetMovieById(ctx context.Context, id int) (dto.GetMovies, error)
}

type movieServiceImpl struct {
	validator *validator.Validate
	repo      repository.MovieRepository
}

func NewMovieService(movieRepo repository.MovieRepository) MovieService {
	return &movieServiceImpl{
		repo:      movieRepo,
		validator: validator.New(),
	}
}

func (service *movieServiceImpl) SaveMovie(ctx context.Context, movieDto dto.SaveMovie) error {
	if err := service.validator.StructCtx(ctx, movieDto); err != nil {
		return ErrInvalidMovieObject
	}

	movie := models.Movie{
		Title:    movieDto.Title,
		Date:     movieDto.Date,
		CoverUrl: movieDto.Cover.Header.Get("cover_url"),
	}

	if err := service.repo.InsertMovie(ctx, movie); err != nil {
		return ErrSavingMovie
	}

	file, err := movieDto.Cover.Open()
	if err != nil {
		return ErrInternalServer
	}
	defer file.Close()

	path, err := filepath.Abs("./static/covers")
	if err != nil {
		return ErrInternalServer
	}

	out, err := os.Create(filepath.Join(path, movieDto.Cover.Filename))
	if err != nil {
		return ErrInternalServer
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		return ErrInternalServer
	}
	return nil
}

func (service *movieServiceImpl) GetMovies(ctx context.Context, page int) ([]dto.GetMovies, error) {
	if page <= 0 {
		return nil, ErrInvalidParams
	}
	page--
	const limit = 5
	offset := limit * page

	movies, err := service.repo.GetMovies(ctx, limit, uint(offset))
	if err == repository.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, ErrInternalServer
	}

	dtoMovies := []dto.GetMovies{}
	for _, movie := range movies {
		dtoMovies = append(dtoMovies, dto.GetMovies{
			Id:       movie.Id,
			Title:    movie.Title,
			Date:     movie.Date,
			CoverUrl: movie.CoverUrl,
		})
	}
	return dtoMovies, nil
}

func (service *movieServiceImpl) GetMovieById(ctx context.Context, id int) (dto.GetMovies, error) {
	movie, err := service.repo.GetMovieById(ctx, id)
	if err == repository.ErrNoRows {
		return dto.GetMovies{}, ErrNotFound
	} else if err != nil {
		return dto.GetMovies{}, ErrInternalServer
	}
	return dto.GetMovies{
		Id:       movie.Id,
		Title:    movie.Title,
		Date:     movie.Date,
		CoverUrl: movie.CoverUrl,
	}, nil
}
