package managers

import (
	"fmt"

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

func TestC2Manager_UpdateNote(t *testing.T) {
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
	n.Title = "TEST2"
	n.Type = "checkbox"

	var en db.Note
	en.ID = 6
	en.OwnerEmail = "test@test.com"
	en.Title = "TEST"
	en.Type = "checkbox"

	cdb.MockUpdateNoteSuc = true
	cdb.MockNote = &en

	res := m.UpdateNote(&n)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_GetUsersNotes(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.Note
	n.ID = 5
	n.OwnerEmail = "test@test.com"
	n.Title = "TEST2"
	n.Type = "checkbox"
	var cblst []db.CheckboxNoteItem
	cdb.MockCheckboxNoteItemList = &cblst

	var en db.Note
	en.ID = 6
	en.OwnerEmail = "test@test.com"
	en.Title = "TEST"
	en.Type = "note"

	var nilst []db.NoteItem
	cdb.MockNoteItemList = &nilst

	var nlst []db.Note
	nlst = append(nlst, n)
	nlst = append(nlst, en)

	cdb.MockNoteList = &nlst

	res := m.GetUsersNotes("test@test.com")
	fmt.Println("notes: ", *res)
	if len(*res) != 2 {
		t.Fail()
	}
}

func TestC2Manager_DeleteNote(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.Note
	n.ID = 5
	n.OwnerEmail = "test@test.com"
	n.Title = "TEST2"
	n.Type = "checkbox"

	cdb.MockNote = &n
	cdb.MockDeleteNoteSuc = true

	res := m.DeleteNote(5, "test@test.com")

	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_GetCheckboxNote(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.Note
	n.ID = 5
	n.OwnerEmail = "test@test.com"
	n.Title = "TEST2"
	n.Type = "checkbox"

	cdb.MockNote = &n

	var ni1 db.CheckboxNoteItem
	ni1.Checked = true
	ni1.ID = 2
	ni1.NoteID = 5
	ni1.Text = "some note item"

	var ni2 db.CheckboxNoteItem
	ni2.Checked = false
	ni2.ID = 2
	ni2.NoteID = 5
	ni2.Text = "some note item 2"

	var nilst []db.CheckboxNoteItem
	nilst = append(nilst, ni1)
	nilst = append(nilst, ni2)

	cdb.MockCheckboxNoteItemList = &nilst

	res := m.GetNote(5)
	fmt.Println("note: ", *res)
	ni := res.NoteCheckboxItems
	fmt.Println("note items: ", ni)
	if res.ID != 5 {
		t.Fail()
	}
}

func TestC2Manager_GetNote(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.Note
	n.ID = 5
	n.OwnerEmail = "test@test.com"
	n.Title = "TEST2"
	n.Type = "note"

	cdb.MockNote = &n

	var ni1 db.NoteItem
	ni1.ID = 2
	ni1.NoteID = 5
	ni1.Text = "some note item"

	var ni2 db.NoteItem
	ni2.ID = 3
	ni2.NoteID = 5
	ni2.Text = "some note item 2"

	var nilst []db.NoteItem
	nilst = append(nilst, ni1)
	nilst = append(nilst, ni2)

	cdb.MockNoteItemList = &nilst

	res := m.GetNote(5)
	fmt.Println("note: ", *res)
	ni := res.NoteTextItems
	fmt.Println("note items: ", ni)
	if res.ID != 5 {
		t.Fail()
	}
}
