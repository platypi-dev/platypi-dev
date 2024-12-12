package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"platypi.dev/release-scraper/controllers"
	"platypi.dev/release-scraper/models"
)

const (
	YYYYMMDD       = "2006-01-01"
	YYYYMMDDhhmmss = "2006-01-02 15:04:05"
)

// var repos = []repo{
// 	{ID: "1", Org: "argo-proj", Repo: "argo-cd"},
// 	{ID: "2", Org: "gruntwork-io", Repo: "boilerplate"},
// }

// func getRepos(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, repos)
// }

// a basic ping function. Could be used for startup
// moved to models/setup.go
// func setupRouter() *gin.Engine {
// 	r := gin.Default()
// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})

// 	return r
// }

func main() {
	// 	We're going to serve the time in different ways. Here are a couple.
	currentTimestamp := time.Now().Local()
	currentShortdate := currentTimestamp.Format(YYYYMMDD)
	currentLongdate := currentTimestamp.Format(YYYYMMDDhhmmss)
	fmt.Println(currentShortdate)
	fmt.Println(currentLongdate)

	// initialize gin
	r := gin.Default()
	// Connect to the db
	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// repo functions
	r.GET("/repos", controllers.FindRepos)
	r.POST("/repo", controllers.AddRepo)
	r.GET("/repo/:id", controllers.FindRepo)
	r.PATCH("/repo/:id", controllers.UpdateRepo)
	r.DELETE("/repo/:id", controllers.DeleteRepo)

	r.Run()
}
