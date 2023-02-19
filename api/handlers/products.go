package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"marketplace/internal/products"
	"net/http"
	"strconv"
)

var (
	productStore = products.NewMemoryProductStore()
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product products.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newp,_:=productStore.CreateProduct(&product)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newp)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var product *products.Product
	product, err = productStore.GetProduct(id)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var updatedProduct *products.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedProduct.ID = id

	go func() {
		productStore.UpdateProduct(updatedProduct)
	}()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("product with %v id was updated.", updatedProduct.ID))
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = productStore.DeleteProduct(id)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("product with %v id was deleted.", id))
}

func GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := productStore.GetAllProducts()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	json.NewEncoder(w).Encode(products)
}
