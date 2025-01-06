package models

type Periods struct {
	Period []string `json:"period"`
}

type Works struct {
	Works []TypeOfWorkModel `json:"work"`
}

type ResultServerResponse struct {
	Cell        []int             `json:"cell"`
	Row         []Row             `json:"row"`
	TypesOfWork []TypeOfWorkModel `json:"typesOfWork"`
}

type TypeOfWorkModel struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Uom    string  `json:"uom"`
	Price  float32 `json:"price"`
	Period string  `json:"period"`
}

type Row struct {
	Id     int `json:"id"`
	CellId int `json:"cellId"`
}
