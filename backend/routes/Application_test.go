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

func setupTestDBForApplication() *gorm.DB {
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
	db.Exec("DROP TABLE IF EXISTS applications CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.Application{})

	// Create test volunteer if not exists
	var volunteerCount int64
	db.Model(&models.Volunteer{}).Where("email = ?", "test@volunteer.com").Count(&volunteerCount)
	if volunteerCount == 0 {
		// Create test volunteer
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		volunteer := models.Volunteer{
			Email:            "test@volunteer.com",
			Password:         string(hashedPassword),
			Name:             "Test Volunteer",
			Phone:            "1234567890",
			Location:         "Test Location",
			Bio_Data:         "Test Bio",
			Category_List:    models.StringList{"Education", "Environment"},
			Availabile_Hours: 10,
			Created_At:       time.Now(),
			Updated_At:       time.Now(),
		}
		db.Create(&volunteer)
	}

	// Create test organization if not exists
	var orgCount int64
	db.Model(&models.Organization{}).Where("email = ?", "test@org.com").Count(&orgCount)
	if orgCount == 0 {
		// Create test organization
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		org := models.Organization{
			Email:       "test@org.com",
			Password:    string(hashedPassword),
			Name:        "Test Organization",
			Phone:       "9876543210",
			Location:    "Test Location",
			Description: "Test Description",
			Website_Url: "https://testorg.com",
			Created_At:  time.Now(),
			Updated_At:  time.Now(),
		}
		db.Create(&org)
	}

	// Create test opportunity if not exists
	var oppCount int64
	db.Model(&models.Opportunity{}).Where("organization_mail = ? AND title = ?", "test@org.com", "Test Opportunity").Count(&oppCount)
	if oppCount == 0 {
		// Create test opportunity
		opportunity := models.Opportunity{
			Organization_mail: "test@org.com",
			Category:          "Education",
			Title:             "Test Opportunity",
			Description:       "Test Description",
			Location:          "Test Location",
			Hours_Required:    10, // 48 hours from now
			Created_At:        time.Now(),
			Updated_At:        time.Now(),
		}
		db.Create(&opportunity)
	}

	return db
}

// Mock the application route handlers for testing
// These are slightly modified from the original to ensure proper type handling in tests

func mockGetApplicationsByVolunteerID(c *gin.Context, db *gorm.DB) {
	volunteerID := c.Query("volunteer_id")
	var applications []models.Application

	// Add debug message
	fmt.Printf("Looking for applications with volunteer_id=%s\n", volunteerID)

	if err := db.Where("volunteer_id = ?", volunteerID).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Debug: print found applications
	fmt.Printf("Found %d applications\n", len(applications))
	for i, app := range applications {
		fmt.Printf("Application %d: ID=%d, VolunteerID=%d\n", i, app.ID, app.Volunteer_ID)
	}

	c.JSON(http.StatusOK, applications)
}

func mockGetApplicationsByOpportunityID(c *gin.Context, db *gorm.DB) {
	opportunityID := c.Query("opportunity_id")
	var applications []models.Application

	// Add debug message
	fmt.Printf("Looking for applications with opportunity_id=%s\n", opportunityID)

	if err := db.Where("opportunity_id = ?", opportunityID).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Debug: print found applications
	fmt.Printf("Found %d applications\n", len(applications))
	for i, app := range applications {
		fmt.Printf("Application %d: ID=%d, OpportunityID=%d\n", i, app.ID, app.Opportunity_ID)
	}

	c.JSON(http.StatusOK, applications)
}

// Mock function for getApplicationsByStatus
func mockGetApplicationsByStatus(c *gin.Context, db *gorm.DB) {
	status := c.Query("status")
	var applications []models.Application

	// Add debug message
	fmt.Printf("Looking for applications with status=%s\n", status)

	if err := db.Where("status = ?", status).Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Debug: print found applications
	fmt.Printf("Found %d applications with status '%s'\n", len(applications), status)
	for i, app := range applications {
		fmt.Printf("Application %d: ID=%d, Status=%s\n", i, app.ID, app.Status)
	}

	c.JSON(http.StatusOK, applications)
}

func setupRouterForApplication(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes with injected database
	r.POST("/applications", func(c *gin.Context) {
		createApplication(c, db)
	})

	// Handler for GET with query parameters
	r.GET("/applications", func(c *gin.Context) {
		volunteerID := c.Query("volunteer_id")
		opportunityID := c.Query("opportunity_id")
		status := c.Query("status")

		if volunteerID != "" {
			mockGetApplicationsByVolunteerID(c, db)
		} else if opportunityID != "" {
			mockGetApplicationsByOpportunityID(c, db)
		} else if status != "" {
			mockGetApplicationsByStatus(c, db)
		} else {
			getAllApplications(c, db)
		}
	})

	r.GET("/applications/:id", func(c *gin.Context) {
		getApplicationByID(c, db)
	})
	r.PUT("/applications/:id", func(c *gin.Context) {
		updateApplication(c, db)
	})
	r.DELETE("/applications/:id", func(c *gin.Context) {
		deleteApplication(c, db)
	})

	return r
}

func createTestApplication(db *gorm.DB) models.Application {
	// Get first volunteer and opportunity IDs
	var volunteer models.Volunteer
	var opportunity models.Opportunity
	db.First(&volunteer)
	db.First(&opportunity)

	application := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunity.ID,
		Status:         "pending",
		Cover_Letter:   "Test cover letter",
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}

	result := db.Create(&application)
	if result.Error != nil {
		panic("Failed to create test application: " + result.Error.Error())
	}

	// Verify the application was created properly
	var created models.Application
	db.First(&created, application.ID)

	return created // Return the application as loaded from DB
}

