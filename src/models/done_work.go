package models

import (
	"github.com/shopspring/decimal"
)

type DoneWorkServerResponse struct {
	Items []DoneWorkModelResponse `json:"items"`
}

type DoneWorkModel struct {
	Id              int             `json:"id"`
	Worker_id       int             `json:"worker_id"`
	Type_of_work_id int             `json:"type_of_work_id"`
	Date            string          `json:"date"`
	Cell_id         int             `json:"cell_id"`
	Row_id          int             `json:"row_id"`
	Count           decimal.Decimal `json:"count"`
	Income          decimal.Decimal `json:"income"`
}

type DoneWorkModelResponse struct {
	Id                int             `json:"id"`
	Worker_id         int             `json:"worker_id"`
	Type_of_work_id   int             `json:"type_of_work_id"`
	Type_of_work_name string          `json:"type_of_work"`
	Date              string          `json:"date"`
	Cell_id           int             `json:"cell_id"`
	Row_id            int             `json:"row_id"`
	Count             decimal.Decimal `json:"count"`
	Income            decimal.Decimal `json:"income"`
}
