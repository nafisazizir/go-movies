package movies

import (
	"context"
	"explore/go-movies/internal/database"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	queries *database.Queries
}

func NewService(queries *database.Queries) *Service {
	return &Service{queries: queries}
}

func (s *Service) RegisterHandlers(router *gin.Engine) {
	router.GET("/movies", s.Get)
}

type apiMovie struct {
	ID          int32
	Title       string
	Description string
	ReleaseDate time.Time
	PosterUrl   string
	AgeRating   string
	TicketPrice int32
}

func (s *Service) Get(c *gin.Context) {
	movies, err := s.queries.GetMovies(context.Background())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	if len(movies) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}
