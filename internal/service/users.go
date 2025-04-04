package service

import (
	"Users/internal/models"
	"Users/internal/repository"
	_ "context"
	_ "errors"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type UserService struct {
	usersRepo *repository.UserRepo
	users     map[string]User
}

func NewUserService(usersRepo *repository.UserRepo) *UserService {
	return &UserService{
		usersRepo: usersRepo,
		users:     make(map[string]User),
	}
}

func (s *UserService) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый формат данных"})
		return
	}

	// Вставляем пользователя в базу данных
	if err := s.usersRepo.InsertUserToDB(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "пользователь успешно создан"})
}

func (s *UserService) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "недопустимый формат данных"})
		return
	}

	// Обновляем пользователя в базе данных
	if err := s.usersRepo.UpdateUserInDB(user.Email, user.Name, user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить пользователя"})
		return
	}
}
