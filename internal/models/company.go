package models

type Company struct {
	Id      string `json:"id" bun:"id,pk"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}
