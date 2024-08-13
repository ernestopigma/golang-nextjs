package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}
