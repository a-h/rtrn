package dataaccess

import "time"

// A RetryPhase defines how the CallbackRequest will handle failure. For
// example, if the
type RetryPhase struct {
	Retries int           `json:"retries"`
	Delay   time.Duration `json:"delay"`
}
