package service

import (
	"Users/internal/models"
	"Users/internal/repository"
	_ "context"
	_ "errors"
	"fmt"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Недопустимый формат данных",
		})
		return
	}

	// Вставляем пользователя в базу данных
	err := s.usersRepo.InsertUserToDB(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Не удалось создать пользователя: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан",
	})
}

func (s *UserService) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Недопустимый формат данных",
		})
		return
	}

	// Обновляем пользователя в базе данных
	err := s.usersRepo.UpdateUserInDB(c.Request.Context(), user.Email, user.Name, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Не удалось обновить пользователя: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь успешно обновлён",
	})
}
