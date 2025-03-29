package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDBOpportunity() *gorm.DB {
	// Use the same PostgreSQL connection as in main.go but with a test database
	dsn := "host=localhost user=postgres password=admin dbname=Helperhub_test port=5432 sslmode=prefer TimeZone=Asia/Shanghai"

	// Configure gorm with minimal logging during tests
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// Open connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	// Clean up existing tables
	db.Exec("DROP TABLE IF EXISTS opportunities CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.Opportunity{})

	// Create test organization (required for foreign key constraint)
	// First, check if organization already exists
	var count int64
	db.Model(&models.Organization{}).Where("email = ?", "test@org.com").Count(&count)

	if count == 0 {
		// Create the test organization if it doesn't exist
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		testOrg := models.Organization{
			Email:       "test@org.com",
			Password:    string(hashedPassword),
			Name:        "Test Organization",
			Phone:       "1234567890",
			Location:    "Test Location",
			Description: "Test Description",
			Website_Url: "https://testorg.com",
			Created_At:  time.Now(),
			Updated_At:  time.Now(),
		}
		db.Create(&testOrg)
	}

	return db
}

func setupRouterOpportunity(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes with injected database
	r.POST("/opportunities/create", func(c *gin.Context) {
		createOpportunity(c, db)
	})
	r.DELETE("/opportunities/delete/:id", func(c *gin.Context) {
		deleteOpportunity(c, db)
	})
	r.PUT("/opportunities/update/:id", func(c *gin.Context) {
		updateOpportunity(c, db)
	})
	r.GET("/opportunities/get/:id", func(c *gin.Context) {
		getOpportunity(c, db)
	})

	return r
}

func createTestOpportunity(db *gorm.DB) models.Opportunity {
	opp := models.Opportunity{
		Organization_mail: "test@org.com",
		Category:          "Education",
		Title:             "Test Opportunity",
		Description:       "Test Description",
		Location:          "Test Location",
		Hours_Required:    5, // 48 hours from now
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	db.Create(&opp)
	return opp
}

func cleanupTestOpportunities(db *gorm.DB) {
	db.Exec("DELETE FROM opportunities")
}

func TestCreateOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Test opportunity data
	opportunity := models.Opportunity{
		Organization_mail: "test@org.com",
		Category:          "Education",
		Title:             "Tutor Children",
		Description:       "Help children with homework",
		Location:          "Local Library",
		Hours_Required:    10, // 24 hours from now
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(opportunity)

	// Create request
	req, _ := http.NewRequest("POST", "/opportunities/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Opportunity
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.NotZero(t, response.ID)
	assert.Equal(t, opportunity.Organization_mail, response.Organization_mail)
	assert.Equal(t, opportunity.Category, response.Category)
	assert.Equal(t, opportunity.Title, response.Title)
	assert.Equal(t, opportunity.Description, response.Description)
	assert.Equal(t, opportunity.Location, response.Location)
	assert.NotZero(t, response.Created_At)
	assert.NotZero(t, response.Updated_At)
}

func TestGetOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)

	// Create request
	req, _ := http.NewRequest("GET", "/opportunities/get/"+fmt.Sprintf("%d", opp.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Opportunity
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, opp.ID, response.ID)
	assert.Equal(t, opp.Organization_mail, response.Organization_mail)
	assert.Equal(t, opp.Title, response.Title)
}

func TestUpdateOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)

	// Updated opportunity data
	updatedOpp := opp
	updatedOpp.Title = "Updated Title"
	updatedOpp.Description = "Updated Description"

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedOpp)

	// Create request
	url := "/opportunities/update/" + fmt.Sprintf("%d", opp.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Opportunity
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, opp.ID, response.ID)
	assert.Equal(t, "Updated Title", response.Title)
	assert.Equal(t, "Updated Description", response.Description)
}

func TestDeleteOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)

	// Create request
	url := "/opportunities/delete/" + fmt.Sprintf("%d", opp.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify opportunity is deleted
	var count int64
	db.Model(&models.Opportunity{}).Where("id = ?", opp.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGetNonExistentOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create request for non-existent ID
	req, _ := http.NewRequest("GET", "/opportunities/get/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInvalidOpportunityData(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/opportunities/create", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
