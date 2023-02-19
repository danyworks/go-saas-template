package api

import (
	"fmt"
	"log"
	"marketplace/api/handlers"
	"net/http"
	"time"
)

var (
	router = http.NewServeMux()
)
func StartServer(address string) {

	log.Print("Initializing Rest Endpoints...")
	
	
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	productEndpoints()
	userEndpoints()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", address),
		// Good practice to set timeouts
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("Error while running server %v", err)
	}
}

func productEndpoints() {

	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		switch r.Method {
		case http.MethodPost:
			handlers.CreateProductHandler(w, r)
		case http.MethodGet:
			handlers.GetProductHandler(w, r)
		case http.MethodPut:
			handlers.UpdateProductHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteProductHandler(w, r)
		default:
			http.Error(w, fmt.Sprintf("Invalid method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})
	router.HandleFunc("/products/all",func(w http.ResponseWriter, r *http.Request) {handlers.GetAllProductsHandler(w,r)}) 
}

func userEndpoints(){

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")
		switch r.Method {
		case http.MethodPost:
			handlers.CreateUserHandler(w, r)
		case http.MethodGet:
			handlers.GetUserHandler(w, r)
		case http.MethodPut:
			handlers.UpdateUserHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteUserHandler(w, r)
		default:
			http.Error(w, fmt.Sprintf("Invalid method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})
	router.HandleFunc("/users/all",func(w http.ResponseWriter, r *http.Request) {handlers.GetAllUsersHandler(w,r)})

}