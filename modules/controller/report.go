package controller

import (
	"net/http"
	"strconv"
	e "tubespbbo/err"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/service"
	"tubespbbo/response"

	"github.com/gofiber/fiber/v2"
)

type ReportParam struct {
	Month string `query:"month"`
	Year  string `query:"year"`
}

func FindReport(c *fiber.Ctx) error {
	_, authErr := ExtractTokenMetadata(c)
	if authErr != nil {
		e.HandleErr(c, authErr)
		return nil
	}
	p := new(ReportParam)
	if err := c.QueryParser(p); err != nil {
		e.HandleErr(c, err)
		return nil
	}
	month, _ := strconv.Atoi(p.Month)
	year, _ := strconv.Atoi(p.Year)
	pm, err := service.FindReport(month, year)
	if err != nil {
		e.HandleErr(c, err)
		return nil
	}
	if pm == nil {
		pm = []dto.ReportDTO{}
	}
	_ = c.JSON(response.HTTPResponse{
		Code: http.StatusOK,
		Data: pm,
	})
	return nil
}
