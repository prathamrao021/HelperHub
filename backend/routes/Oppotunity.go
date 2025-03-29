package routes

import (
	"net/http"
	"time"

	"strconv"

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
	var opportunity models.Opportunity
	if err := db.Where("id = ?", id).First(&opportunity).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Opportunity not found"})
		return
	} else {
		if err := db.Where("id = ?", id).Delete(&opportunity).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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

	var updatedOpportunity map[string]interface{}
	if err := c.ShouldBindJSON(&updatedOpportunity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedOpportunity["updated_at"] = time.Now()

	if err := db.Model(&opportunity).Updates(updatedOpportunity).Error; err != nil {
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

// getLastNExpiredOpportunitiesByOrganization godoc
// @Summary Retrieve the last 'n' expired opportunities for an organization
// @Description Retrieve the last 'n' opportunities where the end_date is less than the current date for a specific organization
// @Tags opportunities
// @Accept json
// @Produce json
// @Param organization_id path uint true "Organization ID"
// @Param n query int true "Number of opportunities"
// @Success 200 {array} models.Opportunity
// @Router /opportunities/organization/{organization_id}/expired [get]
func getLastNExpiredOpportunitiesByOrganization(c *gin.Context, db *gorm.DB) {
	organizationID := c.Param("organization_id")
	nStr := c.Query("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of opportunities"})
		return
	}

	var opportunities []models.Opportunity
	currentDate := time.Now()

	if err := db.Where("organization_mail = ? AND end_date < ?", organizationID, currentDate).
		Order("end_date desc").
		Limit(n).
		Find(&opportunities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunities)
}
