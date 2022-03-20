package controller

import (
	"github.com/gin-gonic/gin"
	"golang.com/m/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	jwtService service.JWTService
}

func NewLoginController(jwtService service.JWTService) LoginController {
	return &loginController{jwtService: jwtService}
}

func (l *loginController) Login(ctx *gin.Context) string {
	username := ctx.PostForm("username")
	token := l.jwtService.GenerateToken(username)
	return token
}