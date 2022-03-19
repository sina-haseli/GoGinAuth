package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.com/m/dto"
	"golang.com/m/entity"
	"golang.com/m/service"
)

type UserController interface {
	GetAllUsers(ctx *gin.Context) entity.User
	CreateUser(ctx *gin.Context) entity.User
}

type userController struct {
	service service.UserService
	dto    dto.UserDto
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

// ShowAccount godoc
// @Summary      Show an User
// @Description  get User By Id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Router       /user/{id} [get]
func (u *userController) GetAllUsers(id *gin.Context) entity.User {
	context, err := strconv.ParseInt(id.Param("id"), 0, 0)
	if err != nil {
		 panic("Failed to parse id")
	}
	return u.service.GetAllUsers(context)
}

// ShowAccount godoc 
// @Summary      Create an User
// @Description  Create User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        dto.UserDto	body    dto.UserDto  true  "User"
// @Router       /user [post]
func (u *userController) CreateUser(ctx *gin.Context) entity.User {
	var user entity.User
	ctx.BindJSON(&user)
	u.service.CreateUser(&user)
	return user
}
