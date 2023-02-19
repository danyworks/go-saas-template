package products

// ProductStore is an interface for a product store that can create, update, delete, and retrieve products
type ProductStore interface {
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id uint64) error
	GetProduct(id uint64) (*Product, error)
	GetAllProducts() ([]*Product, error)
}
