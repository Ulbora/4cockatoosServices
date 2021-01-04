package cocka2db

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var noteID int64

func TestC2DB_AddNote(t *testing.T) {
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

	var n Note
	n.Title = "test note"
	n.Type = "checkbox"
	n.OwnerEmail = "tester2@test.com"

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5
	dbi.Close()
	suc, id := ndb.AddNote(&n)
	noteID = id
	fmt.Println("add note suc: ", suc)
	fmt.Println("add note id: ", id)

	dbi.Close()
	mydb.MockInsertSuccess2 = true
	var nu NoteUsers
	nu.UserEmail = "tester2@test.com"
	nu.NoteID = id
	nusuc := ndb.AddNoteUser(&nu)

	if !suc || id == 0 || !nusuc {
		t.Fail()
	}
}

func TestC2DB_UpdateNote(t *testing.T) {
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

	var n Note
	n.Title = "test note2"
	n.Type = "note"
	//n.OwnerEmail = "tester2@test.com"
	n.ID = noteID

	mydb.MockUpdateSuccess1 = true
	dbi.Close()
	suc := ndb.UpdateNote(&n)
	fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	if !suc {
		t.Fail()
	}
}

func TestC2DB_GetNoteRow(t *testing.T) {
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

	var rtnRow1 ddb.DbRow
	rtnRow1.Row = []string{"4", "test", "checkbox", "tester2@test.com", "2021-01-05T00:00:00Z"}
	mydb.MockRow1 = &rtnRow1

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
	fnote := ndb.GetNote(noteID)
	fmt.Println("get note : ", *fnote)
	// fmt.Println("add note id: ", id)

	if fnote.OwnerEmail != "tester2@test.com" {
		t.Fail()
	}
}

func TestC2DB_GetUsersNotes(t *testing.T) {
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
	var r1 = []string{"4", "test", "checkbox", "tester2@test.com", "2021-01-05T00:00:00Z"}
	var r2 = []string{"4", "test", "checkbox", "tester3@test.com", "2021-01-05T00:00:00Z"}
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
	notes := ndb.GetUsersNotes("tester2@test.com")
	fmt.Println("get note : ", *notes)
	// fmt.Println("add note id: ", id)

	if (*notes)[0].OwnerEmail != "tester2@test.com" {
		t.Fail()
	}
}

func TestC2DB_DeleteNote(t *testing.T) {
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
	suc := ndb.DeleteNote(noteID)
	fmt.Println("delete note suc: ", suc)

	// dsuc := ndb.DeleteUser("tester2@test.com")
	// fmt.Println("delete user suc: ", dsuc)
	// // fmt.Println("add note id: ", id)

	if !suc {
		t.Fail()
	}
}
