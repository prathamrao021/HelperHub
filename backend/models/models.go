package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// StringList is a custom type for a list of strings
type StringList []string

// driver.Valuer interface for StringList (GO -> DB)
func (s StringList) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// sql.Scanner interface for StringList (DB -> GO)
func (s *StringList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, s)
}

// UintList is a custom type for a list of uints
type UintList []uint

// driver.Valuer interface for UintList (GO -> DB)
func (u UintList) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// sql.Scanner interface for UintList (DB -> GO)
func (u *UintList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, u)
}

// User struct
type User struct {
	ID            uint   `gorm:"primaryKey"`
	Email         string `gorm:"unique;not null"`
	Password_Hash string `gorm:"not null"`
	Full_Name     string `gorm:"not null"`
	Role          string `gorm:"not null"`
	Created_At    time.Time
	Updated_At    time.Time
}

// Volunteer struct
type Volunteer struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Email            string     `gorm:"unique;not null" json:"email"`
	Password         string     `gorm:"not null" json:"password"`
	Name             string     `gorm:"not null" json:"name"`
	Phone            string     `gorm:"unique;not null" json:"phone"`
	Location         string     `json:"location"`
	Bio_Data         string     `json:"bio_data"`
	Category_List    StringList `gorm:"type:json;not null" json:"category_list"`
	Availabile_Hours uint       `gorm:"not null" json:"availabile_hours"`
	Created_At       time.Time  `json:"created_at"`
	Updated_At       time.Time  `json:"updated_at"`
}

// Organization struct
type Organization struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Email       string    `gorm:"unique;not null" json:"email"`
	Password    string    `gorm:"not null" json:"password"`
	Name        string    `gorm:"unique;not null" json:"name"`
	Phone       string    `gorm:"not null" json:"phone"`
	Location    string    `gorm:"not null" json:"location"`
	Description string    `gorm:"not null" json:"description"`
	Website_Url string    `gorm:"not null" json:"website_url"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

// Category struct
type Category struct {
	ID         uint   `gorm:"primaryKey"`
	Category   string `gorm:"unique;not null"`
	Created_At time.Time
}

// Opportunity struct
type Opportunity struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Organization_mail string    `gorm:"not null" json:"organization_mail"`
	Category          string    `gorm:"not null" json:"category"`
	Title             string    `gorm:"not null" json:"title"`
	Description       string    `gorm:"not null" json:"description"`
	Location          string    `gorm:"not null" json:"location"`
	Hours_Required    uint      `gorm:"not null" json:"hours_required"`
	Created_At        time.Time `json:"created_at"`
	Updated_At        time.Time `json:"updated_at"`
}

// Application struct
type Application struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Volunteer_ID   uint      `gorm:"not null" json:"volunteer_ID"`
	Opportunity_ID uint      `gorm:"not null" json:"opportunity_ID"`
	Status         string    `gorm:"not null" json:"status"`
	Cover_Letter   string    `gorm:"not null" json:"cover_Letter"`
	Created_At     time.Time `json:"created_At"`
	Updated_At     time.Time `json:"updated_At"`
}

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
