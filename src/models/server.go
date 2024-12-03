package models

// type ServerResponse struct {
// 	Items []Server `json:"items"`
// }

type Server struct {
	ID             string `json:"id"`
	Country        string `json:"country"`
	CountryIconUrl string `json:"countryIconUrl" db:"country_icon_url"`
	Signal         int32  `json:"signal"`
	IpAddress      string `json:"ipAddress" db:"ip_address"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Secret         string `json:"secret"`
	Port           int32  `json:"port"`
	Connections    int32  `json:"connections"`
}
