package mongo

import (
	"testing"
	"github.com/a-h/rtrn/dataaccess_test"
)

func TestThatMongoDBEntitiesCanBeCreated(t *testing.T) {
	da := NewMongoDataAccess("mongodb://localhost:27017")
	dataaccess_test.BaseDataAccessImplementationTest(da, t)
}
