package models

type ServerResponse struct {
	Items []Worker `json:"items"`
}

type Worker struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Firstname  string `json:"firstname"`
	Patronomic string `json:"patronomic"`
}
