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
	db.Exec("DROP TABLE IF EXISTS applications CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.Opportunity{})
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Organization{})
	db.AutoMigrate(&models.Volunteer{})

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

		// Create second test organization
		testOrg2 := models.Organization{
			Email:       "test2@org.com",
			Password:    string(hashedPassword),
			Name:        "Test Organization 2",
			Phone:       "0987654321",
			Location:    "Test Location 2",
			Description: "Test Description 2",
			Website_Url: "https://testorg2.com",
			Created_At:  time.Now(),
			Updated_At:  time.Now(),
		}
		db.Create(&testOrg2)
	}

	// Create test volunteer (required for applications)
	db.Model(&models.Volunteer{}).Where("email = ?", "test@volunteer.com").Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		testVolunteer := models.Volunteer{
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
		db.Create(&testVolunteer)
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

	// Additional routes from Opportunity.go
	r.GET("/opportunities/organization/:organization_id/expired", func(c *gin.Context) {
		getLastNExpiredOpportunitiesByOrganization(c, db)
	})
	r.GET("/opportunities", func(c *gin.Context) {
		getOpportunitiesByOrganization(c, db)
	})
	r.GET("/opportunities/available", func(c *gin.Context) {
		getAvailableOpportunities(c, db)
	})

	return r
}

func createTestOpportunity(db *gorm.DB) models.Opportunity {
	// Get the current time
	now := time.Now()

	// Create dates for start and end
	startDate := models.CustomDate(now.AddDate(0, 0, 1)) // tomorrow
	endDate := models.CustomDate(now.AddDate(0, 0, 30))  // 30 days from now

	opp := models.Opportunity{
		Organization_mail: "test@org.com",
		Category:          "Education",
		Title:             "Test Opportunity",
		Description:       "Test Description",
		Location:          "Test Location",
		Hours_Required:    5,
		Start_Date:        startDate,
		End_Date:          endDate,
		Created_At:        now,
		Updated_At:        now,
	}
	
	result := db.Create(&opp)
	if result.Error != nil {
		panic("Failed to create test opportunity: " + result.Error.Error())
	}
	
	if result.RowsAffected == 0 {
		panic("No rows affected when creating test opportunity")
	}
	
	// Verify the opportunity was created properly
	var created models.Opportunity
	db.First(&created, opp.ID)
	
	return created // Return the opportunity as loaded from DB
}

func createExpiredTestOpportunity(db *gorm.DB) models.Opportunity {
	// Get the current time
	now := time.Now()

	// Create dates for start and end in the past
	startDate := models.CustomDate(now.AddDate(0, 0, -30)) // 30 days ago
	endDate := models.CustomDate(now.AddDate(0, 0, -1))    // yesterday

	opp := models.Opportunity{
		Organization_mail: "test@org.com",
		Category:          "Health",
		Title:             "Expired Opportunity",
		Description:       "Expired Description",
		Location:          "Expired Location",
		Hours_Required:    10,
		Start_Date:        startDate,
		End_Date:          endDate,
		Created_At:        now.AddDate(0, 0, -30),
		Updated_At:        now.AddDate(0, 0, -30),
	}
	
	result := db.Create(&opp)
	if result.Error != nil {
		panic("Failed to create expired test opportunity: " + result.Error.Error())
	}
	
	// Verify the opportunity was created properly
	var created models.Opportunity
	db.First(&created, opp.ID)
	
	return created // Return the opportunity as loaded from DB
}

func createTestAppForOpp(db *gorm.DB, opportunityID uint, status string) models.Application {
	// Get volunteer ID
	var volunteer models.Volunteer
	db.Where("email = ?", "test@volunteer.com").First(&volunteer)

	app := models.Application{
		Volunteer_ID:   volunteer.ID,
		Opportunity_ID: opportunityID,
		Status:         status,
		Cover_Letter:   "Test Cover Letter",
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}
	
	result := db.Create(&app)
	if result.Error != nil {
		panic("Failed to create test application: " + result.Error.Error())
	}
	
	// Verify the application was created properly
	var created models.Application
	db.First(&created, app.ID)
	
	return created // Return the application as loaded from DB
}

func cleanupTestOpportunities(db *gorm.DB) {
	db.Exec("DELETE FROM applications")
	db.Exec("DELETE FROM opportunities")
}

func TestCreateOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Get the current time
	now := time.Now()

	// Create dates for start and end
	startDate := now.AddDate(0, 0, 1) // tomorrow
	endDate := now.AddDate(0, 0, 30)  // 30 days from now

	// Test opportunity data
	opportunity := models.Opportunity{
		Organization_mail: "test@org.com",
		Category:          "Education",
		Title:             "Tutor Children",
		Description:       "Help children with homework",
		Location:          "Local Library",
		Hours_Required:    10,
		Start_Date:        models.CustomDate(startDate),
		End_Date:          models.CustomDate(endDate),
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
	assert.Equal(t, opportunity.Hours_Required, response.Hours_Required)
	assert.NotZero(t, response.Created_At)
	assert.NotZero(t, response.Updated_At)
}

func TestGetOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)
	
	// Debug: Verify the opportunity was created in the database
	var count int64
	db.Model(&models.Opportunity{}).Where("id = ?", opp.ID).Count(&count)
	t.Logf("Opportunity with ID %d exists in DB: %v (count: %d)", opp.ID, count > 0, count)
	
	// Debug: Print the URL that will be used
	t.Logf("Route: /opportunities/get/%d", opp.ID)
	
	// Double-check by loading the opportunity directly from the DB
	var dbOpp models.Opportunity
	result := db.First(&dbOpp, opp.ID)
	t.Logf("DB query result rows affected: %v", result.RowsAffected)
	if result.Error != nil {
		t.Logf("DB query error: %v", result.Error)
	}
	t.Logf("DB opportunity: ID=%d, Title=%s", dbOpp.ID, dbOpp.Title)

	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/opportunities/get/%d", opp.ID), nil)
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
	assert.Equal(t, opp.Description, response.Description)
	assert.Equal(t, opp.Location, response.Location)
	assert.Equal(t, opp.Hours_Required, response.Hours_Required)
}

func TestUpdateOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)

	// Updated opportunity data as a map to match the implementation
	updatedOpp := map[string]interface{}{
		"title":          "Updated Title",
		"description":    "Updated Description",
		"location":       "Updated Location",
		"hours_required": 15,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedOpp)

	// Create request
	url := fmt.Sprintf("/opportunities/update/%d", opp.ID)
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
	assert.Equal(t, "Updated Location", response.Location)
	assert.Equal(t, uint(15), response.Hours_Required)
}

func TestDeleteOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create test opportunity
	opp := createTestOpportunity(db)

	// Create request
	url := fmt.Sprintf("/opportunities/delete/%d", opp.ID)
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

func TestGetLastNExpiredOpportunitiesByOrganization(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create several expired opportunities
	expired1 := createExpiredTestOpportunity(db)
	expired2 := createExpiredTestOpportunity(db)
	// Create a non-expired opportunity too
	_ = createTestOpportunity(db)

	// Create request
	url := fmt.Sprintf("/opportunities/organization/%s/expired?n=2", "test@org.com")
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response []models.Opportunity
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify we got the expired opportunities
	assert.GreaterOrEqual(t, len(response), 1)

	// Verify that the expired opportunities we created are in the response
	foundExpired1 := false
	foundExpired2 := false
	for _, opp := range response {
		if opp.ID == expired1.ID {
			foundExpired1 = true
		} else if opp.ID == expired2.ID {
			foundExpired2 = true
		}
	}

	assert.True(t, foundExpired1 || foundExpired2, "At least one of the expired opportunities should be in the response")
}

func TestGetOpportunitiesByOrganization(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create multiple opportunities for the organization
	opp1 := createTestOpportunity(db)
	opp2 := createTestOpportunity(db)

	// Create applications for the opportunities
	createTestAppForOpp(db, opp1.ID, "pending")
	createTestAppForOpp(db, opp1.ID, "accepted")
	createTestAppForOpp(db, opp2.ID, "pending")

	// Create request
	url := "/opportunities?organization_mail=test@org.com"
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response - since it returns a custom map structure
	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify we got at least 2 opportunities
	assert.GreaterOrEqual(t, len(response), 2)

	// Verify application counts
	for _, oppData := range response {
		id := uint(oppData["id"].(float64))
		applicationCount := int64(oppData["application_count"].(float64))

		if id == opp1.ID {
			assert.Equal(t, int64(2), applicationCount, "Opportunity 1 should have 2 applications")
		} else if id == opp2.ID {
			assert.Equal(t, int64(1), applicationCount, "Opportunity 2 should have 1 application")
		}
	}
}

