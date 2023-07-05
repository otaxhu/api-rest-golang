package rest

// Funcion que inicializa las rutas con handlers definidos en el struct ginMovieHandler.
// Es responsabilidad del programador inicializar las rutas con sus respectivos handlers.
func (ga *ginApp) bindRoutes() {
	ga.app.POST("/movies", ga.movieHandler.PostMovie)
}
