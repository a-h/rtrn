package mysql

import (
	"database/sql"
	"fmt"
	"github.com/a-h/rtrn/dataaccess"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySqlDataAccess struct {
	ConnectionString string
}

// Create a new MySql Data Access
func NewMySqlDataAccess(connectionString string) *MySqlDataAccess {
	return &MySqlDataAccess{ConnectionString: connectionString}
}

func (da MySqlDataAccess) StoreCallbackRequest(request *dataaccess.CallbackRequest) (*dataaccess.CallbackStatus, error) {
	db, err := sql.Open("mysql", da.ConnectionString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO callbacks (url, method) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer insert.Close()

	insertRes, err := insert.Exec(request.URL, request.Method)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	lastId, err := insertRes.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	result := dataaccess.NewCallbackStatus()
	result.ID = fmt.Sprintf("%d", lastId)
	return result, nil
}

func (da MySqlDataAccess) GetCallbackStatus(id string) (*dataaccess.CallbackStatus, error) {
	result := dataaccess.NewCallbackStatus()
	result.ID = id
	return result, nil
}
