package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/game", createGame)
	r.GET("/game/:gameId/player/:playerId", retrieveGame)
	return r
}

func main() {
	r := CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
