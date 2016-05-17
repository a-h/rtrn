package httpservice

import (
	"log"
	"net/http"

	"github.com/a-h/rtrn/dataaccess"
	"github.com/gorilla/mux"
	"os"
	"fmt"
)

func Serve(da dataaccess.DataAccess) {
	address := os.ExpandEnv("${LISTEN_ADDRESS}")
	port := os.ExpandEnv("${LISTEN_PORT}")
	if port == "" {
		port = "8080"
	}
	r := createRoutes(da)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), r))
}

func createRoutes(da dataaccess.DataAccess) *mux.Router {
	r := mux.NewRouter()
	cbh := NewCallbackHandler(da)
	r.Handle("/callback", cbh)
	r.Handle("/callback/{id}", cbh)
	return r
}
