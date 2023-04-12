package main

import (
	"fmt"
	"log"
	"reading/db"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB接続確認
	database, err := db.DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("接続に成功しました")
	database.AutoMigrate(&db.Reading{})
	fmt.Println("マイグレーションに成功しました")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.POST("/create", func(ctx *gin.Context) {
		ctx.HTML(200, "create.html", gin.H{})
	})
	router.Run(":8080")
}
