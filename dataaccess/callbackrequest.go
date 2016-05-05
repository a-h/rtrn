package dataaccess

// CallbackRequest defines the message which will be sent to the remote location.
type CallbackRequest struct {
	// Specifies the location the callback should be sent to
	URL string `json:"url"`
	// One of GET, POST, PUT
	Method string `json:"method"`
	// Data to send via HTTP headers or form post.
	PostData CallbackData `json:"postdata"`
	// The number of times to rety sending the callback, divided up into phases
	// which have their own delay. This allows the construction of policy which
	// attempts rapid retries in the first minute, followed by retrying every 5
	// minutes for an hour before failing.
	RetryPhases []RetryPhase `json:"retryphases"`
	// The name of the Application which is requesting that callbacks are made.
	ApplicationID string `json:"applicationid"`
	// Defines the queue that the message belongs to, for example,
	// a unique queue might be created per customer or order.
	// Example queue names:
	//   HONDA - Ensures ordered messaging for all HONDA vehicles.
	//   HONDA_CIVIC - Ensures ordered messaging for HONDA CIVIC vehicles.
	// If not set, then a random queue is created for the Callback.
	QueueID string `json:"queueid"`
	// The email address to notify if the callback fails.
	NotificationEmailAddress string `json:"notificationemailaddress"`
}

// NewCallbackRequest creates an empty CallbackRequest struct.
func NewCallbackRequest() *CallbackRequest {
	cb := new(CallbackRequest)
	return cb
}
