package db

import (
	"log"
	"time"
	"github.com/2byn-server/model"
)

func InsertNewProduct(newProduct *model.Product) (*model.Product, error) {
	log.Println("adding product")
	db := GetDataBaseInstance()

	product := &model.Product{
		Title: newProduct.Title,
		PhotoURL:   newProduct.PhotoURL,
		Description: newProduct.Description,
		Contacts: newProduct.Contacts,
		Approved: false,
		Date:   time.Now()}

	err := db.Create(&product)

	log.Println(product)

	return product, err.Error
}

func GetNotApprovedProducts() ([]model.Product, error) {
	db := GetDataBaseInstance()
	var products []model.Product
	err := db.Where("approved = ?", false).Find(&products)

	return products, err.Error
}

func GetApprovedProducts() ([]model.Product, error) {
	db := GetDataBaseInstance()
	var products []model.Product
	err := db.Where("approved = ?", true).Find(&products)

	return products, err.Error
}

func GetProductByID(id int64) (model.Product, error) {
	db := GetDataBaseInstance()
	var product model.Product
	err := db.Where("id = ?", id).First(&product)

	return product, err.Error
}

func ApproveProduct(id int) (model.Product, error) {
	db := GetDataBaseInstance()
	var product model.Product
	err := db.Model(&product).Where("id = ?", id).Update("Approved", true)

	return product, err.Error
}

func DeleteProduct(id int) (model.Product, error) {
	db := GetDataBaseInstance()
	var product model.Product
	err := db.Where("id = ?", id).First(&product)
	err = db.Delete(&product);

	return product, err.Error
}
