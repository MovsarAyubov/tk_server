package models

import (
	"github.com/shopspring/decimal"
)

type DoneWorkModel struct {
	WorkerId     int             `json:"workerId"`
	TypeOfWorkId int             `json:"typeOfWorkId"`
	Date         string          `json:"date"`
	CellId       int             `json:"cellId"`
	RowId        int             `json:"rowId"`
	Count        decimal.Decimal `json:"count"`
	Income       decimal.Decimal `json:"income"`
}
