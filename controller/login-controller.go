package controller

import (
	"github.com/gin-gonic/gin"
	"golang.com/m/dto"
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
	var credentials dto.CredentialsDto
	username := ctx.ShouldBindJSON(&credentials)
	if username!= nil {
		panic(username)
	}
	token := l.jwtService.GenerateToken(credentials.Username)
	return token
}