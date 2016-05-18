package mysql

import (
	"github.com/a-h/rtrn/dataaccess_test"
	"testing"
)

func TestThatEntitiesCanBeCreated(t *testing.T) {
	da := NewMySqlDataAccess("rtrn:rtrn@tcp(mariadb:3306)/rtrn")
	dataaccess.BaseDataAccessImplementationTest(da, t)
}
