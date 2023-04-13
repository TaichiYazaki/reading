package main

import (
	"reading/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB接続
	db.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{})
	})
	router.POST("/create", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		price := ctx.PostForm("price")
		review := ctx.PostForm("review")
		impression := ctx.PostForm("impression")
		db.DbCreate(title, price, review, impression)
		ctx.Redirect(302, "/")
	})
	router.Run(":8080")
}
