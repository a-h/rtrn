package mongo

import (
	"github.com/a-h/rtrn/dataaccess_test"
	"testing"
)

func TestThatMongoDBEntitiesCanBeCreated(t *testing.T) {
	da := NewMongoDataAccess("mongodb://mongo:27017")
	dataaccess.BaseDataAccessImplementationTest(da, t)
}
