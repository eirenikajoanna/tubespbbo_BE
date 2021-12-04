package service

import (
	"tubespbbo/db"
	"tubespbbo/mapper"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindItemsShoppingCart(userId int64) (*[]model.ShoppingCart, error) {
	return repository.FindItemsShoppingCart(userId)
}

func FindOneItemShoppingCart(id int64) (*model.ShoppingCart, error) {
	return repository.FindOneItemShoppingCart(id)
}

func CreateOneItemShoppingCart(dto dto.CreateShoppingCart) (*model.ShoppingCart, error) {
	var item *model.ShoppingCart
	item, err := repository.FindOneItemShoppingCartByProductId(dto.ProductId, dto.UserId)
	if err != nil {
		mapper.Map(dto, &item)
		repository.CreateOneItemShoppingCart(item)
	} else {
		item.Quantity = item.Quantity + dto.Quantity
		repository.UpdateOneItemShoppingCart(item)
	}
	return item, nil
}

func UpdateOneItemShoppingCart(updateDto *dto.UpdateShoppingCart, id int64) (*model.ShoppingCart, error) {
	pm, err := repository.FindOneItemShoppingCart(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Quantity >= 0 {
		pm.Quantity = updateDto.Quantity
	}
	pm, err = repository.UpdateOneItemShoppingCart(pm)

	if err != nil {
		return nil, err
	}
	return pm, nil
}

func DeleteOneItemShoppingCart(id int64) (*model.ShoppingCart, error) {
	products, err := repository.FindOneItemShoppingCart(id)
	if err != nil {
		return nil, err
	}
	result := db.Orm.Delete(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
