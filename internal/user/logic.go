package user

import (
	"marketplace/internal/common"
	"marketplace/pkg/utils"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

type User struct {
	Base         common.BaseObject
	ID           uint64 `json:"id,omitempty"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash,omitempty"`
	
	Addresses      []Address   `json:"addresses,omitempty"`
	PaymentOptions PaymentInfo `json:"payment_method,omitempty"`
}

type Address struct {
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

func (user *User) CheckPassword() bool {
	return utils.CompareSHA256Hashes(utils.GenerateSHA256Hash(user.Password), user.PasswordHash)
}

func NewUser(user *User) *User {
	user.Base = *common.NewBaseObject()
	user.ID = utils.GenerateUint64ID()
	user.PasswordHash = utils.GenerateSHA256Hash(user.Password)
	return user
}
