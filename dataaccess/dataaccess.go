package dataaccess

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// The DataAccess interface defines how data is written to the data store.
type DataAccess interface {
	StoreCallbackRequest(request *CallbackRequest) (*CallbackStatus, error)
	GetCallbackStatus(id string) (*CallbackStatus, error)
}

// MongoDataAccess provides access to the data structures.
type MongoDataAccess struct {
	connectionString string
}

// NewMongoDataAccess creates an instance of the MongoDataAccess type.
func NewMongoDataAccess(connectionString string) DataAccess {
	return &MongoDataAccess{connectionString}
}

// GetCallbackStatus returns a Callback status by ID.
func (da MongoDataAccess) GetCallbackStatus(id string) (*CallbackStatus, error) {
	session, err := mgo.Dial(da.connectionString)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB.", err)
		return nil, err
	}
	defer session.Close()

	c := session.DB("rtrn").C("callbacks")

	result := NewCallbackStatus()
	err = c.FindId(id).One(result)

	if err == mgo.ErrNotFound {
		log.Fatalf("Failed to find a callback with id %s.", id)
		return nil, nil
	}

	return result, nil
}

// StoreCallbackRequest stores a callback request and returns the newly created
// request's status.
func (da MongoDataAccess) StoreCallbackRequest(request *CallbackRequest) (*CallbackStatus, error) {
	session, err := mgo.Dial(da.connectionString)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB.", err)
		return nil, err
	}
	defer session.Close()

	c := session.DB("rtrn").C("callbacks")

	result := NewCallbackStatus()
	result.ID = bson.NewObjectId().Hex()
	result.Request = *request

	err = c.Insert(result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}
