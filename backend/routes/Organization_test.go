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

func setupTestDBForOrganization() *gorm.DB {
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
	db.Exec("DROP TABLE IF EXISTS organizations CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.Organization{})

	return db
}

func setupRouterForOrganization(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes with injected database
	r.POST("/organizations/create", func(c *gin.Context) {
		createOrganization(c, db)
	})
	r.DELETE("/organizations/delete/:organization_mail", func(c *gin.Context) {
		deleteOrganization(c, db)
	})
	r.PUT("/organizations/update/:organization_mail", func(c *gin.Context) {
		updateOrganization(c, db)
	})
	r.GET("/organizations/get/:organization_mail", func(c *gin.Context) {
		getOrganization(c, db)
	})
	r.POST("/login/organization", func(c *gin.Context) {
		loginOrganization(c, db)
	})

	return r
}

func createTestOrganization(db *gorm.DB) models.Organization {
	// Create a test organization with hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	
	organization := models.Organization{
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
	
	result := db.Create(&organization)
	if result.Error != nil {
		panic("Failed to create test organization: " + result.Error.Error())
	}
	
	// Verify the organization was created properly
	var created models.Organization
	db.Where("email = ?", organization.Email).First(&created)
	
	return created
}

func cleanupTestOrganizations(db *gorm.DB) {
	db.Exec("DELETE FROM organizations")
}

func TestCreateOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Test organization data
	organization := models.Organization{
		Email:       "new@org.com",
		Password:    "password123",
		Name:        "New Organization",
		Phone:       "9876543210",
		Location:    "New Location",
		Description: "New Description",
		Website_Url: "https://neworg.com",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(organization)

	// Create request
	req, _ := http.NewRequest("POST", "/organizations/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify organization was created in DB
	var createdOrg models.Organization
	result := db.Where("email = ?", organization.Email).First(&createdOrg)
	assert.NoError(t, result.Error, "Organization should be created in database")
	assert.Equal(t, organization.Email, createdOrg.Email)
	assert.Equal(t, organization.Name, createdOrg.Name)
	assert.NotEqual(t, "password123", createdOrg.Password, "Password should be hashed")
}

func TestGetOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization
	org := createTestOrganization(db)

	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/organizations/get/%s", org.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Organization
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify response fields
	assert.Equal(t, org.Email, response.Email)
	assert.Equal(t, org.Name, response.Name)
	assert.Equal(t, org.Phone, response.Phone)
	assert.Equal(t, org.Location, response.Location)
	assert.Equal(t, org.Description, response.Description)
	assert.Equal(t, org.Website_Url, response.Website_Url)
}

func TestUpdateOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization
	org := createTestOrganization(db)

	// Updated organization data
	updatedOrg := models.Organization{
		Name:        "Updated Name",
		Phone:       "5555555555",
		Location:    "Updated Location",
		Description: "Updated Description",
		Website_Url: "https://updated-org.com",
		// Don't update password
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedOrg)

	// Create request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/organizations/update/%s", org.Email), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify organization was updated in DB
	var updatedInDB models.Organization
	result := db.Where("email = ?", org.Email).First(&updatedInDB)
	assert.NoError(t, result.Error, "Organization should exist in database")
	assert.Equal(t, updatedOrg.Name, updatedInDB.Name)
	assert.Equal(t, updatedOrg.Phone, updatedInDB.Phone)
	assert.Equal(t, updatedOrg.Location, updatedInDB.Location)
	assert.Equal(t, updatedOrg.Description, updatedInDB.Description)
	assert.Equal(t, updatedOrg.Website_Url, updatedInDB.Website_Url)
	
	// Password should remain the same since we didn't update it
	assert.Equal(t, org.Password, updatedInDB.Password)
}

func TestUpdateOrganizationWithPassword(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization
	org := createTestOrganization(db)
	originalPassword := org.Password

	// Updated organization data with new password
	updatedOrg := models.Organization{
		Name:     "Password Update Test",
		Password: "newpassword123",
		Phone:    "6666666666",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedOrg)

	// Create request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/organizations/update/%s", org.Email), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify organization was updated in DB
	var updatedInDB models.Organization
	result := db.Where("email = ?", org.Email).First(&updatedInDB)
	assert.NoError(t, result.Error, "Organization should exist in database")
	assert.Equal(t, updatedOrg.Name, updatedInDB.Name)
	
	// Password should be updated and hashed
	assert.NotEqual(t, originalPassword, updatedInDB.Password, "Password should be different after update")
	assert.NotEqual(t, "newpassword123", updatedInDB.Password, "Password should be hashed")
}

func TestDeleteOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization
	org := createTestOrganization(db)

	// Create request
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/organizations/delete/%s", org.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify organization was deleted from DB
	var count int64
	db.Model(&models.Organization{}).Where("email = ?", org.Email).Count(&count)
	assert.Equal(t, int64(0), count, "Organization should be deleted from database")
}

func TestLoginOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization with known plain password
	plainPassword := "loginTestPassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	
	organization := models.Organization{
		Email:       "login@test.org",
		Password:    string(hashedPassword),
		Name:        "Login Test Org",
		Phone:       "7777777777",
		Location:    "Login Location",
		Description: "Login Description",
		Website_Url: "https://login-test.org",
		Created_At:  time.Now(),
		Updated_At:  time.Now(),
	}
	
	db.Create(&organization)

	// Login credentials
	credentials := LoginRequest{
		Email:    organization.Email,
		Password: plainPassword,
		Role:     "organization",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/organization", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify the response has user data
	userMap, exists := response["user"]
	assert.True(t, exists, "Response should contain user data")
	
	// Check user details if it exists
	if exists {
		user, ok := userMap.(map[string]interface{})
		assert.True(t, ok, "User should be a JSON object")
		
		if ok {
			assert.Equal(t, organization.Email, user["Email"])
			assert.Equal(t, organization.Name, user["Name"])
		}
	}
}

func TestLoginOrganizationInvalidPassword(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create test organization
	org := createTestOrganization(db)

	// Wrong login credentials
	credentials := LoginRequest{
		Email:    org.Email,
		Password: "wrongpassword",
		Role:     "organization",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/organization", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Login with wrong password should fail")
}

func TestLoginOrganizationInvalidEmail(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Wrong login credentials
	credentials := LoginRequest{
		Email:    "nonexistent@org.com",
		Password: "anypassword",
		Role:     "organization",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/organization", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Login with nonexistent email should fail")
}

func TestGetNonExistentOrganization(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Create request for non-existent organization
	req, _ := http.NewRequest("GET", "/organizations/get/nonexistent@org.com", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestInvalidOrganizationData(t *testing.T) {
	db := setupTestDBForOrganization()
	router := setupRouterForOrganization(db)
	defer cleanupTestOrganizations(db)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/organizations/create", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}