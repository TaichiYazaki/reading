package main

import (
	"log"
	"os"
	"reading/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB接続
	db.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")
	router.Static("/img", "./img")

	router.GET("/", func(ctx *gin.Context) {
		readings := db.DbFindAll()
		ctx.HTML(200, "index.html", gin.H{
			"readings": readings,
		})
	})
	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{})
	})
	router.GET("/item", func(ctx *gin.Context) {
		ctx.HTML(200, "item.html", gin.H{
			"ok": "success",
		})
	})
	router.POST("/create", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		price := ctx.PostForm("price")
		review := ctx.PostForm("review")
		file := ctx.PostForm("file")
		impression := ctx.PostForm("impression")
		db.DbCreate(title, price, review, file, impression)
		ctx.Redirect(302, "/")
		filePath := "/Users/YAZAKITAICHI/env/go_env/reading/img/" + file
		_, err := os.Create(filePath)
		if err != nil {
			log.Fatal("ファイル確保の失敗")
		}
		ctx.Redirect(302, "/")
	})
	router.Run(":8080")
}
