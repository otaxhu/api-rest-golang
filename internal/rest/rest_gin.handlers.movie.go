package rest

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/otaxhu/api-rest-golang/internal/models/dto"
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
	date, _ := time.Parse("2006-01-02", c.PostForm("date"))
	cover, _ := c.FormFile("cover")
	if cover != nil {
		cover.Filename = uuid.NewString() + filepath.Ext(cover.Filename)
		cover.Header.Set("cover_url", fmt.Sprintf("http://%s/static/covers/%s", c.Request.Host, cover.Filename))
	}
	dtoSaveMovie := dto.SaveMovie{
		Title: c.PostForm("title"),
		Date:  date,
		Cover: cover,
	}
	if err := mh.movieService.SaveMovie(c, dtoSaveMovie); err == service.ErrInvalidMovieObject {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else if err == service.ErrInternalServer {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (mh *ginMovieHandler) GetMovies(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	movies, err := mh.movieService.GetMovies(c, page)
	if err == service.ErrInvalidParams {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else if err == service.ErrNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err == service.ErrInternalServer {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (mh *ginMovieHandler) GetMovieById(c *gin.Context) {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": service.ErrNotFound})
		return
	}
	movie, err := mh.movieService.GetMovieById(c, idParam)
	if err == service.ErrNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else if err == service.ErrInternalServer {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
