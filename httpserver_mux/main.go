package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	urlParams := mux.Vars(r)
	fmt.Println(urlParams)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get")
	case http.MethodPost:
		fmt.Println("Post")
	}

}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware", r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("Zsa")
	mux := mux.NewRouter()
	mux.HandleFunc("/{id:[0-9]+}", Handler).Methods("GET", "POST")

	mux.Use(Middleware)

	http.ListenAndServe(":4000", mux)
}
