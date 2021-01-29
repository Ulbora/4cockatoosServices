package managers

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	db "github.com/Ulbora/cocka2notesServices/cocka2db"
	ml "github.com/Ulbora/go-mail-sender"
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

func TestC2Manager_randStringRunes(t *testing.T) {

	var c2m C2Manager

	pw := c2m.randStringRunes(15)
	fmt.Println("new pw:", pw)
	if pw == "" {
		t.Fail()
	}
}

func TestC2Manager_ResetPassword(t *testing.T) {
	var cdb db.MockC2DB
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cdb.Log = &l

	var c2m C2Manager
	c2m.Db = &cdb
	c2m.Log = &l
	m := c2m.GetNew()

	var n db.MailServer
	n.ID = 5
	n.SenderEmail = "test@test.com"
	n.Password = "dGVzdGVy"

	cdb.MockMailServer = &n
	cdb.MockUpdateUserSuc = true
	var mkusr db.User
	mkusr.Email = "test@test.com"
	cdb.MockUser = &mkusr

	var msender ml.MockSecureSender
	msender.MockSuccess = true
	c2m.MailSender = msender.GetNew()

	res := m.ResetPassword("test@test.com")

	if !res.Success {
		t.Fail()
	}

}
