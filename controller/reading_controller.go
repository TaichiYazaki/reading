package controller

import (
	"fmt"
	"reading/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ReadingIndexPage(ctx *gin.Context) {
	readings := repository.DbFindAll()
	fmt.Println(readings)
	ctx.HTML(200, "index.html", gin.H{
		"readings": readings,
	})
}

func ReadingCreatePage(ctx *gin.Context) {
	ctx.HTML(200, "create.html", gin.H{})
}

func ReadingItemPage(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := repository.DbFindOne(id)
	ctx.HTML(200, "item.html", gin.H{
		"reading": reading,
	})
}

func ReadingEditPage(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading := repository.DbFindOne(id)
	ctx.HTML(200, "edit.html", gin.H{
		"reading": reading,
	})
}
func ReadingCreate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	price := ctx.PostForm("price")
	review := ctx.PostForm("review")
	file, _ := ctx.FormFile("file")
	impression := ctx.PostForm("impression")
	repository.DbCreate(title, price, review, file.Filename, impression)
	filePath := "C:\\Users\\tyazaki\\Documents\\env\\golang\\reading\\img\\" + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	ctx.Redirect(302, "/")
}

func ReadingDelete(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	repository.DbDelete(id)
	ctx.Redirect(302, "/")
}
