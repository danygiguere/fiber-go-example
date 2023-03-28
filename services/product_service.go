package services

import (
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := configs.DBConn.Model(&models.Product{}).Find(&products).Error
	return products, err
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	err := configs.DBConn.Model(&models.Product{}).First(&product, id).Error
	return product, err
}

func CreateProduct(product models.Product) (models.Product, error) {
	err := configs.DBConn.Create(&product).Error
	return product, err
}

func UpdateProduct(product models.Product, id string) error {
	updatedProduct := models.Product{Code: product.Code, Price: product.Price}
	err := configs.DBConn.Model(&product).Where("id = ?", id).Updates(updatedProduct).Error
	return err
}

func DeleteProduct(id string) error {
	err := configs.DBConn.Delete(&models.Product{}, id).Error
	return err
}
