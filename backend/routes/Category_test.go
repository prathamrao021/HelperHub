package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// setupTestDBForCategory initializes a PostgreSQL database for category testing
func setupTestDBForCategory() (*gorm.DB, error) {
	// Use the same PostgreSQL connection as in main.go but with a test database
	dsn := "host=localhost user=postgres password=admin dbname=Helperhub_test port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	
	// Configure gorm with minimal logging during tests
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	// Open connection to PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	// Clean up existing tables
	db.Exec("DROP TABLE IF EXISTS categories CASCADE")

	// Migrate the schema
	err = db.AutoMigrate(&models.Category{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// setupCategoryRouter creates a test router with the test database
func setupCategoryRouter(db *gorm.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Set up the routes for categories
	categoryRouter := router.Group("/categories")
	categoryRouter.POST("/create", func(c *gin.Context) { CreateCategory(c, db) })
	categoryRouter.GET("/get", func(c *gin.Context) { getCategories(c, db) })

	return router
}

func cleanupCategoryTable(db *gorm.DB) {
	db.Exec("DELETE FROM categories")
}

// TestCreateCategory tests the category creation endpoint
func TestCreateCategory(t *testing.T) {
	db, err := setupTestDBForCategory()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer cleanupCategoryTable(db)

	router := setupCategoryRouter(db)

	// Create a test request
	req, _ := http.NewRequest("POST", "/categories/create", nil)

	// Perform the request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Categories created successfully", response["message"])

	// Verify categories were created in the database
	var categories []models.Category
	err = db.Find(&categories).Error
	assert.Nil(t, err)
	
	// Check that we have the expected number of categories
	expectedCategories := []string{
		"Web Development",
		"Graphic Design",
		"Content Writing",
		"Social Media",
		"Teaching",
		"Event Planning",
		"Photography",
		"Translation",
		"First Aid",
		"Project Management",
	}
	assert.Equal(t, len(expectedCategories), len(categories))

	// Check that all expected categories exist
	for _, expectedCategory := range expectedCategories {
		var count int64
		db.Model(&models.Category{}).Where("category = ?", expectedCategory).Count(&count)
		assert.Equal(t, int64(1), count, "Category %s should exist", expectedCategory)
	}
}

// TestGetCategories tests retrieving all categories
func TestGetCategories(t *testing.T) {
	db, err := setupTestDBForCategory()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer cleanupCategoryTable(db)

	router := setupCategoryRouter(db)

	// First create the categories
	CreateCategory(nil, db)

	// Create a test request to get categories
	req, _ := http.NewRequest("GET", "/categories/get", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var categories []models.Category
	err = json.Unmarshal(w.Body.Bytes(), &categories)
	assert.Nil(t, err)

	// Verify we have the expected number of categories
	expectedCategories := []string{
		"Web Development",
		"Graphic Design",
		"Content Writing",
		"Social Media",
		"Teaching",
		"Event Planning",
		"Photography",
		"Translation",
		"First Aid",
		"Project Management",
	}
	assert.Equal(t, len(expectedCategories), len(categories))

	// Check that all expected categories exist in the response
	categoryMap := make(map[string]bool)
	for _, category := range categories {
		categoryMap[category.Category] = true
	}
	
	for _, expectedCategory := range expectedCategories {
		assert.True(t, categoryMap[expectedCategory], "Response should include category: %s", expectedCategory)
	}
}

// TestCreateCategoryIdempotence tests that creating categories multiple times doesn't result in duplicates
func TestCreateCategoryIdempotence(t *testing.T) {
	db, err := setupTestDBForCategory()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer cleanupCategoryTable(db)

	router := setupCategoryRouter(db)

	// Call create categories endpoint twice
	req1, _ := http.NewRequest("POST", "/categories/create", nil)
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusOK, w1.Code)

	req2, _ := http.NewRequest("POST", "/categories/create", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)

	// Verify categories were created without duplicates
	var categories []models.Category
	err = db.Find(&categories).Error
	assert.Nil(t, err)
	
	// Check that we have the expected number of categories (no duplicates)
	expectedCategories := []string{
		"Web Development",
		"Graphic Design",
		"Content Writing",
		"Social Media",
		"Teaching",
		"Event Planning",
		"Photography",
		"Translation",
		"First Aid",
		"Project Management",
	}
	assert.Equal(t, len(expectedCategories), len(categories), "Should not create duplicate categories")
}

// TestCategoryDatabaseError tests error handling when database operations fail
func TestCategoryDatabaseError(t *testing.T) {
	db, err := setupTestDBForCategory()
	if err != nil {
		t.Fatalf("Failed to set up test database: %v", err)
	}
	defer cleanupCategoryTable(db)

	// Close the database connection to force errors
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get DB connection: %v", err)
		return
	}
	sqlDB.Close()

	router := setupCategoryRouter(db)

	// Try to create categories with a closed database
	req, _ := http.NewRequest("POST", "/categories/create", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Log the response for debugging
	t.Logf("Response Status: %d", w.Code)
	t.Logf("Response Body: %s", w.Body.String())

	// Should get an internal server error
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Try to get categories with a closed database
	req2, _ := http.NewRequest("GET", "/categories/get", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// Log the response for debugging
	t.Logf("Response Status: %d", w2.Code)
	t.Logf("Response Body: %s", w2.Body.String())

	// Should get an internal server error
	assert.Equal(t, http.StatusInternalServerError, w2.Code)
}