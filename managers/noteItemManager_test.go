package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_AddNoteItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var ni db.NoteItem

	ni.NoteID = 5
	ni.Text = "test test"

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 3

	res := m.AddNoteItem(&ni)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_UpdateNoteItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var ni db.NoteItem
	ni.ID = 2

	ni.NoteID = 5
	ni.Text = "test test"

	cdb.MockUpdateNoteItemSuc = true

	res := m.UpdateNoteItem(&ni)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_DeleteNoteItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	cdb.MockDeleteNoteItemSuc = true

	res := m.DeleteNoteItem(4)
	if !res.Success {
		t.Fail()
	}
}
