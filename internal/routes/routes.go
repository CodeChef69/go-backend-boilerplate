package routes

import (
	"go-simple-backend/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// Define routes
	api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers(db))
		api.POST("/users", controllers.CreateUser(db))
	}

	return router
}
