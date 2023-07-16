package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/otaxhu/api-rest-golang/internal/service"
	"github.com/otaxhu/api-rest-golang/settings"
)

type ginApp struct {
	serverSettings *settings.Server
	movieHandler   *ginMovieHandler
	app            *gin.Engine
}

func newGinApp(serverSettings *settings.Server, movieService service.MovieService) (*ginApp, error) {
	app := &ginApp{
		serverSettings: serverSettings,
		movieHandler:   newMovieHandler(movieService),
		app:            gin.Default(),
	}
	if err := app.bindRoutes(); err != nil {
		return nil, err
	}
	return app, nil
}

func (ga *ginApp) Start() error {
	return ga.app.Run(fmt.Sprintf(":%d", ga.serverSettings.Port))
}
