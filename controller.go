package main

import (
	"reading/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIndex(ctx *gin.Context) {
	readings := db.DbFindAll()
	ctx.HTML(200, "index.html", gin.H{
		"readings": readings,
	})
}

func getCreate(ctx *gin.Context) {
	ctx.HTML(200, "create.html", gin.H{})
}

func getItem(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := db.DbFindOne(id)
	ctx.HTML(200, "item.html", gin.H{
		"reading": reading,
	})
}

func getEdit(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := db.DbFindOne(id)
	ctx.HTML(200, "edit.html", gin.H{
		"reading": reading,
	})
}
func postCreate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	price := ctx.PostForm("price")
	review := ctx.PostForm("review")
	file, _ := ctx.FormFile("file")
	impression := ctx.PostForm("impression")
	db.DbCreate(title, price, review, file.Filename, impression)
	filePath := "/Users/YAZAKITAICHI/env/go_env/reading/img/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	ctx.Redirect(302, "/")
}

func delete(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	db.DbDelete(id)
	ctx.Redirect(302, "/")
}
