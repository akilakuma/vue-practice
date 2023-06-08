package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"hotel-server/controller"
	"io"
	"net/http"
)

func getRoom(c *gin.Context) {
	c.JSON(http.StatusOK, controller.GetRoomList())
}

func getRestaurant(c *gin.Context) {
	c.JSON(http.StatusOK, controller.GetRestaurantList())
}

func userLogin(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("password")
	c.JSON(http.StatusOK, controller.UserLogin(email, pwd))
}

func register(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}
	email := gjson.GetBytes(jsonData, "email").String()
	pwd := gjson.GetBytes(jsonData, "password").String()
	controller.RegisterUser(email, pwd)

	c.Status(http.StatusOK)
}
