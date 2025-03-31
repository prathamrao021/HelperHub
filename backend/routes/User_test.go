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

func setupTestDBForUser() *gorm.DB {
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
	db.Exec("DROP TABLE IF EXISTS users CASCADE")

	// Auto migrate required models
	db.AutoMigrate(&models.User{})

	return db
}

func setupRouterForUser(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register routes with injected database
	r.POST("/users/create", func(c *gin.Context) {
		createUser(c, db)
	})
	r.DELETE("/users/delete/:username", func(c *gin.Context) {
		deleteUser(c, db)
	})
	r.PUT("/users/update/:username", func(c *gin.Context) {
		updateUser(c, db)
	})
	r.GET("/users/get/:username", func(c *gin.Context) {
		getUser(c, db)
	})

	return r
}

func createTestUser(db *gorm.DB) models.User {
	// Create a test user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{
		Email:         "test@example.com",
		Password_Hash: string(hashedPassword),
		Full_Name:     "Test User",
		Role:          "volunteer",
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
	}

	result := db.Create(&user)
	if result.Error != nil {
		panic("Failed to create test user: " + result.Error.Error())
	}

	// Verify the user was created properly
	var created models.User
	db.Where("email = ?", user.Email).First(&created)

	return created // Return the user as loaded from DB
}

func cleanupTestUsers(db *gorm.DB) {
	db.Exec("DELETE FROM users")
}

func TestCreateUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Test user data
	user := models.User{
		Email:         "new@example.com",
		Password_Hash: "securepassword", // Will be hashed by the handler
		Full_Name:     "New User",
		Role:          "organization",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(user)

	// Create request
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response models.User
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, user.Email, response.Email)
	assert.Equal(t, user.Full_Name, response.Full_Name)
	assert.Equal(t, user.Role, response.Role)
	assert.NotEqual(t, user.Password_Hash, response.Password_Hash) // Password should be hashed
	assert.NotZero(t, response.Created_At)
	assert.NotZero(t, response.Updated_At)

	// Verify password was hashed properly
	err := bcrypt.CompareHashAndPassword([]byte(response.Password_Hash), []byte("securepassword"))
	assert.NoError(t, err, "Password should be hashed correctly")
}

func TestGetUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Create test user
	user := createTestUser(db)

	// Create request - Note: The route expects a username but the model uses email
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/get/%s", user.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response struct {
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Verify response fields
	assert.Equal(t, "User Data Sent.", response.Message)
	assert.Equal(t, user.Email, response.Data.Email)
	assert.Equal(t, user.Full_Name, response.Data.Full_Name)
	assert.Equal(t, user.Role, response.Data.Role)
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Create test user
	user := createTestUser(db)

	// Updated user data
	updatedUser := user
	updatedUser.Email = "updated@example.com"
	updatedUser.Full_Name = "Updated Name"
	updatedUser.Role = "admin"
	updatedUser.Password_Hash = "newpassword"

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedUser)

	// Create request - Note: The route expects a username but the model uses email
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/users/update/%s", user.Email), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response to check the update confirmation message
	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "User data updated successfully", response["message"])

	// Verify user was updated in the database
	var updatedUserFromDB models.User
	db.Where("email = ?", updatedUser.Email).First(&updatedUserFromDB)

	assert.Equal(t, "Updated Name", updatedUserFromDB.Full_Name)
	assert.Equal(t, "admin", updatedUserFromDB.Role)
	
	// Verify password was hashed properly
	err := bcrypt.CompareHashAndPassword([]byte(updatedUserFromDB.Password_Hash), []byte("newpassword"))
	assert.NoError(t, err, "Password should be hashed correctly")
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Create test user
	user := createTestUser(db)

	// Create request - Note: The route expects a username but the model uses email
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/users/delete/%s", user.Email), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify user is deleted
	var count int64
	db.Model(&models.User{}).Where("email = ?", user.Email).Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestGetNonExistentUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Create request for non-existent user
	req, _ := http.NewRequest("GET", "/users/get/nonexistent@example.com", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions - this should return an internal server error as per the getUser function
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateNonExistentUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Updated user data
	updatedUser := models.User{
		Email:         "nonexistent@example.com",
		Password_Hash: "password123",
		Full_Name:     "Non Existent",
		Role:          "volunteer",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(updatedUser)

	// Create request
	req, _ := http.NewRequest("PUT", "/users/update/nonexistent@example.com", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestInvalidUserData(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Invalid JSON
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDuplicateUser(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Create first user
	user := createTestUser(db)

	// Try to create another user with the same email
	duplicateUser := user
	duplicateUser.Full_Name = "Different Name" // Different name, same email

	// Convert to JSON
	jsonData, _ := json.Marshal(duplicateUser)

	// Create request
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions - this should fail with an internal server error (due to unique constraint)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateUserWithMissingFields(t *testing.T) {
	db := setupTestDBForUser()
	router := setupRouterForUser(db)
	defer cleanupTestUsers(db)

	// Test user data with missing required fields
	incompleteUser := struct {
		Email string `json:"email"`
		// Missing Password_Hash, Full_Name, and Role
	}{
		Email: "incomplete@example.com",
	}

	// Convert to JSON
	jsonData, _ := json.Marshal(incompleteUser)

	// Create request
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Assertions - this should fail with a bad request error
	assert.Equal(t, http.StatusBadRequest, w.Code)
}