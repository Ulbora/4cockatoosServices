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
