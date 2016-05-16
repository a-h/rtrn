package dataaccess


// The DataAccess interface defines how data is written to the data store.
type DataAccess interface {
	StoreCallbackRequest(request *CallbackRequest) (*CallbackStatus, error)
	GetCallbackStatus(id string) (*CallbackStatus, error)
}
