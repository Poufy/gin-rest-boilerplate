package models

type LoginCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"` // The password field should not be returned in the response
}
