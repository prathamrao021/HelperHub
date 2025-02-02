package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prathamrao021/HelperHub/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=User port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})
	fmt.Println("Database connection successfully opened")
}

func Caller(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	initDB()
	r := gin.Default()
	r.GET("/ping", Caller)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
