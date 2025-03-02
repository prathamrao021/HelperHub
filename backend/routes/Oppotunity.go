package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"gorm.io/gorm"
)

// createOpportunity godoc
// @Summary Create a new opportunity
// @Description Create a new opportunity with the provided details
// @Tags opportunities
// @Accept json
// @Produce json
// @Param opportunity body models.Opportunity true "Opportunity data"
// @Success 200 {object} models.Opportunity
// @Router /opportunities/create [post]
func createOpportunity(c *gin.Context, db *gorm.DB) {
	var opportunity models.Opportunity
	if err := c.ShouldBindJSON(&opportunity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	opportunity.Created_At = time.Now()
	opportunity.Updated_At = time.Now()

	if err := db.Create(&opportunity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunity)
}

// deleteOpportunity godoc
// @Summary Delete an existing opportunity
// @Description Delete an existing opportunity by ID
// @Tags opportunities
// @Accept json
// @Produce json
// @Param id path uint true "Opportunity ID"
// @Success 200 {object} map[string]string
// @Router /opportunities/delete/{id} [delete]
func deleteOpportunity(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&models.Opportunity{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Opportunity deleted successfully"})
}

// updateOpportunity godoc
// @Summary Update an existing opportunity
// @Description Update an existing opportunity with the provided details
// @Tags opportunities
// @Accept json
// @Produce json
// @Param id path uint true "Opportunity ID"
// @Param opportunity body models.Opportunity true "Opportunity data"
// @Success 200 {object} models.Opportunity
// @Router /opportunities/update/{id} [put]
func updateOpportunity(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var opportunity models.Opportunity

	if err := db.Where("id = ?", id).First(&opportunity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Opportunity not found"})
		return
	}

	if err := c.ShouldBindJSON(&opportunity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	opportunity.Updated_At = time.Now()

	if err := db.Save(&opportunity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunity)
}

// getOpportunity godoc
// @Summary Get an existing opportunity
// @Description Get an existing opportunity by ID
// @Tags opportunities
// @Accept json
// @Produce json
// @Param id path uint true "Opportunity ID"
// @Success 200 {object} models.Opportunity
// @Router /opportunities/get/{id} [get]
func getOpportunity(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var opportunity models.Opportunity

	if err := db.Where("id = ?", id).First(&opportunity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Opportunity not found"})
		return
	}

	c.JSON(http.StatusOK, opportunity)
}
