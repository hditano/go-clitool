package initializers

import (
	"fmt"
	utils "github/hditano/clitool/utils/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnection() {
	os.Remove("database.db")
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func DbMigrate() {
	DB.AutoMigrate(&utils.Customer{})
}

func AddData() {
	customer := DB.Create(&utils.Customer{
		First_name: "Hernan",
		Last_name:  "Di Tano",
		Email:      "hditano@gmail.com",
		Address:    "Imbramowska 3/46",
	})
	fmt.Print(customer)
}
