package controllers

import (
	"net/http"
	"time"

	"github.com/Xuduoteng/gomall/internal/services"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (userController *UserController) CreateUser(ctx *gin.Context) {
	var request createUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	user, err := userService.CreateUser(request.Username, request.Password, request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	response := createUserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

type LoginByUsernamePasswordRequest struct {
	Usernmae string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginByUsernamePasswordResponse struct {
	Usernmae  string `json:"username" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	BearToken string `json:"token"`
}

// @Router /users/loginByUsernamePassword [post]
// @Description Login By Username Password
// @Tags User
// @Param data body LoginByUsernamePasswordRequest true "username、password"
func (userController *UserController) LoginByUsernamePassword(ctx *gin.Context) {

	var request LoginByUsernamePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if request.Usernmae == "" || request.Password == "" || request.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error request"})
		return
	}

	token, err := userService.LoginByUsernamePassword(request.Usernmae, request.Password, request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	bear_header_token := "Bearer " + token

	response := LoginByUsernamePasswordResponse{
		Usernmae:  request.Usernmae,
		Email:     request.Email,
		BearToken: bear_header_token,
	}
	ctx.JSON(http.StatusOK, response)
}
