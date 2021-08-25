package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	cl "apiStore/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func ConnectionServer() {
	fmt.Println("Hi")
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},

		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	//1.Query Buyers
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		allbuyers := cl.AllBuyers()
		json.NewEncoder(w).Encode(allbuyers)

	})

	//2.Query Buyer general info

	//3. Post Buyer, Treansactions, Productos
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

		allbuyers := cl.AllBuyers()
		json.NewEncoder(w).Encode(allbuyers)

		fmt.Println(cl.AllBuyers())

	})

	http.ListenAndServe(":4000", r)
}
