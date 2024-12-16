package models

type ResultServerResponse struct {
	Cell        []int             `json:"cell"`
	Row         []Row             `json:"row"`
	TypesOfWork []TypeOfWorkModel `json:"typesOfWork"`
}

type TypeOfWorkModel struct {
	TypeOfWorkId int    `json:"typeOfWorkId"`
	NameOfWOrk   string `json:"nameOfWork"`
	UOM          string `json:"uom"`
	PriceOfWork  int    `json:"priceOfWork"`
	PeriodOfWork string `json:"periodOfWork"`
}

type Row struct {
	RowId  int `json:"rowId"`
	CellId int `json:"cellId"`
}
