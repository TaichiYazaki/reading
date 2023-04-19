package main

import (
	"reading/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB接続
	db.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/styles", "./styles")
	router.Static("/img", "./img")
	// マルチパートフォームが利用できるメモリの制限を設定する(デフォルトは 32 MiB)
	router.MaxMultipartMemory = 8 << 20

	router.GET("/", func(ctx *gin.Context) {
		readings := db.DbFindAll()
		ctx.HTML(200, "index.html", gin.H{
			"readings": readings,
		})
	})
	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{})
	})
	router.GET("/item/:id", func(ctx *gin.Context) {
		var id int
		getId := ctx.Param("id")
		id, _ = strconv.Atoi(getId)
		reading := db.DbFindOne(id)
		ctx.HTML(200, "item.html", gin.H{
			"reading": reading,
		})
	})
	router.GET("/edit/:id", func(ctx *gin.Context) {
		var id int
		getId := ctx.Param("id")
		id, _ = strconv.Atoi(getId)
		reading := db.DbFindOne(id)
		ctx.HTML(200, "edit.html", gin.H{
			"reading": reading,
		})
	})
	router.POST("/create", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		price := ctx.PostForm("price")
		review := ctx.PostForm("review")
		file, _ := ctx.FormFile("file")
		impression := ctx.PostForm("impression")
		db.DbCreate(title, price, review, file.Filename, impression)
		filePath := "/Users/YAZAKITAICHI/env/go_env/reading/img/" + file.Filename
		ctx.SaveUploadedFile(file, filePath)
		ctx.Redirect(302, "/")
	})
	router.POST("/update/:id", func(ctx *gin.Context) {
		getId := ctx.Param("id")
		id, _ := strconv.Atoi(getId)
		title := ctx.PostForm("title")
		price := ctx.PostForm("price")
		review := ctx.PostForm("review")
		impression := ctx.PostForm("impression")
		db.DbUpdate(id, title, price, review, impression)
		ctx.Redirect(302, "/")
	})
	router.GET("/delete/:id", func(ctx *gin.Context) {
		getId := ctx.Param("id")
		id, _ := strconv.Atoi(getId)
		db.DbDelete(id)
		ctx.Redirect(302, "/")
	})
	router.Run(":8080")
}
