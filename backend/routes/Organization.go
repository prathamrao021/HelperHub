package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// createOrganization godoc
// @Summary Create a new organization
// @Description Create a new organization with the provided details
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization body models.Organization true "Organization data"
// @Success 200 {object} models.Organization
// @Router /organizations/create [post]
func createOrganization(c *gin.Context, db *gorm.DB) {
	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(organization.Password_Hash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	organization.Password_Hash = string(hashedPassword)
	organization.Created_At = time.Now()
	organization.Updated_At = time.Now()

	if err := db.Create(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization data created successfully"})
}

// deleteOrganization godoc
// @Summary Delete an existing organization
// @Description Delete an existing organization by username
// @Tags organizations
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Organization
// @Router /organizations/delete/{organization} [delete]
func deleteOrganization(c *gin.Context, db *gorm.DB) {
	username := c.Param("username")

	if err := db.Where("username = ?", username).Delete(&models.Organization{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// updateOrganization godoc
// @Summary Update an existing organization
// @Description Update an existing organization by username
// @Tags organizations
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param organization body models.Organization true "Organization data"
// @Success 200 {object} models.Organization
// @Router /organizations/update/{organization} [put]
func updateOrganization(c *gin.Context, db *gorm.DB) {
	username := c.Param("username")

	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if organization.Password_Hash != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(organization.Password_Hash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		organization.Password_Hash = string(hashedPassword)
	}

	organization.Updated_At = time.Now()

	if err := db.Where("username = ?", username).Updates(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization data updated successfully"})
}

// getOrganization godoc
// @Summary Get an existing organization
// @Description Get an existing organization by username
// @Tags organizations
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} models.Organization
// @Router /organizations/get/{organization} [get]
func getOrganization(c *gin.Context, db *gorm.DB) {
	username := c.Param("username")

	var organization models.Organization
	if err := db.Where("username = ?", username).First(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organization)
}
