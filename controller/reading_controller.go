package controller

import (
	"log"
	"reading/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

const imgPath = "C:\\Users\\tyazaki\\Documents\\env\\golang\\reading\\img\\"

func ReadingIndexPage(ctx *gin.Context) {
	readings, err := repository.DbFindAll()
	if err != nil{
		log.Fatal(err)
	}
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
	reading, err := repository.DbFindOne(id)
	if err != nil{
		log.Fatal(err)
	}
	ctx.HTML(200, "item.html", gin.H{
		"reading": reading,
	})
}

func ReadingEditPage(ctx *gin.Context) {
	var id int
	getId := ctx.Param("id")
	id, _ = strconv.Atoi(getId)
	reading, err := repository.DbFindOne(id)
	if err != nil {
		log.Fatal(err)
	}
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
	err:= repository.DbCreate(title, price, review, file.Filename, impression)
	if err !=nil {
		log.Fatal(err)
	}
	filePath := imgPath + file.Filename
	ctx.SaveUploadedFile(file, filePath)
	ctx.Redirect(302, "/")
}

func ReadingDelete(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	repository.DbDelete(id)
	ctx.Redirect(302, "/")
}
