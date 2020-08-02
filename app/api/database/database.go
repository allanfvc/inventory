package database

import (
	"fmt"

	"github.com/allanfvc/inventory/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open("sqlite3", "inventory.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&models.Product{})
}

func CloseDatabase() {
	if DBConn != nil {
		DBConn.Close()
	}
}
