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

	c.JSON(http.StatusOK, volunteer)
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
	var volunteer models.Volunteer

	if err := db.Where("email = ?", mail).First(&volunteer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		if err := db.Where("email = ?", mail).Delete(&volunteer).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Volunteer data deleted successfully"})
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Volunteer data not found"})

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

	if err := db.Where("email = ?", mail).First(&volunteer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
		return
	}

	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if password, exists := updatedData["password"]; exists && password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		updatedData["password"] = string(hashedPassword)
	}

	updatedData["updated_at"] = time.Now()

	if err := db.Model(&volunteer).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, volunteer)
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

// getVolunteerStats godoc
// @Summary Retrieve the total number of jobs and hours worked for a volunteer
// @Description Retrieve the total number of jobs and hours worked for a volunteer based on accepted applications
// @Tags volunteers
// @Accept json
// @Produce json
// @Param volunteer_id path uint true "Volunteer ID"
// @Success 200 {object} map[string]interface{}
// @Router /volunteers/{volunteer_id}/stats [get]
func getVolunteerStats(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Param("volunteer_id")

	var totalJobs int64
	var totalHoursWorked int64

	// Query to count the total number of jobs and sum the hours worked
	if err := db.Table("applications").
		Joins("join opportunities on applications.opportunity_id = opportunities.id").
		Where("applications.volunteer_id = ? AND applications.status = ?", volunteerID, "Accepted").
		Count(&totalJobs).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.Table("applications").
		Select("SUM(opportunities.hours_required)").
		Joins("join opportunities on applications.opportunity_id = opportunities.id").
		Where("applications.volunteer_id = ? AND applications.status = ?", volunteerID, "Accepted").
		Scan(&totalHoursWorked).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the stats
	c.JSON(http.StatusOK, gin.H{
		"total_jobs":         totalJobs,
		"total_hours_worked": totalHoursWorked,
	})
}
