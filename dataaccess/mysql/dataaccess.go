package mysql

import "github.com/a-h/rtrn/dataaccess"

type MySqlDataAccess struct {
	ConnectionString string
}

// Create a new MySql Data Access
func NewMySqlDataAccess(connectionString string) *MySqlDataAccess {
	return &MySqlDataAccess{ConnectionString: connectionString}
}

func (da MySqlDataAccess) StoreCallbackRequest(request *dataaccess.CallbackRequest) (*dataaccess.CallbackStatus, error) {
	result := dataaccess.NewCallbackStatus()
	result.ID = "123"
	return result, nil
}

func (da MySqlDataAccess) GetCallbackStatus(id string) (*dataaccess.CallbackStatus, error) {
	return nil, nil
}
