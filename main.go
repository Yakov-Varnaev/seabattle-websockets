package main

import (
	"fmt"
	"net/http"

	"github.com/Yakov-Varnaev/seabattle-websockets/client"
	"github.com/gin-gonic/gin"
)

var field [3][3]string = [3][3]string{}

func CreateRouter() *gin.Engine {
	hub := client.NewHub()
	go hub.Run()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", func(c *gin.Context) {
		client.ServeWs(hub, c.Writer, c.Request)
	})
	return r
}

func main() {
	fmt.Println("We are here")
	r := CreateRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
