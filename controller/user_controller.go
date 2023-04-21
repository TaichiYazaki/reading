package controller

import (
	"log"
	"reading/model"
	"reading/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUpPage(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", gin.H{})
}

func UserCreate(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	h, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash := string(h)
	form := model.User{Name: name, Email: email, Password: password}
	if err := ctx.ShouldBind(&form); err != nil {
		log.Println("-----", err)
		ctx.HTML(200, "signup.html", gin.H{
			"err": "空欄を埋めてください",
		})
		return
	}
	repository.UserCreate(name, email, hash)
	ctx.Redirect(302, "/")

}
