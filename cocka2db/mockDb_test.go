package cocka2db

import (
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestMockC2DB_AddUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockAddUserSuc = true

	var u User
	suc := cdb.AddUser(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_UpdateUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockUpdateUserSuc = true

	var u User
	suc := cdb.UpdateUser(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_GetUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var u User
	cdb.MockUser = &u
	u.WebEnabled = true

	fu := cdb.GetUser("test")
	if !fu.WebEnabled {
		t.Fail()
	}
}

func TestMockC2DB_DeleteUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockDeleteUserSuc = true

	suc := cdb.DeleteUser("test")
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_AddNote(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockAddNoteSuc = true
	cdb.MockAddNoteID = 4

	var u Note
	suc, id := cdb.AddNote(&u)
	if !suc && id != 0 {
		t.Fail()
	}
}

func TestMockC2DB_UpdateNote(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockUpdateNoteSuc = true

	var u Note
	suc := cdb.UpdateNote(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_GetNote(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var u Note
	u.Title = "testttt"
	cdb.MockNote = &u

	fu := cdb.GetNote(8)
	if fu.Title != "testttt" {
		t.Fail()
	}
}

func TestMockC2DB_GetNoteList(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var u Note
	u.Title = "testttt"
	var nlst []Note
	cdb.MockNote = &u
	nlst = append(nlst, u)
	cdb.MockNoteList = &nlst

	fu := cdb.GetUsersNotes("test@test.com")
	if (*fu)[0].Title != "testttt" {
		t.Fail()
	}
}

func TestMockC2DB_DeleteNote(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockDeleteNoteSuc = true

	suc := cdb.DeleteNote(4)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_AddNoteItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockAddNoteItemSuc = true
	cdb.MockNoteItemID = 4

	var u NoteItem
	suc, id := cdb.AddNoteItem(&u)
	if !suc && id != 0 {
		t.Fail()
	}
}

func TestMockC2DB_UpdateNoteItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockUpdateNoteItemSuc = true

	var u NoteItem
	suc := cdb.UpdateNoteItem(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_GetNoteItemList(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var u NoteItem
	u.Text = "testttt"

	var nilst []NoteItem
	nilst = append(nilst, u)
	cdb.MockNoteItemList = &nilst

	fu := cdb.GetNoteItemList(8)
	if (*fu)[0].Text != "testttt" {
		t.Fail()
	}
}

func TestMockC2DB_DeleteNoteItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockDeleteNoteItemSuc = true

	suc := cdb.DeleteNoteItem(4)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_AddCheckboxItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockAddCheckboxNoteItemSuc = true
	cdb.MockCheckboxNoteItemID = 4

	var u CheckboxNoteItem
	suc, id := cdb.AddCheckboxItem(&u)
	if !suc && id != 0 {
		t.Fail()
	}
}

func TestMockC2DB_UpdateCheckboxItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockUpdateCheckboxNoteItemSuc = true

	var u CheckboxNoteItem
	suc := cdb.UpdateCheckboxItem(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_GetCheckboxItemList(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var u CheckboxNoteItem
	u.Text = "testttt"

	var nilst []CheckboxNoteItem
	nilst = append(nilst, u)
	cdb.MockCheckboxNoteItemList = &nilst

	fu := cdb.GetCheckboxItemList(8)
	if (*fu)[0].Text != "testttt" {
		t.Fail()
	}
}

func TestMockC2DB_DeleteCheckboxItem(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockDeleteCheckboxNoteItemSuc = true

	suc := cdb.DeleteCheckboxItem(4)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_AddNoteUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockAddNoteUserSuc = true

	var u NoteUsers
	suc := cdb.AddNoteUser(&u)
	if !suc {
		t.Fail()
	}
}

func TestMockC2DB_DeleteNoteUser(t *testing.T) {
	var cdb MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	cdb.MockDeleteNoteUserSuc = true

	i := cdb.GetNew()

	var u NoteUsers
	suc := i.DeleteNoteUser(&u)
	if !suc {
		t.Fail()
	}
}
