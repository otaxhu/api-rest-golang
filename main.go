package main

import (
	"log"

	"github.com/otaxhu/api-rest-golang/internal/repository"
	"github.com/otaxhu/api-rest-golang/internal/rest"
	"github.com/otaxhu/api-rest-golang/internal/service"
	"github.com/otaxhu/api-rest-golang/settings"
)

func main() {
	dbSettings, err := settings.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	serverSettings, err := settings.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	movieRepo, err := repository.NewMovieRepository(dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	movieService := service.NewMovieService(movieRepo)
	app, err := rest.New(serverSettings, movieService)
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
