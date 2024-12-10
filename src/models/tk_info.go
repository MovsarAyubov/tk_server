package models

type TKInfoServerResponse struct {
	CellId      []int             `json:"cellId"`
	RowId       []int             `json:"rowId"`
	TypesOfWork []TypeOfWorkModel `json:"typesOfWork"`
}

type TypeOfWorkModel struct {
	TypeOfWorkId int    `json:"typeOfWorkId"`
	NameOfWOrk   string `json:"nameOfWork"`
	UIM          string `json:"uim"`
	PriceOfWork  int    `json:"priceOfWork"`
	PeriodOfWork string `json:"periodOfWork"`
}
