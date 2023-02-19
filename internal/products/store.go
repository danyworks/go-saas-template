package products

import (
	"fmt"
	"log"
	"marketplace/internal/common"
	"marketplace/pkg/cron"
	"marketplace/pkg/utils"
	"sync"
)

// MemoryProductStore is an implementation of ProductStore that stores products in memory
type MemoryProductStore struct {
	mu       sync.Mutex
	products map[uint64]*Product
}

// NewMemoryProductStore creates a new MemoryProductStore with an empty map of products
func NewMemoryProductStore() *MemoryProductStore {
	productStore := &MemoryProductStore{
		products: make(map[uint64]*Product),
	}
	go cron.CronJob("0 * * * *", func() {
		UploadProducts(productStore)
	})
	return productStore
}

// Upload Products to the cloud
func UploadProducts(ps *MemoryProductStore) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	products, err := ps.GetAllProducts()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("Uploading products...\n len: %v , \n items: %v", len(products), products)
}

// CreateProduct adds a new product to the store
func (s *MemoryProductStore) CreateProduct(product *Product) (*Product,error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product.Base = *common.NewBaseObject()
	product.ID = utils.GenerateUint64ID()
	// Add product to store
	s.products[product.ID] = product
	return s.products[product.ID],nil
}

// UpdateProduct updates an existing product in the store
func (s *MemoryProductStore) UpdateProduct(product *Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if _, ok := s.products[product.ID]; !ok {
		return fmt.Errorf("product with ID %v does not exist", product.ID)
	}
	product.Base.BaseObjectUpdated()
	// Update product in store
	s.products[product.ID] = product
	return nil
}

// DeleteProduct removes a product from the store
func (s *MemoryProductStore) DeleteProduct(id uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if _, ok := s.products[id]; !ok {
		return fmt.Errorf("product with ID %v does not exist", id)
	}
	s.products[id].Base.BaseObjectDeleted()
	// Remove product from store
	delete(s.products, id)
	return nil
}

// GetProduct retrieves a product from the store by ID
func (s *MemoryProductStore) GetProduct(id uint64) (*Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if product exists
	if product, ok := s.products[id]; ok {
		return product, nil
	} else {
		return nil, fmt.Errorf("product with ID %v does not exist", id)
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
