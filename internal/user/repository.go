package user

// UserStore is an interface for a User store that can create, update, delete, and retrieve Users
type UserStore interface {
	CreateUser(User *User) error
	UpdateUser(User *User) error
	DeleteUser(id uint64) error
	GetUser(id uint64) (*User, error)
	GetAllUsers() ([]*User, error)
}
