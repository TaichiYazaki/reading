package main

import (
	"reading/controller"
	"reading/repository"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	repository.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")
	router.Static("/img", "./img")
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20
	router.GET("/", controller.GetIndex)
	router.GET("/create", controller.GetCreate)
	router.GET("/item/:id", controller.GetItem)
	router.GET("/edit/:id", controller.GetEdit)
	router.GET("/delete/:id", controller.Delete)
	router.POST("/create", controller.PostCreate)
	return router
}

func main() {
	r := router()
	r.Run()
}
