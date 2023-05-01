package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var Users = []User{
	{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
	{ID: 2, FirstName: "Jane", LastName: "Doe", Email: "jane.doe@example.com"},
}
