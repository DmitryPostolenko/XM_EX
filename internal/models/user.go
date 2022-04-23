package models

type User struct {
	Id       string `json:"id" bun:"id,pk"`
	Username string `json:"username"`
	Password string `json:"password"`
}
