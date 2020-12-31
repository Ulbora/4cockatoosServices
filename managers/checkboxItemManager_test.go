package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_AddCheckboxItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var ni db.CheckboxNoteItem
	ni.Checked = true
	ni.NoteID = 5
	ni.Text = "test test"

	cdb.MockAddCheckboxNoteItemSuc = true
	cdb.MockCheckboxNoteItemID = 3

	res := m.AddCheckboxItem(&ni)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_UpdateCheckboxItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var ni db.CheckboxNoteItem
	ni.ID = 2
	ni.Checked = true
	ni.NoteID = 5
	ni.Text = "test test"

	cdb.MockUpdateCheckboxNoteItemSuc = true

	res := m.UpdateCheckboxItem(&ni)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_DeleteCheckboxItem(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	cdb.MockDeleteCheckboxNoteItemSuc = true

	res := m.DeleteCheckboxItem(4)
	if !res.Success {
		t.Fail()
	}
}
