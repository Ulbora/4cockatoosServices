package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_AddNote(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.Note
	n.OwnerEmail = "test@test.com"
	n.Title = "TEST"
	n.Type = "checkbox"

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 4
	cdb.MockAddNoteUserSuc = true

	res := m.AddNote(&n)
	if !res.Success {
		t.Fail()
	}
}
