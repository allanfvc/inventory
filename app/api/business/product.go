package business

import (
	"fmt"

	"github.com/allanfvc/inventory/api/database"
	"github.com/allanfvc/inventory/api/models"
)

func SaveProduct(product *models.Product) {
	database.DBConn.Save(product)
}

func ListProducts(name string) []models.Product {
	var products []models.Product
	if name != "" {
		query := fmt.Sprint("%%", name, "%%")
		database.DBConn.Where("name LIKE ? ", query).Find(&products)
	} else {
		database.DBConn.Find(&products)
	}
	return products
}
