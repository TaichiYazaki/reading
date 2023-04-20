package main

import (
	"reading/db"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	db.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")
	router.Static("/img", "./img")
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20
	router.GET("/", getIndex)
	router.GET("/create", getCreate)
	router.GET("/item/:id", getItem)
	router.GET("/edit/:id", getEdit)
	router.GET("/delete/:id", delete)
	router.POST("/create", postCreate)
	return router
}

func main() {
	r := router()
	r.Run()
}
