package dataaccess

// CallbackData stores the larger data used in HTTP request / reponses.
type CallbackData struct {
	// HTTP headers to apply to the callback
	Headers map[string]string `json:"headers"`
	// Form post data (used when the Method is POST or PUT)
	Data []byte `json:"data"`
}
