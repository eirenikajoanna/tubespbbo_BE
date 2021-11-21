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

func FindTransactionDetail(c *fiber.Ctx) error {
	AuthCheck(c)
	pm, err := service.FindTransactionDetail()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.TransactionDetailDTO
	mapper.Map(pm, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneTransactionDetail(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	pm, err := service.FindOneTransactionDetail(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDetailDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateTransactionDetail(c *fiber.Ctx) error {
	AuthCheck(c)
	createDto := new(dto.CreateTransactionDetailDTO)
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

	var pm model.TransactionDetail
	mapper.Map(createDto, &pm)

	err = service.CreateTransactionDetail(&pm)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: pm,
	})
	return nil
}

func UpdateTransactionDetail(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdateTransactionDetailDTO)
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

	pm, err := service.UpdateTransactionDetail(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDetailDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteTransactionDetail(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	pm, err := service.DeleteTransactionDetail(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.TransactionDetailDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
