package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/a-h/rtrn/dataaccess"
)

type MockDataAccess struct {
	Statuses []dataaccess.CallbackStatus
}

func NewMockDataAccess() MockDataAccess {
	da := MockDataAccess{
		Statuses: []dataaccess.CallbackStatus{},
	}
	return da
}

func (da MockDataAccess) StoreCallbackRequest(request *dataaccess.CallbackRequest) (*dataaccess.CallbackStatus, error) {
	status := dataaccess.CallbackStatus{}
	da.Statuses = append(da.Statuses, status)
	return &status, nil
}

func (da MockDataAccess) GetCallbackStatus(id string) (*dataaccess.CallbackStatus, error) {
	for _, status := range da.Statuses {
		if status.ID == id {
			return &status, nil
		}
	}

	return nil, nil
}

func TestTheMockDataReturnsValues(t *testing.T) {
	da := NewMockDataAccess()
	status := &dataaccess.CallbackStatus{
		ID: "t",
	}
	da.Statuses = append(da.Statuses, *status)

	result, _ := da.GetCallbackStatus("t")

	if result.ID != "t" {
		t.Error("Failed to receive the expected value.")
	}
}

func TestCallbackHandlerPostData(t *testing.T) {
	callbackRequestJSON := `{
		"url": "http://example.com/callback",
		"method": "GET",
		"postdata": {
		"headers": {
		 "x-header-one": "x-header-one-value"
		},
		"data": null
		},
		"retryphases": [
		{
		 "retries": 1,
		 "delay": 300000000000
		}
		],
		"applicationid": "test",
		"queueid": "queueid value",
		"notificationemailaddress": "a-h@github.com"
	}`
	sr := strings.NewReader(callbackRequestJSON)
	r, _ := http.NewRequest(http.MethodPost, "/callback", sr)

	w := httptest.NewRecorder()

	da := NewMockDataAccess()
	createRoutes(da).ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected %d status code, received %d", http.StatusOK, w.Code)
	}
}

func TestThatCallbacksCanBeFound(t *testing.T) {
	sr := strings.NewReader("")
	r, _ := http.NewRequest("GET", "/callback/uwey12rin", sr)

	w := httptest.NewRecorder()

	da := NewMockDataAccess()
	status := &dataaccess.CallbackStatus{
		ID: "uwey12rin",
	}
	da.Statuses = append(da.Statuses, *status)
	createRoutes(da).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("The service didn't respond with HTTP 200. The response was %d", w.Code)
	}

	//TODO: Validate the JSON?
}

func TestThatA404ErrorIsGivenWhenACallbackIsNotFoundInTheDatabase(t *testing.T) {
	sr := strings.NewReader("")
	r, _ := http.NewRequest("GET", "/callback/uwey12rin", sr)

	w := httptest.NewRecorder()

	da := NewMockDataAccess()
	routes := createRoutes(da)
	routes.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Error("An incorrect error code was provided. ", w.Code)
	}
}