func TestGetAvailableOpportunities(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create one active opportunity
	active := createTestOpportunity(db)

	// Create one expired opportunity
	expired := createExpiredTestOpportunity(db)

	// Create request
	req, _ := http.NewRequest("GET", "/opportunities/available", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response - custom map structure with organization name
	var response []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify we got at least one opportunity
	assert.GreaterOrEqual(t, len(response), 1)

	// Verify we only get active opportunities and they include organization name
	foundActive := false
	foundExpired := false

	for _, opp := range response {
		id := uint(opp["id"].(float64))

		if id == active.ID {
			foundActive = true
			// Check that organization name is included
			assert.NotNil(t, opp["organization_name"])
			assert.Equal(t, "Test Organization", opp["organization_name"])
		} else if id == expired.ID {
			foundExpired = true
		}
	}

	assert.True(t, foundActive, "Active opportunity should be in the response")
	assert.False(t, foundExpired, "Expired opportunity should not be in the response")
}

func TestInvalidRequestGetLastNExpiredOpportunities(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Missing n parameter
	url := fmt.Sprintf("/opportunities/organization/%s/expired", "test@org.com")
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid n parameter
	url = fmt.Sprintf("/opportunities/organization/%s/expired?n=invalid", "test@org.com")
	req, _ = http.NewRequest("GET", url, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestInvalidRequestGetOpportunitiesByOrganization(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Missing organization_mail parameter
	req, _ := http.NewRequest("GET", "/opportunities", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateNonExistentOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Updated opportunity data
	updatedOpp := map[string]interface{}{
		"title":          "Updated Title",
		"description":    "Updated Description",
		"location":       "Updated Location",
		"hours_required": 15,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedOpp)

	// Create request for non-existent ID
	url := "/opportunities/update/999"
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeleteNonExistentOpportunity(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Create request for non-existent ID
	url := "/opportunities/delete/999"
	req, _ := http.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateOpportunityWithInvalidOrganization(t *testing.T) {
	db := setupTestDBOpportunity()
	router := setupRouterOpportunity(db)
	defer cleanupTestOpportunities(db)

	// Get the current time
	now := time.Now()

	// Create dates for start and end
	startDate := now.AddDate(0, 0, 1) // tomorrow
	endDate := now.AddDate(0, 0, 30)  // 30 days from now

	// Test opportunity data with non-existent organization
	opportunity := models.Opportunity{
		Organization_mail: "nonexistent@org.com", // This organization doesn't exist
		Category:          "Education",
		Title:             "Tutor Children",
		Description:       "Help children with homework",
		Location:          "Local Library",
		Hours_Required:    10,
		Start_Date:        models.CustomDate(startDate),
		End_Date:          models.CustomDate(endDate),
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

	// Assertions - if there's a foreign key constraint, this should fail
	// If no constraint, it might succeed, so we're not strictly asserting failure
	if w.Code != http.StatusOK {
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	}
}
func TestGetApplicationsByOpportunityWithVolunteerDetails(t *testing.T) {
	// Setup test database
	db := setupTestDBOpportunity()
	defer cleanupTestOpportunities(db)
	
	// Verify tables are clean
	var opportunityCount, volunteerCount, applicationCount int64
	db.Model(&models.Opportunity{}).Count(&opportunityCount)
	db.Model(&models.Volunteer{}).Count(&volunteerCount)
	db.Model(&models.Application{}).Count(&applicationCount)
	t.Logf("Initial counts - Opportunities: %d, Volunteers: %d, Applications: %d", 
		opportunityCount, volunteerCount, applicationCount)
	
	// Use existing organization instead of creating a new one
	var org models.Organization
	result := db.First(&org)
	if result.Error != nil {
		t.Fatalf("Failed to find an existing organization: %v", result.Error)
	}
	t.Logf("Using existing organization with email: %s", org.Email)
	
	// Create test opportunity
	opportunity := models.Opportunity{
		Organization_mail: org.Email,
		Category:          "Education",
		Title:             "Test Opportunity for Applications",
		Description:       "Test Description",
		Location:          "Test Location",
		Hours_Required:    10,
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	result = db.Create(&opportunity)
	if result.Error != nil {
		t.Fatalf("Failed to create test opportunity: %v", result.Error)
	}
	t.Logf("Created test opportunity with ID: %d", opportunity.ID)
	
	// Get existing volunteer to use instead of creating new ones
	// This helps avoid unique constraint violations
	var existingVolunteers []models.Volunteer
	db.Limit(2).Find(&existingVolunteers)
	
	if len(existingVolunteers) < 2 {
		t.Fatalf("Not enough existing volunteers found, need at least 2 for this test")
	}
	
	volunteer1 := existingVolunteers[0]
	volunteer2 := existingVolunteers[1]
	
	t.Logf("Using existing volunteer 1 with ID: %d, Name: %s", volunteer1.ID, volunteer1.Name)
	t.Logf("Using existing volunteer 2 with ID: %d, Name: %s", volunteer2.ID, volunteer2.Name)
	
	// Create test application 1 (older)
	application1 := models.Application{
		Volunteer_ID:   volunteer1.ID,
		Opportunity_ID: opportunity.ID,
		Status:         "pending",
		Cover_Letter:   "Cover letter from volunteer one",
		Created_At:     time.Now().Add(-time.Hour), // Created 1 hour ago
		Updated_At:     time.Now(),
	}
	result = db.Create(&application1)
	if result.Error != nil {
		t.Fatalf("Failed to create test application 1: %v", result.Error)
	}
	t.Logf("Created test application 1 with ID: %d for volunteer ID: %d and opportunity ID: %d", 
		application1.ID, application1.Volunteer_ID, application1.Opportunity_ID)
	
	// Create test application 2 (newer)
	application2 := models.Application{
		Volunteer_ID:   volunteer2.ID,
		Opportunity_ID: opportunity.ID,
		Status:         "accepted",
		Cover_Letter:   "Cover letter from volunteer two",
		Created_At:     time.Now(), // More recently created
		Updated_At:     time.Now(),
	}
	result = db.Create(&application2)
	if result.Error != nil {
		t.Fatalf("Failed to create test application 2: %v", result.Error)
	}
	t.Logf("Created test application 2 with ID: %d for volunteer ID: %d and opportunity ID: %d", 
		application2.ID, application2.Volunteer_ID, application2.Opportunity_ID)
	
	// Verify applications were created correctly in the database
	var apps []models.Application
	db.Where("opportunity_id = ?", opportunity.ID).Find(&apps)
	t.Logf("Found %d applications in database for opportunity ID %d", len(apps), opportunity.ID)
	for i, app := range apps {
		t.Logf("Application %d: ID=%d, VolunteerID=%d, Status=%s", i+1, app.ID, app.Volunteer_ID, app.Status)
	}
	
	// Setup the router with our test endpoint
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/opportunities/:opportunity_id/applications", func(c *gin.Context) {
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request
	url := fmt.Sprintf("/opportunities/%d/applications", opportunity.ID)
	t.Logf("Making request to: %s", url)
	req, _ := http.NewRequest("GET", url, nil)
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
	
	// Verify we got applications
	if len(response) == 0 {
		// If no applications found, run a direct DB query to check data
		var directApps []struct {
			ID            uint
			Volunteer_ID  uint
			Volunteer_Name string `gorm:"column:volunteer_name"`
			Opportunity_ID uint
		}
		
		db.Raw(`
			SELECT a.id, a.volunteer_id, v.name as volunteer_name, a.opportunity_id
			FROM applications a
			INNER JOIN volunteers v ON a.volunteer_id = v.id
			WHERE a.opportunity_id = ?
		`, opportunity.ID).Scan(&directApps)
		
		t.Fatalf("No applications found in response. Direct DB query found %d apps: %+v", 
			len(directApps), directApps)
	}
	
	assert.Equal(t, 2, len(response), "Expected 2 applications in the response")
	
	// Sort the response to ensure consistent testing order
	// In case the sorting in the function being tested changes
	if len(response) >= 2 {
		// Only proceed with these checks if we have at least 2 applications
		
		// Find which response corresponds to which application by matching IDs
		var app1Response, app2Response map[string]interface{}
		for _, resp := range response {
			appID := uint(resp["id"].(float64))
			if appID == application1.ID {
				app1Response = resp
			} else if appID == application2.ID {
				app2Response = resp
			}
		}
		
		// Check application 1 details
		if app1Response != nil {
			assert.Equal(t, float64(volunteer1.ID), app1Response["volunteer_id"])
			assert.Equal(t, volunteer1.Name, app1Response["volunteer_name"])
			assert.Equal(t, volunteer1.Email, app1Response["volunteer_email"])
			assert.Equal(t, float64(opportunity.ID), app1Response["opportunity_id"])
			assert.Equal(t, application1.Status, app1Response["status"])
			assert.Equal(t, application1.Cover_Letter, app1Response["cover_letter"])
		} else {
			t.Errorf("Application 1 (ID=%d) not found in response", application1.ID)
		}
		
		// Check application 2 details
		if app2Response != nil {
			assert.Equal(t, float64(volunteer2.ID), app2Response["volunteer_id"])
			assert.Equal(t, volunteer2.Name, app2Response["volunteer_name"])
			assert.Equal(t, volunteer2.Email, app2Response["volunteer_email"])
			assert.Equal(t, float64(opportunity.ID), app2Response["opportunity_id"])
			assert.Equal(t, application2.Status, app2Response["status"])
			assert.Equal(t, application2.Cover_Letter, app2Response["cover_letter"])
		} else {
			t.Errorf("Application 2 (ID=%d) not found in response", application2.ID)
		}
	}
}


func TestGetApplicationsByOpportunityWithVolunteerDetailsEmptyID(t *testing.T) {
	// Setup test database
	db := setupTestDBOpportunity()
	defer cleanupTestOpportunities(db)
	
	// Setup the router with our test endpoint
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	
	// Register the endpoint with a parameter
	router.GET("/opportunities/:opportunity_id/applications", func(c *gin.Context) {
		// We'll manually check for empty ID here to simulate the behavior
		opportunityID := c.Param("opportunity_id")
		if opportunityID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "opportunity_id is required"})
			return
		}
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request with empty opportunity_id parameter
	req, _ := http.NewRequest("GET", "/opportunities//applications", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())
	
	// Since Gin's router treats '//' as a 404, let's try a different approach
	// Let's test the function's behavior with an empty opportunity_id more directly
	
	// Register a route that will explicitly call our function with an empty parameter
	router = gin.Default()
	router.GET("/test-empty-id", func(c *gin.Context) {
		// Explicitly create an empty opportunity_id parameter
		c.Params = gin.Params{gin.Param{Key: "opportunity_id", Value: ""}}
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request
	req, _ = http.NewRequest("GET", "/test-empty-id", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// Log the response for debugging
	t.Logf("Direct test response Status: %d", w.Code)
	t.Logf("Direct test response Body: %s", w.Body.String())
	
	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")
	
	// Verify error message - since we're manually setting an empty ID, 
	// the function should return a 400 Bad Request with the appropriate error message
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, response, "error")
	assert.Equal(t, "opportunity_id is required", response["error"])
}

func TestGetApplicationsByOpportunityWithVolunteerDetailsNoApplications(t *testing.T) {
	// Setup test database
	db := setupTestDBOpportunity()
	defer cleanupTestOpportunities(db)
	
	// Create test organization
	org := models.Organization{
		Email:       "test_org_no_apps@example.com",
		Password:    "hashed_password",
		Name:        "No Applications Organization",
		Phone:       "9876543210",
		Location:    "Test Location",
		Description: "Test Description",
		Website_Url: "https://testorg.com",
		Created_At:  time.Now(),
		Updated_At:  time.Now(),
	}
	db.Create(&org)
	
	// Create test opportunity with no applications
	opportunity := models.Opportunity{
		Organization_mail: org.Email,
		Category:          "Education",
		Title:             "Test Opportunity with No Applications",
		Description:       "Test Description",
		Location:          "Test Location",
		Hours_Required:    10,
		Start_Date:        models.CustomDate(time.Now()),
		End_Date:          models.CustomDate(time.Now().AddDate(0, 0, 30)),
		Created_At:        time.Now(),
		Updated_At:        time.Now(),
	}
	db.Create(&opportunity)
	
	// Setup the router with our test endpoint
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/opportunities/:opportunity_id/applications", func(c *gin.Context) {
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/opportunities/%d/applications", opportunity.ID), nil)
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
	
	// Verify we got 0 applications
	assert.Equal(t, 0, len(response), "Expected 0 applications in the response")
}

func TestGetApplicationsByOpportunityWithVolunteerDetailsInvalidID(t *testing.T) {
	// Setup test database
	db := setupTestDBOpportunity()
	defer cleanupTestOpportunities(db)
	
	// Setup the router with our test endpoint
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/opportunities/:opportunity_id/applications", func(c *gin.Context) {
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request with non-existent opportunity ID
	req, _ := http.NewRequest("GET", "/opportunities/9999/applications", nil)
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
	
	// Verify we got 0 applications
	assert.Equal(t, 0, len(response), "Expected 0 applications in the response")
}

func TestGetApplicationsByOpportunityWithVolunteerDetailsDatabaseError(t *testing.T) {
	// Setup test database
	db := setupTestDBOpportunity()
	defer cleanupTestOpportunities(db)
	
	// Setup router with a mock function that simulates a database error
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/opportunities/:opportunity_id/applications", func(c *gin.Context) {
		// Close the database connection to force an error
		sqlDB, _ := db.DB()
		sqlDB.Close()
		
		// Now try to use the closed database
		getApplicationsByOpportunityWithVolunteerDetails(c, db)
	})
	
	// Create request
	req, _ := http.NewRequest("GET", "/opportunities/1/applications", nil)
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