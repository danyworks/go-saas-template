package user

import (
	"marketplace/internal/logic"
)

type User struct {
	Base      logic.BaseObject
	ID        uint64 `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`

	Addresses      []Address   `json:"addresses,omitempty"`
	PasswordHash   string      `json:"password_hash,omitempty"`
	PaymentOptions PaymentInfo `json:"payment_method,omitempty"`
}

type Address struct {
	Base    logic.BaseObject
	ID      uint64 `json:"id"`
	Line1   string `json:"line1"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type PaymentInfo struct {
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv,omitempty"`
}
