package rest

// Funcion que inicializa las rutas con handlers definidos en el struct ginMovieHandler.
// Es responsabilidad del programador inicializar las rutas con sus respectivos handlers.
func (ga *ginApp) bindRoutes() error {

	movieRoutes := ga.app.Group("/movies")

	movieRoutes.POST("", ga.movieHandler.PostMovie)
	movieRoutes.GET("", ga.movieHandler.GetMovies)
	movieRoutes.GET("/:id", ga.movieHandler.GetMovieById)
	movieRoutes.DELETE("/:id", ga.movieHandler.DeleteMovie)
	movieRoutes.PUT("/:id", ga.movieHandler.PutMovie)

	ga.app.Static("/static", "./static")
	return nil
}
