package dataaccess

import "time"

// CallbackAttempt provides the result of a Callback.
type CallbackAttempt struct {
	// 0 if the endpoint was unavailable or did not provide a response within
	// the server timeout.
	StatusCode int `json:"statuscode"`
	// The data received in the HTTP reponse.
	Response CallbackData `json:"response"`
	// The time that the callback was made.
	Time time.Time `json:"time"`
	// The duration that the server took to respond to the callback request.
	Duration time.Duration `json:"duration"`
}
