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

	// Routes for volunteer management
	volunteerRouter := router.Group("/volunteers")
	volunteerRouter.POST("/create", func(c *gin.Context) { createVolunteer(c, db) })
	volunteerRouter.DELETE("/delete/:volunteer_mail", func(c *gin.Context) { deleteVolunteer(c, db) })
	volunteerRouter.PUT("/update/:volunteer_mail", func(c *gin.Context) { updateVolunteer(c, db) })
	volunteerRouter.GET("/get/:volunteer_mail", func(c *gin.Context) { getVolunteer(c, db) })

	// Routes for organization management
	organizationRouter := router.Group("/organizations")
	organizationRouter.POST("/create", func(c *gin.Context) { createOrganization(c, db) })
	organizationRouter.DELETE("/delete/:organization_mail", func(c *gin.Context) { deleteOrganization(c, db) })
	organizationRouter.PUT("/update/:organization_mail", func(c *gin.Context) { updateOrganization(c, db) })
	organizationRouter.GET("/get/:organization_mail", func(c *gin.Context) { getOrganization(c, db) })

	// Routes for category management
	categoryRouter := router.Group("/categories")
	categoryRouter.POST("/create", func(c *gin.Context) { CreateCategory(c, db) })
	categoryRouter.GET("/get", func(c *gin.Context) { getCategories(c, db) })

	// Routes for opportunity management
	opportunityRouter := router.Group("/opportunities")
	opportunityRouter.POST("/create", func(c *gin.Context) { createOpportunity(c, db) })
	opportunityRouter.DELETE("/delete/:id", func(c *gin.Context) { deleteOpportunity(c, db) })
	opportunityRouter.PUT("/update/:id", func(c *gin.Context) { updateOpportunity(c, db) })
	opportunityRouter.GET("/get/:id", func(c *gin.Context) { getOpportunity(c, db) })

}
