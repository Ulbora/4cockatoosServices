package cocka2db

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestC2DB_AddUser(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	// rtnRow.Row = []string{"1", "test2"}
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var dbi ddb.Database = &mydb

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	var u User
	u.Email = "tester@test.com"

	mydb.MockInsertSuccess1 = true
	dbi.Close()
	suc := ndb.AddUser(&u)
	fmt.Println("add user suc: ", suc)

	if !suc {
		t.Fail()
	}
}

func TestC2DB_AddDubUser(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1", "test2"}
	mydb.MockTestRow = &rtnRow

	var dbi ddb.Database = &mydb

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	var u User
	u.Email = "tester@test.com"

	dbi.Close()
	suc := ndb.AddUser(&u)
	fmt.Println("add dup user suc: ", suc)

	if suc {
		t.Fail()
	}
}

func TestC2DB_AddDubUser2(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1d", "test2"}
	mydb.MockTestRow = &rtnRow

	var dbi ddb.Database = &mydb

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	var u User
	u.Email = "tester@test.com"

	dbi.Close()
	suc := ndb.AddUser(&u)
	fmt.Println("add dup user suc: ", suc)

	if suc {
		t.Fail()
	}
}

func TestC2DB_UpdateUser(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"
	var dbi ddb.Database = &mydb

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	var u User
	u.Email = "tester@test.com"
	u.Password = "test"
	u.WebEnabled = true

	mydb.MockUpdateSuccess1 = true

	dbi.Close()
	suc := ndb.UpdateUser(&u)
	fmt.Println("update user suc: ", suc)

	if !suc {
		t.Fail()
	}
}

func TestC2DB_GetUser(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"
	var dbi ddb.Database = &mydb

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var rtnRow2 ddb.DbRow
	rtnRow2.Row = []string{"tester@test.com", "test", "true"}
	mydb.MockRow1 = &rtnRow2

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	var u User
	u.Email = "tester@test.com"
	u.Password = "test"
	u.WebEnabled = true

	dbi.Close()
	fu := ndb.GetUser("tester@test.com")
	fmt.Println("user found: ", fu)

	if !fu.WebEnabled || fu.Password != "test" {
		t.Fail()
	}
}

func TestC2DB_DeleteUser(t *testing.T) {
	var mydb mdb.MyDBMock
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"
	var dbi ddb.Database = &mydb

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	mydb.MockDeleteSuccess1 = true

	dbi.Close()
	suc := ndb.DeleteUser("tester@test.com")
	fmt.Println("delete user suc: ", suc)

	if !suc {
		t.Fail()
	}
}
