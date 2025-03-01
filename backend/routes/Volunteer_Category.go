package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"gorm.io/gorm"
)

// CreateVolunteerCategory godoc
// @Summary Create volunteer categories
// @Description Create volunteer categories in the database
// @Tags volunteer_categories
// @Accept json
// @Produce json
// @Param volunteer_category body models.Volunteer_Category true "Volunteer_Category data"
// @Success 200 {object} map[string]string
// @Router /volunteer_categories/create [post]
func createVolunteerCategory(c *gin.Context, db *gorm.DB) {
	var volunteer_category models.Volunteer_Category
	if err := c.ShouldBindJSON(&volunteer_category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&volunteer_category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data created successfully"})
}

// deleteVolunteerCategory godoc
// @Summary Delete an existing volunteer category
// @Description Delete an existing volunteer category by category
// @Tags volunteer_categories
// @Accept json
// @Produce json
// @Param volunteer_id path string true "Volunteer ID"
// @Success 200 {object} models.Volunteer_Category
// @Router /volunteer_categories/delete/{volunteer_id} [delete]
func deleteVolunteerCategory(c *gin.Context, db *gorm.DB) {
	var volunteer_category models.Volunteer_Category
	if err := c.ShouldBindJSON(&volunteer_category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("volunteer_id = ?", volunteer_category.Volunteer_ID).Delete(&models.Volunteer_Category{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data deleted successfully"})
}

// getVolunteerCategory godoc
// @Summary Get categories for a volunteer
// @Description Get all categories associated with a volunteer by volunteer_id
// @Tags volunteer_categories
// @Accept json
// @Produce json
// @Param volunteer_id path uint true "Volunteer ID"
// @Success 200 {array} uint
// @Router /volunteer_categories/get/{volunteer_id} [get]
func getVolunteerCategory(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")
	var volunteerCategory models.Volunteer_Category

	// Retrieve the Volunteer_Category record for the given volunteer_id
	if err := db.Where("volunteer_id = ?", volunteerID).First(&volunteerCategory).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer categories not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Return the associated category IDs
	c.JSON(http.StatusOK, volunteerCategory.Category_ID)
}

// updateVolunteerCategory godoc
// @Summary Update an existing volunteer category
// @Description Update an existing volunteer category by volunteer_id
// @Tags volunteer_categories
// @Accept json
// @Produce json
// @Param volunteer_category body models.Volunteer_Category true "Volunteer_Category data"
// @Success 200 {object} models.Volunteer_Category
// @Router /volunteer_categories/update [put]
func updateVolunteerCategory(c *gin.Context, db *gorm.DB) {
	var volunteerCategory models.Volunteer_Category
	if err := c.ShouldBindJSON(&volunteerCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the Volunteer_Category record for the given volunteer_id
	if err := db.Where("volunteer_id = ?", volunteerCategory.Volunteer_ID).Updates(&volunteerCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data updated successfully"})
}
