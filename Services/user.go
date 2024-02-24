package Services

import (
	"github.com/gin-gonic/gin"
	Log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"webservice-pattern/Constants"
	"webservice-pattern/Entities"
	"webservice-pattern/Repositories"
	"webservice-pattern/Utilites"
)

type UserService interface {
	GetAllUser(c *gin.Context)
	GetUserById(c *gin.Context)
	AddUserData(c *gin.Context)
	UpdateUserData(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserServiceImplement struct {
	userRepository Repositories.UserRepository
}

func (u UserServiceImplement) UpdateUserData(c *gin.Context) {
	defer Utilites.PanicHandler(c)

	Log.Info("start to execute program update user data by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	var request Entities.User
	if err := c.ShouldBindJSON(&request); err != nil {
		Log.Error("Happened error when mapping request from FE. Error", err)
		Utilites.PanicException(Constants.InvalidRequest)
	}

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		Log.Error("Happened error when get data from database. Error", err)
		Utilites.PanicException(Constants.DataNotFound)
	}

	data.RoleID = request.RoleID
	data.Email = request.Email
	data.Name = request.Password
	data.Status = request.Status
	_, err = u.userRepository.Save(&data)
	if err != nil {
		return
	}

	if err != nil {
		Log.Error("Happened error when updating data to database. Error", err)
		Utilites.PanicException(Constants.UnknownError)
	}

	c.JSON(http.StatusOK, Utilites.BuildResponse(Constants.Success, data))
}

func (u UserServiceImplement) GetUserById(c *gin.Context) {
	defer Utilites.PanicHandler(c)

	Log.Info("start to execute program get user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	data, err := u.userRepository.FindUserById(userID)
	if err != nil {
		Log.Error("Happened error when get data from database. Error", err)
		Utilites.PanicException(Constants.DataNotFound)
	}

	c.JSON(http.StatusOK, Utilites.BuildResponse(Constants.Success, data))
}

func (u UserServiceImplement) AddUserData(c *gin.Context) {
	defer Utilites.PanicHandler(c)

	Log.Info("start to execute program add data user")
	var request Entities.User
	if err := c.ShouldBindJSON(&request); err != nil {
		Log.Error("Happened error when mapping request from FE. Error", err)
		Utilites.PanicException(Constants.InvalidRequest)
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 15)
	request.Password = string(hash)

	data, err := u.userRepository.Save(&request)
	if err != nil {
		Log.Error("Happened error when saving data to database. Error", err)
		Utilites.PanicException(Constants.UnknownError)
	}

	c.JSON(http.StatusOK, Utilites.BuildResponse(Constants.Success, data))
}

func (u UserServiceImplement) GetAllUser(c *gin.Context) {
	defer Utilites.PanicHandler(c)

	Log.Info("start to execute get all data user")

	data, err := u.userRepository.FindAllUser()
	if err != nil {
		Log.Error("Happened Error when find all user data. Error: ", err)
		Utilites.PanicException(Constants.UnknownError)
	}

	c.JSON(http.StatusOK, Utilites.BuildResponse(Constants.Success, data))
}

func (u UserServiceImplement) DeleteUser(c *gin.Context) {
	defer Utilites.PanicHandler(c)

	Log.Info("start to execute delete data user by id")
	userID, _ := strconv.Atoi(c.Param("userID"))

	err := u.userRepository.DeleteUserById(userID)
	if err != nil {
		Log.Error("Happened Error when try delete data user from DB. Error:", err)
		Utilites.PanicException(Constants.UnknownError)
	}

	c.JSON(http.StatusOK, Utilites.BuildResponse(Constants.Success, Utilites.Null()))
}

func UserServiceInit(userRepository Repositories.UserRepository) *UserServiceImplement {
	return &UserServiceImplement{
		userRepository: userRepository,
	}
}
