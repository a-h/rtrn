package dataaccess

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCallbackStatusSerialization(t *testing.T) {
	request := createExample()
	callbackData := CallbackData{}
	callbackDate := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	timeTaken := time.Millisecond * 400
	cbs := CallbackStatus{"id_value",
		request,
		[]CallbackAttempt{CallbackAttempt{404, callbackData, callbackDate, timeTaken}},
		[]CallbackLock{CallbackLock{}},
		true}

	actual, err := json.MarshalIndent(cbs, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	expected := `{
 "id": "id_value",
 "request": {
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
 },
 "attempts": [
  {
   "statuscode": 404,
   "response": {
    "headers": null,
    "data": null
   },
   "time": "2016-01-01T00:00:00Z",
   "duration": 400000000
  }
 ],
 "locks": [
  {
   "acquired": "0001-01-01T00:00:00Z",
   "timeout": "0001-01-01T00:00:00Z",
   "completed": false,
   "machinename": ""
  }
 ],
 "success": true
}`

	if expected != string(actual) {
		t.Error("The expected didn't match the actual.", expected, string(actual))
	}
}

func createExample() CallbackRequest {
	var cb = NewCallbackRequest()
	cb.ApplicationID = "test"
	cb.Method = "GET"
	cb.NotificationEmailAddress = "a-h@github.com"
	headers := make(map[string]string)
	headers["x-header-one"] = "x-header-one-value"
	cb.PostData = CallbackData{headers, nil}
	cb.QueueID = "queueid value"
	retryPhases := []RetryPhase{RetryPhase{1, time.Minute * 5}}
	cb.RetryPhases = retryPhases
	cb.URL = "http://example.com/callback"
	return *cb
}

func TestCallbackRequestSerialization(t *testing.T) {
	cb := createExample()

	actual, err := json.MarshalIndent(cb, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	expected := `{
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

	if expected != string(actual) {
		t.Error("The expected didn't match the actual.", expected, string(actual))
	}
}
