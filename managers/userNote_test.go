package managers

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_AddUserToNote(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	// var uo db.User
	// uo.Email = "test2@test.com"
	// uo.Password = "test"
	// uo.WebEnabled = true

	var u db.User
	u.Email = "test@test.com"
	u.Password = "test"
	u.WebEnabled = true

	var fn db.Note
	fn.ID = 4

	cdb.MockUser = &u
	cdb.MockNote = &fn

	cdb.MockAddNoteUserSuc = true

	var nu NoteUsers
	nu.OwnerEmail = "test2@test.com"
	nu.UserEmail = "test@test.com"
	nu.NoteID = 4

	res := m.AddUserToNote(&nu)
	if res.Success {
		t.Fail()
	}
}
