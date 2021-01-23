package cocka2db

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
	"testing"
)

func TestC2DB_GetMailServerInfo(t *testing.T) {
	var mydb mdb.MyDBMock
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	// rtnRow.Row = []string{"1", "test2"}
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var rtnRows ddb.DbRows
	var r1 = []string{"4", "somehost", "0", "25", "0", "tester2@test.com", "test", "tester2@test.com"}

	var val [][]string
	val = append(val, r1)

	rtnRows.Rows = val
	mydb.MockRows1 = &rtnRows

	// var rtnRow1 ddb.DbRow
	// rtnRow1.Row = []string{"4", "test", "checkbox", "tester2@test.com"}
	// mydb.MockRow1 = &rtnRow1

	var dbi ddb.Database = &mydb

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	// var u User
	// u.Email = "tester2@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note2"
	// n.Type = "note"
	// //n.OwnerEmail = "tester2@test.com"
	// n.ID = noteId

	//mydb.MockInsertSuccess1 = true
	dbi.Close()
	msrv := ndb.GetMailServerInfo()
	fmt.Println("get mail server : ", *msrv)
	// fmt.Println("add note id: ", id)

	if msrv.SenderEmail != "tester2@test.com" {
		t.Fail()
	}
}
