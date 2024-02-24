package Controllers

import (
	"github.com/gin-gonic/gin"
	"webservice-pattern/Services"
)

type UserController interface {
	GetAllUserData(c *gin.Context)
	AddUserData(c *gin.Context)
	GetUserById(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserControllerImplement struct {
	svc Services.UserService
}

func (u UserControllerImplement) GetAllUserData(c *gin.Context) {
	u.svc.GetAllUser(c)
}

func (u UserControllerImplement) AddUserData(c *gin.Context) {
	u.svc.AddUserData(c)
}

func (u UserControllerImplement) GetUserById(c *gin.Context) {
	u.svc.GetUserById(c)
}

func (u UserControllerImplement) UpdateUserData(c *gin.Context) {
	u.svc.UpdateUserData(c)
}

func (u UserControllerImplement) DeleteUser(c *gin.Context) {
	u.svc.DeleteUser(c)
}

func UserControllerInit(userService Services.UserService) *UserControllerImplement {
	return &UserControllerImplement{
		svc: userService,
	}
}
