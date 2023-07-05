package rest

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/otaxhu/api-rest-golang/internal/models"
	"github.com/otaxhu/api-rest-golang/internal/service"
)

type ginMovieHandler struct {
	movieService service.MovieService
}

func newMovieHandler(movieService service.MovieService) *ginMovieHandler {
	return &ginMovieHandler{
		movieService,
	}
}

func (mh *ginMovieHandler) PostMovie(c *gin.Context) {
	title, exists := c.GetPostForm("title")
	if !exists || strings.TrimSpace(title) == "" {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("the title field is empty"))
	}
	date, err := time.Parse("2006-01-02", c.PostForm("date"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	cover, err := c.FormFile("cover")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newFilename := uuid.NewString() + filepath.Ext(cover.Filename)

	cover.Filename = newFilename

	coverUrl := fmt.Sprintf("http://%s/static/covers/%s", c.Request.Host, newFilename)

	path, err := filepath.Abs("./static/covers")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := c.SaveUploadedFile(cover, fmt.Sprintf("%s/%s", path, newFilename)); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := mh.movieService.SaveMovie(c, models.Movie{
		Title:    title,
		Date:     date,
		CoverUrl: coverUrl,
	}); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
