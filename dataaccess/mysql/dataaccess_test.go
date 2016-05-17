package mysql

import (
	"testing"
	"github.com/a-h/rtrn/dataaccess_test"
)

func TestThatEntitiesCanBeCreated(t *testing.T) {
	da := NewMySqlDataAccess("mariadb")
	dataaccess_test.BaseDataAccessImplementationTest(da, t)
}