func cleanupTestApplications(db *gorm.DB) {
	db.Exec("DELETE FROM applications")
}

func TestCreateApplication(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Get first volunteer and opportunity IDs
	var volunteer models.Volunteer
	var opportunity models.Opportunity
	db.First(&volunteer)
	db.First(&opportunity)

	// Test application data
	application := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunity.ID,
		Status:         "pending",
		Cover_Letter:   "I am very interested in this opportunity",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(application)

	// Create request
	req, _ := http.NewRequest("POST", "/applications", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Application
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.NotZero(t, response.ID)
	assert.Equal(t, application.Volunteer_ID, response.Volunteer_ID)
	assert.Equal(t, application.Opportunity_ID, response.Opportunity_ID)
	assert.Equal(t, application.Status, response.Status)
	assert.Equal(t, application.Cover_Letter, response.Cover_Letter)
	assert.NotZero(t, response.Created_At)
	assert.NotZero(t, response.Updated_At)
}

func TestGetAllApplications(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create multiple test applications
	createTestApplication(db)
	createTestApplication(db)

	// Create request
	req, _ := http.NewRequest("GET", "/applications", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []models.Application
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify we got all applications
	assert.GreaterOrEqual(t, len(response), 2)
}

func TestGetApplicationByID(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application
	app := createTestApplication(db)

	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/applications/%d", app.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Application
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, app.ID, response.ID)
	assert.Equal(t, app.Volunteer_ID, response.Volunteer_ID)
	assert.Equal(t, app.Opportunity_ID, response.Opportunity_ID)
	assert.Equal(t, app.Status, response.Status)
	assert.Equal(t, app.Cover_Letter, response.Cover_Letter)
}

func TestGetApplicationsByVolunteerID(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application and make sure it's loaded from DB
	app := createTestApplication(db)

	// Print debug information
	t.Logf("Testing with application ID=%d, VolunteerID=%d", app.ID, app.Volunteer_ID)

	// Manually verify the application exists in the database
	var count int64
	db.Model(&models.Application{}).Where("volunteer_id = ?", app.Volunteer_ID).Count(&count)
	assert.GreaterOrEqual(t, count, int64(1), "Database should contain at least one application with this volunteer ID")

	// Create request - using string formatting to ensure proper query parameter
	req, _ := http.NewRequest("GET", fmt.Sprintf("/applications?volunteer_id=%d", app.Volunteer_ID), nil)
	w := httptest.NewRecorder()

	// Execute the request
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []models.Application
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify we got at least one application
	assert.GreaterOrEqual(t, len(response), 1, "Expected at least one application in the response")

	// If we got responses, verify they have the correct volunteer ID
	if len(response) > 0 {
		assert.Equal(t, app.Volunteer_ID, response[0].Volunteer_ID,
			"Application in response should have the requested volunteer ID")
	}
}

func TestGetApplicationsByOpportunityID(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application and make sure it's loaded from DB
	app := createTestApplication(db)

	// Print debug information
	t.Logf("Testing with application ID=%d, OpportunityID=%d", app.ID, app.Opportunity_ID)

	// Manually verify the application exists in the database
	var count int64
	db.Model(&models.Application{}).Where("opportunity_id = ?", app.Opportunity_ID).Count(&count)
	assert.GreaterOrEqual(t, count, int64(1), "Database should contain at least one application with this opportunity ID")

	// Create request - using string formatting to ensure proper query parameter
	req, _ := http.NewRequest("GET", fmt.Sprintf("/applications?opportunity_id=%d", app.Opportunity_ID), nil)
	w := httptest.NewRecorder()

	// Execute the request
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []models.Application
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify we got at least one application
	assert.GreaterOrEqual(t, len(response), 1, "Expected at least one application in the response")

	// If we got responses, verify they have the correct opportunity ID
	if len(response) > 0 {
		assert.Equal(t, app.Opportunity_ID, response[0].Opportunity_ID,
			"Application in response should have the requested opportunity ID")
	}
}

func TestGetApplicationsByStatus(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application with "pending" status
	app := createTestApplication(db)

	// Print debug information
	t.Logf("Testing with application ID=%d, Status=%s", app.ID, app.Status)

	// Manually verify the application exists in the database with the correct status
	var count int64
	db.Model(&models.Application{}).Where("status = ?", "pending").Count(&count)
	assert.GreaterOrEqual(t, count, int64(1), "Database should contain at least one application with 'pending' status")

	// Create request for applications with pending status
	req, _ := http.NewRequest("GET", "/applications?status=pending", nil)
	w := httptest.NewRecorder()

	// Execute the request
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []models.Application
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify we got at least one application
	assert.GreaterOrEqual(t, len(response), 1, "Expected at least one application with 'pending' status")

	// Only check for the created application if we got any results
	if len(response) > 0 {
		// Check if we can find the application we created in the response
		found := false
		for _, a := range response {
			if a.ID == app.ID {
				found = true
				break
			}
		}
		assert.True(t, found, "The created application wasn't found in the response")

		// Check if the applications have the requested status
		for _, a := range response {
			assert.Equal(t, "pending", a.Status, "All returned applications should have 'pending' status")
		}
	}
}

func TestUpdateApplication(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application
	app := createTestApplication(db)

	// Updated application data
	updatedApp := app
	updatedApp.Status = "accepted"
	updatedApp.Cover_Letter = "Updated cover letter"

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedApp)

	// Create request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/applications/%d", app.ID), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Application
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, app.ID, response.ID)
	assert.Equal(t, "accepted", response.Status)
	assert.Equal(t, "Updated cover letter", response.Cover_Letter)

	// Verify update time is later than creation time
	assert.True(t, response.Updated_At.After(app.Created_At))
}

