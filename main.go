package main

import (
	"log"
	"net/http"

	"github.com/dmps/PerkBoxTest/handlers"
	"github.com/dmps/PerkBoxTest/storage"
)

func main() {
	// this creates the backend storage system
	db := storage.NewInMemoryDB()
	// this creates a new http.ServeMux, which is used to register handlers to execute in response to routes
	mux := http.NewServeMux()
	// get the value of a key
	mux.Handle("/get", handlers.GetKey(db))
	// set the value of a key
	mux.Handle("/update", handlers.UpdateKey(db))
	mux.Handle("/list", handlers.List(db))

	log.Printf("serving on port 8080")

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since ServeMux implements a ServeHTTP function, it is also an http.Handler,
	// so we can pass it here.
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
