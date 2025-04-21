package routes

import (
	"net/http"
	"strconv"
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
// @Param volunteer_id path uint true "Volunteer ID"
// @Success 200 {array} models.Application
// @Router /applications/volunteer/{volunteer_id} [get]
func getApplicationsByVolunteerID(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")
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
// @Param opportunity_id path uint true "Opportunity ID"
// @Success 200 {array} models.Application
// @Router /applications/opportunity/{opportunity_id} [get]
func getApplicationsByOpportunityID(c *gin.Context, db *gorm.DB) {
	opportunityID := c.Param("opportunity_id")
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
// @Param status path string true "Status"
// @Success 200 {array} models.Application
// @Router /applications/status/{status} [get]
func getApplicationsByStatus(c *gin.Context, db *gorm.DB) {
	status := c.Param("status")
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
	var application models.Application

	if err := db.Where("id = ?", id).First(&application).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	} else {
		if err := db.Where("id = ?", id).Delete(&application).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application deleted successfully"})
}

// getLastNApprovedApplications godoc
// @Summary Retrieve the last 'n' approved applications for a volunteer
// @Description Retrieve the last 'n' applications with the status "Approved" for a specific volunteer
// @Tags applications
// @Accept json
// @Produce json
// @Param volunteer_id path uint true "Volunteer ID"
// @Param n query int true "Number of applications"
// @Success 200 {array} models.Application
// @Router /applications/volunteer/{volunteer_id}/approved [get]
func getLastNApprovedApplications(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")
	nStr := c.Query("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of applications"})
		return
	}

	var applications []models.Application
	if err := db.Where("volunteer_id = ? AND status = ?", volunteerID, "Approved").Order("created_at desc").Limit(n).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// getLastNAcceptedOpportunitiesForVolunteer godoc
// @Summary Retrieve the last 'n' opportunities for a volunteer where the application was accepted and end_date < current date
// @Description Retrieve the last 'n' opportunities for a volunteer where the application was accepted and end_date < current date
// @Tags opportunities
// @Accept json
// @Produce json
// @Param volunteer_id path uint true "Volunteer ID"
// @Param n query int true "Number of opportunities"
// @Success 200 {array} models.Opportunity
// @Router /opportunities/volunteer/{volunteer_id}/accepted-expired [get]
func getLastNAcceptedOpportunitiesForVolunteer(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")
	nStr := c.Query("n")
	n, err := strconv.Atoi(nStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of opportunities"})
		return
	}

	var opportunities []models.Opportunity
	currentDate := time.Now()

	// Query to join Applications and Opportunities
	if err := db.Table("applications").
		Select("opportunities.*").
		Joins("join opportunities on applications.opportunity_id = opportunities.id").
		Where("applications.volunteer_id = ? AND applications.status = ? AND opportunities.end_date < ?", volunteerID, "Accepted", currentDate).
		Order("opportunities.end_date desc").
		Limit(n).
		Scan(&opportunities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunities)
}

// getApplicationsByVolunteerWithDetails godoc
// @Summary Retrieve applications by volunteer with opportunity and organization details
// @Description Retrieve all applications for a volunteer with opportunity title and organization name
// @Tags applications
// @Accept json
// @Produce json
// @Param volunteer_id path uint true "Volunteer ID"
// @Success 200 {array} map[string]interface{}
// @Router /applications/volunteer/{volunteer_id} [get]
func getApplicationsByVolunteerWithDetails(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")

	var results []map[string]interface{}

	// Query with inner joins to get data from all three tables and filter by volunteer ID
	if err := db.Table("applications").
		Select(`
            applications.id,
            applications.volunteer_id,
            applications.opportunity_id,
            opportunities.title as opportunity_title,
            organizations.name as organization_name,
            applications.status,
            applications.cover_letter,
            applications.created_at,
            applications.updated_at
        `).
		Joins("INNER JOIN opportunities ON applications.opportunity_id = opportunities.id").
		Joins("INNER JOIN organizations ON opportunities.organization_mail = organizations.email").
		Where("applications.volunteer_id = ?", volunteerID).
		Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
