package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hotel-server/model"
	"net/http"
)

// AddErrorFields 中間件函式
//func AddErrorFields() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// 先呼叫下一個處理程序取得回傳的 JSON 物件
//		c.Next()
//
//		// 只在處理成功時添加欄位
//		if c.Writer.Status() == http.StatusOK {
//			// 取得回傳的 JSON 物件
//			responseData, exists := c.Get("response")
//			if exists {
//				// 在 JSON 物件中添加欄位
//				responseData.(gin.H)["error_msg"] = ""
//				responseData.(gin.H)["error_code"] = 0
//			}
//		}
//	}
//}

func main() {

	dbErr := model.InitDB(model.EnvVariable{
		OrderDBStr: "root:qwe123@tcp(127.0.0.1:3306)/hotel?charset=utf8&parseTime=True&loc=Local",
	})
	if dbErr != nil {
		panic(dbErr.Error())
	}

	r := gin.Default()

	// 處理cross domain問題
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	// 使用middleware，幫api加上error_msg和error_code
	//r.Use(AddErrorFields())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 取得房間列表
	r.GET("/room-list", getRoom)
	// 取得餐廳列表
	r.GET("/restaurant-list", getRestaurant)

	// 會員
	// 註冊
	r.GET("/login", userLogin)
	r.POST("/register", register)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
