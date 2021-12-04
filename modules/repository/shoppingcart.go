package repository

import (
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindItemsShoppingCart(userId int64) (*[]model.ShoppingCart, error) {
	var item []model.ShoppingCart
	result := db.Orm.Where("user_id = ?", userId).Preload("Product").Find(&item).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func FindOneItemShoppingCart(id int64) (*model.ShoppingCart, error) {
	var pm model.ShoppingCart
	result := db.Orm.First(&pm, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pm, nil
}

func FindOneItemShoppingCartByProductId(productId int64, userId int64) (*model.ShoppingCart, error) {
	var pm model.ShoppingCart
	result := db.Orm.Where("product_id = ? AND user_id = ?", productId, userId).Preload("Product").First(&pm).Not("deleted_at = ?", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	return &pm, nil
}

func CreateOneItemShoppingCart(pm *model.ShoppingCart) error {
	result := db.Orm.Create(&pm)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateOneItemShoppingCart(pm *model.ShoppingCart) (*model.ShoppingCart, error) {
	result := db.Orm.Save(&pm)
	if result.Error != nil {
		return nil, result.Error
	}
	return pm, nil
}

func DeleteOneItemShoppingCart(pm *model.ShoppingCart) (*model.ShoppingCart, error) {
	result := db.Orm.Delete(&pm)
	if result.Error != nil {
		return nil, result.Error
	}

	return pm, nil
}
