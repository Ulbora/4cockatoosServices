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
