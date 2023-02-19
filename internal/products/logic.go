package products

import (
	"fmt"
	"sync"
)

// Product is a struct representing a product in the store
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint64  `json:"quantity"`
}

// ProductStore is an interface for a product store that can create, update, delete, and retrieve products
type ProductStore interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
	GetProduct(id string) (*Product, error)
	GetAllProducts() ([]*Product, error)
}

// MemoryProductStore is an implementation of ProductStore that stores products in memory
type MemoryProductStore struct {
	mu       sync.Mutex
	products map[string]*Product
}

// CreateProduct adds a new product to the store
func (s *MemoryProductStore) CreateProduct(product *Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product already exists
	if _, ok := s.products[product.ID]; ok {
		return fmt.Errorf("product with ID %s already exists", product.ID)
	}

	// Add product to store
	s.products[product.ID] = product
	return nil
}

// UpdateProduct updates an existing product in the store
func (s *MemoryProductStore) UpdateProduct(product *Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if _, ok := s.products[product.ID]; !ok {
		return fmt.Errorf("product with ID %s does not exist", product.ID)
	}

	// Update product in store
	s.products[product.ID] = product
	return nil
}

// DeleteProduct removes a product from the store
func (s *MemoryProductStore) DeleteProduct(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if _, ok := s.products[id]; !ok {
		return fmt.Errorf("product with ID %s does not exist", id)
	}

	// Remove product from store
	delete(s.products, id)
	return nil
}

// GetProduct retrieves a product from the store by ID
func (s *MemoryProductStore) GetProduct(id string) (*Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if product, ok := s.products[id]; ok {
		return product, nil
	} else {
		return nil, fmt.Errorf("product with ID %s does not exist", id)
	}
}

// GetAllProducts retrieves all products from the store
func (s *MemoryProductStore) GetAllProducts() ([]*Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Convert map to slice of products
	products := make([]*Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}

	return products, nil
}

// NewMemoryProductStore creates a new MemoryProductStore with an empty map of products
func NewMemoryProductStore() *MemoryProductStore {
	return &MemoryProductStore{
		products: make(map[string]*Product),
	}
}
