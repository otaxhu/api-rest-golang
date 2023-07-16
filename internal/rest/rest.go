package rest

import (
	"fmt"

	"github.com/otaxhu/api-rest-golang/internal/service"
	"github.com/otaxhu/api-rest-golang/settings"
)

type RestApp interface {
	Start() error
	bindRoutes() error
}

func New(serverSettings *settings.Server, movieService service.MovieService) (RestApp, error) {
	switch serverSettings.Framework {
	case "gin":
		return newGinApp(serverSettings, movieService)
	default:
		return nil, fmt.Errorf("rest.go\nNew(): the %s framework does not have a RestApp implementation", serverSettings.Framework)
	}
}
