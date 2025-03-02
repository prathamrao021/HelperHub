package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// createVolunteer godoc
// @Summary Create a new volunteer
// @Description Create a new volunteer with the provided details
// @Tags volunteers
// @Accept json
// @Produce json
// @Param volunteer body models.Volunteer true "Volunteer data"
// @Success 200 {object} models.Volunteer
// @Router /volunteers/create [post]
func createVolunteer(c *gin.Context, db *gorm.DB) {
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(volunteer.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	volunteer.Password = string(hashedPassword)
	volunteer.Created_At = time.Now()
	volunteer.Updated_At = time.Now()

	if err := db.Create(&volunteer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data created successfully"})
}

// deleteVolunteer godoc
// @Summary Delete an existing volunteer
// @Description Delete an existing volunteer by username
// @Tags volunteers
// @Accept json
// @Produce json
// @Param volunteer_mail path string true "Email"
// @Success 200 {object} models.Volunteer
// @Router /volunteers/delete/{volunteer_mail} [delete]
func deleteVolunteer(c *gin.Context, db *gorm.DB) {
	mail := c.Param("volunteer_mail")

	if err := db.Where("email = ?", mail).Delete(&models.Volunteer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data deleted successfully"})
}

// updateVolunteer godoc
// @Summary Update an existing volunteer
// @Description Update an existing volunteer with the provided details
// @Tags volunteers
// @Accept json
// @Produce json
// @Param volunteer_mail path string true "Email"
// @Param volunteer body models.Volunteer true "Volunteer data"
// @Success 200 {object} models.Volunteer
// @Router /volunteers/update/{volunteer_mail} [put]
func updateVolunteer(c *gin.Context, db *gorm.DB) {
	mail := c.Param("volunteer_mail")
	var volunteer models.Volunteer
	if err := c.ShouldBindJSON(&volunteer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password if it is provided
	if volunteer.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(volunteer.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		volunteer.Password = string(hashedPassword)
	}

	volunteer.Updated_At = time.Now()

	if err := db.Where("email = ?", mail).Updates(&volunteer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer data updated successfully"})
}

// getVolunteer godoc
// @Summary Get an existing volunteer
// @Description Get an existing volunteer by username
// @Tags volunteers
// @Accept json
// @Produce json
// @Param volunteer_mail path string true "Email"
// @Success 200 {object} models.Volunteer
// @Router /volunteers/get/{volunteer_mail} [get]
func getVolunteer(c *gin.Context, db *gorm.DB) {
	mail := c.Param("volunteer_mail")
	var volunteer models.Volunteer

	if err := db.Where("email = ?", mail).First(&volunteer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, volunteer)
}

// loginVolunteer godoc
// @Summary Login a volunteer
// @Description Login a volunteer with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.Volunteer
// @Router /login/volunteer [post]
func loginVolunteer(c *gin.Context, db *gorm.DB) {
	var credentials models.LoginRequest
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var volunteer models.Volunteer
	if err := db.Where("email = ?", credentials.Email).First(&volunteer).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(volunteer.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": volunteer})
}
