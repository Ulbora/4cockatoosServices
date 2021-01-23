package managers

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
)

func TestC2Manager_AddUser(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var u db.User
	u.Email = "test@test.com"
	u.Password = "test"

	cdb.MockAddUserSuc = true

	res := m.AddUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_UpdateUser(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var u db.User
	u.Email = "test@test.com"
	u.Password = "test"
	u.WebEnabled = true

	cdb.MockUpdateUserSuc = true
	var ou db.User
	ou.Email = "test@test.com"
	cdb.MockUser = &ou

	res := m.UpdateUser(&u)
	if !res.Success {
		t.Fail()
	}
}

func TestC2Manager_GetUser(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var en db.User
	en.Email = "test@test.com"

	cdb.MockUser = &en

	res := m.GetUser("test@test.com")
	fmt.Println("user: ", *res)
	if res.Email != "test@test.com" {
		t.Fail()
	}
}

func TestC2Manager_Login(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var en db.User
	en.Email = "test@test.com"
	suc, hpw := c2m.hashPassword("test")
	if suc {
		en.Password = hpw
	}

	cdb.MockUser = &en

	res := m.Login("test@test.com", "test")
	fmt.Println("user: ", *res)
	if res.Email != "test@test.com" || !res.Success {
		t.Fail()
	}
}
