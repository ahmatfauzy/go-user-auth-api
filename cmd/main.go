package main

import (
	"auth-api/internal/config"
	"auth-api/internal/domain"
	"auth-api/internal/delivery/http/handler"
	"auth-api/internal/repository"
	"auth-api/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("▶️ Mulai eksekusi...")

	cfg := config.LoadConfig()
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(db)
	authUc := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret, cfg.JWTExpireHours)
	authHandler := handler.NewAuthHandler(authUc)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)
		api.POST("/refresh", authHandler.RefreshToken)
		api.POST("/logout", authHandler.Logout)
	}

	router.Run(":" + cfg.AppPort)
}