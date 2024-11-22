package services

import (
	"errors"

	"github.com/Xuduoteng/gomall/internal/models"
	"github.com/Xuduoteng/gomall/internal/pkg/utils"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (userService *UserService) CreateUser(username string, password string, email string) (models.User, error) {
	existingUser := models.User{}
	res := db.First(&existingUser, "email = ?", email)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		} else {
			return models.User{}, errors.New("database query error: " + res.Error.Error())
		}
	}
	if res.RowsAffected != 0 {
		return models.User{}, errors.New("email has been registered")
	}

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, errors.New("hash password error")
	}

	// add it to the database
	user := models.User{
		Username: username,
		Password: string(hashed_password),
		Email:    email,
	}
	if err := db.Create(&user).Error; err != nil {
		return models.User{}, errors.New("create user error")
	}

	return user, nil
}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string, email string) (string, error) {
	user := models.User{
		Email: email,
	}
	res := db.First(&user, "email = ?", email)
	if res.Error != nil || res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("password error")
	}
	claims := utils.Claims{
		Username: user.Username,
		Uid:      user.ID,
	}
	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return "", errors.New("generate token error")
	}
	return token, nil
}

func (UserService *UserService) GetUser(UserId int) *models.User {

	User := models.User{}
	res := db.First(&User, UserId).Select("id, name, status")

	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}
	return &User
}

func (UserService *UserService) UpdateUser(data map[string]interface{}) bool {

	User := models.User{}
	User.ID = uint(data["UserId"].(int))

	res := db.Model(&User).Updates(data)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}

func (UserService *UserService) DeleteUser(UserId int) bool {

	User := models.User{}
	res := db.Delete(&User, UserId)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
