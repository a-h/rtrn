package dataaccess

import "time"

// CallbackLock allows the background services running on multiple servers
// to acquire a temporary lock against the callback request. The lock is held
// for a timeout period which allows the background service to attempt to
// complete the callback. If the callback server does not respond within the
// lock timeout, the attempt is logged as a failure, the lock cancelled, and
// a retry may occur.
//
// If the attempt is not completed and the lock cancelled within the lock
// timeout period, (for example, if the background service is terminated
// unexpectedly), another background service instance can attempt to execute
// the callback.
type CallbackLock struct {
	// The time that the lock was acquired.
	Acquired time.Time `json:"acquired"`
	// The time that the lock expires.
	Timeout time.Time `json:"timeout"`
	// Set to true when the lock is no longer required.
	Completed bool `json:"completed"`
	// The name of the machine which has acquired the lock.
	MachineName string `json:"machinename"`
}
