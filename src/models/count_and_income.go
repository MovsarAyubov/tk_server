package models

import "github.com/shopspring/decimal"

type CountAndIncomeResponse struct {
	Items []CountAndIncome `json:"items"`
}

type CountAndIncome struct {
	Count  decimal.Decimal `json:"count"`
	Income decimal.Decimal `json:"income"`
}
