package httpservice

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/a-h/rtrn/dataaccess"
	"github.com/gorilla/mux"
)

// The CallbackHandler handles requests for /callback
type CallbackHandler struct {
	DataAccess dataaccess.DataAccess
}

// NewCallbackHandler creates an instance of the callback handler.
func NewCallbackHandler(da dataaccess.DataAccess) *CallbackHandler {
	return &CallbackHandler{da}
}

func (cbh CallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling callback.")
	if r.Method == http.MethodPost {
		handleCallbackPost(w, r, cbh)
	} else {
		handleCallbackGet(w, r, cbh)
	}
}

func handleCallbackGet(w http.ResponseWriter, r *http.Request, cbh CallbackHandler) {
	log.Printf("Handling callback get.")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	log.Printf("Received id value %s", id)

	if !ok || id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Printf("Retrieving callback request %s", id)

	callbackResponse, err := cbh.DataAccess.GetCallbackStatus(id)

	if callbackResponse == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal("Failed to retrieve the callback request.", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(callbackResponse); err != nil {
		log.Fatal("Failed to marshal the callback response.", err)
	}
}

func handleCallbackPost(w http.ResponseWriter, r *http.Request, cbh CallbackHandler) {
	log.Print("Handling callback post.")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal("Failed to read the HTTP post body.", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = r.Body.Close(); err != nil {
		log.Fatal("Failed to close the reader.", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Retrieve the callback request from the body.
	cbr := dataaccess.NewCallbackRequest()
	if err = json.Unmarshal(body, &cbr); err != nil {
		log.Fatal("Failed to Unmarshal the callback request.", err)
		w.WriteHeader(422) // unprocessable entity
		return
	}

	callbackResponse, err := cbh.DataAccess.StoreCallbackRequest(cbr)

	if err != nil {
		log.Fatal("Failed to store the callback request.", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(callbackResponse); err != nil {
		log.Fatal("Failed to marshal the callback response.", err)
	}
}
