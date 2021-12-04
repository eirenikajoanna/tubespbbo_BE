package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindProduct() (*[]model.Product, error) {
	var products []model.Product
	result := db.Orm.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func FindOneProduct(id int64) (*model.Product, error) {
	var product model.Product
	result := db.Orm.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no product found")
	}

	return &product, nil
}

func CreateProduct(product *model.Product) error {
	result := db.Orm.Create(product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateProduct(product *model.Product) (*model.Product, error) {
	result := db.Orm.Save(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	product, _ = FindOneProduct(product.Id)
	return product, nil
}

func DeleteProduct(product *model.Product) (*model.Product, error) {
	result := db.Orm.Delete(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}
