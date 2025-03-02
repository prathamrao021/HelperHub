package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"gorm.io/gorm"
)

// createApplication godoc
// @Summary Create a new application
// @Description Create a new application with the provided details
// @Tags applications
// @Accept json
// @Produce json
// @Param application body models.Application true "Application data"
// @Success 200 {object} models.Application
// @Router /applications [post]
func createApplication(c *gin.Context, db *gorm.DB) {
	var application models.Application
	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.Created_At = time.Now()
	application.Updated_At = time.Now()

	if err := db.Create(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, application)
}

// getAllApplications godoc
// @Summary Retrieve all applications (Admin-only)
// @Description Retrieve all applications
// @Tags applications
// @Accept json
// @Produce json
// @Success 200 {array} models.Application
// @Router /applications [get]
func getAllApplications(c *gin.Context, db *gorm.DB) {
	var applications []models.Application
	if err := db.Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// getApplicationByID godoc
// @Summary Retrieve an application by ID
// @Description Retrieve an application by ID
// @Tags applications
// @Accept json
// @Produce json
// @Param id path uint true "Application ID"
// @Success 200 {object} models.Application
// @Router /applications/{id} [get]
func getApplicationByID(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var application models.Application

	if err := db.Where("id = ?", id).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	c.JSON(http.StatusOK, application)
}

// getApplicationsByVolunteerID godoc
// @Summary Retrieve applications by volunteer
// @Description Retrieve applications by volunteer ID
// @Tags applications
// @Accept json
// @Produce json
// @Param volunteer_id query uint true "Volunteer ID"
// @Success 200 {array} models.Application
// @Router /applications [get]
func getApplicationsByVolunteerID(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Query("volunteer_id")
	var applications []models.Application

	if err := db.Where("volunteer_id = ?", volunteerID).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// getApplicationsByOpportunityID godoc
// @Summary Retrieve applications by opportunity
// @Description Retrieve applications by opportunity ID
// @Tags applications
// @Accept json
// @Produce json
// @Param opportunity_id query uint true "Opportunity ID"
// @Success 200 {array} models.Application
// @Router /applications [get]
func getApplicationsByOpportunityID(c *gin.Context, db *gorm.DB) {
	opportunityID := c.Query("opportunity_id")
	var applications []models.Application

	if err := db.Where("opportunity_id = ?", opportunityID).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// getApplicationsByStatus godoc
// @Summary Retrieve applications by status
// @Description Retrieve applications by status
// @Tags applications
// @Accept json
// @Produce json
// @Param status query string true "Status"
// @Success 200 {array} models.Application
// @Router /applications [get]
func getApplicationsByStatus(c *gin.Context, db *gorm.DB) {
	status := c.Query("status")
	var applications []models.Application

	if err := db.Where("status = ?", status).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// updateApplication godoc
// @Summary Update application details
// @Description Update application details (Cover Letter for Volunteers, Status for Organizations)
// @Tags applications
// @Accept json
// @Produce json
// @Param id path uint true "Application ID"
// @Param application body models.Application true "Application data"
// @Success 200 {object} models.Application
// @Router /applications/{id} [put]
func updateApplication(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var application models.Application

	if err := db.Where("id = ?", id).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	if err := c.ShouldBindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application.Updated_At = time.Now()

	if err := db.Save(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, application)
}

// deleteApplication godoc
// @Summary Delete an application
// @Description Delete an application by ID
// @Tags applications
// @Accept json
// @Produce json
// @Param id path uint true "Application ID"
// @Success 200 {object} map[string]string
// @Router /applications/{id} [delete]
func deleteApplication(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&models.Application{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}
