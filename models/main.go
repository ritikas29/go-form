package models


type User struct {
	Username string  `json:"username"`
	Password string `json:"-"`
	Address string
	Phone string `json:"phone"`
	Email string  `json:"email"`
	IsAdmin bool
	
}