package main

import (
	"all_inclusive_app/logger"
	"log"
	"net/http"

	"all_inclusive_app/config"
	"all_inclusive_app/db"
	_ "all_inclusive_app/docs" // Этот пакет содержит сгенерированные файлы Swagger
	"all_inclusive_app/repository"
	"all_inclusive_app/server"
	"all_inclusive_app/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title All Inclusive App API
// @version 1.0
// @description API для All Inclusive приложения с использованием Go, Gin, GORM, JWT и Swagger
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	config.ReadConfigApp("config/config.json")
	config.ReadConfigDB("config/configDB.json")
	config.ReadConfigLogger("config/configLogger.json")

	logger.InitLogger()
	dbSettings := db.InitDB()

	repositorySettings := repository.NewRepository(dbSettings)

	serviceSettings := service.NewService(repositorySettings)

	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	handler := server.NewHandler(serviceSettings, r)

	handler.AllRoutes()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "page not found"})
	})

	err := r.Run(config.Configure.PortRun)
	if err != nil {
		log.Fatal("fatal to run route, Error:", err.Error())
	}

}

// func authMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")

// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, HTTPError{Message: "Missing token"})
// 			c.Abort()
// 			return
// 		}

// 		claims := &models.Claims{}
// 		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 			return []byte("your-256-bit-secret"), nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, HTTPError{Message: "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("email", claims.Email)
// 		c.Next()
// 	}
// }
