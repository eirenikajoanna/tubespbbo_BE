package controller

import (
	"net/http"
	"strconv"
	"tubespbbo/domains/payment/dto"
	"tubespbbo/domains/payment/model"
	"tubespbbo/domains/payment/service"
	e "tubespbbo/err"
	"tubespbbo/mapper"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

func FindPaymentMethod(c *fiber.Ctx) error {
	pm, err := service.FindPaymentMethod()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PaymentMethodDTO
	mapper.Map(pm, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOnePaymentMethod(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	pm, err := service.FindOnePaymentMethod(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentMethodDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreatePaymentMethod(c *fiber.Ctx) error {
	createDto := new(dto.CreatePaymentMethodDTO)
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

	var pm model.PaymentMethod
	mapper.Map(createDto, &pm)

	err = service.CreatePaymentMethod(&pm)
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
