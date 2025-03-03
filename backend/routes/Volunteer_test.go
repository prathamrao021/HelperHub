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

// LoginRequest struct for testing
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func setupTestDBForVolunteer() *gorm.DB {
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
	db.Exec("DROP TABLE IF EXISTS volunteers CASCADE")
	db.Exec("DROP TABLE IF EXISTS categories CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.Volunteer{})

	return db
}

func setupRouterForVolunteer(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes with injected database
	r.POST("/volunteers/create", func(c *gin.Context) {
		createVolunteer(c, db)
	})
	r.DELETE("/volunteers/delete/:volunteer_mail", func(c *gin.Context) {
		deleteVolunteer(c, db)
	})
	r.PUT("/volunteers/update/:volunteer_mail", func(c *gin.Context) {
		updateVolunteer(c, db)
	})
	r.GET("/volunteers/get/:volunteer_mail", func(c *gin.Context) {
		getVolunteer(c, db)
	})
	r.POST("/login/volunteer", func(c *gin.Context) {
		loginVolunteer(c, db)
	})

	return r
}

func createTestVolunteer(db *gorm.DB) models.Volunteer {
	// Create a test volunteer with hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	
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
	
	result := db.Create(&volunteer)
	if result.Error != nil {
		panic("Failed to create test volunteer: " + result.Error.Error())
	}
	
	// Verify the volunteer was created properly
	var created models.Volunteer
	db.Where("email = ?", volunteer.Email).First(&created)
	
	return created
}

func cleanupTestVolunteers(db *gorm.DB) {
	db.Exec("DELETE FROM volunteers")
}

func TestCreateVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Test volunteer data
	volunteer := models.Volunteer{
		Email:            "new@volunteer.com",
		Password:         "password123",
		Name:             "New Volunteer",
		Phone:            "9876543210",
		Location:         "New Location",
		Bio_Data:         "New Bio",
		Category_List:    models.StringList{"Health", "Technology"},
		Availabile_Hours: 15,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(volunteer)

	// Create request
	req, _ := http.NewRequest("POST", "/volunteers/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify volunteer was created in DB
	var createdVolunteer models.Volunteer
	result := db.Where("email = ?", volunteer.Email).First(&createdVolunteer)
	assert.NoError(t, result.Error, "Volunteer should be created in database")
	assert.Equal(t, volunteer.Email, createdVolunteer.Email)
	assert.Equal(t, volunteer.Name, createdVolunteer.Name)
	assert.NotEqual(t, "password123", createdVolunteer.Password, "Password should be hashed")
}

func TestGetVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer
	volunteer := createTestVolunteer(db)

	// Create request
	req, _ := http.NewRequest("GET", fmt.Sprintf("/volunteers/get/%s", volunteer.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.Volunteer
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Response should be valid JSON")

	// Verify response fields
	assert.Equal(t, volunteer.Email, response.Email)
	assert.Equal(t, volunteer.Name, response.Name)
	assert.Equal(t, volunteer.Phone, response.Phone)
	assert.Equal(t, volunteer.Location, response.Location)
	assert.Equal(t, volunteer.Bio_Data, response.Bio_Data)
	assert.Equal(t, volunteer.Availabile_Hours, response.Availabile_Hours)
	
	// Check category list
	assert.Equal(t, len(volunteer.Category_List), len(response.Category_List))
	for i, category := range volunteer.Category_List {
		assert.Equal(t, category, response.Category_List[i])
	}
}

func TestUpdateVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer
	volunteer := createTestVolunteer(db)

	// Updated volunteer data
	updatedVolunteer := models.Volunteer{
		Name:             "Updated Name",
		Phone:            volunteer.Phone, // Keep the same phone to avoid unique constraint
		Location:         "Updated Location",
		Bio_Data:         "Updated Bio",
		Category_List:    models.StringList{"Health", "Education"},
		Availabile_Hours: 20,
		// Don't update password
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedVolunteer)

	// Create request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/volunteers/update/%s", volunteer.Email), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify volunteer was updated in DB
	var updatedInDB models.Volunteer
	result := db.Where("email = ?", volunteer.Email).First(&updatedInDB)
	assert.NoError(t, result.Error, "Volunteer should exist in database")
	assert.Equal(t, updatedVolunteer.Name, updatedInDB.Name)
	assert.Equal(t, updatedVolunteer.Location, updatedInDB.Location)
	assert.Equal(t, updatedVolunteer.Bio_Data, updatedInDB.Bio_Data)
	assert.Equal(t, updatedVolunteer.Availabile_Hours, updatedInDB.Availabile_Hours)
	
	// Password should remain the same since we didn't update it
	assert.Equal(t, volunteer.Password, updatedInDB.Password)
}

func TestUpdateVolunteerWithPassword(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer
	volunteer := createTestVolunteer(db)
	originalPassword := volunteer.Password

	// Updated volunteer data with new password
	updatedVolunteer := models.Volunteer{
		Name:             "Password Update Test",
		Phone:            volunteer.Phone, // Keep the same phone to avoid unique constraint
		Password:         "newpassword123",
		Availabile_Hours: 25,
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedVolunteer)

	// Create request
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/volunteers/update/%s", volunteer.Email), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify volunteer was updated in DB
	var updatedInDB models.Volunteer
	result := db.Where("email = ?", volunteer.Email).First(&updatedInDB)
	assert.NoError(t, result.Error, "Volunteer should exist in database")
	assert.Equal(t, updatedVolunteer.Name, updatedInDB.Name)
	
	// Password should be updated and hashed
	assert.NotEqual(t, originalPassword, updatedInDB.Password, "Password should be different after update")
	assert.NotEqual(t, "newpassword123", updatedInDB.Password, "Password should be hashed")
}

func TestDeleteVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer
	volunteer := createTestVolunteer(db)

	// Create request
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/volunteers/delete/%s", volunteer.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify volunteer was deleted from DB
	var count int64
	db.Model(&models.Volunteer{}).Where("email = ?", volunteer.Email).Count(&count)
	assert.Equal(t, int64(0), count, "Volunteer should be deleted from database")
}

func TestLoginVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer with known plain password
	plainPassword := "loginTestPassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	
	volunteer := models.Volunteer{
		Email:            "login@test.com",
		Password:         string(hashedPassword),
		Name:             "Login Test",
		Phone:            "5555555555",
		Location:         "Login Location",
		Bio_Data:         "Login Bio",
		Category_List:    models.StringList{"Education"},
		Availabile_Hours: 5,
		Created_At:       time.Now(),
		Updated_At:       time.Now(),
	}
	
	db.Create(&volunteer)

	// Login credentials
	credentials := LoginRequest{
		Email:    volunteer.Email,
		Password: plainPassword,
		Role:     "volunteer",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/volunteer", bytes.NewBuffer(jsonData))
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
			// Check for case-sensitivity in JSON keys
			email, emailExists := user["email"] // Try lowercase first
			if !emailExists {
				email, emailExists = user["Email"] // Then try with uppercase first letter
			}
			assert.True(t, emailExists, "Email field should exist in response")
			assert.Equal(t, volunteer.Email, email)
			
			name, nameExists := user["name"] // Try lowercase first
			if !nameExists {
				name, nameExists = user["Name"] // Then try with uppercase first letter
			}
			assert.True(t, nameExists, "Name field should exist in response")
			assert.Equal(t, volunteer.Name, name)
		}
	}
}

func TestLoginVolunteerInvalidPassword(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create test volunteer
	volunteer := createTestVolunteer(db)

	// Wrong login credentials
	credentials := LoginRequest{
		Email:    volunteer.Email,
		Password: "wrongpassword",
		Role:     "volunteer",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/volunteer", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Login with wrong password should fail")
}

func TestLoginVolunteerInvalidEmail(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Wrong login credentials
	credentials := LoginRequest{
		Email:    "nonexistent@email.com",
		Password: "anypassword",
		Role:     "volunteer",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(credentials)

	// Create request
	req, _ := http.NewRequest("POST", "/login/volunteer", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Login with nonexistent email should fail")
}

func TestGetNonExistentVolunteer(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Create request for non-existent volunteer
	req, _ := http.NewRequest("GET", "/volunteers/get/nonexistent@email.com", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestInvalidVolunteerData(t *testing.T) {
	db := setupTestDBForVolunteer()
	router := setupRouterForVolunteer(db)
	defer cleanupTestVolunteers(db)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/volunteers/create", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}