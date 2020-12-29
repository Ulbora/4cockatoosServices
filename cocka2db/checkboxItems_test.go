package cocka2db

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var noteID2 int64
var noteItemID int64

func TestC2DB_AddCheckboxItem(t *testing.T) {
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
	// u.Email = "tester3@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester3@test.com"

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID2 = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	// dbi.Close()
	// // mydb.MockInsertSuccess2 = true
	// var nu NoteUsers
	// nu.UserEmail = "tester3@test.com"
	// nu.NoteID = id
	// nusuc := ndb.AddNoteUser(&nu)

	dbi.Close()
	var ni CheckboxNoteItem
	ni.NoteID = 5 // id
	ni.Text = "beans"
	ni.Checked = true

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	nisuc, nid := ndb.AddCheckboxItem(&ni)
	noteItemID = nid

	fmt.Println("add CheckboxItem suc: ", nisuc)

	if !nisuc || nid == 0 {
		t.Fail()
	}
}

func TestC2DB_UpdateCheckboxItem(t *testing.T) {
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
	// u.Email = "tester3@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester3@test.com"

	// // mydb.MockInsertSuccess1 = true
	// // mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID2 = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	// dbi.Close()
	// // mydb.MockInsertSuccess2 = true
	// var nu NoteUsers
	// nu.UserEmail = "tester3@test.com"
	// nu.NoteID = id
	// nusuc := ndb.AddNoteUser(&nu)

	dbi.Close()
	var ni CheckboxNoteItem
	ni.ID = noteItemID
	ni.NoteID = noteID2
	ni.Text = "bread"
	ni.Checked = false

	mydb.MockUpdateSuccess1 = true
	nisuc := ndb.UpdateCheckboxItem(&ni)

	if !nisuc {
		t.Fail()
	}
}

func TestC2DB_GetCheckboxItemList(t *testing.T) {
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
	var r1 = []string{"4", "bread", "true", "5"}
	var r2 = []string{"4", "test", "false", "5"}
	var val [][]string
	val = append(val, r1)
	val = append(val, r2)
	rtnRows.Rows = val
	mydb.MockRows1 = &rtnRows

	var dbi ddb.Database = &mydb

	var cdb C2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l
	cdb.DB = dbi
	dbi.Connect()

	ndb := cdb.GetNew()

	// var u User
	// u.Email = "tester3@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester3@test.com"

	// // mydb.MockInsertSuccess1 = true
	// // mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID2 = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	// dbi.Close()
	// // mydb.MockInsertSuccess2 = true
	// var nu NoteUsers
	// nu.UserEmail = "tester3@test.com"
	// nu.NoteID = id
	// nusuc := ndb.AddNoteUser(&nu)

	dbi.Close()
	// var ni CheckboxNoteItem
	// ni.ID = noteItemID
	// ni.NoteID = noteID2
	// ni.Text = "bread"
	// ni.Checked = false

	nilist := ndb.GetCheckboxItemList(noteID2)
	fmt.Println("add CheckboxItem list: ", nilist)

	if (*nilist)[0].Text != "bread" {
		t.Fail()
	}
}

func TestC2DB_DeleteCheckboxItem(t *testing.T) {
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
	// u.Email = "tester3@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester3@test.com"

	// // mydb.MockInsertSuccess1 = true
	// // mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID2 = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	// dbi.Close()
	// // mydb.MockInsertSuccess2 = true
	// var nu NoteUsers
	// nu.UserEmail = "tester3@test.com"
	// nu.NoteID = id
	// nusuc := ndb.AddNoteUser(&nu)

	dbi.Close()
	// var ni CheckboxNoteItem
	// ni.ID = noteItemID
	// ni.NoteID = noteID2
	// ni.Text = "bread"
	// ni.Checked = false

	mydb.MockDeleteSuccess1 = true
	dsuc := ndb.DeleteCheckboxItem(6)

	if !dsuc {
		t.Fail()
	}
}
