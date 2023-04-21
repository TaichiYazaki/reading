package main

import (
	"reading/controller"
	"reading/repository"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	repository.ReadingDbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")
	router.Static("/img", "./img")
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20
	router.GET("/", controller.ReadingIndexPage)
	router.GET("/create", controller.ReadingCreatePage)
	router.GET("/item/:id", controller.ReadingItemPage)
	router.GET("/edit/:id", controller.ReadingEditPage)
	router.GET("/delete/:id", controller.ReadingDelete)
	router.POST("/create", controller.ReadingCreate)

	user := router.Group("/user")
	{
		user.GET("/signup", controller.UserSignUpPage)
	}
	return router
}

func main() {
	r := router()
	r.Run()
}
