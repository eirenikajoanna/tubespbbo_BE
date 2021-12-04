package controller

import (
	"net/http"
	"strconv"
	e "tubespbbo/err"
	"tubespbbo/mapper"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/service"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func FindOneUserShoppingCart(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	cart, err := service.FindItemsShoppingCart(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.ShoppingCart
	mapper.Map(cart, &DTOs)

	if cart == nil {
		DTOs = []dto.ShoppingCart{}
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func CreateShoppingCart(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	createDto := new(dto.CreateShoppingCart)
	err := c.BodyParser(createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	cart, err := service.CreateOneItemShoppingCart(*createDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: cart,
	})
	return nil
}

func UpdateOneItemShoppingCart(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdateShoppingCart)
	err = c.BodyParser(updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	err = validate.Validate(&updateDto)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	pm, err := service.UpdateOneItemShoppingCart(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ShoppingCart
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteOneItemShoppingCart(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	pm, err := service.DeleteOneItemShoppingCart(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ShoppingCart
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
