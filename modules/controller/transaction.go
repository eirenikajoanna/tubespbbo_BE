package controller

import (
	"net/http"
	"strconv"
	e "tubespbbo/err"
	"tubespbbo/mapper"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/service"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func FindTransaction(c *fiber.Ctx) error {
	AuthCheck(c)
	pm, err := service.FindTransaction()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.TransactionDTO
	mapper.Map(pm, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneTransaction(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	pm, err := service.FindOneTransaction(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateTransaction(c *fiber.Ctx) error {
	AuthCheck(c)
	createDto := new(dto.CreateTransactionDTO)
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

	var transaction model.Transaction
	mapper.Map(createDto, &transaction)
	transaction.Status = "Belum Terverifikasi"

	err = service.CreateTransaction(&transaction)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: transaction,
	})
	return nil
}

func UpdateTransaction(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdateTransactionDTO)
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

	pm, err := service.UpdateTransaction(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteTransaction(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	pm, err := service.DeleteTransaction(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
