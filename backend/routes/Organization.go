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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(organization.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	organization.Password = string(hashedPassword)
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
// @Description Delete an existing organization by organization_mail
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization_mail path string true "Email"
// @Success 200 {object} models.Organization
// @Router /organizations/delete/{organization_mail} [delete]
func deleteOrganization(c *gin.Context, db *gorm.DB) {
	mail := c.Param("organization_mail")

	if err := db.Where("email = ?", mail).Delete(&models.Organization{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Organization data deleted successfully"})
}

// updateOrganization godoc
// @Summary Update an existing organization
// @Description Update an existing organization by organization_mail
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization_mail path string true "Email"
// @Param organization body models.Organization true "Organization data"
// @Success 200 {object} models.Organization
// @Router /organizations/update/{organization_mail} [put]
func updateOrganization(c *gin.Context, db *gorm.DB) {
	mail := c.Param("organization_mail")

	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if organization.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(organization.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		organization.Password = string(hashedPassword)
	}

	organization.Updated_At = time.Now()

	if err := db.Where("email = ?", mail).Updates(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization data updated successfully"})
}

// getOrganization godoc
// @Summary Get an existing organization
// @Description Get an existing organization by organization_mail
// @Tags organizations
// @Accept json
// @Produce json
// @Param organization_mail path string true "Email"
// @Success 200 {object} models.Organization
// @Router /organizations/get/{organization_mail} [get]
func getOrganization(c *gin.Context, db *gorm.DB) {
	mail := c.Param("organization_mail")

	var organization models.Organization
	if err := db.Where("email = ?", mail).First(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, organization)
}

// loginOrganization godoc
// @Summary Login an organization
// @Description Login an organization with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginRequest
// @Router /login/organization [post]
func loginOrganization(c *gin.Context, db *gorm.DB) {
	var credentials models.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var organization models.Organization
	if err := db.Where("email = ?", credentials.Email).First(&organization).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(organization.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": organization})
}
