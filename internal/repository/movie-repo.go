package repository

import (
	"context"
	"fmt"

	"github.com/otaxhu/api-rest-golang/database"
	"github.com/otaxhu/api-rest-golang/internal/models"
	"github.com/otaxhu/api-rest-golang/settings"
)

type MovieRepository interface {
	InsertMovie(ctx context.Context, movie models.Movie) error
}

func NewMovieRepository(dbSettings *settings.Database) (MovieRepository, error) {
	switch dbSettings.Driver {
	case "mysql":
		conn, err := database.NewSqlConnection(dbSettings)
		if err != nil {
			return nil, err
		}
		return newMysqlMovieRepo(conn), nil
	default:
		return nil, fmt.Errorf("movie-repo.go\nNew(): the \"%s\" driver does not have a repository implementation", dbSettings.Driver)
	}
}
