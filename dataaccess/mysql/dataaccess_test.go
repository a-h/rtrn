package mysql

import (
	"testing"
	"github.com/a-h/rtrn/dataaccess_test"
)

func TestThatEntitiesCanBeCreated(t *testing.T) {
	da := NewMySqlDataAccess("rtrn:rtrn@tcp(mariadb:3306)/rtrn")
	dataaccess.BaseDataAccessImplementationTest(da, t)
}
