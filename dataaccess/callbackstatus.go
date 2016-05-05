package dataaccess

// CallbackStatus provides the status of a callback, including all
// attempts.
type CallbackStatus struct {
	ID string `bson:"_id" json:"id"`
	// Information about the initial request.
	Request CallbackRequest `json:"request"`
	// A list of all attempts which have been made.
	Attempts []CallbackAttempt `json:"attempts"`
	// All of the locks associated with the callback.
	Locks []CallbackLock `json:"locks"`
	// Whether any of the attempts are classified as successful.
	Succcess bool `json:"success"`
}
