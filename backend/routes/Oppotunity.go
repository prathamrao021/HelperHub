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

// getOpportunitiesWithApplicationCount godoc
// @Summary Retrieve all opportunities for an organization with application counts
// @Description Retrieve all opportunities for a specific organization, including the number of applications each opportunity has received
// @Tags opportunities
// @Accept json
// @Produce json
// @Param organization_mail query string true "Organization Mail"
// @Success 200 {array} map[string]interface{}
// @Router /opportunities [get]
func getOpportunitiesByOrganization(c *gin.Context, db *gorm.DB) {
	organizationMail := c.Query("organization_mail")
	if organizationMail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization_id is required"})
		return
	}

	var opportunities []map[string]interface{}

	// Query to retrieve opportunities with application counts
	if err := db.Table("opportunities").
		Select("opportunities.*, COUNT(applications.id) AS application_count").
		Joins("LEFT JOIN applications ON applications.opportunity_id = opportunities.id").
		Where("opportunities.organization_mail = ?", organizationMail).
		Group("opportunities.id").
		Find(&opportunities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunities)
}

// func getOpportunitiesByOrganization(c *gin.Context, db *gorm.DB) {
// 	stringOrgID := c.Query("organization_id")

// 	if stringOrgID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "organization_id is required"})
// 		return
// 	}

// 	pageStr := c.DefaultQuery("page", "1")
// 	pageSizeStr := c.DefaultQuery("page_size", "10")

// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil || page < 1 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
// 		return
// 	}

// 	pageSize, err := strconv.Atoi(pageSizeStr)
// 	if err != nil || pageSize < 1 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
// 		return
// 	}

// 	offset := (page - 1) * pageSize

// 	var opportunities []models.Opportunity
// 	if err := db.Table("opportunities").
// 		Select("opportunities.*, COUNT(applications.id) AS application_count").
// 		Joins("LEFT JOIN applications ON applications.opportunity_id = opportunities.id").
// 		Where("opportunities.organization_mail = ?", stringOrgID).
// 		Group("opportunities.id").
// 		Offset(offset).
// 		Limit(pageSize).
// 		Find(&opportunities).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, opportunities)
// }

// getAvailableOpportunities godoc
// @Summary Retrieve available volunteer opportunities
// @Description Retrieve available volunteer opportunities, excluding expired ones, and include organization name
// @Tags opportunities
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /opportunities/available [get]
func getAvailableOpportunities(c *gin.Context, db *gorm.DB) {
	currentDate := time.Now()

	var opportunities []map[string]interface{}

	// Query to retrieve available opportunities
	if err := db.Table("opportunities").
		Select("opportunities.*, organizations.name AS organization_name").
		Joins("INNER JOIN organizations ON opportunities.organization_mail = organizations.email").
		Where("opportunities.end_date >= ?", currentDate).
		Order("opportunities.start_date ASC").
		Find(&opportunities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, opportunities)
}

// getOpportunityWithStats godoc
// @Summary Get opportunity details with application statistics
// @Description Retrieve details of a specific opportunity including counts of applications by status
// @Tags opportunities
// @Accept json
// @Produce json
// @Param id path uint true "Opportunity ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /opportunities/{id} [get]
func getOpportunityWithStats(c *gin.Context, db *gorm.DB) {
	opportunityID := c.Param("opportunity_id")

	// Get the basic opportunity details
	var opportunity models.Opportunity
	if err := db.First(&opportunity, opportunityID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Opportunity not found"})
		return
	}

	// Get application counts by status
	var totalApplications int64
	var pendingApplications int64
	var acceptedApplications int64
	var rejectedApplications int64

	// Count total applications
	db.Model(&models.Application{}).Where("opportunity_id = ?", opportunityID).Count(&totalApplications)

	// Count pending applications
	db.Model(&models.Application{}).Where("opportunity_id = ? AND status = ?", opportunityID, "Pending").Count(&pendingApplications)

	// Count accepted applications
	db.Model(&models.Application{}).Where("opportunity_id = ? AND status = ?", opportunityID, "Accepted").Count(&acceptedApplications)

	// Count rejected applications
	db.Model(&models.Application{}).Where("opportunity_id = ? AND status = ?", opportunityID, "Rejected").Count(&rejectedApplications)

	// Create a map for the response
	response := map[string]interface{}{
		"id":                    opportunity.ID,
		"organization_id":       opportunity.Organization_mail,
		"title":                 opportunity.Title,
		"description":           opportunity.Description,
		"location":              opportunity.Location,
		"hours_required":        opportunity.Hours_Required,
		"start_date":            opportunity.Start_Date,
		"end_date":              opportunity.End_Date,
		"created_at":            opportunity.Created_At,
		"updated_at":            opportunity.Updated_At,
		"category":              opportunity.Category,
		"total_applications":    totalApplications,
		"pending_applications":  pendingApplications,
		"accepted_applications": acceptedApplications,
		"rejected_applications": rejectedApplications,
	}

	c.JSON(http.StatusOK, response)
}
