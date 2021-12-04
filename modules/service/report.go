package service

import (
	"sort"
	"time"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindReport(month int, year int) ([]dto.ReportDTO, error) {
	pm, err := repository.FindTransactionOneMonth(month, year)
	if err != nil {
		return nil, err
	}

	var DTOs []dto.ReportDTO
	for i := 0; i < len(*pm); i++ {
		pb := dto.ReportDTO{
			Id:        (*pm)[i].Id,
			Date:      (*pm)[i].CreatedAt.Format("2006-01-02"),
			Type:      "Debit",
			Amount:    CalculateTotalTransaction(&(*pm)[i]),
			CreatedAt: (*pm)[i].CreatedAt,
		}
		DTOs = append(DTOs, pb)
	}

	ex, err := repository.FindExpenseOneMonth(month, year)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(*ex); i++ {
		pb := dto.ReportDTO{
			Id:        (*ex)[i].Id,
			Date:      (*ex)[i].ReleaseDate,
			Type:      "Credit",
			Amount:    (*ex)[i].TotalAmount,
			CreatedAt: (*ex)[i].CreatedAt,
		}
		DTOs = append(DTOs, pb)
	}
	sort.Slice(DTOs, func(i, j int) bool {
		layoutFormat := "2006-01-02"
		dateI, _ := time.Parse(layoutFormat, DTOs[i].Date)
		dateJ, _ := time.Parse(layoutFormat, DTOs[j].Date)
		return dateI.Before(dateJ)
	})
	return DTOs, nil
}

func CalculateTotalTransaction(transaction *model.Transaction) float32 {
	total := 0
	for i := 0; i < len(transaction.TransactionDetails); i++ {
		total += int(transaction.TransactionDetails[i].Quantity) * int(transaction.TransactionDetails[i].Product.Price)
	}
	return float32(total)
}