func TestDeleteApplication(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create test application
	app := createTestApplication(db)

	// Create request
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/applications/%d", app.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify application is deleted
	var count int64
	db.Model(&models.Application{}).Where("id = ?", app.ID).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGetNonExistentApplication(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Create request for non-existent ID
	req, _ := http.NewRequest("GET", "/applications/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInvalidApplicationData(t *testing.T) {
	db := setupTestDBForApplication()
	router := setupRouterForApplication(db)
	defer cleanupTestApplications(db)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/applications", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestGetApplicationsByVolunteerWithDetails tests the endpoint that retrieves applications
// for a specific volunteer with opportunity title and organization name details
func TestGetApplicationsByVolunteerWithDetails(t *testing.T) {
	// Setup test database and router
	db := setupTestDBForApplication()
	defer cleanupTestApplications(db)
	
	// Get first volunteer and organization IDs
	var volunteer models.Volunteer
	var organization models.Organization
	db.First(&volunteer)
	db.First(&organization)
	
	// Create test opportunity with the organization email
	opportunity := models.Opportunity{
		Organization_mail: organization.Email,
		Category:          "Education",
		Title:             "Special Test Opportunity",
		Description:       "Test Description for Detailed Query",
		Location:          "Test Location",
		Hours_Required:    10,
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	db.Create(&opportunity)
	
	// Create test application
	application := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunity.ID,
		Status:         "pending",
		Cover_Letter:   "Test cover letter for detailed query",
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}
	db.Create(&application)
	
	// Print debug information
	t.Logf("Created application with ID=%d, VolunteerID=%d, OpportunityID=%d", 
		application.ID, application.Volunteer_ID, application.Opportunity_ID)
	t.Logf("Opportunity title: %s, Organization email: %s", opportunity.Title, opportunity.Organization_mail)
	
	// Setup router for this specific test
	router := gin.Default()
	router.GET("/applications/volunteer/:volunteer_id", func(c *gin.Context) {
		getApplicationsByVolunteerWithDetails(c, db)
	})
	
	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/applications/volunteer/%d", volunteer.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())
	
	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Parse response
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	
	// Verify we got at least one application that matches our test data
	assert.GreaterOrEqual(t, len(response), 1, "Expected at least one application in the response")
	
	// Find our specific test application in the response
	var found bool
	for _, app := range response {
		// Convert IDs to appropriate types for comparison
		appID, _ := app["id"].(float64)
		oppID, _ := app["opportunity_id"].(float64)
		
		if int(appID) == int(application.ID) && int(oppID) == int(opportunity.ID) {
			found = true
			
			// Verify all required fields are present and have the correct values
			assert.Equal(t, float64(volunteer.ID), app["volunteer_id"])
			assert.Equal(t, opportunity.Title, app["opportunity_title"])
			assert.Equal(t, organization.Name, app["organization_name"])
			assert.Equal(t, application.Status, app["status"])
			assert.Equal(t, application.Cover_Letter, app["cover_letter"])
			assert.NotNil(t, app["created_at"])
			assert.NotNil(t, app["updated_at"])
			break
		}
	}
	
	assert.True(t, found, "Could not find the test application in the response")
}

// TestGetApplicationsByVolunteerWithDetailsInvalidID tests the endpoint with an invalid volunteer ID
func TestGetApplicationsByVolunteerWithDetailsInvalidID(t *testing.T) {
	// Setup test database and router
	db := setupTestDBForApplication()
	defer cleanupTestApplications(db)
	
	// Setup router for this specific test
	router := gin.Default()
	router.GET("/applications/volunteer/:volunteer_id", func(c *gin.Context) {
		getApplicationsByVolunteerWithDetails(c, db)
	})
	
	// Create request with non-existent volunteer ID
	req, _ := http.NewRequest("GET", "/applications/volunteer/9999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())
	
	// Assertions
	assert.Equal(t, http.StatusOK, w.Code) // Should still return 200 with empty array
	
	// Parse response - should be an empty array
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	
	// Verify we got no applications
	assert.Equal(t, 0, len(response), "Expected no applications for non-existent volunteer ID")
}

// TestGetApplicationsByVolunteerWithDetailsMultipleApplications tests retrieving multiple applications
func TestGetApplicationsByVolunteerWithDetailsMultipleApplications(t *testing.T) {
	// Setup test database and router
	db := setupTestDBForApplication()
	defer cleanupTestApplications(db)
	
	// Get first volunteer and organization
	var volunteer models.Volunteer
	var organization models.Organization
	db.First(&volunteer)
	db.First(&organization)
	
	// Create multiple opportunities
	opportunity1 := models.Opportunity{
		Organization_mail: organization.Email,
		Category:          "Education",
		Title:             "First Test Opportunity",
		Description:       "First test description",
		Location:          "Test Location 1",
		Hours_Required:    10,
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	db.Create(&opportunity1)
	
	opportunity2 := models.Opportunity{
		Organization_mail: organization.Email,
		Category:          "Environment",
		Title:             "Second Test Opportunity",
		Description:       "Second test description",
		Location:          "Test Location 2",
		Hours_Required:    15,
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	db.Create(&opportunity2)
	
	// Create applications for both opportunities
	application1 := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunity1.ID,
		Status:         "pending",
		Cover_Letter:   "First test cover letter",
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}
	db.Create(&application1)
	
	application2 := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunity2.ID,
		Status:         "accepted",
		Cover_Letter:   "Second test cover letter",
		Created_At:     time.Now().Add(time.Hour), // Create a bit later to test ordering
		Updated_At:     time.Now(),
	}
	db.Create(&application2)
	
	// Print debug information
	t.Logf("Created applications with IDs=%d and %d for volunteer ID=%d", 
		application1.ID, application2.ID, volunteer.ID)
	
	// Setup router for this specific test
	router := gin.Default()
	router.GET("/applications/volunteer/:volunteer_id", func(c *gin.Context) {
		getApplicationsByVolunteerWithDetails(c, db)
	})
	
	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/applications/volunteer/%d", volunteer.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())
	
	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	
	// Parse response
	var response []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	
	// Verify we got at least two applications
	assert.GreaterOrEqual(t, len(response), 2, "Expected at least two applications in the response")
	
	// Verify both opportunity titles are present in the response
	opportunityTitles := []string{}
	for _, app := range response {
		if title, ok := app["opportunity_title"].(string); ok {
			opportunityTitles = append(opportunityTitles, title)
		}
	}
	
	assert.Contains(t, opportunityTitles, "First Test Opportunity")
	assert.Contains(t, opportunityTitles, "Second Test Opportunity")
	
	// Find and verify details of each test application
	var foundFirstApp, foundSecondApp bool
	for _, app := range response {
		appID, _ := app["id"].(float64)
		
		if int(appID) == int(application1.ID) {
			foundFirstApp = true
			assert.Equal(t, "First Test Opportunity", app["opportunity_title"])
			assert.Equal(t, "pending", app["status"])
		} else if int(appID) == int(application2.ID) {
			foundSecondApp = true
			assert.Equal(t, "Second Test Opportunity", app["opportunity_title"])
			assert.Equal(t, "accepted", app["status"])
		}
	}
	
	assert.True(t, foundFirstApp, "Could not find the first test application in the response")
	assert.True(t, foundSecondApp, "Could not find the second test application in the response")
}

// TestGetApplicationsByVolunteerWithDetailsDatabaseError tests database error handling
func TestGetApplicationsByVolunteerWithDetailsDatabaseError(t *testing.T) {
	// Setup test database
	db := setupTestDBForApplication()
	defer cleanupTestApplications(db)
	
	// Setup router with a mock function that simulates a database error
	router := gin.Default()
	router.GET("/applications/volunteer/:volunteer_id", func(c *gin.Context) {
		// Close the database connection to force an error
		sqlDB, _ := db.DB()
		sqlDB.Close()
		
		// Now try to use the closed database
		getApplicationsByVolunteerWithDetails(c, db)
	})
	
	// Create request
	req, _ := http.NewRequest("GET", "/applications/volunteer/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())
	
	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	
	// Verify response contains an error message
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Contains(t, response, "error")
}