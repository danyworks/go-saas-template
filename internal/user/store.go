package user

import (
	"fmt"
	"log"
	"marketplace/internal/common"
	"marketplace/pkg/cron"
	"marketplace/pkg/utils"
	"sync"
)

// MemoryUserStore is an implementation of UserStore that stores Users in memory
type MemoryUserStore struct {
	mu    sync.Mutex
	Users map[uint64]*User
}

// NewMemoryUserStore creates a new MemoryUserStore with an empty map of Users
func NewMemoryUserStore() *MemoryUserStore {
	UserStore := &MemoryUserStore{
		Users: make(map[uint64]*User),
	}
	go cron.CronJob("0 * * * *", func() {
		UploadUsers(UserStore)
	})
	return UserStore
}

// UploadUsers to the cloud
func UploadUsers(ps *MemoryUserStore) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	Users, err := ps.GetAllUsers()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Uploading Users...\n len: %v , \n items: %v", len(Users), Users)
}

// CreateUser adds a new User to the store
func (s *MemoryUserStore) CreateUser(user *User) (*User,error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if User already exists
	if _, ok := s.Users[user.ID]; ok {
		return nil,fmt.Errorf("User with ID %d already exists", user.ID)
	}
	user.Base = *common.NewBaseObject()
	user.PasswordHash = utils.GenerateSHA256Hash(user.Password)
	user.ID = utils.GenerateUint64ID()

	// Add User to store
	s.Users[user.ID] = user
	return s.Users[user.ID],nil
}

// UpdateUser updates an existing User in the store
func (s *MemoryUserStore) UpdateUser(user *User) (*User,error)  {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if User exists
	if _, ok := s.Users[user.ID]; !ok {
		return nil,fmt.Errorf("User with ID %v does not exist", user.ID)
	}
	user.Base.BaseObjectUpdated()
	// Update User in store
	s.Users[user.ID] = user
	return s.Users[user.ID],nil
}

// DeleteUser removes a User from the store
func (s *MemoryUserStore) DeleteUser(id uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if User exists
	if _, ok := s.Users[id]; !ok {
		return fmt.Errorf("User with ID %v does not exist", id)
	}

	s.Users[id].Base.BaseObjectDeleted()
	// Remove User from store
	delete(s.Users, id)
	return nil
}

// GetUser retrieves a User from the store by ID
func (s *MemoryUserStore) GetUser(id uint64) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if User exists
	if User, ok := s.Users[id]; ok {
		return User, nil
	} else {
		return nil, fmt.Errorf("User with ID %v does not exist", id)
	}
}

// GetAllUsers retrieves all Users from the store
func (s *MemoryUserStore) GetAllUsers() ([]*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Convert map to slice of Users
	Users := make([]*User, 0, len(s.Users))
	for _, User := range s.Users {
		Users = append(Users, User)
	}

	return Users, nil
}
