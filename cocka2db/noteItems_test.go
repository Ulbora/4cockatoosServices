package cocka2db

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ddb "github.com/Ulbora/dbinterface"
	mdb "github.com/Ulbora/dbinterface_mysql"
)

var noteID4 int64
var noteItemID4 int64

func TestC2DB_AddNoteItem(t *testing.T) {
	var mydb mdb.MyDBMock
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1", "test2"}
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
	// u.Email = "tester4@test.com"
	// suc := ndb.AddUser(&u)
	// fmt.Println("add user suc: ", suc)

	// var n Note
	// n.Title = "test note"
	// n.Type = "checkbox"
	// n.OwnerEmail = "tester4@test.com"

	// mydb.MockInsertSuccess1 = true
	// mydb.MockInsertID1 = 5
	// dbi.Close()
	// suc, id := ndb.AddNote(&n)
	// noteID4 = id
	// fmt.Println("add note suc: ", suc)
	// fmt.Println("add note id: ", id)

	// dbi.Close()
	// // mydb.MockInsertSuccess2 = true
	// var nu NoteUsers
	// nu.UserEmail = "tester4@test.com"
	// nu.NoteID = id
	// nusuc := ndb.AddNoteUser(&nu)
	// fmt.Println("add nusuc: ", nusuc)

	dbi.Close()
	var ni NoteItem
	ni.NoteID = 6 // id
	ni.Text = "beans"

	mydb.MockInsertSuccess1 = true
	mydb.MockInsertID1 = 5

	nisuc, nid := ndb.AddNoteItem(&ni)
	noteItemID4 = nid

	fmt.Println("add NoteItem suc: ", nisuc)

	if !nisuc || nid == 0 {
		t.Fail()
	}
}

func TestC2DB_UpdateNoteItem(t *testing.T) {
	var mydb mdb.MyDBMock
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1", "test2"}
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
	var ni NoteItem
	ni.ID = noteItemID4
	ni.NoteID = noteID4
	ni.Text = "bread"

	mydb.MockUpdateSuccess1 = true
	nisuc := ndb.UpdateNoteItem(&ni)

	if !nisuc {
		t.Fail()
	}
}

func TestC2DB_GetNoteItemList(t *testing.T) {
	var mydb mdb.MyDBMock
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1", "test2"}
	rtnRow.Row = []string{}
	mydb.MockTestRow = &rtnRow

	var rtnRows ddb.DbRows
	var r1 = []string{"4", "bread", "5"}
	var r2 = []string{"4", "test", "5"}
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

	nilist := ndb.GetNoteItemList(noteID4)
	fmt.Println("add NoteItem list: ", nilist)

	if (*nilist)[0].Text != "bread" {
		t.Fail()
	}
}

func TestC2DB_DeleteNoteItem(t *testing.T) {
	var mydb mdb.MyDBMock
	// var mydb mdb.MyDB
	mydb.Host = "localhost:3306"
	mydb.User = "admin"
	mydb.Password = "admin"
	mydb.Database = "cocka2_notes"

	var rtnRow ddb.DbRow
	rtnRow.Row = []string{"1", "test2"}
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
	dsuc := ndb.DeleteNoteItem(noteItemID4)

	if !dsuc {
		t.Fail()
	}
}
