package controller

import "github.com/gin-gonic/gin"

func UserSignUpPage(ctx *gin.Context) {
	ctx.HTML(200, "signup.html", gin.H{})
}
