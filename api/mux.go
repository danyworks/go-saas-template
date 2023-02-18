package api

import (
	"fmt"
	"log"
	"net/http"
	"time"
)



func StartServer(address string)  {

        log.Print("Initializing Rest Endpoints...")
        router := http.NewServeMux()

		router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})


        srv := &http.Server{
                Addr: fmt.Sprintf(":%v",address),
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