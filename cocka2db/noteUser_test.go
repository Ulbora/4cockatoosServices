package cocka2db

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

func TestC2DB_AddNoteUser(t *testing.T) {
	var mydb mdb.MyDBMock
	//var mydb mdb.MyDB
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

	// var u User
	// u.Email = "tester2@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester2@test.com"

	//mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	dbi.Close()
	mydb.MockInsertSuccess1 = true
	var nu NoteUsers
	nu.UserEmail = "tester2@test.com"
	nu.NoteID = 5
	nusuc := ndb.AddNoteUser(&nu)

	if !nusuc {
		t.Fail()
	}
}

func TestC2DB_DeleteNoteUser(t *testing.T) {
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

	mydb.MockDeleteSuccess1 = true
	dbi.Close()

	var nu NoteUsers
	nu.UserEmail = "test"
	nu.NoteID = 5
	suc := ndb.DeleteNoteUser(&nu)
	fmt.Println("delete note user suc: ", suc)

	// dsuc := ndb.DeleteUser("tester2@test.com")
	// fmt.Println("delete user suc: ", dsuc)
	// // fmt.Println("add note id: ", id)

	if !suc {
		t.Fail()
	}
}

func TestC2DB_GetNoteUserList(t *testing.T) {
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
	var r1 = []string{"tester2@test.com"}
	var r2 = []string{"tester3@test.com"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
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
	users := ndb.GetNoteUserList(4, "tester2@test.com")
	fmt.Println("get users : ", *users)
	// fmt.Println("add note id: ", id)

	if (*users)[0] != "tester2@test.com" {
		t.Fail()
	}
}
