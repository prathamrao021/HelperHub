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
	ID               uint   `gorm:"primaryKey"`
	Phone            string `gorm:"unique;not null"`
	Location         string `gorm:"not null"`
	Bio_Data         string `gorm:"not null"`
	Availabile_Hours int    `gorm:"not null"`
	Created_At       time.Time
	Updated_At       time.Time
}
