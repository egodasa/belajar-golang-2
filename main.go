package main

import (
  "fmt"
	"contoh_web/api"
	"github.com/gorilla/mux"
  "log"
  "net/http"
)

func main() {
  router := mux.NewRouter()

	router.HandleFunc("/api/ambulan", api.GetAmbulan).Methods("GET")
	router.HandleFunc("/api/ambulan/{id_ambulan}", api.GetAmbulanById).Methods("GET")
	router.HandleFunc("/api/ambulan", api.InsertAmbulan).Methods("POST")
	router.HandleFunc("/api/ambulan/{id_ambulan}", api.UpdateAmbulan).Methods("PUT")
	router.HandleFunc("/api/ambulan/{id_ambulan}", api.DeleteAmbulan).Methods("DELETE")
  router.Use(loggingMiddleware)
  
  
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func loggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      // Do stuff here
      log.Println(r.RequestURI)
      // Call the next handler, which can be another middleware in the chain, or the final handler.
      next.ServeHTTP(w, r)
  })
}
