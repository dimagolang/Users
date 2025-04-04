package http_server

import (
	"Users/config"
	"Users/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Server структура сервера с роутингом и зависимостями
type Server struct {
	userService *service.UserService
	cfg         config.Config
}

// NewServer создает экземпляр HTTP-сервера с настройкой роутинга
func NewServer(userService *service.UserService, cfg config.Config) *Server {
	return &Server{
		userService: userService,
		cfg:         cfg,
	}
}

// Run запускает сервер
func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Роуты для рейсов
	r.POST("/post-user", s.userService.CreateUser)
	r.PUT("/update-users", s.userService.UpdateUser)

	log.Printf("Server is running on port %s...", s.cfg.ServerPort)
	if err := r.Run(":" + s.cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
