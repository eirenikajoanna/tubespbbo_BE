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

func FindPayment(c *fiber.Ctx) error {
	AuthCheck(c)
	pm, err := service.FindPayment()
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTOs []dto.PaymentDTO
	mapper.Map(pm, &DTOs)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTOs,
	})
	return nil
}

func FindOnePayment(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	pm, err := service.FindOnePayment(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func CreatePayment(c *fiber.Ctx) error {
	AuthCheck(c)
	createDto := new(dto.CreatePaymentDTO)
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

	_, err = service.FindOnePaymentMethod(createDto.PaymentMethodId)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var payment model.Payment
	mapper.Map(createDto, &payment)
	payment.Status = "Belum Terverifikasi"

	err = service.CreatePayment(&payment)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: payment,
	})
	return nil
}

func UpdatePayment(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	updateDto := new(dto.UpdatePaymentDTO)
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

	pm, err := service.UpdatePayment(updateDto, id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}

func DeletePayment(c *fiber.Ctx) error {
	AuthCheck(c)
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	pm, err := service.DeletePayment(id)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}

	var DTO dto.PaymentDTO
	mapper.Map(pm, &DTO)

	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: DTO,
	})
	return nil
}
