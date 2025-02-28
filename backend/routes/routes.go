package routes

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	// Routes for user management
	userRouter := router.Group("/users")
	userRouter.POST("/create", func(c *gin.Context) { createUser(c, db) })
	userRouter.DELETE("/delete/:username", func(c *gin.Context) { deleteUser(c, db) })
	userRouter.PUT("/update/:username", func(c *gin.Context) { updateUser(c, db) })
	userRouter.GET("/get/:username", func(c *gin.Context) { getUser(c, db) })
}
