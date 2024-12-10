package models

type ServerResponse struct {
	Items []Worker `json:"items"`
}

type Worker struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Firstname  string `json:"firstName"`
	Patronomic string `json:"patronomic"`
}
