package service

import (
	"context"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/otaxhu/api-rest-golang/internal/models"
	"github.com/otaxhu/api-rest-golang/internal/models/dto"
	"github.com/otaxhu/api-rest-golang/internal/repository"
)

type MovieService interface {
	SaveMovie(ctx context.Context, movie dto.SaveMovie) error
	GetMovies(ctx context.Context, page int) ([]dto.GetMovie, error)
	GetMovieById(ctx context.Context, id int) (dto.GetMovie, error)
	DeleteMovie(ctx context.Context, id int) error
	UpdateMovie(ctx context.Context, movie dto.UpdateMovie) error
}

type movieServiceImpl struct {
	validator  *validator.Validate
	movieRepo  repository.MovieRepository
	coversRepo repository.CoversRepository
}

func NewMovieService(movieRepo repository.MovieRepository, coversRepo repository.CoversRepository) MovieService {
	return &movieServiceImpl{
		movieRepo:  movieRepo,
		validator:  validator.New(),
		coversRepo: coversRepo,
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

	tx, err := service.movieRepo.InsertMovie(ctx, movie)
	if err != nil {
		return ErrInternalServer
	}
	if err := service.coversRepo.SaveMovieCover(movieDto.Cover); err != nil {
		tx.Rollback()
		return ErrInternalServer
	}
	return tx.Commit()
}

func (service *movieServiceImpl) GetMovies(ctx context.Context, page int) ([]dto.GetMovie, error) {
	if page <= 0 {
		return nil, ErrInvalidParams
	}
	page--
	const limit = 5
	offset := limit * page

	movies, err := service.movieRepo.GetMovies(ctx, limit, uint(offset))
	if err == repository.ErrNoRows {
		return nil, ErrNoEntries
	} else if err != nil {
		return nil, ErrInternalServer
	}

	dtoMovies := []dto.GetMovie{}
	for _, movie := range movies {
		dtoMovies = append(dtoMovies, dto.GetMovie{
			Id:       movie.Id,
			Title:    movie.Title,
			Date:     movie.Date,
			CoverUrl: movie.CoverUrl,
		})
	}
	return dtoMovies, nil
}

func (service *movieServiceImpl) GetMovieById(ctx context.Context, id int) (dto.GetMovie, error) {
	movie, err := service.movieRepo.GetMovieById(ctx, id)
	if err == repository.ErrNoRows {
		return dto.GetMovie{}, ErrNotFound
	} else if err != nil {
		return dto.GetMovie{}, ErrInternalServer
	}
	return dto.GetMovie{
		Id:       movie.Id,
		Title:    movie.Title,
		Date:     movie.Date,
		CoverUrl: movie.CoverUrl,
	}, nil
}

func (service *movieServiceImpl) DeleteMovie(ctx context.Context, id int) error {
	dbMovie, err := service.GetMovieById(ctx, id)
	if err != nil {
		return err
	}
	tx, err := service.movieRepo.DeleteMovie(ctx, id)
	if err == repository.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return ErrInternalServer
	}
	if err := service.coversRepo.DeleteCover(dbMovie.CoverUrl); err != nil {
		tx.Rollback()
		return ErrInternalServer
	}
	return tx.Commit()
}

func (service *movieServiceImpl) UpdateMovie(ctx context.Context, movieDto dto.UpdateMovie) error {
	dbMovie, err := service.GetMovieById(ctx, movieDto.Id)
	if err != nil {
		return err
	}
	coverUrl := ""
	if movieDto.Cover != nil {
		coverUrl = movieDto.Cover.Header.Get("cover_url")
	}
	movie := models.Movie{
		Id:       movieDto.Id,
		Title:    movieDto.Title,
		Date:     movieDto.Date,
		CoverUrl: coverUrl,
	}
	if strings.TrimSpace(movie.Title) == "" {
		movie.Title = dbMovie.Title
	}
	var dateZeroValue time.Time
	if movie.Date == dateZeroValue {
		movie.Date = dbMovie.Date
	}
	if movie.CoverUrl == "" {
		movie.CoverUrl = dbMovie.CoverUrl
	}
	tx, err := service.movieRepo.UpdateMovie(ctx, movie)
	if err == repository.ErrNoRows {
		return ErrNotFound
	} else if err != nil {
		return ErrInternalServer
	}
	if movieDto.Cover == nil {
		return tx.Commit()
	}
	if err := service.coversRepo.DeleteCover(dbMovie.CoverUrl); err != nil {
		tx.Rollback()
		return ErrInternalServer
	}
	if err := service.coversRepo.SaveMovieCover(movieDto.Cover); err != nil {
		tx.Rollback()
		return ErrInternalServer
	}
	return tx.Commit()
}
