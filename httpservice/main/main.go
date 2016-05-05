package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/a-h/rtrn/dataaccess"
	"github.com/gorilla/mux"
)

var connectionString = flag.String("connectionString", "mongodb://localhost:27017",
	"The MongoDB connection string used to store data.")

func main() {
	flag.Parse()

	r := createRoutes(dataaccess.NewMongoDataAccess(*connectionString))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createRoutes(da dataaccess.DataAccess) *mux.Router {
	r := mux.NewRouter()
	cbh := NewCallbackHandler(da)
	r.Handle("/callback", cbh)
	r.Handle("/callback/{id}", cbh)
	return r
}
