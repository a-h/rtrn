package mongo

import (
	"github.com/a-h/rtrn/dataaccess"
	"reflect"
	"testing"
	"time"
)

func TestThatMongoDBEntitiesCanBeCreated(t *testing.T) {
	da := NewMongoDataAccess("mongodb://localhost:27017")
	request := &dataaccess.CallbackRequest{
		URL:    "http://example.com/test/123",
		Method: "GET",
		PostData: dataaccess.CallbackData{
			Headers: make(map[string]string),
			Data:    []byte{},
		},
		RetryPhases: []dataaccess.RetryPhase{
			dataaccess.RetryPhase{
				Retries: 1,
				Delay:   time.Minute * 5,
			},
		},
		ApplicationID:            "UnitTest",
		QueueID:                  "TestQueue",
		NotificationEmailAddress: "a-h@github.com",
	}

	stored, _ := da.StoreCallbackRequest(request)
	retrieved, _ := da.GetCallbackStatus(stored.ID)

	if stored.ID != retrieved.ID {
		t.Errorf("The stored ID %s and retrieved Id %s didn't match.", stored.ID, retrieved.ID)
	}

	if len(stored.Attempts) != len(retrieved.Attempts) {
		t.Error("The number of attempts didn't match.")
	}

	if len(stored.Locks) != len(retrieved.Locks) {
		t.Error("The number of locks didn't match.")
	}

	if stored.Request.ApplicationID != retrieved.Request.ApplicationID {
		t.Error("The application id's didn't match.")
	}

	if stored.Succcess != retrieved.Succcess {
		t.Error("The success didn't match.")
	}

	if stored.Request.Method != retrieved.Request.Method {
		t.Error("The request methods didn't match.")
	}

	if !reflect.DeepEqual(stored.Request, retrieved.Request) {
		t.Error("The request objects didn't match.")
	}

	if !reflect.DeepEqual(stored.Locks, retrieved.Locks) {
		t.Error("The lock objects didn't match.")
	}

	if !reflect.DeepEqual(stored, retrieved) {
		t.Error("The objects are not deeply equal.")
	}
}
