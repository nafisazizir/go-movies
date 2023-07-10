package main

import (
	"explore/go-movies/api/movies"
	"explore/go-movies/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to Movie API!")
}

func main() {
	// Instantiates the database
	postgres, err := database.NewPostgres(os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiate the author service
	queries := database.New(postgres.DB)
	movieService := movies.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()
	movieService.RegisterHandlers(router)
	router.GET("/", getHome)

	// Start the server
	router.Run()
}
