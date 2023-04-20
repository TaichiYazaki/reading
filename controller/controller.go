package controller

import (
	"reading/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	readings := repository.DbFindAll()
	ctx.HTML(200, "index.html", gin.H{
		"readings": readings,
	})
}

func GetCreate(ctx *gin.Context) {
	ctx.HTML(200, "create.html", gin.H{})
}

func GetItem(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := repository.DbFindOne(id)
	ctx.HTML(200, "item.html", gin.H{
		"reading": reading,
	})
}

func GetEdit(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := repository.DbFindOne(id)
	ctx.HTML(200, "edit.html", gin.H{
		"reading": reading,
	})
}
func PostCreate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	price := ctx.PostForm("price")
	review := ctx.PostForm("review")
	file, _ := ctx.FormFile("file")
	impression := ctx.PostForm("impression")
	repository.DbCreate(title, price, review, file.Filename, impression)
	filePath := "/Users/YAZAKITAICHI/env/go_env/reading/img/" + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	ctx.Redirect(302, "/")
}

func Delete(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	repository.DbDelete(id)
	ctx.Redirect(302, "/")
}
