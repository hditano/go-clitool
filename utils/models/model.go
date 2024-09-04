package utils

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	First_name string
	Last_name  string
	Email      string
	Address    string
}
