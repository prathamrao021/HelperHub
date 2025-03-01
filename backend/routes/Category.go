package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"gorm.io/gorm"
)

// CreateCategory godoc
// @Summary Create static categories
// @Description Create static categories in the database
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /categories/create [post]
func CreateCategory(c *gin.Context, db *gorm.DB) {
	// Define the static categories
	categories := []string{
		"Web Development",
		"Graphic Design",
		"Content Writing",
		"Social Media",
		"Teaching",
		"Event Planning",
		"Photography",
		"Translation",
		"First Aid",
		"Project Management",
	}

	// Iterate over the categories and insert them into the database if they do not already exist
	for _, category := range categories {
		var existingCategory models.Category
		if err := db.Where("category = ?", category).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				newCategory := models.Category{
					Category:   category,
					Created_At: time.Now(),
				}
				if err := db.Create(&newCategory).Error; err != nil {
					if c != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					}
					return
				}
			} else {
				if c != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}
				return
			}
		}
	}

	if c != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Categories created successfully"})
	}
}

// getCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories/get [get]
func getCategories(c *gin.Context, db *gorm.DB) {
	var categories []models.Category
	if err := db.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
