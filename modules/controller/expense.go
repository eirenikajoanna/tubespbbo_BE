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

func FindExpense(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	pm, err := service.FindExpense()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.ExpenseDTO
	mapper.Map(pm, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOneExpense(c *fiber.Ctx) error {
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
	pm, err := service.FindOneExpense(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ExpenseDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreateExpense(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	createDto := new(dto.CreateExpenseDTO)
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

	var expense model.Expense
	mapper.Map(createDto, &expense)

	err = service.CreateExpense(&expense)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: expense,
	})
	return nil
}

func UpdateExpense(c *fiber.Ctx) error {
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
	updateDto := new(dto.UpdateExpenseDTO)
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

	pm, err := service.UpdateExpense(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ExpenseDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeleteExpense(c *fiber.Ctx) error {
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

	pm, err := service.DeleteExpense(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.ExpenseDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
