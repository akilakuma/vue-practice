package main

import (
	"hotel-server/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/room-list", func(c *gin.Context) {
		c.JSON(http.StatusOK, controller.GetRoomList())
	})
	r.GET("/restaurant-list", func(c *gin.Context) {
		c.JSON(http.StatusOK, controller.GetRestaurantList())
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
