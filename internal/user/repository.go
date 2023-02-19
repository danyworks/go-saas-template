package user

type UserRepository interface {
	Create(user User) (User, error)
	GetByID(id uint64) (User, error)
	Update(user User) (User, error)
	Delete(id uint64) error
	List() ([]User, error)
}
