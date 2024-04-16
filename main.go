package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var field [3][3]string = [3][3]string{}

func CreateRouter() *gin.Engine {
	hub := newHub()
	go hub.run()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	return r
}

func main() {
	r := CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
