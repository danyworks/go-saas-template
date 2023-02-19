package products

import (
	"marketplace/internal/common"
	"marketplace/pkg/utils"
)

// Product is a struct representing a product in the store
type Product struct {
	Base        common.BaseObject
	ID          uint64   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Images      []string `json:"image_urls"`
	Stock       uint64   `json:"quantity"`
}

func NewProduct(product *Product) *Product {
	product.Base = *common.NewBaseObject()
	product.ID = utils.GenerateUint64ID()
	return product
}
